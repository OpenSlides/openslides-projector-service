package database

import (
	"encoding/json"
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-projector-service/pkg/models"
	"github.com/rs/zerolog/log"
)

type subscription[V any] struct {
	db            *Datastore
	updateChannel chan map[string]map[string]string
	load          func() error
	Channel       V
	Unsubscribe   func()
}

func (s *subscription[V]) Reload() error {
	return s.load()
}

func (q *query[T, PT]) Subscribe() *subscription[<-chan map[string]map[string]interface{}] {
	notifyChannel := make(chan map[string]map[string]interface{})
	updateChannel := make(chan map[string]map[string]string)
	listener := queryChangeListener{
		fqids:   q.fqids,
		fields:  q.Fields,
		channel: updateChannel,
	}
	q.datastore.change.AddListener <- &listener
	go func() {
		for update := range updateChannel {
			next := make(map[string]map[string]interface{})
			for fqid, fields := range update {
				if fields == nil {
					next[fqid] = nil
					continue
				}

				next[fqid] = make(map[string]interface{})
				for key, val := range fields {
					var fieldData interface{}
					err := json.Unmarshal([]byte(val), &fieldData)
					if err != nil {
						log.Error().Err(err).Msgf("parsing subscription field `%s` for fqid `%s` with value %s", key, fqid, val)
					}

					next[fqid][key] = fieldData
				}
			}
			notifyChannel <- next
		}

		q.datastore.change.RemoveListener <- updateChannel
	}()

	return &subscription[<-chan map[string]map[string]interface{}]{q.datastore, updateChannel, func() error { return nil }, notifyChannel, func() {
		q.datastore.change.RemoveListener <- updateChannel
	}}
}

func (q *query[T, PT]) SubscribeOne(model PT) (*subscription[<-chan []string], error) {
	notifyChannel := make(chan []string, 1)
	updateChannel := make(chan map[string]map[string]string)
	listener := queryChangeListener{
		fqids:   q.fqids,
		fields:  q.Fields,
		channel: updateChannel,
	}
	if len(q.subquerys) != 0 {
		listener.fqids = []string{}
	}

	q.datastore.change.AddListener <- &listener

	load := func() error {
		if len(q.subquerys) != 0 {
			listener.fqids = []string{}
		} else {
			listener.fqids = q.fqids
		}

		data, err := q.GetOne()
		if err != nil {
			return fmt.Errorf("failed to fetch data from db: %w", err)
		}
		*model = *data
		notifyChannel <- []string{}
		return nil
	}

	if len(q.fqids) > 0 {
		if err := load(); err != nil {
			return nil, err
		}
	}

	go func() {
		for update := range updateChannel {
			if len(q.fqids) == 0 {
				continue
			}

			updatedAny := false
			if obj, ok := update[q.fqids[0]]; ok {
				if obj == nil {
					close(notifyChannel)
					break
				}

				if err := model.Update(obj); err != nil {
					log.Error().Err(err).Msg("updating subscribed model failed")
				}

				updatedAny = true
			}

			for field, subQuery := range q.subquerys {
				update, err := q.recursiveUpdateSubqueries(model.GetRelatedModelsAccessor(), field, subQuery, update)
				updatedAny = updatedAny || update
				if err != nil {
					log.Err(err).Msg("Could not update subscribed subqueries")
					continue
				}
			}

			if updatedAny {
				notifyChannel <- slices.Collect(maps.Keys(update[q.fqids[0]]))
			}
		}

		q.datastore.change.RemoveListener <- updateChannel
	}()

	return &subscription[<-chan []string]{q.datastore, updateChannel, load, notifyChannel, func() {
		q.datastore.change.RemoveListener <- updateChannel
	}}, nil
}

func (q *query[T, PT]) recursiveUpdateSubqueries(el *models.RelatedModelsAccessor, field string, subQuery *recursiveSubqueryList, update map[string]map[string]string) (bool, error) {
	subQuery.fqids = slices.DeleteFunc(subQuery.fqids, func(fqid string) bool {
		return !slices.Contains(el.GetFqids(field), fqid)
	})

	added := []string{}
	updatedAny := false
	for _, fqid := range el.GetFqids(field) {
		if !slices.Contains(subQuery.fqids, fqid) {
			added = append(added, fqid)
		} else {
			p := strings.Split(fqid, "/")
			id, err := strconv.Atoi(p[1])
			if err != nil {
				return false, err
			}

			model := el.GetRelated(field, id)
			if model != nil {
				if obj, ok := update[fqid]; ok {
					updatedAny = true
					err := model.Update(obj)
					if err != nil {
						return false, err
					}
				}

				for sField, nextSubQuery := range subQuery.subquerys {
					updated, err := q.recursiveUpdateSubqueries(model, sField, nextSubQuery, update)
					updatedAny = updated || updatedAny
					if err != nil {
						return false, err
					}
				}
			}
		}
	}

	subDsResult, err := q.datastore.getFull(added)
	if err != nil {
		return false, err
	}

	if len(subDsResult) != 0 {
		updatedAny = true
		for _, dsResult := range subDsResult {
			model, err := el.SetRelatedJSON(field, dsResult)
			if err != nil {
				log.Err(err).Msgf("Failed to parse JSON (%s)", string(dsResult))
			} else if model != nil {
				for sField, nextSubQuery := range subQuery.subquerys {
					err := q.recursiveLoadSubqueries(model, sField, nextSubQuery)
					if err != nil {
						return false, err
					}
				}
			}
		}
	}

	return len(subDsResult) != 0 || updatedAny, nil
}

func (q *query[T, PT]) SubscribeField(field interface{}) (*subscription[<-chan struct{}], error) {
	notifyChannel := make(chan struct{})
	updateChannel := make(chan map[string]map[string]string)
	listener := queryChangeListener{
		fqids:   q.fqids,
		fields:  q.Fields,
		channel: updateChannel,
	}
	q.datastore.change.AddListener <- &listener

	val := reflect.ValueOf(field)
	if val.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("value passed to SubscribeField must be a pointer")
	}

	data, err := q.GetOne()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch inital data from db: %w", err)
	}

	go func() {
		if data != nil {
			val.Elem().Set(reflect.ValueOf(data.Get(q.Fields[0])))
			notifyChannel <- struct{}{}
		}

		for update := range updateChannel {
			if obj, ok := update[q.fqids[0]]; ok {
				if val, ok := obj[q.Fields[0]]; ok {
					err := json.Unmarshal([]byte(val), field)
					if err != nil {
						log.Error().Err(err).Msg("parsing subscription field")
					}

					notifyChannel <- struct{}{}
				}
			}
		}

		q.datastore.change.RemoveListener <- updateChannel
	}()

	return &subscription[<-chan struct{}]{q.datastore, updateChannel, func() error { return nil }, notifyChannel, func() {
		q.datastore.change.RemoveListener <- updateChannel
	}}, nil
}

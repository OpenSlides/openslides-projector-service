package database

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/OpenSlides/openslides-projector-service/pkg/models"
	"github.com/iancoleman/strcase"
	"github.com/rs/zerolog/log"
)

type baseModelPtr[T any] interface {
	models.IBaseModel
	*T
}

type recursiveSubqueryList struct {
	subquerys map[string]*recursiveSubqueryList
	Fields    []string
	fqids     []string
}

func (q *recursiveSubqueryList) With(idField string, fields []string) *recursiveSubqueryList {
	sq, ok := q.subquerys[idField]
	if !ok {
		sq = &recursiveSubqueryList{}
	}
	sq.Fields = fields
	sq.subquerys = map[string]*recursiveSubqueryList{}
	q.subquerys[idField] = sq

	return sq
}

type query[T any, PT baseModelPtr[T]] struct {
	collection PT
	datastore  *Datastore
	Fields     []string
	fqids      []string
	subquerys  map[string]*recursiveSubqueryList
}

func (q *query[T, PT]) SetFqids(ids ...string) *query[T, PT] {
	q.fqids = append([]string{}, ids...)

	return q
}

func (q *query[T, PT]) With(idField string, fields []string) *query[T, PT] {
	sq, ok := q.subquerys[idField]
	if !ok {
		sq = &recursiveSubqueryList{}
	}
	sq.Fields = fields
	sq.subquerys = map[string]*recursiveSubqueryList{}

	q.subquerys[idField] = sq

	return q
}

func (q *query[T, PT]) GetSubquery(idField string) *recursiveSubqueryList {
	sq, ok := q.subquerys[idField]
	if !ok {
		sq = &recursiveSubqueryList{}
	}

	return sq
}

// Sets the fqids of the query by plain ids
func (q *query[T, PT]) SetIds(ids ...int) *query[T, PT] {
	typeName := strcase.ToSnake(reflect.ValueOf(q.collection).Elem().Type().Name())

	q.fqids = []string{}
	for _, id := range ids {
		q.fqids = append(q.fqids, fmt.Sprintf("%s/%d", typeName, id))
	}

	return q
}

// Sets the fqids of the query by plain ids
func (q *query[T, PT]) SetFields(fields ...string) *query[T, PT] {
	q.Fields = append([]string{}, fields...)

	return q
}

func (q *query[T, PT]) GetOne() (PT, error) {
	result, err := q.resultStructs()
	if err != nil {
		return nil, fmt.Errorf("getting single result as struct: %w", err)
	}

	for _, fqid := range q.fqids {
		if val, ok := result[fqid]; ok {
			return val, nil
		}
	}

	return nil, nil
}

func (q *query[T, PT]) Get() (map[string]PT, error) {
	result, err := q.resultStructs()
	if err != nil {
		return nil, fmt.Errorf("getting results as structs: %w", err)
	}

	return result, nil
}

func (q *query[T, PT]) GetMaps() (map[string]map[string]interface{}, error) {
	result, err := q.resultMaps()
	if err != nil {
		return nil, fmt.Errorf("getting results as maps: %w", err)
	}

	return result, nil
}

func (q *query[T, PT]) resultMaps() (map[string]map[string]interface{}, error) {
	data, err := q.datastore.getKeys(q.fqids, q.Fields)
	if err != nil {
		return nil, err
	}

	results := map[string]map[string]interface{}{}
	for fqid, row := range data {
		result := map[string]interface{}{}
		if err := json.Unmarshal(row, &result); err != nil {
			return nil, err
		}
		// TODO: Fix oversized fields
		results[fqid] = result
	}

	return results, nil
}

func (q *query[T, PT]) resultStructs() (map[string]PT, error) {
	dsResults, err := q.datastore.getFull(q.fqids)
	if err != nil {
		return nil, err
	}

	results := map[string]PT{}
	for fqid, dsResult := range dsResults {
		var el PT
		err := json.Unmarshal(dsResult, &el)
		if err != nil {
			return nil, err
		}
		for field, subQuery := range q.subquerys {
			err := q.recursiveLoadSubqueries(el.GetRelatedModelsAccessor(), field, subQuery)
			if err != nil {
				return nil, err
			}
		}
		results[fqid] = el
	}

	return results, nil
}

func (q *query[T, PT]) recursiveLoadSubqueries(el *models.RelatedModelsAccessor, field string, subQuery *recursiveSubqueryList) error {
	subQuery.fqids = el.GetFqids(field)
	subDsResult, err := q.datastore.getFull(subQuery.fqids)
	if err != nil {
		return err
	}

	if len(subDsResult) == 0 {
		el.SetRelated(field, nil)
	} else {
		for _, dsResult := range subDsResult {
			model, err := el.SetRelatedJSON(field, dsResult)
			if err != nil {
				log.Err(err).Msgf("Failed to parse JSON (%s)", string(dsResult))
			} else if model != nil {
				for sField, nextSubQuery := range subQuery.subquerys {
					err := q.recursiveLoadSubqueries(model, sField, nextSubQuery)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

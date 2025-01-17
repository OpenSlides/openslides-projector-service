package projector

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"slices"
	"strconv"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
	"github.com/OpenSlides/openslides-projector-service/pkg/projector/slide"
	"github.com/rs/zerolog/log"
)

type projector struct {
	ctxCancel      context.CancelFunc
	db             *datastore.Datastore
	slideRouter    *slide.SlideRouter
	projector      *models.Projector
	listeners      []chan *ProjectorUpdateEvent
	Content        string
	Projections    map[int]template.HTML
	AddListener    chan chan *ProjectorUpdateEvent
	RemoveListener chan (<-chan *ProjectorUpdateEvent)
}

type ProjectorUpdateEvent struct {
	Event string
	Data  string
}

func newProjector(id int, db *datastore.Datastore) (*projector, error) {
	projectorQuery := datastore.Collection(db, &models.Projector{}).SetIds(id)
	data, err := projectorQuery.GetOne()
	if err != nil {
		return nil, fmt.Errorf("error fetching projector from db %w", err)
	}

	if data == nil {
		return nil, fmt.Errorf("projector not found")
	}

	ctx, cancel := context.WithCancel(context.Background())
	p := &projector{
		ctxCancel:      cancel,
		db:             db,
		projector:      data,
		slideRouter:    slide.New(ctx, db),
		Projections:    make(map[int]template.HTML),
		AddListener:    make(chan chan *ProjectorUpdateEvent),
		RemoveListener: make(chan (<-chan *ProjectorUpdateEvent)),
	}

	if err = p.updateFullContent(); err != nil {
		return nil, fmt.Errorf("error generating projector content: %w", err)
	}

	go p.subscribeProjector()

	if len(p.projector.CurrentProjectionIDs) > 0 {
		initListener := make(chan *ProjectorUpdateEvent)
		p.AddListener <- initListener
		for event := range initListener {
			if event.Event == "projection-updated" {
				break
			}
		}
		p.RemoveListener <- initListener
	}

	return p, nil
}

func (p *projector) subscribeProjector() {
	defer p.ctxCancel()
	// TODO: Subscribe on projector settings updates
	// Ignore e.g. projection defaults and [...]_projection_ids
	// If header active: Meeting name + description need to be subscribed
	projectorSub, err := datastore.Collection(p.db, &models.Projector{}).SetIds(p.projector.ID).SubscribeOne(p.projector)
	if err != nil {
		log.Fatal().Err(err).Msg("could not open projector subscription")
	}
	defer projectorSub.Unsubscribe()

	projectionUpdate, projections, unsubscibeProjections, err := p.getProjectionSubscription()
	if err != nil {
		log.Fatal().Err(err).Msg("could not open projection subscription")
	}
	defer unsubscibeProjections()

	for {
		select {
		case listener := <-p.AddListener:
			p.listeners = append(p.listeners, listener)
			listener <- &ProjectorUpdateEvent{
				Event: "connected",
				Data:  "",
			}
		case listener := <-p.RemoveListener:
			i := slices.IndexFunc(p.listeners, func(el chan *ProjectorUpdateEvent) bool { return el == listener })
			if i > -1 {
				close(p.listeners[i])
				p.listeners[i] = p.listeners[len(p.listeners)-1]
				p.listeners = p.listeners[:len(p.listeners)-1]
			}
		case updatedFields, ok := <-projectorSub.Channel:
			if !ok {
				p.sendToAll(&ProjectorUpdateEvent{"deleted", ""})
				return
			}

			updatedData := map[string]interface{}{}
			for _, field := range updatedFields {
				updatedData[field] = p.projector.Get(field)
			}
			encodedData, err := json.Marshal(updatedData)
			if err != nil {
				log.Error().Err(err).Msg("could not encode projector data")
			} else {
				p.sendToAll(&ProjectorUpdateEvent{"settings", string(encodedData)})
			}

			if err = p.updateFullContent(); err != nil {
				log.Error().Err(err).Msg("error generating projector content after settings update")
			}
		case data, ok := <-projectionUpdate:
			if !ok {
				return
			}

			p.processProjectionUpdate(data, projections)
		}
	}
}

func (p *projector) processProjectionUpdate(updated []int, projections map[int]string) {
	if updated == nil {
		return
	}

	updatedProjections := map[int]string{}
	for _, projectionId := range updated {
		if projection, ok := projections[projectionId]; ok {
			p.Projections[projectionId] = template.HTML(projection)
			updatedProjections[projectionId] = projection
		} else {
			delete(p.Projections, projectionId)
			defer p.sendToAll(&ProjectorUpdateEvent{"projection-deleted", strconv.Itoa(projectionId)})
		}
	}

	if len(updatedProjections) > 0 {
		eventContent, err := json.Marshal(updatedProjections)
		if err != nil {
			log.Error().Err(err).Msg("failed to encode update event")
		} else {
			p.sendToAll(&ProjectorUpdateEvent{"projection-updated", string(eventContent)})
		}
	}

	if err := p.updateFullContent(); err != nil {
		log.Error().Err(err).Msg("failed to generate projector content")
	}
}

func (p *projector) sendToAll(event *ProjectorUpdateEvent) {
	for _, listener := range p.listeners {
		listener <- event
	}
}

func (p *projector) updateFullContent() error {
	tmpl, err := template.ParseFiles("templates/projector.html")
	if err != nil {
		return fmt.Errorf("error reading projector template %w", err)
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"Projector":   p.projector,
		"Projections": p.Projections,
	})
	if err != nil {
		return fmt.Errorf("error generating projector template %w", err)
	}

	p.Content = content.String()

	return nil
}

func (p *projector) getProjectionSubscription() (<-chan []int, map[int]string, func(), error) {
	updateChannel := make(chan []int)
	projections := make(map[int]string)
	addProjection := make(chan int)
	removeProjection := make(chan int)
	var projectionIDs []int
	sub, err := datastore.Collection(p.db, &models.Projector{}).SetIds(p.projector.ID).SetFields("current_projection_ids").SubscribeField(&projectionIDs)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to subscibe projection ids: %w", err)
	}

	go func() {
		defer sub.Unsubscribe()

		projectionChannel := p.slideRouter.SubscribeContent(addProjection, removeProjection)
		for {
			select {
			case <-sub.Channel:
				updated := []int{}
				for id := range projections {
					if !slices.Contains(projectionIDs, id) {
						updated = append(updated, id)
						removeProjection <- id
						delete(projections, id)
					}
				}

				for _, id := range projectionIDs {
					if _, ok := projections[id]; !ok {
						addProjection <- id
					}
				}

				if len(updated) > 0 || len(projectionIDs) == 0 {
					updateChannel <- updated
				}
			case update := <-projectionChannel:
				projections[update.ID] = update.Content
				updateChannel <- []int{update.ID}
			}
		}
	}()

	return updateChannel, projections, func() {
		// TODO: Unsubscribe
	}, nil
}

package slide

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
)

type projectionRequest struct {
	Projection *models.Projection
	DB         *datastore.Datastore
}

type projectionUpdate struct {
	ID      int
	Content string
}

type slideHandler func(context.Context, *projectionRequest) (<-chan string, error)

type SlideRouter struct {
	ctx    context.Context
	db     *datastore.Datastore
	Routes map[string]slideHandler
}

func New(ctx context.Context, db *datastore.Datastore) *SlideRouter {
	routes := make(map[string]slideHandler)
	routes["topic"] = TopicSlideHandler
	routes["current_list_of_speakers"] = CurrentListOfSpeakersSlideHandler
	routes["current_speaker_chyron"] = CurrentSpeakerChyronSlideHandler

	return &SlideRouter{
		ctx:    ctx,
		db:     db,
		Routes: routes,
	}
}

func (r *SlideRouter) SubscribeContent(addProjection <-chan int, removeProjection <-chan int) <-chan *projectionUpdate {
	updateChannel := make(chan *projectionUpdate)
	closeChannels := make(map[int]chan struct{})

	go func() {
		for {
			select {
			case <-r.ctx.Done():
				close(updateChannel)
				for id := range closeChannels {
					closeChannels[id] <- struct{}{}
					close(closeChannels[id])
				}

				return
			case id := <-addProjection:
				if _, ok := closeChannels[id]; !ok {
					closeChannels[id] = make(chan struct{})
					go r.subscribeProjection(id, updateChannel, closeChannels[id])
				}
			case id := <-removeProjection:
				closeChannels[id] <- struct{}{}
				close(closeChannels[id])
				delete(closeChannels, id)
			}
		}
	}()

	return updateChannel
}

func (r *SlideRouter) subscribeProjection(id int, updateChannel chan<- *projectionUpdate, closeSubscription <-chan struct{}) {
	projection, err := datastore.Collection(r.db, &models.Projection{}).SetIds(id).SetFields("id", "content_object_id", "type").GetOne()
	if err != nil {
		log.Error().Err(err).Msg("getting projection type and content object from db")
		return
	}

	projectionType := getProjectionType(projection)
	if handler, ok := r.Routes[projectionType]; ok {
		projectionChan, err := handler(r.ctx, &projectionRequest{
			Projection: projection,
			DB:         r.db,
		})

		if err != nil {
			log.Error().Err(err).Msg("failed initialize projection handler")
			return
		}

		for {
			select {
			case <-closeSubscription:
				return
			case projectionContent, ok := <-projectionChan:
				if !ok {
					return
				}

				updateChannel <- &projectionUpdate{
					ID:      id,
					Content: projectionContent,
				}
			}
		}
	} else {
		log.Warn().Msgf("unknown projection type %s", projectionType)
		updateChannel <- &projectionUpdate{
			ID:      id,
			Content: "",
		}
	}
}

func getProjectionType(projection *models.Projection) string {
	if projection.Type != nil {
		return *projection.Type
	}

	collection, _, found := strings.Cut(projection.ContentObjectID, "/")
	if found {
		return collection
	}

	return "unknown"
}

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

	return &SlideRouter{
		ctx:    ctx,
		db:     db,
		Routes: routes,
	}
}

func (r *SlideRouter) SubscribeContent(addProjection <-chan int, removeProjection <-chan int) <-chan *projectionUpdate {
	updateChannel := make(chan *projectionUpdate)
	contextCancel := make(map[int]context.CancelFunc)

	go func() {
		for {
			select {
			case <-r.ctx.Done():
				close(updateChannel)
				return
			case id := <-addProjection:
				if _, ok := contextCancel[id]; !ok {
					ctx, cancel := context.WithCancel(r.ctx)
					contextCancel[id] = cancel
					go r.subscribeProjection(ctx, id, updateChannel)
				}
			case id := <-removeProjection:
				contextCancel[id]()
				delete(contextCancel, id)
			}
		}
	}()

	return updateChannel
}

func (r *SlideRouter) subscribeProjection(ctx context.Context, id int, updateChannel chan<- *projectionUpdate) {
	projection, err := datastore.Collection(r.db, &models.Projection{}).SetIds(id).SetFields("id", "content_object_id", "type").GetOne()
	if err != nil {
		log.Error().Err(err).Msg("getting projection type and content object from db")
		return
	}

	projectionType := getProjectionType(projection)
	if handler, ok := r.Routes[projectionType]; ok {
		projectionChan, err := handler(ctx, &projectionRequest{
			Projection: projection,
			DB:         r.db,
		})

		if err != nil {
			log.Error().Err(err).Msg("failed initialize projection handler")
			return
		}

		for {
			select {
			case <-ctx.Done():
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

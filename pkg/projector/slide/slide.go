package slide

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-projector-service/pkg/database"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
)

type projectionRequest struct {
	ContentObjectID *int
	Projection      *models.Projection
	DB              *database.Datastore
	Fetch           *dsfetch.Fetch
}

type projectionUpdate struct {
	ID      int
	Content string
}

type slideHandler func(context.Context, *projectionRequest) (interface{}, error)

type SlideRouter struct {
	ctx    context.Context
	db     *database.Datastore
	ds     flow.Flow
	Routes map[string]slideHandler
}

func New(ctx context.Context, db *database.Datastore, ds flow.Flow) *SlideRouter {
	routes := make(map[string]slideHandler)
	routes["topic"] = TopicSlideHandler
	// routes["current_list_of_speakers"] = CurrentListOfSpeakersSlideHandler
	// routes["current_speaker_chyron"] = CurrentSpeakerChyronSlideHandler

	return &SlideRouter{
		ctx:    ctx,
		db:     db,
		ds:     ds,
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
	projection, err := database.Collection(r.db, &models.Projection{}).SetIds(id).SetFields("id", "content_object_id", "type").GetOne()
	if err != nil {
		log.Error().Err(err).Msg("getting projection type and content object from db")
		return
	}

	projectionType, contentObjectID := getProjectionType(projection)
	if handler, ok := r.Routes[projectionType]; ok {
		r.db.NewContext(ctx, func(fetch *dsfetch.Fetch) {
			var cId *int
			if contentObjectID != 0 {
				cId = &contentObjectID
			}

			projectionContent, err := handler(ctx, &projectionRequest{
				ContentObjectID: cId,
				Projection:      projection,
				DB:              r.db,
				Fetch:           fetch,
			})

			if err != nil {
				log.Error().Err(err).Msg("failed executing projection handler")
			}

			tmpl, err := template.ParseFiles(fmt.Sprintf("templates/slides/%s.html", projectionType))
			if err != nil {
				log.Error().Err(err).Msgf("could not load %s template", projectionType)
			}

			var content bytes.Buffer
			err = tmpl.Execute(&content, projectionContent)
			if err != nil {
				log.Error().Err(err).Msgf("could not execute %s template", projectionType)
			}

			updateChannel <- &projectionUpdate{
				ID:      id,
				Content: content.String(),
			}
		})
	} else {
		log.Warn().Msgf("unknown projection type %s", projectionType)
		updateChannel <- &projectionUpdate{
			ID:      id,
			Content: "",
		}
	}
}

func getProjectionType(projection *models.Projection) (string, int) {
	if projection.Type != nil {
		return *projection.Type, 0
	}

	collection, id, found := strings.Cut(projection.ContentObjectID, "/")
	if found {
		nId, _ := strconv.Atoi(id)
		return collection, nId
	}

	return "unknown", 0
}

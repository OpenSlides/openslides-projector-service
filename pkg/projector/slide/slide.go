package slide

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/leonelquinteros/gotext"
	"github.com/rs/zerolog/log"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-projector-service/pkg/database"
)

type projectionRequest struct {
	ContentObjectID *int
	Projection      *dsmodels.Projection
	Fetch           *dsmodels.Fetch
	Locale          *gotext.Locale
}

type projectionUpdate struct {
	ID      int
	Content string
}

type slideHandler func(context.Context, *projectionRequest) (map[string]any, error)

type SlideRouter struct {
	ctx    context.Context
	db     *database.Datastore
	ds     flow.Flow
	locale *gotext.Locale
	Routes map[string]slideHandler
}

func New(ctx context.Context, db *database.Datastore, ds flow.Flow, locale *gotext.Locale) *SlideRouter {
	routes := make(map[string]slideHandler)
	routes["assignment"] = AssignmentSlideHandler
	routes["current_los"] = CurrentListOfSpeakersSlideHandler
	routes["current_speaker_chyron"] = CurrentSpeakerChyronSlideHandler
	routes["projector_countdown"] = ProjectorCountdownSlideHandler
	routes["projector_message"] = ProjectorMessageSlideHandler
	routes["topic"] = TopicSlideHandler

	return &SlideRouter{
		ctx:    ctx,
		db:     db,
		ds:     ds,
		locale: locale,
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
	r.db.NewContext(ctx, func(fetch *dsmodels.Fetch) {
		projection, err := fetch.Projection(id).First(ctx)
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Error().Err(err).Msg("getting projection from db")
			}

			return
		}

		projectionType, contentObjectID := getProjectionType(&projection)
		if handler, ok := r.Routes[projectionType]; ok {
			var cId *int
			if contentObjectID != 0 {
				cId = &contentObjectID
			}

			projectionContent, err := handler(ctx, &projectionRequest{
				ContentObjectID: cId,
				Projection:      &projection,
				Fetch:           fetch,
				Locale:          r.locale,
			})

			if projectionContent == nil {
				updateChannel <- &projectionUpdate{
					ID:      id,
					Content: "",
				}
				return
			}

			if err != nil {
				log.Error().Err(err).Msg("failed executing projection handler")
				return
			}

			tmpl, err := template.ParseFiles(fmt.Sprintf("templates/slides/%s.html", projectionType))
			if err != nil {
				log.Error().Err(err).Msgf("could not load %s template", projectionType)
				return
			}

			var content bytes.Buffer
			projectionContent["Loc"] = r.locale
			err = tmpl.Execute(&content, projectionContent)
			if err != nil {
				log.Error().Err(err).Msgf("could not execute %s template", projectionType)
				return
			}

			updateChannel <- &projectionUpdate{
				ID:      id,
				Content: content.String(),
			}
		} else {
			log.Warn().Msgf("unknown projection type %s", projectionType)
			updateChannel <- &projectionUpdate{
				ID:      id,
				Content: "",
			}
		}
	})
}

func getProjectionType(projection *dsmodels.Projection) (string, int) {
	collection, id, found := strings.Cut(projection.ContentObjectID, "/")
	if projection.Type != "" {
		collection = projection.Type
	}

	if found {
		nId, _ := strconv.Atoi(id)
		return collection, nId
	}

	return "unknown", 0
}

package slide

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"runtime/debug"
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
	routes["agenda_item_list"] = AgendaItemListSlideHandler
	routes["assignment"] = AssignmentSlideHandler
	routes["current_los"] = ListOfSpeakersSlideHandler
	routes["current_speaker_chyron"] = CurrentSpeakerChyronSlideHandler
	routes["current_speaking_structure_level"] = CurrentSpeakingStructureLevelSlideHandler
	routes["current_structure_level_list"] = CurrentStructureLevelListSlideHandler
	routes["home"] = HomeSlideHandler
	routes["list_of_speakers"] = ListOfSpeakersSlideHandler
	routes["meeting_mediafile"] = MeetingMediafileSlideHandler
	routes["motion"] = MotionSlideHandler
	routes["motion_block"] = MotionBlockSlideHandler
	routes["poll"] = PollSlideHandler
	routes["projector_countdown"] = ProjectorCountdownSlideHandler
	routes["projector_message"] = ProjectorMessageSlideHandler
	routes["topic"] = TopicSlideHandler
	routes["wifi_access_data"] = WifiAccessDataSlideHandler

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
	onError := func(err error, msg string) {
		log.Error().Err(err).Msg(msg)

		updateChannel <- &projectionUpdate{
			ID:      id,
			Content: "",
		}
	}

	r.db.NewContext(ctx, func(fetch *dsmodels.Fetch) {
		projection, err := fetch.Projection(id).First(ctx)
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				onError(err, fmt.Sprintf("getting projection %d from db", id))
			}

			return
		}

		projectionType, contentObjectID := getProjectionType(&projection)

		defer func() {
			if r := recover(); r != nil {
				var ok bool
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("pkg: %v", r)
				}

				onError(err, fmt.Sprintf("panic in slide handler: %s (%d)\n%s", projectionType, id, string(debug.Stack())))
			}
		}()

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

			if err != nil {
				onError(err, fmt.Sprintf("failed executing projection handler %s for %d", projectionType, id))
				return
			}

			if projectionContent == nil {
				updateChannel <- &projectionUpdate{
					ID:      id,
					Content: "",
				}
				return
			}

			templateName := projectionType
			if val, ok := projectionContent["_template"]; ok {
				templateName = val.(string)
			}

			tmplName := fmt.Sprintf("%s.html", templateName)
			tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
				"RenderIndex": func(i int) int {
					return i + 1
				},
				"Loc": func() *gotext.Locale {
					return r.locale
				},
			}).ParseFiles(fmt.Sprintf("templates/slides/%s.html", templateName))
			if err != nil {
				onError(err, fmt.Sprintf("could not load %s template", projectionType))
				return
			}

			var content bytes.Buffer
			err = tmpl.Lookup(tmplName).Execute(&content, projectionContent)
			if err != nil {
				onError(err, fmt.Sprintf("could not execute %s template for projection %d", projectionType, id))
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

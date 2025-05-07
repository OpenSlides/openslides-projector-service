package projector

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"slices"
	"strconv"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-projector-service/pkg/database"
	"github.com/OpenSlides/openslides-projector-service/pkg/projector/slide"
	"github.com/rs/zerolog/log"
)

type projector struct {
	ctxCancel      context.CancelFunc
	db             *database.Datastore
	slideRouter    *slide.SlideRouter
	projector      *dsmodels.Projector
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

func newProjector(parentCtx context.Context, id int, db *database.Datastore, ds flow.Flow) (*projector, error) {
	ctx, cancel := context.WithCancel(parentCtx)

	data, err := db.Fetch.Projector(id).First(ctx)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error fetching projector from db %w", err)
	}

	p := &projector{
		ctxCancel:      cancel,
		db:             db,
		projector:      &data,
		slideRouter:    slide.New(ctx, db, ds),
		Projections:    make(map[int]template.HTML),
		AddListener:    make(chan chan *ProjectorUpdateEvent),
		RemoveListener: make(chan (<-chan *ProjectorUpdateEvent)),
	}

	if err = p.updateFullContent(); err != nil {
		return nil, fmt.Errorf("error generating projector content: %w", err)
	}

	go p.subscribeProjector(ctx)

	if len(p.projector.CurrentProjectionIDs) > 0 {
		initListener := make(chan *ProjectorUpdateEvent)
		p.AddListener <- initListener
		updateCnt := 0
		for event := range initListener {
			if event.Event == "projection-updated" {
				updateCnt++
				if updateCnt >= len(p.projector.CurrentProjectionIDs) {
					break
				}
			}
		}
		p.RemoveListener <- initListener
	}

	return p, nil
}

type ProjectorSettings struct {
	Name                   string
	IsInternal             bool
	Scale                  int
	Scroll                 int
	Width                  int
	AspectRatioNumerator   int
	AspectRatioDenominator int
	Color                  string
	BackgroundColor        string
	HeaderBackgroundColor  string
	HeaderFontColor        string
	HeaderH1Color          string
	ChyronBackgroundColor  string
	ChyronBackgroundColor2 string
	ChyronFontColor        string
	ChyronFontColor2       string
	ShowHeaderFooter       bool
	ShowTitle              bool
	ShowLogo               bool
	ShowClock              bool
}

func (p *projector) subscribeProjector(ctx context.Context) {
	defer p.ctxCancel()
	// TODO: If header active: Meeting name + description need to be subscribed
	p.db.NewContext(ctx, func(f *dsmodels.Fetch) {
		var projectorSettings ProjectorSettings
		f.Projector_Name(p.projector.ID).Lazy(&projectorSettings.Name)
		f.Projector_IsInternal(p.projector.ID).Lazy(&projectorSettings.IsInternal)
		f.Projector_Scale(p.projector.ID).Lazy(&projectorSettings.Scale)
		f.Projector_Scroll(p.projector.ID).Lazy(&projectorSettings.Scroll)
		f.Projector_Width(p.projector.ID).Lazy(&projectorSettings.Width)
		f.Projector_AspectRatioNumerator(p.projector.ID).Lazy(&projectorSettings.AspectRatioNumerator)
		f.Projector_AspectRatioDenominator(p.projector.ID).Lazy(&projectorSettings.AspectRatioDenominator)
		f.Projector_Color(p.projector.ID).Lazy(&projectorSettings.Color)
		f.Projector_BackgroundColor(p.projector.ID).Lazy(&projectorSettings.BackgroundColor)
		f.Projector_HeaderBackgroundColor(p.projector.ID).Lazy(&projectorSettings.HeaderBackgroundColor)
		f.Projector_HeaderFontColor(p.projector.ID).Lazy(&projectorSettings.HeaderFontColor)
		f.Projector_HeaderH1Color(p.projector.ID).Lazy(&projectorSettings.HeaderH1Color)
		f.Projector_ChyronBackgroundColor(p.projector.ID).Lazy(&projectorSettings.ChyronBackgroundColor)
		f.Projector_ChyronBackgroundColor2(p.projector.ID).Lazy(&projectorSettings.ChyronBackgroundColor2)
		f.Projector_ChyronFontColor(p.projector.ID).Lazy(&projectorSettings.ChyronFontColor)
		f.Projector_ChyronFontColor2(p.projector.ID).Lazy(&projectorSettings.ChyronFontColor2)
		f.Projector_ShowHeaderFooter(p.projector.ID).Lazy(&projectorSettings.ShowHeaderFooter)
		f.Projector_ShowTitle(p.projector.ID).Lazy(&projectorSettings.ShowTitle)
		f.Projector_ShowLogo(p.projector.ID).Lazy(&projectorSettings.ShowLogo)
		f.Projector_ShowClock(p.projector.ID).Lazy(&projectorSettings.ShowClock)

		f.Projector_Name(p.projector.ID).Lazy(&p.projector.Name)
		f.Projector_IsInternal(p.projector.ID).Lazy(&p.projector.IsInternal)
		f.Projector_Scale(p.projector.ID).Lazy(&p.projector.Scale)
		f.Projector_Scroll(p.projector.ID).Lazy(&p.projector.Scroll)
		f.Projector_Width(p.projector.ID).Lazy(&p.projector.Width)
		f.Projector_AspectRatioNumerator(p.projector.ID).Lazy(&p.projector.AspectRatioNumerator)
		f.Projector_AspectRatioDenominator(p.projector.ID).Lazy(&p.projector.AspectRatioDenominator)
		f.Projector_Color(p.projector.ID).Lazy(&p.projector.Color)
		f.Projector_BackgroundColor(p.projector.ID).Lazy(&p.projector.BackgroundColor)
		f.Projector_HeaderBackgroundColor(p.projector.ID).Lazy(&p.projector.HeaderBackgroundColor)
		f.Projector_HeaderFontColor(p.projector.ID).Lazy(&p.projector.HeaderFontColor)
		f.Projector_HeaderH1Color(p.projector.ID).Lazy(&p.projector.HeaderH1Color)
		f.Projector_ChyronBackgroundColor(p.projector.ID).Lazy(&p.projector.ChyronBackgroundColor)
		f.Projector_ChyronBackgroundColor2(p.projector.ID).Lazy(&p.projector.ChyronBackgroundColor2)
		f.Projector_ChyronFontColor(p.projector.ID).Lazy(&p.projector.ChyronFontColor)
		f.Projector_ChyronFontColor2(p.projector.ID).Lazy(&p.projector.ChyronFontColor2)
		f.Projector_ShowHeaderFooter(p.projector.ID).Lazy(&p.projector.ShowHeaderFooter)
		f.Projector_ShowTitle(p.projector.ID).Lazy(&p.projector.ShowTitle)
		f.Projector_ShowLogo(p.projector.ID).Lazy(&p.projector.ShowLogo)
		f.Projector_ShowClock(p.projector.ID).Lazy(&p.projector.ShowClock)

		err := f.Execute(ctx)
		var doesNotExist dsfetch.DoesNotExistError
		if errors.As(err, &doesNotExist) {
			p.sendToAll(&ProjectorUpdateEvent{"deleted", ""})
			p.ctxCancel()
			return
		} else if err != nil {
			log.Error().Err(err).Msg("failed to update projector data")
			return
		}

		encodedData, err := json.Marshal(projectorSettings)
		if err != nil {
			log.Error().Err(err).Msg("could not encode projector data")
		} else {
			p.sendToAll(&ProjectorUpdateEvent{"settings", string(encodedData)})
		}

		if err = p.updateFullContent(); err != nil {
			log.Error().Err(err).Msg("error generating projector content after settings update")
		}
	})

	projectionUpdate, projections, err := p.getProjectionSubscription(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("could not open projection subscription")
	}

	for {
		select {
		case <-ctx.Done():
			return
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
		select {
		case listener <- event:
		default:
			// TODO: Check if handling makes sense
			log.Error().Msg("could not send a projection: listener queue is full")
		}
	}
}

func (p *projector) updateFullContent() error {
	tmpl, err := template.ParseFiles("templates/projector-content.html")
	if err != nil {
		return fmt.Errorf("error reading projector template %w", err)
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]any{
		"Projector":   p.projector,
		"Projections": p.Projections,
	})
	if err != nil {
		return fmt.Errorf("error generating projector template %w", err)
	}

	p.Content = content.String()

	return nil
}

func (p *projector) getProjectionSubscription(ctx context.Context) (<-chan []int, map[int]string, error) {
	updateChannel := make(chan []int)
	projections := make(map[int]string)
	addProjection := make(chan int)
	removeProjection := make(chan int)

	projectionChannel := p.slideRouter.SubscribeContent(addProjection, removeProjection)
	go func() {
		defer close(updateChannel)
		defer close(addProjection)
		defer close(removeProjection)

		p.db.NewContext(ctx, func(f *dsmodels.Fetch) {
			projectionIDs, err := f.Projector_CurrentProjectionIDs(p.projector.ID).Value(ctx)
			if err != nil {
				log.Error().Err(err).Msg("failed to subscibe projection ids")
			}

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
		})

		for {
			select {
			case <-ctx.Done():
				return
			case update := <-projectionChannel:
				projections[update.ID] = update.Content
				updateChannel <- []int{update.ID}
			}
		}
	}()

	return updateChannel, projections, nil
}

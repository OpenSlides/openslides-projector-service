package projector

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"runtime/debug"
	"slices"
	"strconv"
	"time"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-projector-service/pkg/database"
	"github.com/OpenSlides/openslides-projector-service/pkg/projector/slide"
	"github.com/leonelquinteros/gotext"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
)

type ProjectorPreviewSettings struct {
	Scale                  int    `json:"scale"`
	Scroll                 int    `json:"scroll"`
	Width                  int    `json:"width"`
	AspectRatioNumerator   int    `json:"aspect_ratio_numerator"`
	AspectRatioDenominator int    `json:"aspect_ratio_denominator"`
	Color                  string `json:"color"`
	BackgroundColor        string `json:"background_color"`
	HeaderBackgroundColor  string `json:"header_background_color"`
	HeaderFontColor        string `json:"header_font_color"`
	HeaderH1Color          string `json:"header_h1_color"`
	ChyronBackgroundColor  string `json:"chyron_background_color"`
	ChyronBackgroundColor2 string `json:"chyron_background_color2"`
	ChyronFontColor        string `json:"chyron_font_color"`
	ChyronFontColor2       string `json:"chyron_font_color_2"`
	ShowHeaderFooter       bool   `json:"show_header_footer"`
	ShowTitle              bool   `json:"show_title"`
	ShowLogo               bool   `json:"show_logo"`
	ShowClock              bool   `json:"show_clock"`
}

type ProjectorSettings struct {
	MeetingName            string
	MeetingDescription     string
	MeetingLogo            int
	HeaderImage            int
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
	Theme                  dsmodels.Theme
}

type projector struct {
	ctxCancel          context.CancelFunc
	db                 *database.Datastore
	slideRouter        *slide.SlideRouter
	projector          *dsmodels.Projector
	pSettings          *ProjectorSettings
	pSettingsOverwrite *ProjectorPreviewSettings
	listeners          []chan *ProjectorUpdateEvent
	locale             *gotext.Locale
	Content            string
	Projections        map[int]template.HTML
	ProjectionsHash    map[int]uint64
	AddListener        chan chan *ProjectorUpdateEvent
	RemoveListener     chan (<-chan *ProjectorUpdateEvent)
}

type ProjectorUpdateEvent struct {
	Event string
	Data  string
}

func newProjector(parentCtx context.Context, id int, lang language.Tag, db *database.Datastore, ds flow.Flow) (*projector, error) {
	ctx, cancel := context.WithCancel(parentCtx)

	data, err := db.Fetch.Projector(id).First(ctx)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error fetching projector from db %w", err)
	}

	langName, _ := lang.Base()
	locale := gotext.NewLocale("locale", langName.String())
	p := &projector{
		ctxCancel:       cancel,
		db:              db,
		projector:       &data,
		pSettings:       &ProjectorSettings{},
		slideRouter:     slide.New(ctx, db, ds, locale),
		locale:          locale,
		Projections:     make(map[int]template.HTML),
		ProjectionsHash: make(map[int]uint64),
		AddListener:     make(chan chan *ProjectorUpdateEvent),
		RemoveListener:  make(chan (<-chan *ProjectorUpdateEvent)),
	}

	p.initProjector(ctx)

	return p, nil
}

func projectorPreview(ctx context.Context, id int, lang language.Tag, db *database.Datastore, ds flow.Flow, settings ProjectorPreviewSettings) (string, error) {
	ctx, cancel := context.WithCancel(ctx)

	data, err := db.Fetch.Projector(id).First(ctx)
	if err != nil {
		cancel()
		return "", fmt.Errorf("error fetching projector from db %w", err)
	}

	langName, _ := lang.Base()
	locale := gotext.NewLocale("locale", langName.String())
	p := &projector{
		ctxCancel:          cancel,
		db:                 db,
		projector:          &data,
		pSettings:          &ProjectorSettings{},
		pSettingsOverwrite: &settings,
		slideRouter:        slide.New(ctx, db, ds, locale),
		locale:             locale,
		Projections:        make(map[int]template.HTML),
		ProjectionsHash:    make(map[int]uint64),
		AddListener:        make(chan chan *ProjectorUpdateEvent),
		RemoveListener:     make(chan (<-chan *ProjectorUpdateEvent)),
	}

	p.initProjector(ctx)
	content := p.Content

	cancel()
	return content, nil
}

func (p *projector) initProjector(ctx context.Context) {
	p.locale.AddDomain("default")
	go p.subscribeProjector(ctx)

	initListener := make(chan *ProjectorUpdateEvent)
	p.AddListener <- initListener
	updateCnt := 0
	for event := range initListener {
		if event.Event == "projection-updated" {
			updateCnt++
			if updateCnt >= len(p.projector.CurrentProjectionIDs) {
				break
			}
		} else if len(p.projector.CurrentProjectionIDs) == 0 {
			break
		}
	}
	p.RemoveListener <- initListener
}

func (p *projector) subscribeProjector(ctx context.Context) {
	defer p.ctxCancel()
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}

			log.Err(err).Msgf("panic on projector: %d\n%s", p.projector.ID, string(debug.Stack()))
		}
	}()

	p.subscribeSettings(ctx)

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
				Data:  strconv.Itoa(int(time.Now().Unix())),
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

func (p *projector) subscribeSettings(ctx context.Context) {
	p.db.NewContext(ctx, func(f *dsmodels.Fetch) {
		f.Projector_Name(p.projector.ID).Lazy(&p.pSettings.Name)
		f.Projector_IsInternal(p.projector.ID).Lazy(&p.pSettings.IsInternal)
		if p.pSettingsOverwrite == nil {
			f.Projector_Scale(p.projector.ID).Lazy(&p.pSettings.Scale)
			f.Projector_Scroll(p.projector.ID).Lazy(&p.pSettings.Scroll)
			f.Projector_Width(p.projector.ID).Lazy(&p.pSettings.Width)
			f.Projector_AspectRatioNumerator(p.projector.ID).Lazy(&p.pSettings.AspectRatioNumerator)
			f.Projector_AspectRatioDenominator(p.projector.ID).Lazy(&p.pSettings.AspectRatioDenominator)
			f.Projector_Color(p.projector.ID).Lazy(&p.pSettings.Color)
			f.Projector_BackgroundColor(p.projector.ID).Lazy(&p.pSettings.BackgroundColor)
			f.Projector_HeaderBackgroundColor(p.projector.ID).Lazy(&p.pSettings.HeaderBackgroundColor)
			f.Projector_HeaderFontColor(p.projector.ID).Lazy(&p.pSettings.HeaderFontColor)
			f.Projector_HeaderH1Color(p.projector.ID).Lazy(&p.pSettings.HeaderH1Color)
			f.Projector_ChyronBackgroundColor(p.projector.ID).Lazy(&p.pSettings.ChyronBackgroundColor)
			f.Projector_ChyronBackgroundColor2(p.projector.ID).Lazy(&p.pSettings.ChyronBackgroundColor2)
			f.Projector_ChyronFontColor(p.projector.ID).Lazy(&p.pSettings.ChyronFontColor)
			f.Projector_ChyronFontColor2(p.projector.ID).Lazy(&p.pSettings.ChyronFontColor2)
			f.Projector_ShowHeaderFooter(p.projector.ID).Lazy(&p.pSettings.ShowHeaderFooter)
			f.Projector_ShowTitle(p.projector.ID).Lazy(&p.pSettings.ShowTitle)
			f.Projector_ShowLogo(p.projector.ID).Lazy(&p.pSettings.ShowLogo)
			f.Projector_ShowClock(p.projector.ID).Lazy(&p.pSettings.ShowClock)
		} else {
			p.pSettings.Scale = p.pSettingsOverwrite.Scale
			p.pSettings.Scroll = p.pSettingsOverwrite.Scroll
			p.pSettings.Width = p.pSettingsOverwrite.Width
			p.pSettings.AspectRatioNumerator = p.pSettingsOverwrite.AspectRatioNumerator
			p.pSettings.AspectRatioDenominator = p.pSettingsOverwrite.AspectRatioDenominator
			p.pSettings.Color = p.pSettingsOverwrite.Color
			p.pSettings.BackgroundColor = p.pSettingsOverwrite.BackgroundColor
			p.pSettings.HeaderBackgroundColor = p.pSettingsOverwrite.HeaderBackgroundColor
			p.pSettings.HeaderFontColor = p.pSettingsOverwrite.HeaderFontColor
			p.pSettings.HeaderH1Color = p.pSettingsOverwrite.HeaderH1Color
			p.pSettings.ChyronBackgroundColor = p.pSettingsOverwrite.ChyronBackgroundColor
			p.pSettings.ChyronBackgroundColor2 = p.pSettingsOverwrite.ChyronBackgroundColor2
			p.pSettings.ChyronFontColor = p.pSettingsOverwrite.ChyronFontColor
			p.pSettings.ChyronFontColor2 = p.pSettingsOverwrite.ChyronFontColor2
			p.pSettings.ShowHeaderFooter = p.pSettingsOverwrite.ShowHeaderFooter
			p.pSettings.ShowTitle = p.pSettingsOverwrite.ShowTitle
			p.pSettings.ShowLogo = p.pSettingsOverwrite.ShowLogo
			p.pSettings.ShowClock = p.pSettingsOverwrite.ShowClock
		}

		f.Meeting_Name(p.projector.MeetingID).Lazy(&p.pSettings.MeetingName)
		f.Meeting_Description(p.projector.MeetingID).Lazy(&p.pSettings.MeetingDescription)

		var logo dsfetch.Maybe[int]
		f.Meeting_LogoProjectorMainID(p.projector.MeetingID).Lazy(&logo)
		var header dsfetch.Maybe[int]
		f.Meeting_LogoProjectorHeaderID(p.projector.MeetingID).Lazy(&header)

		var themeId int
		f.Organization_ThemeID(1).Lazy(&themeId)

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

		if val, set := logo.Value(); set {
			p.pSettings.MeetingLogo = val
		}

		if val, set := header.Value(); set {
			p.pSettings.HeaderImage = val
		}

		p.pSettings.Theme, err = f.Theme(themeId).First(ctx)
		if err != nil {
			log.Error().Err(err).Msg("failed to load theme")
			return
		}

		encodedData, err := json.Marshal(p.pSettings)
		if err != nil {
			log.Error().Err(err).Msg("could not encode projector data")
		} else {
			p.sendToAll(&ProjectorUpdateEvent{"settings", string(encodedData)})
		}

		if err = p.updateFullContent(); err != nil {
			log.Error().Err(err).Msg("error generating projector content after settings update")
		}
	})
}

func (p *projector) processProjectionUpdate(updated []int, projections map[int]string) {
	if updated == nil {
		return
	}

	updatedProjections := map[int]string{}
	deletionOccured := false
	for _, projectionId := range updated {
		if projection, ok := projections[projectionId]; ok {
			newHash := djb2(projection)
			oldHash, exists := p.ProjectionsHash[projectionId]

			if !exists || oldHash != newHash {
				p.Projections[projectionId] = template.HTML(projection)
				p.ProjectionsHash[projectionId] = newHash
				updatedProjections[projectionId] = projection
			}
		} else {
			delete(p.Projections, projectionId)
			delete(p.ProjectionsHash, projectionId)
			defer p.sendToAll(&ProjectorUpdateEvent{"projection-deleted", strconv.Itoa(projectionId)})
			deletionOccured = true
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

	if len(updatedProjections) > 0 || deletionOccured {
		if err := p.updateFullContent(); err != nil {
			log.Error().Err(err).Msg("failed to generate projector content")
		}
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
		"Projector":   p.pSettings,
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
				log.Error().Err(err).Msg("failed to subscribe projection ids")
				return
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
				if update != nil {
					projections[update.ID] = update.Content
					updateChannel <- []int{update.ID}
				}
			}
		}
	}()

	return updateChannel, projections, nil
}

func djb2(str string) uint64 {
	var hash uint64 = 5381
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + uint64(str[i])
	}
	return hash
}

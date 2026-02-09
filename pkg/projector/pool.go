package projector

import (
	"context"
	"fmt"
	"sync"

	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-projector-service/pkg/database"
	"golang.org/x/text/language"
)

type ProjectorPool struct {
	ctx        context.Context
	mu         sync.Mutex
	projectors map[string]*projector
	db         *database.Datastore
	ds         flow.Flow
}

func NewProjectorPool(ctx context.Context, db *database.Datastore, ds flow.Flow) *ProjectorPool {
	return &ProjectorPool{
		ctx:        ctx,
		db:         db,
		ds:         ds,
		projectors: make(map[string]*projector),
	}
}

func (pool *ProjectorPool) readOrCreateProjector(id int, lang language.Tag) (*projector, error) {
	projectorId := fmt.Sprintf("%d_%s", id, lang)
	if projector, ok := pool.projectors[projectorId]; ok {
		return projector, nil
	}

	pool.mu.Lock()
	defer pool.mu.Unlock()
	if projector, ok := pool.projectors[projectorId]; ok {
		return projector, nil
	}

	projector, err := newProjector(pool.ctx, id, lang, pool.db, pool.ds)
	if err != nil {
		return nil, fmt.Errorf("error creating new projector: %w", err)
	}

	pool.projectors[projectorId] = projector
	return projector, nil
}

func (pool *ProjectorPool) GetProjectorContent(id int, lang language.Tag) (*string, error) {
	projector, err := pool.readOrCreateProjector(id, lang)
	if err != nil {
		return nil, fmt.Errorf("error retrieving projector content: %w", err)
	}

	return &projector.Content, nil
}

func (pool *ProjectorPool) GetProjectorPreview(id int, lang language.Tag, settings ProjectorPreviewSettings) (*string, error) {
	content, err := projectorPreview(pool.ctx, id, lang, pool.db, pool.ds, settings)
	if err != nil {
		return nil, fmt.Errorf("error retrieving projector preview content: %w", err)
	}

	return &content, err
}

func (pool *ProjectorPool) SubscribeProjectorContent(ctx context.Context, id int, lang language.Tag) (<-chan *ProjectorUpdateEvent, error) {
	projector, err := pool.readOrCreateProjector(id, lang)
	if err != nil {
		return nil, fmt.Errorf("error retrieving projector channel: %w", err)
	}

	channel := make(chan *ProjectorUpdateEvent, 10)
	projector.AddListener <- channel
	go func() {
		<-ctx.Done()
		projector.RemoveListener <- channel
	}()

	return channel, nil
}

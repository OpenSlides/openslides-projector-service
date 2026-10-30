package projector

import (
	"context"
	"fmt"
	"time"

	syncMu "sync"

	sync "github.com/zolstein/sync-map"

	"github.com/OpenSlides/openslides-projector-service/pkg/database"
	"golang.org/x/text/language"
)

type ProjectorPool struct {
	ctx        context.Context
	mu         syncMu.Mutex
	projectors sync.Map[string, *projector]
	db         *database.Datastore
}

func NewProjectorPool(ctx context.Context, cleanupInt time.Duration, db *database.Datastore) *ProjectorPool {
	pool := &ProjectorPool{
		ctx: ctx,
		db:  db,
	}

	go pool.poolCleanup(ctx, cleanupInt)

	return pool
}

func (pool *ProjectorPool) poolCleanup(ctx context.Context, cleanupInt time.Duration) {
	ticker := time.NewTicker(cleanupInt)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			pool.projectors.Range(func(id string, val *projector) bool {
				if len(val.listeners) == 0 {
					val.ctxCancel()
					pool.projectors.Delete(id)
				}

				return true
			})
		}
	}
}

func (pool *ProjectorPool) readOrCreateProjector(id int, lang language.Tag) (*projector, error) {
	projectorId := fmt.Sprintf("%d_%s", id, lang)
	if projector, ok := pool.projectors.Load(projectorId); ok {
		return projector, nil
	}

	pool.mu.Lock()
	defer pool.mu.Unlock()
	if projector, ok := pool.projectors.Load(projectorId); ok {
		return projector, nil
	}

	projector, err := newProjector(pool.ctx, id, lang, pool.db)
	if err != nil {
		return nil, fmt.Errorf("error creating new projector: %w", err)
	}

	pool.projectors.Store(projectorId, projector)
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
	content, err := projectorPreview(pool.ctx, id, lang, pool.db, settings)
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

	channel := make(chan *ProjectorUpdateEvent, 25)
	projector.AddListener <- channel
	go func() {
		<-ctx.Done()
		projector.RemoveListener <- channel
	}()

	return channel, nil
}

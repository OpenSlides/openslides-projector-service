package projector

import (
	"context"
	"fmt"
	"sync"

	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-projector-service/pkg/database"
)

type ProjectorPool struct {
	ctx        context.Context
	mu         sync.Mutex
	projectors map[int]*projector
	db         *database.Datastore
	ds         flow.Flow
}

func NewProjectorPool(ctx context.Context, db *database.Datastore, ds flow.Flow) *ProjectorPool {
	return &ProjectorPool{
		ctx:        ctx,
		db:         db,
		ds:         ds,
		projectors: make(map[int]*projector),
	}
}

func (pool *ProjectorPool) readOrCreateProjector(id int) (*projector, error) {
	if projector, ok := pool.projectors[id]; ok {
		return projector, nil
	}

	pool.mu.Lock()
	defer pool.mu.Unlock()
	if projector, ok := pool.projectors[id]; ok {
		return projector, nil
	}

	projector, err := newProjector(pool.ctx, id, pool.db, pool.ds)
	if err != nil {
		return nil, fmt.Errorf("error creating new projector: %w", err)
	}

	pool.projectors[id] = projector
	return projector, nil
}

func (pool *ProjectorPool) GetProjectorContent(id int) (*string, error) {
	projector, err := pool.readOrCreateProjector(id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving projector content: %w", err)
	}

	return &projector.Content, nil
}

func (pool *ProjectorPool) SubscribeProjectorContent(ctx context.Context, id int) (<-chan *ProjectorUpdateEvent, error) {
	projector, err := pool.readOrCreateProjector(id)
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

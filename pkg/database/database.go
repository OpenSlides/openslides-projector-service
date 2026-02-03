package database

import (
	"context"
	"sync"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-go/datastore/flow"
)

type Datastore struct {
	mu          sync.RWMutex
	ctx         context.Context
	ds          flow.Flow
	dsListeners []*dsChangeListener
	Fetch       *dsmodels.Fetch
}

func New(addr string, redisAddr string, dsFlow flow.Flow) (*Datastore, error) {
	ctx := context.Background()
	ds := Datastore{
		ctx:   context.Background(),
		ds:    dsFlow,
		Fetch: dsmodels.New(dsFlow),
	}
	go dsFlow.Update(ctx, func(m map[dskey.Key][]byte, err error) {
		hasCanceled := false
		ds.mu.RLock()
		for _, listener := range ds.dsListeners {
			if listener.ctx.Err() != nil {
				hasCanceled = true
				continue
			}

			for key := range m {
				if _, ok := listener.keys[key]; ok {
					listener.handler()
					break
				}
			}
		}
		ds.mu.RUnlock()

		if hasCanceled {
			ds.mu.Lock()
			defer ds.mu.Unlock()

			n := 0
			for _, listener := range ds.dsListeners {
				if listener.ctx.Err() == nil {
					ds.dsListeners[n] = listener
					n++
				}
			}
			ds.dsListeners = ds.dsListeners[:n]
		}
	})

	return &ds, nil
}

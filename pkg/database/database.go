package database

import (
	"context"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/flow"
)

type Datastore struct {
	ctx         context.Context
	ds          flow.Flow
	dsListeners []*dsChangeListener
	Fetch       *dsfetch.Fetch
}

func New(addr string, redisAddr string, dsFlow flow.Flow) (*Datastore, error) {
	ctx := context.Background()
	ds := Datastore{
		ctx:   context.Background(),
		ds:    dsFlow,
		Fetch: dsfetch.New(dsFlow),
	}
	go dsFlow.Update(ctx, func(m map[dskey.Key][]byte, err error) {
		for _, listener := range ds.dsListeners {
			for key := range m {
				if _, ok := listener.keys[key]; ok {
					listener.handler()
					break
				}
			}
		}
	})

	return &ds, nil
}

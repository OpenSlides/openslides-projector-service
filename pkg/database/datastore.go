package database

import (
	"context"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-go/datastore/dsrecorder"
)

type dsChangeListener struct {
	ctx     context.Context
	keys    map[dskey.Key]struct{}
	handler func()
}

func (db *Datastore) NewContext(ctx context.Context, handler func(*dsmodels.Fetch)) {
	recorder := dsrecorder.New(db.ds)
	fetch := dsmodels.New(recorder)

	handler(fetch)
	listener := dsChangeListener{
		ctx:  ctx,
		keys: recorder.Keys(),
	}

	listener.handler = func() {
		recorder.Reset()
		handler(fetch)
		listener.keys = recorder.Keys()
	}

	db.mu.Lock()
	defer db.mu.Unlock()
	db.dsListeners = append(db.dsListeners, &listener)
}

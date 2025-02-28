package database

import (
	"context"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsrecorder"
)

type dsChangeListener struct {
	ctx     context.Context
	keys    map[dskey.Key]struct{}
	handler func()
}

func (db *Datastore) NewContext(ctx context.Context, handler func(*dsfetch.Fetch)) {
	recorder := dsrecorder.New(db.ds)
	fetch := dsfetch.New(recorder)

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

	db.dsListeners = append(db.dsListeners, &listener)
}

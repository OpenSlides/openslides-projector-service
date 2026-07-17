package database

import (
	"context"
	"time"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-go/datastore/dsrecorder"
	"github.com/OpenSlides/openslides-go/throttle"
)

type dsChangeListener struct {
	ctx       context.Context
	keys      map[dskey.Key]struct{}
	handler   func()
	throttler *throttle.Throttler
}

func (l *dsChangeListener) Next() {
	l.throttler.Run(func() {
		l.handler()
	})
}

func (db *Datastore) NewContext(ctx context.Context, handler func(*dsmodels.Fetch)) {
	recorder := dsrecorder.New(db.Flow)
	fetch := dsmodels.New(recorder)

	throttler := throttle.New(ctx, 20*time.Millisecond)

	listener := dsChangeListener{
		ctx:       ctx,
		keys:      recorder.Keys(),
		throttler: throttler,
	}

	listener.handler = func() {
		recorder.Reset()
		handler(fetch)
		listener.keys = recorder.Keys()
	}

	listener.Next()

	db.mu.Lock()
	defer db.mu.Unlock()
	db.dsListeners = append(db.dsListeners, &listener)
}

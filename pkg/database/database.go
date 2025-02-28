package database

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Datastore struct {
	ctx         context.Context
	pool        *pgxpool.Pool
	redis       *redis.Client
	change      *changeListenerServer
	ds          flow.Flow
	dsListeners []*dsChangeListener
}

func New(addr string, redisAddr string, dsFlow flow.Flow) (*Datastore, error) {
	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	ds := Datastore{pool: pool, redis: rdb, ctx: context.Background(), ds: dsFlow}
	go ds.setupRedisListener()
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

func Collection[T any, PT baseModelPtr[T]](ds *Datastore, coll *T) *query[T, PT] {
	return &query[T, PT]{
		collection: coll,
		datastore:  ds,
		subquerys:  map[string]*recursiveSubqueryList{},
	}
}

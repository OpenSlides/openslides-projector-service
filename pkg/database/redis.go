package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type queryChangeListener struct {
	fqids   []string
	fields  []string
	channel chan map[string]map[string]string
}

type changeListenerServer struct {
	AddListener    chan *queryChangeListener
	RemoveListener chan (<-chan map[string]map[string]string)
}

func (ds *Datastore) setupRedisListener() {
	ds.change = &changeListenerServer{
		AddListener:    make(chan *queryChangeListener),
		RemoveListener: make(chan (<-chan map[string]map[string]string)),
	}
	c := ds.change

	changeSource := make(chan map[string]map[string]string)
	go setupChangeListener(ds.ctx, ds.redis, changeSource)

	listeners := make([]*queryChangeListener, 0)
	defer func() {
		for _, listener := range listeners {
			close(listener.channel)
		}
		close(changeSource)
	}()

	for {
		select {
		case <-ds.ctx.Done():
			return
		case listener := <-c.AddListener:
			listeners = append(listeners, listener)
		case channel := <-c.RemoveListener:
			for i, listener := range listeners {
				if listener.channel == channel {
					listeners[i] = listeners[len(listeners)-1]
					listeners = listeners[:len(listeners)-1]
					close(listener.channel)
					break
				}
			}
		case changeMap, ok := <-changeSource:
			if !ok {
				return
			}

			for _, listener := range listeners {
				listenerChanged := map[string]map[string]string{}
				if len(listener.fqids) != 0 {
					for _, fqid := range listener.fqids {
						if val, ok := changeMap[fqid]; ok {
							listenerChanged[fqid] = val
						}
					}
				} else {
					listenerChanged = changeMap
				}

				if len(listenerChanged) > 0 {
					select {
					case listener.channel <- listenerChanged:
					default:
					}
				}
			}
		}
	}
}

func setupChangeListener(ctx context.Context, conn *redis.Client, channel chan map[string]map[string]string) {
	id := "$"

	for ctx.Err() == nil {
		nextID, data, err := getNextRedisUpdate(ctx, conn, id)
		if err != nil {
			if nextID != "" {
				id = nextID
			}

			log.Error().Err(err).Msg("reading update from message bus")
			time.Sleep(5 * time.Second)
			continue
		}
		if data != nil {
			channel <- data
		}
		id = nextID
	}
}

func getNextRedisUpdate(ctx context.Context, conn *redis.Client, id string) (string, map[string]map[string]string, error) {
	reply, err := conn.XRead(ctx, &redis.XReadArgs{
		Block:   0,
		Streams: []string{"ModifiedFields", id},
		Count:   10,
	}).Result()
	if err != nil {
		return "", nil, fmt.Errorf("redis `XREAD %s BLOCK 0 STREAMS %s: %w", "ModifiedFields", id, err)
	}

	if reply == nil {
		// This happens, when the redis command times out.
		return id, nil, nil
	}

	id, data, err := parseMessageBus(reply)
	if err != nil {
		return id, nil, fmt.Errorf("parsing message bus: %w", err)
	}

	return id, data, nil
}

func parseMessageBus(streams []redis.XStream) (string, map[string]map[string]string, error) {
	data := map[string]map[string]string{}

	nextID := ""
	for _, stream := range streams {
		for _, message := range stream.Messages {
			nextID = message.ID
			for k, v := range message.Values {
				sepIndex := strings.LastIndex(k, "/")
				if sepIndex < 0 {
					log.Error().Err(fmt.Errorf("invalid fqid %s", k)).Msg("parsing message bus entry")
				}

				val := v.(string)
				if _, ok := data[k[:sepIndex]]; !ok {
					data[k[:sepIndex]] = map[string]string{}
				}

				// Element is marked as deleted
				if data[k[:sepIndex]] == nil {
					continue
				}

				// Mark item as deleted
				if k[sepIndex+1:] == "id" && val == "null" {
					data[k[:sepIndex]] = nil
				} else {
					data[k[:sepIndex]][k[sepIndex+1:]] = val
				}
			}
		}
	}

	if nextID == "" {
		return "", nil, fmt.Errorf("did not receive new id")
	}

	return nextID, data, nil
}

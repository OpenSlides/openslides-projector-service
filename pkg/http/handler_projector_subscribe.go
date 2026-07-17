package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/OpenSlides/openslides-go/throttle"
	"github.com/rs/zerolog/log"
)

func (s *projectorHttp) ProjectorSubscribeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeResponse(w, `{"error": true, "msg": "Projector id invalid"}`)
			return
		}

		content, err := s.projector.SubscribeProjectorContent(r.Context(), id, getRequestLanguage(r))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeResponse(w, `{"error": true, "msg": "Error reading projector content"}`)
			return
		}

		if content == nil {
			w.WriteHeader(http.StatusNotFound)
			writeResponse(w, `{"error": true, "msg": "Projector not found"}`)
			return
		}

		needsInit := r.URL.Query().Get("init") == "1"
		var projectorContent string
		if needsInit {
			projectorContentRaw, err := s.projector.GetProjectorContent(id, getRequestLanguage(r))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				writeResponse(w, `{"error": true, "msg": "Error reading projector content"}`)
				return
			}

			currentContent, err := json.Marshal(projectorContentRaw)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				writeResponse(w, `{"error": true, "msg": "Error encoding projector content"}`)
				return
			}
			projectorContent = string(currentContent)
		}

		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		if needsInit {
			if _, err := fmt.Fprintf(w, "event: projector-replace\ndata: %s\n\n", projectorContent); err != nil {
				log.Err(err).Msg("error sending event")
			}
		}

		throttler := throttle.New(r.Context(), 50*time.Millisecond)
		if f, ok := w.(http.Flusher); ok {
			throttler.Run(func() {
				f.Flush()
			})
		} else {
			log.Warn().Msg("connection lost during initialization")
			return
		}

		defer func() {
			f, ok := w.(http.Flusher)
			if !ok {
				log.Warn().Msg("connection lost or flusher unavailable, stopping stream")
				return
			}
			f.Flush()
		}()

		for {
			select {
			case <-r.Context().Done():
				return
			case event := <-content:
				if _, err := fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event.Event, event.Data); err != nil {
					log.Err(err).Msg("error sending event")
				}

				f, ok := w.(http.Flusher)
				if !ok {
					log.Warn().Msg("connection lost or flusher unavailable, stopping stream")
					return
				}
				throttler.Run(func() {
					f.Flush()
				})
			}
		}
	}
}

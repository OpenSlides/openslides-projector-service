package http

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *projectorHttp) ProjectorSubscribeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Check if user can access projector

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, `{"error": true, "msg": "Projector id invalid"}`)
			return
		}

		content, err := s.Projector.SubscribeProjectorContent(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error": true, "msg": "Error reading projector content"}`)
			return
		}

		if content == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, `{"error": true, "msg": "Projector not found"}`)
			return
		}

		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.(http.Flusher).Flush()

		for {
			select {
			case event := <-content:
				fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event.Event, event.Data)
				w.(http.Flusher).Flush()
			case <-r.Context().Done():
				return
			}
		}
	}
}

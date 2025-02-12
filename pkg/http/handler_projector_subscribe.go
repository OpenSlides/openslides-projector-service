package http

import (
	"encoding/json"
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

		content, err := s.projector.SubscribeProjectorContent(r.Context(), id)
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

		needsInit := r.URL.Query().Get("init") == "1"
		var projectorContent string
		if needsInit {
			projectorContentRaw, err := s.projector.GetProjectorContent(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, `{"error": true, "msg": "Error reading projector content"}`)
				return
			}

			currentContent, err := json.Marshal(projectorContentRaw)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, `{"error": true, "msg": "Error encoding projector content"}`)
				return
			}
			projectorContent = string(currentContent)
		}

		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		if needsInit {
			fmt.Fprintf(w, "event: projector-replace\ndata: %s\n\n", projectorContent)
		}
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

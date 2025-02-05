package http

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *projectorHttp) ProjectorGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, `{"error": true, "msg": "Projector id invalid"}`)
			return
		}

		content, err := s.projector.GetProjectorContent(id)
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

		fmt.Fprintln(w, *content)
	}
}

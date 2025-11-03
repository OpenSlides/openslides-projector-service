package http

import (
	"net/http"
)

func (s *projectorHttp) HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		writeResponse(w, `{"healthy": true}`)
	}
}

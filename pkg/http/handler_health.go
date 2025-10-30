package http

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (s *projectorHttp) HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("lol")
		w.Header().Set("Content-Type", "application/json")
		writeResponse(w, `{"healthy": true}`)
	}
}

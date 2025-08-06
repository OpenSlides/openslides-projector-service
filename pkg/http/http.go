package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-go/auth"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-go/environment"
	"github.com/OpenSlides/openslides-go/redis"
	"github.com/OpenSlides/openslides-projector-service/pkg/database"
	"github.com/OpenSlides/openslides-projector-service/pkg/projector"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
)

type ProjectorConfig struct {
	RestricterUrl string
}

type projectorHttp struct {
	ctx       context.Context
	serverMux *http.ServeMux
	db        *database.Datastore
	ds        flow.Flow
	projector *projector.ProjectorPool
	cfg       ProjectorConfig
	auth      *auth.Auth
}

func New(ctx context.Context, cfg ProjectorConfig, serverMux *http.ServeMux, db *database.Datastore, ds flow.Flow) {
	projectorPool := projector.NewProjectorPool(ctx, db, ds)

	lookup := new(environment.ForProduction)
	redis := redis.New(lookup)
	authService, authBackground, err := auth.New(lookup, redis)
	if err != nil {
		log.Err(err).Msg("auth error")
	}

	go authBackground(ctx, func(e error) {
		log.Err(e).Msg("auth background error")
	})

	handler := projectorHttp{
		ctx:       ctx,
		serverMux: serverMux,
		db:        db,
		ds:        ds,
		projector: projectorPool,
		auth:      authService,
		cfg:       cfg,
	}
	handler.registerRoutes(cfg)
}

func writeResponse(w http.ResponseWriter, resp string) {
	if _, err := fmt.Fprintln(w, resp); err != nil {
		log.Err(err).Msg("writing response")
	}
}

func (s *projectorHttp) registerRoutes(cfg ProjectorConfig) {
	s.serverMux.HandleFunc("/system/projector/health", s.HealthHandler())
	s.serverMux.Handle("/system/projector/get/{id}", http.HandlerFunc(s.ProjectorGetHandler()))
	s.serverMux.Handle("/system/projector/subscribe/{id}", http.HandlerFunc(s.ProjectorSubscribeHandler()))
}

var languageMatcher = language.NewMatcher([]language.Tag{
	language.English,
	language.German,
})

func getRequestLanguage(r *http.Request) language.Tag {
	lang, _ := r.Cookie("lang")
	accept := r.Header.Get("Accept-Language")
	tag, _ := language.MatchStrings(languageMatcher, lang.String(), accept)

	return tag
}

func authMiddleware(next http.Handler, auth *auth.Auth, cfg ProjectorConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := auth.Authenticate(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			writeResponse(w, `{"error": true, "msg": "authenticate request failed"}`)
			return
		}

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeResponse(w, `{"error": true, "msg": "Projector id invalid"}`)
			return
		}

		// TODO: Listen for permission changes
		body := []byte(fmt.Sprintf(`[{"collection": "projector", "ids":[%d], "fields": {"id": null}}]`, id))
		userID := auth.FromContext(ctx)
		restrictUrl := fmt.Sprintf("%s?user_id=%d&single=1", cfg.RestricterUrl, userID)
		req, err := http.NewRequest("POST", restrictUrl, bytes.NewReader(body))
		if err != nil {
			writeResponse(w, `{"error": true, "msg": "creating restriction request failed"}`)
			return
		}

		req.Header = http.Header{
			"Content-Type": {"application/json"},
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeResponse(w, `{"error": true, "msg": "restriction request failed"}`)
			return
		}

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			writeResponse(w, `{"error": true, "msg": "restriction request failed"}`)
			return
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil || !strings.Contains(string(b), fmt.Sprintf(`"projector/%d/id":%d`, id, id)) {
			w.WriteHeader(http.StatusUnauthorized)
			writeResponse(w, `{"error": true, "msg": "permissions denied"}`)
			return
		}

		if err := resp.Body.Close(); err != nil {
			log.Err(err).Msg("error closing response body")
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

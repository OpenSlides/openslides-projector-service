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

func (s *projectorHttp) registerRoutes(cfg ProjectorConfig) {
	s.serverMux.HandleFunc("/system/projector/health", s.HealthHandler())
	s.serverMux.Handle("/system/projector/get/{id}", authMiddleware(http.HandlerFunc(s.ProjectorGetHandler()), s.auth, cfg))
	s.serverMux.Handle("/system/projector/subscribe/{id}", authMiddleware(http.HandlerFunc(s.ProjectorSubscribeHandler()), s.auth, cfg))
}

func authMiddleware(next http.Handler, auth *auth.Auth, cfg ProjectorConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := auth.Authenticate(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, `{"error": true, "msg": "authenticate request failed"}`)
			return
		}

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, `{"error": true, "msg": "Projector id invalid"}`)
			return
		}

		// TODO: Listen for permission changes
		body := []byte(fmt.Sprintf(`[{"collection": "projector", "ids":[%d], "fields": {"id": null}}]`, id))
		userID := auth.FromContext(ctx)
		restrictUrl := fmt.Sprintf("%s?user_id=%d&single=1", cfg.RestricterUrl, userID)
		req, err := http.NewRequest("POST", restrictUrl, bytes.NewReader(body))
		if err != nil {
			fmt.Fprintln(w, `{"error": true, "msg": "creating restriction request failed"}`)
			return
		}

		req.Header = http.Header{
			"Content-Type": {"application/json"},
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error": true, "msg": "restriction request failed"}`)
			return
		}

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprintln(w, `{"error": true, "msg": "restriction request failed"}`)
			return
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil || !strings.Contains(string(b), fmt.Sprintf(`"projector/%d/id":%d`, id, id)) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, `{"error": true, "msg": "permissions denied"}`)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

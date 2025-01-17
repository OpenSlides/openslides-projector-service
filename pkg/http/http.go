package http

import (
	"net/http"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/projector"
)

type ProjectorHttp struct {
	ServerMux *http.ServeMux
	DS        *datastore.Datastore
	Projector *projector.ProjectorPool
}

func (s *ProjectorHttp) RegisterRoutes() {
	s.ServerMux.HandleFunc("/system/projector/health", s.HealthHandler())
	s.ServerMux.HandleFunc("/system/projector/get/{id}", s.ProjectorGetHandler())
	s.ServerMux.HandleFunc("/system/projector/subscribe/{id}", s.ProjectorSubscribeHandler())
}

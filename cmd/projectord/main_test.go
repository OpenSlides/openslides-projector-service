package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	projectorHttp "github.com/OpenSlides/openslides-projector-service/pkg/http"
)

func TestHandlerHealth(t *testing.T) {
	serverMux := http.NewServeMux()
	projectorHandler := projectorHttp.ProjectorHttp{
		ServerMux: serverMux,
	}
	projectorHandler.RegisterRoutes()
	rec := httptest.NewRecorder()

	projectorHandler.ServerMux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/system/projector/health", strings.NewReader(``)))

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "{\"healthy\": true}\n" {
		t.Errorf("expected body %q, got %q", `{"healthy": true}`, rec.Body.String())
	}
}

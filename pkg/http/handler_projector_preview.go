package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/OpenSlides/openslides-projector-service/pkg/projector"
)

func (s *projectorHttp) ProjectorPreviewHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeResponse(w, `{"error": true, "msg": "Projector id invalid"}`)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var settings projector.ProjectorPreviewSettings
		if err := decoder.Decode(&settings); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeResponse(w, fmt.Sprintf(`{"error": true, "msg": "Could not parse json", "error": "%s"}`, err.Error()))
			return
		}

		projectorContent, err := s.projector.GetProjectorPreview(id, getRequestLanguage(r), settings)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeResponse(w, `{"error": true, "msg": "Error reading projector content"}`)
			return
		}

		if projectorContent == nil {
			w.WriteHeader(http.StatusNotFound)
			writeResponse(w, `{"error": true, "msg": "Projector not found"}`)
			return
		}

		tmpl, err := template.ParseFiles("templates/projector.html")
		if err != nil {
			writeResponse(w, `{"error": true, "msg": "Error providing projector content"}`)
			return
		}

		var content bytes.Buffer
		if err := tmpl.Execute(&content, map[string]any{
			"ProjectorContent": template.HTML(*projectorContent),
		}); err != nil {
			writeResponse(w, `{"error": true, "msg": "Error providing projector content"}`)
			return
		}

		writeResponse(w, content.String())
	}
}

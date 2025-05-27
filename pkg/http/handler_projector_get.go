package http

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
)

func (s *projectorHttp) ProjectorGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeResponse(w, `{"error": true, "msg": "Projector id invalid"}`)
			return
		}

		projectorContent, err := s.projector.GetProjectorContent(id, getRequestLanguage(r))
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

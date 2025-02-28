package http

import (
	"bytes"
	"fmt"
	"html/template"
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

		projectorContent, err := s.projector.GetProjectorContent(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error": true, "msg": "Error reading projector content"}`)
			return
		}

		if projectorContent == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, `{"error": true, "msg": "Projector not found"}`)
			return
		}

		tmpl, err := template.ParseFiles("templates/projector.html")
		if err != nil {
			fmt.Fprintln(w, `{"error": true, "msg": "Error providing projector content"}`)
		}

		var content bytes.Buffer
		if err := tmpl.Execute(&content, map[string]interface{}{
			"ProjectorContent": template.HTML(*projectorContent),
		}); err != nil {
			fmt.Fprintln(w, `{"error": true, "msg": "Error providing projector content"}`)
		}

		fmt.Fprintln(w, content.String())
	}
}

package slide

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
)

func TopicSlideHandler(ctx context.Context, req *projectionRequest) (string, error) {
	tmpl, err := template.ParseFiles("templates/slides/topic.html")
	if err != nil {
		return "", fmt.Errorf("could not load topic template", err)
	}

	if req.ContentObjectID == nil {
		return "", fmt.Errorf("no topic id provided for slide %w", err)
	}

	topic, err := req.Fetch.Topic(*req.ContentObjectID).Value(ctx)
	if err != nil {
		return "", fmt.Errorf("could not load topic %w", err)
	}

	agendaItem, err := topic.AgendaItem().Value(ctx)
	if err != nil {
		return "", fmt.Errorf("could not load agenda item %w", err)
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"AgendaItem": agendaItem,
		"Topic":      topic,
		"Text":       template.HTML(topic.Text),
	})
	if err != nil {
		return "", fmt.Errorf("could not execute topic template")
	}

	return content.String(), nil
}

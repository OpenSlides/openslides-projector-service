package slide

import (
	"context"
	"fmt"
	"html/template"
)

func TopicSlideHandler(ctx context.Context, req *projectionRequest) (any, error) {
	if req.ContentObjectID == nil {
		return "", fmt.Errorf("no topic id provided for slide")
	}

	topic, err := req.Fetch.Topic(*req.ContentObjectID).Preload("AgendaItem").Load(ctx)
	if err != nil {
		return "", fmt.Errorf("could not load topic %w", err)
	}

	agendaItem := topic.AgendaItem().Get()
	return map[string]any{
		"AgendaItem": agendaItem,
		"Topic":      topic,
		"Text":       template.HTML(topic.Text),
	}, nil
}

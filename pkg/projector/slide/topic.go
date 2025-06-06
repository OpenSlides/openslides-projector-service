package slide

import (
	"context"
	"fmt"
	"html/template"
)

func TopicSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no topic id provided for slide")
	}

	t := req.Fetch.Topic(*req.ContentObjectID)
	topic, err := t.Preload(t.AgendaItem()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load topic %w", err)
	}

	return map[string]any{
		"AgendaItem": topic.AgendaItem,
		"Topic":      topic,
		"Text":       template.HTML(topic.Text),
	}, nil
}

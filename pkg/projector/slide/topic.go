package slide

import (
	"context"
	"fmt"
	"html/template"
)

func TopicSlideHandler(ctx context.Context, req *projectionRequest) (interface{}, error) {
	if req.ContentObjectID == nil {
		return "", fmt.Errorf("no topic id provided for slide")
	}

	topic, err := req.Fetch.Topic(*req.ContentObjectID).Value(ctx)
	if err != nil {
		return "", fmt.Errorf("could not load topic %w", err)
	}

	agendaItem, err := topic.AgendaItem().Value(ctx)
	if err != nil {
		return "", fmt.Errorf("could not load agenda item %w", err)
	}

	return map[string]interface{}{
		"AgendaItem": agendaItem,
		"Topic":      topic,
		"Text":       template.HTML(topic.Text),
	}, nil
}

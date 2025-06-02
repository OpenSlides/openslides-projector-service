package slide

import (
	"context"
	"fmt"
	"html/template"
)

func ProjectorMessageSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no topic id provided for slide")
	}

	message, err := req.Fetch.ProjectorMessage(*req.ContentObjectID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load projector message %w", err)
	}

	return map[string]any{
		"Message": template.HTML(message.Message),
	}, nil
}

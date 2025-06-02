package slide

import (
	"context"
	"encoding/json"
	"fmt"
)

type projectorCountdownOptions struct {
	Fullscreen  bool   `json:"fullscreen"`
	DisplayType string `json:"displayType"`
}

func ProjectorCountdownSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no topic id provided for slide")
	}

	countdown, err := req.Fetch.ProjectorCountdown(*req.ContentObjectID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load projector message %w", err)
	}

	options := projectorCountdownOptions{
		Fullscreen:  false,
		DisplayType: "",
	}
	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse slide options: %w", err)
		}
	}

	return map[string]any{
		"Countdown": countdown,
		"Options":   options,
	}, nil
}

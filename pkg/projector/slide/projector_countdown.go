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
	warningTime, err := req.Fetch.Meeting_ProjectorCountdownWarningTime(countdown.MeetingID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load projector countdown warning time %w", err)
	}
	return map[string]any{
		"Countdown":    countdown,
		"Options":      options,
		"WarningTime":  warningTime,
		"ProjectionID": req.Projection.ID,
	}, nil
}

package slide

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func PollSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	poll, err := req.Fetch.Poll(*req.ContentObjectID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll id %w", err)
	}

	title, err := viewmodels.GetContentObjectField[string](ctx, req.Fetch, "title", poll.ContentObjectID)
	if err != nil {
		return nil, fmt.Errorf("could not load poll content object title %w", err)
	}

	return map[string]any{
		"Title": title,
		"Poll":  poll,
	}, nil
}

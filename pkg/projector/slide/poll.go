package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

type pollSlideOptions struct {
	SingleVotes  bool   `json:"single_votes"`
}

func PollSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	poll, err := req.Fetch.Poll(*req.ContentObjectID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll id %w", err)
	}

	options := pollSlideOptions{
		SingleVotes:  false,
	}
	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse slide options: %w", err)
		}
	}

	if options.SingleVotes {
		return pollSingleVotesSlideHandler(ctx, req, poll)
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

func pollSingleVotesSlideHandler(ctx context.Context, req *projectionRequest, poll dsmodels.Poll) (map[string]any, error) {
	return map[string]any{
		"_template": "poll_single_vote",
	}, nil
}

package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

type pollSlideOptions struct {
	SingleVotes bool `json:"single_votes"`
}

func PollSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	options := pollSlideOptions{
		SingleVotes: false,
	}
	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse slide options: %w", err)
		}
	}

	if options.SingleVotes {
		return pollSingleVotesSlideHandler(ctx, req)
	}

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

func pollSingleVotesSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	pQ := req.Fetch.Poll()
	poll, err := req.Fetch.Poll(*req.ContentObjectID).Preload(pQ.OptionList()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll id %w", err)
	}

	title, err := viewmodels.GetContentObjectField[string](ctx, req.Fetch, "title", poll.ContentObjectID)
	if err != nil {
		return nil, fmt.Errorf("could not load poll content object title %w", err)
	}

	pollOption := dsmodels.Option{}
	if len(poll.OptionList) > 0 {
		pollOption = poll.OptionList[0]
	}

	var entitledUsersAtStop []dsmodels.MeetingUser
	if err := json.Unmarshal(poll.EntitledUsersAtStop, &entitledUsersAtStop); err != nil {
		return nil, fmt.Errorf("parse los id: %w", err)
	}

	numEntitledUsers := len(entitledUsersAtStop)
	pollMethod := map[string]bool{
		"Yes":     strings.Contains(poll.Pollmethod, "Y"),
		"No":      strings.Contains(poll.Pollmethod, "N"),
		"Abstain": strings.Contains(poll.Pollmethod, "A"),
	}

	return map[string]any{
		"_template":        "poll_single_vote",
		"_fullHeight":      true,
		"Title":            title,
		"Poll":             poll,
		"PollMethod":       pollMethod,
		"PollOption":       pollOption,
		"EntitledUsers":    entitledUsersAtStop,
		"NumEntitledUsers": numEntitledUsers,
	}, nil
}

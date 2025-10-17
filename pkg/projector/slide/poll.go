package slide

import (
	"context"
	"encoding/json"
	"fmt"
)

type pollSlideOptions struct {
	SingleVotes bool `json:"single_votes"`
}

func PollSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	pollID := *req.ContentObjectID

	var options pollSlideOptions
	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse slide options: %w", err)
		}
	}

	var pollState string
	var pollTitle string
	var pollLiveVotingEnabled bool
	req.Fetch.Poll_State(pollID).Lazy(&pollState)
	req.Fetch.Poll_Title(pollID).Lazy(&pollTitle)
	req.Fetch.Poll_LiveVotingEnabled(pollID).Lazy(&pollLiveVotingEnabled)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("could not load poll base info %w", err)
	}

	if pollState != "published" && (pollState != "started" && pollLiveVotingEnabled) {
		state := req.Locale.Get("No results yet")
		if pollState == "finished" {
			state = req.Locale.Get("Counting of votes is in progress ...")
		}

		if pollState == "started" && !pollLiveVotingEnabled {
			state = req.Locale.Get("Voting in progress")
		}

		return map[string]any{
			"Title": pollTitle,
			"State": state,
		}, nil
	}

	if options.SingleVotes {
		return pollSingleVotesSlideHandler(ctx, req)
	}

	poll, err := req.Fetch.Poll(pollID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll %w", err)
	}

	if len(poll.OptionIDs) == 1 || poll.Pollmethod == "Y" {
		return pollChartSlideHandler(ctx, req)
	}

	return map[string]any{
		"_fullHeight": true,
	}, nil
}

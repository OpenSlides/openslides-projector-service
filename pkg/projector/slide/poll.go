package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

type pollSlideOptions struct {
	SingleVotes bool `json:"single_votes"`
}

func PollSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	pollID := *req.ContentObjectID

	var options pollSlideOptions
	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse poll slide options: %w", err)
		}
	}

	var pollState string
	var pollPublished bool
	var pollTitle string
	var pollLiveVotingEnabled bool
	req.Fetch.Poll_State(pollID).Lazy(&pollState)
	req.Fetch.Poll_Published(pollID).Lazy(&pollPublished)
	req.Fetch.Poll_Title(pollID).Lazy(&pollTitle)
	req.Fetch.Poll_LiveVotingEnabled(pollID).Lazy(&pollLiveVotingEnabled)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("could not load poll base info %w", err)
	}

	showResults := pollPublished || ((pollState == "created" || pollState == "started") && pollLiveVotingEnabled)
	if !showResults {
		state := req.Locale.Get("No results yet")
		if pollState == "finished" {
			state = req.Locale.Get("Counting of votes is in progress ...")
		}

		if pollState == "started" && !pollLiveVotingEnabled {
			state = req.Locale.Get("Voting in progress")
		}

		if options.SingleVotes {
			return map[string]any{
				"_template": "poll_single_vote",
				"Title":     pollTitle,
				"State":     state,
			}, nil
		}

		return map[string]any{
			"Title": pollTitle,
			"State": state,
		}, nil
	}

	if options.SingleVotes {
		var isAnonymized bool
		req.Fetch.Poll_Anonymized(pollID).Lazy(&isAnonymized)
		if err := req.Fetch.Execute(ctx); err != nil {
			return nil, fmt.Errorf("could not check if poll is anonymized: %w", err)
		}

		if !isAnonymized {
			return pollSingleVotesSlideHandler(ctx, req)
		}
	}

	poll, err := req.Fetch.Poll(pollID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll %w", err)
	}

	if viewmodels.Poll_ShouldShowChart(poll) {
		return pollChartSlideHandler(ctx, req)
	}

	return pollTableSlideHandler(ctx, req)
}

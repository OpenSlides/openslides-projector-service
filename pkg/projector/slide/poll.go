package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

type pollSlideOptions struct {
	SingleVotes bool `json:"single_votes"`
}

type pollSlideProjectionOptionData struct {
	Color      string
	Icon       string
	Name       string
	TotalVotes decimal.Decimal
	PercVotes  decimal.Decimal
}

type pollSlideChartProjectionData struct {
	TotalValidvotes decimal.Decimal
	PercValidvotes  decimal.Decimal
	Options         []pollSlideProjectionOptionData
}

func PollSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	var options pollSlideOptions
	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse slide options: %w", err)
		}
	}

	poll, err := req.Fetch.Poll(*req.ContentObjectID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll id %w", err)
	}

	if poll.State != "published" && (poll.State != "started" && poll.LiveVotingEnabled) {
		state := req.Locale.Get("No results yet")
		if poll.State == "finished" {
			state = req.Locale.Get("Counting of votes is in progress ...")
		}

		if poll.State == "started" && !poll.LiveVotingEnabled {
			state = req.Locale.Get("Voting in progress")
		}

		return map[string]any{
			"Title": poll.Title,
			"State": state,
		}, nil
	}

	if options.SingleVotes {
		return pollSingleVotesSlideHandler(ctx, req)
	}

	template := "poll_chart"
	data := pollSlideChartProjectionData{}

	return map[string]any{
		"_template":   template,
		"_fullHeight": true,
		"Poll":        poll,
		"Data":        data,
	}, nil
}

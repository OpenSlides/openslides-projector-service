package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
	"github.com/shopspring/decimal"
)

type pollSlideOptions struct {
	SingleVotes bool `json:"single_votes"`
}

type pollSlideTableOption struct {
	Name         string
	TotalYes     decimal.Decimal
	TotalNo      decimal.Decimal
	TotalAbstain decimal.Decimal
	PercYes      decimal.Decimal
	PercNo       decimal.Decimal
	PercAbstain  decimal.Decimal
}

type pollSlideTableSum struct {
	Name  string
	Total decimal.Decimal
	Perc  string
}

type pollSlideTable struct {
	DisplayPercAbstain bool
	Options            []pollSlideTableOption
	Sums               []pollSlideTableSum
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

	showResults := pollState == "published" || ((pollState == "created" || pollState == "started") && pollLiveVotingEnabled)
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
		req.Fetch.Poll_IsPseudoanonymized(pollID).Lazy(&isAnonymized)
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

	if len(poll.OptionIDs) == 1 || (poll.Pollmethod == "Y" && !strings.HasPrefix(poll.ContentObjectID, "assignment")) {
		return pollChartSlideHandler(ctx, req)
	}

	pQ := req.Fetch.Poll()
	poll, err = req.Fetch.Poll(pollID).Preload(pQ.OptionList()).Preload(pQ.GlobalOption()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll %w", err)
	}

	userMap, err := viewmodels.User_MeetingUserMap(ctx, req.Fetch, poll.MeetingID)
	if err != nil {
		return nil, fmt.Errorf("could not load user map %w", err)
	}
	data := pollSlideTable{
		Options: []pollSlideTableOption{},
		Sums:    []pollSlideTableSum{},
	}

	for _, option := range poll.OptionList {
		onehundredPercentBase := viewmodels.Poll_OneHundredPercentBase(poll, &option)
		name, err := viewmodels.Option_OptionLabel(ctx, req.Fetch, req.Locale, &option, userMap)
		if err != nil {
			return nil, err
		}

		optData := pollSlideTableOption{
			Name:         name,
			TotalYes:     option.Yes,
			TotalNo:      option.No,
			TotalAbstain: option.Abstain,
		}

		if !onehundredPercentBase.IsZero() {
			optData.PercYes = optData.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			optData.PercNo = optData.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			optData.PercAbstain = optData.TotalAbstain.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		}

		data.Options = append(data.Options, optData)
	}

	data.DisplayPercAbstain = strings.Contains(poll.OnehundredPercentBase, "A") ||
		poll.OnehundredPercentBase == "cast" ||
		poll.OnehundredPercentBase == "valid"

	pollMethod := map[string]bool{
		"Yes":     strings.Contains(poll.Pollmethod, "Y"),
		"No":      strings.Contains(poll.Pollmethod, "N"),
		"Abstain": strings.Contains(poll.Pollmethod, "A"),
	}

	if poll.GlobalOption != nil && !poll.GlobalOption.Null() {
		globalOption, _ := poll.GlobalOption.Value()
		if poll.GlobalYes && !globalOption.Yes.IsZero() {
			data.Sums = append(data.Sums, pollSlideTableSum{
				Name:  req.Locale.Get("General approval"),
				Total: globalOption.Yes,
			})
		}

		if poll.GlobalNo && !globalOption.No.IsZero() {
			data.Sums = append(data.Sums, pollSlideTableSum{
				Name:  req.Locale.Get("General rejection"),
				Total: globalOption.No,
			})
		}

		if poll.GlobalAbstain && !globalOption.Abstain.IsZero() {
			data.Sums = append(data.Sums, pollSlideTableSum{
				Name:  req.Locale.Get("General abstain"),
				Total: globalOption.Abstain,
			})
		}
	}

	data.Sums = append(data.Sums, pollSlideTableSum{
		Name:  req.Locale.Get("Valid votes"),
		Total: poll.Votesvalid,
	})

	if !poll.Votesinvalid.IsZero() {
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Invalid votes"),
			Total: poll.Votesvalid,
		})
	}

	if !poll.Votescast.IsZero() && poll.Type == "analog" {
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Total votes cast"),
			Total: poll.Votescast,
		})
	}

	onehundredPercentBase := viewmodels.Poll_OneHundredPercentBase(poll, nil)
	if !onehundredPercentBase.IsZero() {
		for i, sum := range data.Sums {
			data.Sums[i].Perc = sum.Total.Div(onehundredPercentBase).Mul(decimal.NewFromInt(100)).Round(3).String()
		}
	}

	switch poll.OnehundredPercentBase {
	case "entitled":
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Entitled users"),
			Total: onehundredPercentBase,
			Perc:  "100",
		})
	case "entitled_present":
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Entitled present users"),
			Total: onehundredPercentBase,
			Perc:  "100",
		})
	}

	return map[string]any{
		"_fullHeight": true,
		"Title":       pollTitle,
		"Data":        data,
		"Base":        poll.OnehundredPercentBase,
		"Method":      poll.Pollmethod,
		"Methods":     pollMethod,
	}, nil
}

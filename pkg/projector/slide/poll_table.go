package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
	"github.com/shopspring/decimal"
)

type pollSlideTableOption struct {
	ID           int
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

func pollTableSlideHandler(ctx context.Context, req *projectionRequest, templateData map[string]any) (map[string]any, error) {
	pollID := *req.ContentObjectID
	pQ := req.Fetch.Poll(pollID)
	poll, err := req.Fetch.Poll(pollID).Preload(pQ.OptionList()).Preload(pQ.Config()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll %w", err)
	}

	var data pollSlideTable
	hasGlobalAbstain := false
	pollMethod := map[string]struct{}{}
	switch poll.Config.(type) {
	case *dsmodels.PollConfigRatingApproval:
		config := poll.Config.(*dsmodels.PollConfigRatingApproval)
		data, err = pollRatingApprovalTable(ctx, req, poll, *config)
		if err != nil {
			return nil, fmt.Errorf("could parse rating approval table: %w", err)
		}

		hasGlobalAbstain = config.AllowAbstain
		pollMethod["Yes"] = struct{}{}
		pollMethod["No"] = struct{}{}
		if config.AllowAbstain {
			pollMethod["Abstain"] = struct{}{}
		}

	case *dsmodels.PollConfigRatingScore:
		config := poll.Config.(*dsmodels.PollConfigRatingScore)
		data.DisplayPercAbstain = config.OnehundredPercentBase == "cast" ||
			config.OnehundredPercentBase == "valid"
		pollMethod["Yes"] = struct{}{}
		hasGlobalAbstain = config.MinOptionsAmount == 0

	case *dsmodels.PollConfigSelection:
		config := poll.Config.(*dsmodels.PollConfigSelection)
		data, err = pollSelectionTable(ctx, req, poll, *config)
		if err != nil {
			return nil, fmt.Errorf("could parse rating approval table: %w", err)
		}

		hasGlobalAbstain = config.MinOptionsAmount == 0
		if config.StrikeOut {
			pollMethod["No"] = struct{}{}
		} else {
			pollMethod["Yes"] = struct{}{}
		}
	}

	if hasGlobalAbstain {
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("General abstain"),
			Total: decimal.Decimal{}, // TODO: Fill
		})
	}

	/*
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Valid votes"),
			Total: poll.Votesvalid,
		})

		if !poll.Votesinvalid.IsZero() {
			data.Sums = append(data.Sums, pollSlideTableSum{
				Name:  req.Locale.Get("Invalid votes"),
				Total: poll.Votesinvalid,
			})
		}

		if !poll.Votescast.IsZero() && poll.Type == "analog" {
			data.Sums = append(data.Sums, pollSlideTableSum{
				Name:  req.Locale.Get("Total votes cast"),
				Total: poll.Votescast,
			})
		}

		onehundredPercentBase := viewmodels.Poll_OneHundredPercentBase(poll, nil)
		if !onehundredPercentBase.IsZero() && (poll.GlobalOption.Null() || poll.OnehundredPercentBase[0] != 'Y') {
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

		sortResult, err := req.Fetch.Meeting_AssignmentPollSortPollResultByVotes(poll.MeetingID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not fetch meeting poll sort option: %w", err)
		}

		if sortResult {
			slices.SortFunc(data.Options, func(a, b pollSlideTableOption) int {
				return b.TotalYes.Cmp(a.TotalYes)
			})
		}
	*/

	templateData["_fullHeight"] = true
	templateData["Poll"] = poll
	templateData["Data"] = data
	templateData["Methods"] = pollMethod
	return templateData, nil
}

func pollRatingApprovalTable(
	ctx context.Context,
	req *projectionRequest,
	poll dsmodels.Poll,
	config dsmodels.PollConfigRatingApproval,
) (pollSlideTable, error) {
	data := pollSlideTable{
		Options: []pollSlideTableOption{},
		Sums:    []pollSlideTableSum{},
		DisplayPercAbstain: config.OnehundredPercentBase == "yes_no_abstain" ||
			config.OnehundredPercentBase == "cast" ||
			config.OnehundredPercentBase == "valid",
	}

	// TODO: We need a specialized method to parse for
	var result viewmodels.PollResultRatingApproval
	if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
		return data, fmt.Errorf("parse approval poll result %w", err)
	}

	for _, option := range poll.OptionList {
		onehundredPercentBase := viewmodels.Poll_OneHundredPercentBase(poll, &option)
		name, err := viewmodels.Option_OptionLabel(ctx, req.Fetch, req.Locale, &option)
		if err != nil {
			return data, err
		}

		optData := pollSlideTableOption{
			ID:           option.ID,
			Name:         name,
			TotalYes:     decimal.Decimal{},
			TotalNo:      decimal.Decimal{},
			TotalAbstain: decimal.Decimal{},
		}

		if !onehundredPercentBase.IsZero() {
			optData.PercYes = optData.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			optData.PercNo = optData.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			optData.PercAbstain = optData.TotalAbstain.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		}

		data.Options = append(data.Options, optData)
	}

	return data, nil
}

func pollRatingScoreTable(
	ctx context.Context,
	req *projectionRequest,
	poll dsmodels.Poll,
	config dsmodels.PollConfigRatingScore,
) (pollSlideTable, error) {
	data := pollSlideTable{
		Options: []pollSlideTableOption{},
		Sums:    []pollSlideTableSum{},
	}

	var result viewmodels.PollResultRatingScore
	if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
		return data, fmt.Errorf("parse approval poll result %w", err)
	}

	return data, nil
}

func pollSelectionTable(
	ctx context.Context,
	req *projectionRequest,
	poll dsmodels.Poll,
	config dsmodels.PollConfigSelection,
) (pollSlideTable, error) {
	data := pollSlideTable{
		Options:            []pollSlideTableOption{},
		Sums:               []pollSlideTableSum{},
		DisplayPercAbstain: config.OnehundredPercentBase == "cast" || config.OnehundredPercentBase == "valid",
	}

	// TODO: We need a specialized method to parse for
	var result viewmodels.PollResultSelection
	if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
		return data, fmt.Errorf("parse approval poll result %w", err)
	}

	onehundredPercentBase := viewmodels.Poll_OneHundredPercentBase(poll, nil)
	for _, option := range poll.OptionList {
		name, err := viewmodels.Option_OptionLabel(ctx, req.Fetch, req.Locale, &option)
		if err != nil {
			return data, err
		}

		optData := pollSlideTableOption{
			ID:   option.ID,
			Name: name,
		}

		if config.StrikeOut {
			optData.TotalNo = result.Options[strconv.Itoa(option.ID)]
			if !onehundredPercentBase.IsZero() {
				optData.PercNo = optData.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			}
		} else {
			optData.TotalYes = result.Options[strconv.Itoa(option.ID)]
			if !onehundredPercentBase.IsZero() {
				optData.PercYes = optData.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			}
		}

		data.Options = append(data.Options, optData)
	}

	if config.AllowNota {
		if config.StrikeOut {
			data.Sums = append(data.Sums, pollSlideTableSum{
				Name:  req.Locale.Get("General approval"),
				Total: result.Nota,
			})
		} else {
			data.Sums = append(data.Sums, pollSlideTableSum{
				Name:  req.Locale.Get("General rejection"),
				Total: result.Nota,
			})
		}
	}

	data.Sums = append(data.Sums, pollSlideTableSum{
		Name:  req.Locale.Get("Valid votes"),
		Total: decimal.NewFromInt(int64(result.TotalBallots - result.Invalid)),
	})

	if result.Invalid > 0 {
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Invalid votes"),
			Total: decimal.NewFromInt(int64(result.Invalid)),
		})
	}

	if result.TotalBallots > 0 && poll.Visibility == "manually" {
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Total votes cast"),
			Total: decimal.NewFromInt(int64(result.TotalBallots)),
		})
	}

	return data, nil
}

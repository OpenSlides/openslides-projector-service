package slide

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
	"github.com/shopspring/decimal"
)

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

func pollTableSlideHandler(ctx context.Context, req *projectionRequest, templateData map[string]any) (map[string]any, error) {
	pollID := *req.ContentObjectID
	pQ := req.Fetch.Poll(pollID)
	poll, err := req.Fetch.Poll(pollID).Preload(pQ.OptionList()).Preload(pQ.Config()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll %w", err)
	}

	data := pollSlideTable{
		Options: []pollSlideTableOption{},
		Sums:    []pollSlideTableSum{},
	}

	for _, option := range poll.OptionList {
		onehundredPercentBase := viewmodels.Poll_OneHundredPercentBase(poll, &option)
		name, err := viewmodels.Option_OptionLabel(ctx, req.Fetch, req.Locale, &option)
		if err != nil {
			return nil, err
		}

		optData := pollSlideTableOption{
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

	hasGlobalAbstain := false
	hasGlobalYes := false
	hasGlobalNo := false
	pollMethod := map[string]struct{}{}
	switch poll.Config.(type) {
	case *dsmodels.PollConfigRatingApproval:
		config := poll.Config.(*dsmodels.PollConfigRatingApproval)
		data.DisplayPercAbstain = config.OnehundredPercentBase == "yes_no_abstain" ||
			config.OnehundredPercentBase == "cast" ||
			config.OnehundredPercentBase == "valid"
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
		data.DisplayPercAbstain = config.OnehundredPercentBase == "cast" ||
			config.OnehundredPercentBase == "valid"
		hasGlobalAbstain = config.MinOptionsAmount == 0
		if config.StrikeOut {
			pollMethod["No"] = struct{}{}
			if config.AllowNota {
				hasGlobalYes = true
			}
		} else {
			pollMethod["Yes"] = struct{}{}
			if config.AllowNota {
				hasGlobalNo = true
			}
		}
	}

	if hasGlobalYes {
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("General approval"),
			Total: decimal.Decimal{}, // TODO: Fill
		})
	} else if hasGlobalNo {
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("General rejection"),
			Total: decimal.Decimal{}, // TODO: Fill
		})
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

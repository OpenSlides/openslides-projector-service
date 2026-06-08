package slide

import (
	"context"
	"fmt"

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
	poll, err := req.Fetch.Poll(pollID).Preload(pQ.OptionList()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll %w", err)
	}

	/*
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

			// TODO: Set values from result
			optData := pollSlideTableOption{
				Name:         name,
				TotalYes:     decimal.Decimal{},
				TotalNo:      decimal.Decimal{},
				TotalAbstain: decimal.Decimal{},
			}

			// TODO: Handling of Strikout vote
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
			if poll.GlobalYes && poll.Pollmethod != "N" {
				data.Sums = append(data.Sums, pollSlideTableSum{
					Name:  req.Locale.Get("General approval"),
					Total: globalOption.Yes,
				})
			}

			if poll.GlobalNo {
				data.Sums = append(data.Sums, pollSlideTableSum{
					Name:  req.Locale.Get("General rejection"),
					Total: globalOption.No,
				})
			}

			if poll.GlobalAbstain {
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
	return templateData, nil
}

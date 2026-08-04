package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/i18n"
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
	pollMethod := map[string]struct{}{}
	switch poll.Config.(type) {
	case *dsmodels.PollConfigRatingApproval:
		config := poll.Config.(*dsmodels.PollConfigRatingApproval)
		data, err = pollRatingApprovalTable(ctx, req, poll, *config)
		if err != nil {
			return nil, fmt.Errorf("could parse rating approval table: %w", err)
		}

		pollMethod["Yes"] = struct{}{}
		pollMethod["No"] = struct{}{}
		if config.AllowAbstain {
			pollMethod["Abstain"] = struct{}{}
		}

	case *dsmodels.PollConfigRatingScore:
		config := poll.Config.(*dsmodels.PollConfigRatingScore)
		data, err = pollRatingScoreTable(ctx, req, poll, *config)
		if err != nil {
			return nil, fmt.Errorf("could parse rating approval table: %w", err)
		}

		pollMethod["Yes"] = struct{}{}

	case *dsmodels.PollConfigSelection:
		config := poll.Config.(*dsmodels.PollConfigSelection)
		data, err = pollSelectionTable(ctx, req, poll, *config)
		if err != nil {
			return nil, fmt.Errorf("could parse rating approval table: %w", err)
		}

		if config.StrikeOut {
			pollMethod["No"] = struct{}{}
		} else {
			pollMethod["Yes"] = struct{}{}
		}
	}

	/* TODO: Readd if available
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
	*/

	var configID dsfetch.Maybe[int]
	if strings.HasPrefix(poll.ContentObjectID, "assignment") {
		configID, err = req.Fetch.Meeting_AssignmentPollConfigID(poll.MeetingID).Value(ctx)
	} else if strings.HasPrefix(poll.ContentObjectID, "motion") {
		configID, err = req.Fetch.Meeting_MotionPollConfigID(poll.MeetingID).Value(ctx)
	} else if strings.HasPrefix(poll.ContentObjectID, "topic") {
		configID, err = req.Fetch.Meeting_TopicPollConfigID(poll.MeetingID).Value(ctx)
	}

	if err != nil {
		return nil, fmt.Errorf("could not fetch meeting poll config id: %w", err)
	}

	if val, isSet := configID.Value(); isSet {
		sortResult, err := req.Fetch.MeetingPollDefault_SortResultByVotes(val).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not fetch meeting poll sort option: %w", err)
		}

		if sortResult {
			slices.SortFunc(data.Options, func(a, b pollSlideTableOption) int {
				return b.TotalYes.Cmp(a.TotalYes)
			})
		}
	}

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

	var result viewmodels.PollResultRatingApproval
	if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
		return data, fmt.Errorf("parse approval poll result %w", err)
	}

	for _, option := range poll.OptionList {
		onehundredPercentBase := result.OneHundredPercentBase(&config, &option)
		name, err := viewmodels.Option_OptionLabel(ctx, req.Fetch, req.Locale, &option)
		if err != nil {
			return data, err
		}

		optResult := result.Options[strconv.Itoa(option.ID)]
		optData := pollSlideTableOption{
			ID:           option.ID,
			Name:         name,
			TotalYes:     optResult.Yes,
			TotalNo:      optResult.No,
			TotalAbstain: optResult.Abstain,
		}

		if !onehundredPercentBase.IsZero() {
			optData.PercYes = optData.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			optData.PercNo = optData.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			optData.PercAbstain = optData.TotalAbstain.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		}

		data.Options = append(data.Options, optData)
	}

	data.Sums = append(data.Sums, pollTableSums(req.Locale, &result, poll.Visibility == "manually")...)

	return data, nil
}

func pollRatingScoreTable(
	ctx context.Context,
	req *projectionRequest,
	poll dsmodels.Poll,
	config dsmodels.PollConfigRatingScore,
) (pollSlideTable, error) {
	data := pollSlideTable{
		Options:            []pollSlideTableOption{},
		Sums:               []pollSlideTableSum{},
		DisplayPercAbstain: config.OnehundredPercentBase == "cast" || config.OnehundredPercentBase == "valid",
	}

	var result viewmodels.PollResultRatingScore
	if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
		return data, fmt.Errorf("parse approval poll result %w", err)
	}

	data.Sums = append(data.Sums, pollTableSums(req.Locale, &result, poll.Visibility == "manually")...)

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

	var result viewmodels.PollResultSelection
	if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
		return data, fmt.Errorf("parse approval poll result %w", err)
	}

	onehundredPercentBase := result.OneHundredPercentBase(&config)
	total := decimal.NewFromInt(result.VotesValid()).Sub(result.Abstain)
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
			optData.TotalYes = total.Sub(result.Options[strconv.Itoa(option.ID)])
		} else {
			optData.TotalYes = result.Options[strconv.Itoa(option.ID)]
		}

		if !onehundredPercentBase.IsZero() {
			optData.PercYes = optData.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
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

	if config.MinOptionsAmount == 0 {
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("General abstain"),
			Total: result.Abstain,
		})
	}

	data.Sums = append(data.Sums, pollTableSums(req.Locale, &result, poll.Visibility == "manually")...)

	return data, nil
}

func pollTableSums(locale *i18n.ProjectorLocale, result viewmodels.PollResult, showCast bool) []pollSlideTableSum {
	sums := []pollSlideTableSum{}
	sums = append(sums, pollSlideTableSum{
		Name:  locale.Get("Valid votes"),
		Total: decimal.NewFromInt(result.VotesValid()),
	})

	if result.VotesInvalid() > 0 {
		sums = append(sums, pollSlideTableSum{
			Name:  locale.Get("Invalid votes"),
			Total: decimal.NewFromInt(result.VotesInvalid()),
		})
	}

	if result.VotesCast() > 0 && showCast {
		sums = append(sums, pollSlideTableSum{
			Name:  locale.Get("Total votes cast"),
			Total: decimal.NewFromInt(result.VotesInvalid()),
		})
	}

	return sums
}

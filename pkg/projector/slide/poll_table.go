package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-go/datastore/dstypes"
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
	poll, err := req.Fetch.Poll(pollID).Preload(pQ.OptionList().ContentObject()).Preload(pQ.Config()).First(ctx)
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

		templateData["Base"] = config.OnehundredPercentBase
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

		templateData["Base"] = config.OnehundredPercentBase
		pollMethod["Yes"] = struct{}{}

	case *dsmodels.PollConfigSelection:
		config := poll.Config.(*dsmodels.PollConfigSelection)
		data, err = pollSelectionTable(ctx, req, poll, *config)
		if err != nil {
			return nil, fmt.Errorf("could parse rating approval table: %w", err)
		}

		templateData["Base"] = config.OnehundredPercentBase
		if config.StrikeOut {
			pollMethod["No"] = struct{}{}
		} else {
			pollMethod["Yes"] = struct{}{}
		}
	}

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
		sortResult, err := req.Fetch.MeetingPollSetting_SortResultByVotes(val).Value(ctx)
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

	var onehundredPercentBase decimal.Decimal
	resultOptions := viewmodels.RemapPollResultOptions(result.Options, poll.OptionList)
	for _, option := range poll.OptionList {
		onehundredPercentBase = result.OneHundredPercentBase(&config, &option)
		name, err := viewmodels.Option_OptionLabel(ctx, req.Fetch, req.Locale, &option)
		if err != nil {
			return data, err
		}

		optResult := resultOptions[option.ID]
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

	if config.OnehundredPercentBase == dstypes.RatingApprovalOnehundredPercentBasesEntitled ||
		config.OnehundredPercentBase == dstypes.RatingApprovalOnehundredPercentBasesEntitledPresent {
		onehundredPercentBase = result.OneHundredPercentBase(&config, nil)
	} else {
		onehundredPercentBase = decimal.Decimal{}
	}
	data.Sums = append(data.Sums, pollTableSums(req.Locale, &result, poll.Visibility == "manually", onehundredPercentBase)...)

	switch config.OnehundredPercentBase {
	case dstypes.RatingApprovalOnehundredPercentBasesEntitled:
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Entitled users"),
			Total: decimal.NewFromInt(result.VotesEntitledUsers()),
			Perc:  "100",
		})
	case dstypes.RatingApprovalOnehundredPercentBasesEntitledPresent:
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Entitled present users"),
			Total: decimal.NewFromInt(result.VotesEntitledUsersPresent()),
			Perc:  "100",
		})
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
		Options:            []pollSlideTableOption{},
		Sums:               []pollSlideTableSum{},
		DisplayPercAbstain: config.OnehundredPercentBase == "cast" || config.OnehundredPercentBase == "valid",
	}

	var result viewmodels.PollResultRatingScore
	if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
		return data, fmt.Errorf("parse approval poll result %w", err)
	}

	data.Sums = append(data.Sums, pollTableSums(req.Locale, &result, poll.Visibility == "manually", decimal.Decimal{})...)

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

	resultOptions := viewmodels.RemapPollResultOptions(result.Options, poll.OptionList)
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
			optData.TotalYes = total.Sub(resultOptions[option.ID])
		} else {
			optData.TotalYes = resultOptions[option.ID]
		}

		if !onehundredPercentBase.IsZero() {
			optData.PercYes = optData.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		}

		data.Options = append(data.Options, optData)
	}

	displayGeneralPercent := config.OnehundredPercentBase != dstypes.SelectionOnehundredPercentBasesDisabled &&
		config.OnehundredPercentBase != dstypes.SelectionOnehundredPercentBasesNoGeneral
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

		if displayGeneralPercent && !onehundredPercentBase.IsZero() {
			data.Sums[len(data.Sums)-1].Perc = result.Nota.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100)).String()
		}
	}

	if poll.AllowEmpty {
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("General abstain"),
			Total: result.Empty,
		})

		if displayGeneralPercent && !onehundredPercentBase.IsZero() {
			data.Sums[len(data.Sums)-1].Perc = result.Empty.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100)).String()
		}
	}

	if config.OnehundredPercentBase != dstypes.SelectionOnehundredPercentBasesCast {
		valid := decimal.NewFromInt(result.VotesValid())
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Valid votes"),
			Total: valid,
		})

		if !onehundredPercentBase.IsZero() {
			data.Sums[len(data.Sums)-1].Perc = valid.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100)).String()
		}
	}

	if result.VotesInvalid() > 0 {
		invalid := decimal.NewFromInt(result.VotesInvalid())
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Invalid votes"),
			Total: invalid,
		})

		if !onehundredPercentBase.IsZero() {
			data.Sums[len(data.Sums)-1].Perc = invalid.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100)).String()
		}
	}

	if result.VotesCast() > 0 && config.OnehundredPercentBase == dstypes.SelectionOnehundredPercentBasesCast {
		cast := decimal.NewFromInt(result.VotesCast())
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Total votes cast"),
			Total: cast,
		})

		if !onehundredPercentBase.IsZero() {
			data.Sums[len(data.Sums)-1].Perc = cast.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100)).String()
		}
	}

	switch config.OnehundredPercentBase {
	case dstypes.SelectionOnehundredPercentBasesEntitled:
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Entitled users"),
			Total: onehundredPercentBase,
			Perc:  "100",
		})
	case dstypes.SelectionOnehundredPercentBasesEntitledPresent:
		data.Sums = append(data.Sums, pollSlideTableSum{
			Name:  req.Locale.Get("Entitled present users"),
			Total: onehundredPercentBase,
			Perc:  "100",
		})
	}

	return data, nil
}

func pollTableSums(locale *i18n.ProjectorLocale, result viewmodels.PollResult, showCast bool, base decimal.Decimal) []pollSlideTableSum {
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
			Total: decimal.NewFromInt(result.VotesCast()),
		})
	}

	if !base.IsZero() {
		for i, amount := range sums {
			sums[i].Perc = amount.Total.DivRound(base, 5).Mul(decimal.NewFromInt(100)).String()
		}
	}

	return sums
}

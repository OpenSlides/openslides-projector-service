package slide

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"slices"
	"strconv"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-go/datastore/dstypes"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
	"github.com/shopspring/decimal"
)

type pollSlideProjectionOptionData struct {
	Type         rune
	Color        template.CSS
	Icon         string
	Name         string
	TotalVotes   decimal.Decimal
	PercVotes    string
	DisplayPerc  bool
	GlobalOption bool
}

type pollSlideChartProjectionData struct {
	TotalValidvotes decimal.Decimal
	PercValidvotes  string
	ResultTitle     string
	ChartData       string
	EntitledUsers   int
	Options         []pollSlideProjectionOptionData
}

func pollChartSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	pollID := *req.ContentObjectID
	pQ := req.Fetch.Poll(pollID)
	poll, err := req.Fetch.Poll(pollID).Preload(pQ.OptionList().ContentObject()).Preload(pQ.Config()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll %w", err)
	}

	if len(poll.Result) == 0 {
		return nil, errors.New("poll result empty")
	}

	data := pollSlideChartProjectionData{
		Options: []pollSlideProjectionOptionData{},
	}

	var onehundredPercentBase decimal.Decimal

	switch poll.Config.(type) {
	case *dsmodels.PollConfigApproval:
		var result viewmodels.PollResultApproval
		if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
			return nil, fmt.Errorf("parse approval poll result %w", err)
		}

		config := poll.Config.(*dsmodels.PollConfigApproval)
		data.Options = append(data.Options, pollSlideProjectionOptionData{
			Type:       'Y',
			Color:      "--theme-yes",
			Icon:       "check_circle",
			Name:       req.Locale.Get("Yes"),
			TotalVotes: result.Yes,
			DisplayPerc: config.OnehundredPercentBase == dstypes.ApprovalOnehundredPercentBasesYesNo ||
				config.OnehundredPercentBase == dstypes.ApprovalOnehundredPercentBasesValid,
		})

		data.Options = append(data.Options, pollSlideProjectionOptionData{
			Type:       'N',
			Color:      "--theme-no",
			Icon:       "cancel",
			Name:       req.Locale.Get("No"),
			TotalVotes: result.No,
			DisplayPerc: config.OnehundredPercentBase == dstypes.ApprovalOnehundredPercentBasesYesNo ||
				config.OnehundredPercentBase == dstypes.ApprovalOnehundredPercentBasesValid,
		})

		if config.AllowAbstain {
			data.Options = append(data.Options, pollSlideProjectionOptionData{
				Type:        'A',
				Color:       "--theme-abstain",
				Icon:        "circle",
				Name:        req.Locale.Get("Abstain"),
				TotalVotes:  result.Abstain,
				DisplayPerc: config.OnehundredPercentBase == dstypes.ApprovalOnehundredPercentBasesValid,
			})
		}

		onehundredPercentBase = result.OneHundredPercentBase(config)

		data.TotalValidvotes = decimal.NewFromInt(int64(result.TotalBallots - result.Invalid))
		if !onehundredPercentBase.IsZero() && config.OnehundredPercentBase != "yes_no" && config.OnehundredPercentBase != "yes_no_abstain" {
			data.PercValidvotes = data.TotalValidvotes.Div(onehundredPercentBase).Mul(decimal.NewFromInt(100)).Round(3).String()
		}
	case *dsmodels.PollConfigSelection:
		var result viewmodels.PollResultSelection
		if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
			return nil, fmt.Errorf("parse approval poll result %w", err)
		}

		config := poll.Config.(*dsmodels.PollConfigSelection)
		onehundredPercentBase = result.OneHundredPercentBase(config)

		for _, option := range poll.OptionList {
			optionLabel, err := viewmodels.Option_OptionLabel(ctx, req.Fetch, req.Locale, &option)
			if err != nil {
				return nil, fmt.Errorf("parsing option label: %w", err)
			}

			data.Options = append(data.Options, pollSlideProjectionOptionData{
				Icon:        "circle",
				Name:        optionLabel,
				TotalVotes:  result.Options[strconv.Itoa(option.ID)],
				DisplayPerc: true,
			})
		}

		slices.SortStableFunc(data.Options, func(a pollSlideProjectionOptionData, b pollSlideProjectionOptionData) int {
			return b.TotalVotes.Cmp(a.TotalVotes)
		})

		if config.AllowNota {
			if config.StrikeOut {
				data.Options = append(data.Options, pollSlideProjectionOptionData{
					Name:         req.Locale.Get("General approval"),
					TotalVotes:   result.Nota,
					GlobalOption: true,
				})
			} else {
				data.Options = append(data.Options, pollSlideProjectionOptionData{
					Name:         req.Locale.Get("General rejection"),
					TotalVotes:   result.Nota,
					GlobalOption: true,
				})
			}
		}

		if config.MinOptionsAmount == 0 {
			data.Options = append(data.Options, pollSlideProjectionOptionData{
				Name:         req.Locale.Get("General abstain"),
				TotalVotes:   result.Abstain,
				GlobalOption: true,
			})
		}

		data.TotalValidvotes = decimal.NewFromInt(int64(result.TotalBallots - result.Invalid))
		if !onehundredPercentBase.IsZero() {
			data.PercValidvotes = data.TotalValidvotes.Div(onehundredPercentBase).Mul(decimal.NewFromInt(100)).Round(3).String()
		}
	default:
		return nil, fmt.Errorf("chart slide not implemented for this config type: %w", err)
	}

	type chartDataEntry struct {
		Color string  `json:"color,omitempty"`
		Val   float64 `json:"val"`
	}

	chartData := []chartDataEntry{}
	for i, option := range data.Options {
		if option.GlobalOption {
			continue
		}

		// TODO:Hide abstain from chart for `yes_no` 100%-base

		chartData = append(chartData, chartDataEntry{
			Color: string(option.Color),
			Val:   option.TotalVotes.InexactFloat64(),
		})

		if !onehundredPercentBase.IsZero() && option.DisplayPerc {
			data.Options[i].PercVotes = option.TotalVotes.Div(onehundredPercentBase).Mul(decimal.NewFromInt(100)).Round(3).String()
		}
	}

	chartDataJSON, err := json.Marshal(chartData)
	if err != nil {
		return nil, fmt.Errorf("could not marshal chart data json %w", err)
	}
	data.ChartData = string(chartDataJSON)

	return map[string]any{
		"_template":   "poll_chart",
		"_fullHeight": true,
		"Poll":        poll,
		"Data":        data,
	}, nil
}

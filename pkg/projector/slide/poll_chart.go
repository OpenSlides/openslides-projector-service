package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"slices"
	"strings"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
	"github.com/shopspring/decimal"
)

type pollSlideProjectionOptionData struct {
	Color       template.CSS
	Icon        string
	Name        string
	TotalVotes  decimal.Decimal
	PercVotes   string
	DisplayPerc bool
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
	poll, err := req.Fetch.Poll(pollID).Preload(pQ.OptionList()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll %w", err)
	}

	data := pollSlideChartProjectionData{
		Options: []pollSlideProjectionOptionData{},
	}
	onehundredPercentBase := viewmodels.Poll_OneHundredPercentBase(poll, nil)
	if len(poll.OptionList) == 1 {
		opt := poll.OptionList[0]
		coID, hasCoID := opt.ContentObjectID.Value()
		if opt.Text != "" {
			data.ResultTitle = opt.Text
		} else if hasCoID && strings.HasPrefix(coID, "user") {
			data.ResultTitle = fmt.Sprintf("%s %s", "FIRST NAME", "LAST NAME")
			// TODO: Structure Level
		}

		if strings.Contains(poll.Pollmethod, "Y") {
			data.Options = append(data.Options, pollSlideProjectionOptionData{
				Color:       "--theme-yes",
				Icon:        "check_circle",
				Name:        req.Locale.Get("Yes"),
				TotalVotes:  opt.Yes,
				DisplayPerc: strings.Contains(poll.OnehundredPercentBase, "Y"),
			})
		}

		if strings.Contains(poll.Pollmethod, "N") {
			data.Options = append(data.Options, pollSlideProjectionOptionData{
				Color:       "--theme-no",
				Icon:        "cancel",
				Name:        req.Locale.Get("No"),
				TotalVotes:  opt.No,
				DisplayPerc: strings.Contains(poll.OnehundredPercentBase, "N"),
			})
		}

		if strings.Contains(poll.Pollmethod, "A") {
			data.Options = append(data.Options, pollSlideProjectionOptionData{
				Color:       "--theme-abstain",
				Icon:        "circle",
				Name:        req.Locale.Get("Abstain"),
				TotalVotes:  opt.Abstain,
				DisplayPerc: strings.Contains(poll.OnehundredPercentBase, "A"),
			})
		}
	} else {
		for _, opt := range poll.OptionList {
			data.Options = append(data.Options, pollSlideProjectionOptionData{
				Icon:        "circle",
				Name:        opt.Text,
				TotalVotes:  opt.Yes,
				DisplayPerc: true,
			})
		}

		slices.SortStableFunc(data.Options, func(a pollSlideProjectionOptionData, b pollSlideProjectionOptionData) int {
			return b.TotalVotes.Cmp(a.TotalVotes)
		})
	}

	type chartDataEntry struct {
		Color string  `json:"color,omitempty"`
		Val   float64 `json:"val"`
	}

	chartData := []chartDataEntry{}
	for i, option := range data.Options {
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

	data.TotalValidvotes = poll.Votesvalid
	if !onehundredPercentBase.IsZero() {
		data.PercValidvotes = poll.Votesvalid.Div(onehundredPercentBase).Mul(decimal.NewFromInt(100)).Round(3).String()
	}

	return map[string]any{
		"_template":   "poll_chart",
		"_fullHeight": true,
		"Poll":        poll,
		"Data":        data,
	}, nil
}

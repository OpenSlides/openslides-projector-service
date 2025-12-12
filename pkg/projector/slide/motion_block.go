package slide

import (
	"context"
	"fmt"
	"slices"
	"strings"
)

func MotionBlockSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no motion block id provided for slide")
	}

	b := req.Fetch.MotionBlock(*req.ContentObjectID)
	block, err := b.Preload(b.MotionList().Recommendation()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load motion block %w", err)
	}

	var maxColumns int
	req.Fetch.Meeting_MotionsBlockSlideColumns(block.MeetingID).Lazy(&maxColumns)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("could not load meeting settings: %w", err)
	}

	numMotions := len(block.MotionIDs)

	type motionListEntry struct {
		Number              string
		Title               string
		Recommendation      string
		RecommendationColor string
	}
	motionList := []motionListEntry{}
	for _, motion := range block.MotionList {
		recoName := ""
		recoColor := ""
		if reco, hasReco := motion.Recommendation.Value(); hasReco {
			recoName = reco.RecommendationLabel
			recoColor = reco.CssClass
		}
		motionList = append(motionList, motionListEntry{
			Number:              motion.Number,
			Title:               motion.Title,
			Recommendation:      recoName,
			RecommendationColor: recoColor,
		})
	}

	slices.SortFunc(motionList, func(a, b motionListEntry) int {
		if a.Number != "" || b.Number != "" {
			return strings.Compare(a.Number, b.Number)
		}

		return strings.Compare(a.Title, b.Title)
	})

	return map[string]any{
		"MotionBlock": block,
		"Motions":     motionList,
		"NumMotions":  numMotions,
		"MaxColumns":  maxColumns,
	}, nil
}

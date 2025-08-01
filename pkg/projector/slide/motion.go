package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"maps"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

type motionSlideMode string

const (
	motionTextOriginal      motionSlideMode = "original"
	motionTextChanged       motionSlideMode = "changed"
	motionTextDiff          motionSlideMode = "diff"
	motionTextFinal         motionSlideMode = "agreed"
	motionTextModifiedFinal motionSlideMode = "modified_final_version"
)

type motionSlideOptions struct {
	Mode motionSlideMode `json:"mode"`
}

type motionSlideCommonData struct {
	Mode          string
	Motion        dsmodels.Motion
	ShowSidebox   bool
	ShowText      bool
	LineNumbering string
	LineLength    int
	Preamble      string
}

func (m *motionSlideCommonData) templateData(additional map[string]any) map[string]any {
	data := map[string]any{
		"Motion":                    m.Motion,
		"MotionText":                template.HTML(m.Motion.Text),
		"Reason":                    template.HTML(m.Motion.Reason),
		"IsParagraphBasedAmendment": !m.Motion.LeadMotionID.Null(),
		"ShowSidebox":               m.ShowSidebox,
		"LineLength":                m.LineLength,
		"LineNumbering":             m.LineNumbering,
		"ShowText":                  m.ShowText,
		"Preamble":                  m.Preamble,
	}
	maps.Copy(data, additional)
	return data
}

func MotionSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no motion id provided for slide")
	}

	mQ := req.Fetch.Motion(*req.ContentObjectID)
	motion, err := mQ.First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load motion block %w", err)
	}

	options := motionSlideOptions{
		Mode: motionTextOriginal,
	}

	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse slide options: %w", err)
		}
	}

	data := motionSlideCommonData{
		Mode:   string(options.Mode),
		Motion: motion,
	}

	req.Fetch.Meeting_MotionsDefaultLineNumbering(motion.MeetingID).Lazy(&data.LineNumbering)
	req.Fetch.Meeting_MotionsEnableSideboxOnProjector(motion.MeetingID).Lazy(&data.ShowSidebox)
	req.Fetch.Meeting_MotionsEnableTextOnProjector(motion.MeetingID).Lazy(&data.ShowText)
	req.Fetch.Meeting_MotionsLineLength(motion.MeetingID).Lazy(&data.LineLength)
	req.Fetch.Meeting_MotionsPreamble(motion.MeetingID).Lazy(&data.Preamble)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("could fetch motion slide data: %w", err)
	}

	switch options.Mode {
	case motionTextChanged:
		return motionTextChangedSlide(ctx, &data)
	case motionTextDiff:
		return motionTextDiffSlide(ctx, &data)
	case motionTextFinal:
		return motionTextFinalSlide(ctx, &data)
	case motionTextModifiedFinal:
		return motionTextModifiedFinalSlide(ctx, &data)
	}

	return motionTextOriginalSlide(ctx, &data)
}

func motionTextOriginalSlide(ctx context.Context, req *motionSlideCommonData) (map[string]any, error) {
	return req.templateData(map[string]any{}), nil
}

func motionTextChangedSlide(ctx context.Context, req *motionSlideCommonData) (map[string]any, error) {
	return req.templateData(map[string]any{}), nil
}

func motionTextDiffSlide(ctx context.Context, req *motionSlideCommonData) (map[string]any, error) {
	return req.templateData(map[string]any{}), nil
}

func motionTextFinalSlide(ctx context.Context, req *motionSlideCommonData) (map[string]any, error) {
	return req.templateData(map[string]any{}), nil
}

func motionTextModifiedFinalSlide(ctx context.Context, req *motionSlideCommonData) (map[string]any, error) {
	return req.templateData(map[string]any{}), nil
}

package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"maps"
	"slices"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
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
	ProjectionReq      *projectionRequest
	Mode               string
	Motion             *dsmodels.Motion
	ShowSidebox        bool
	ShowText           bool
	HideMetaBackground bool
	Submitters         string
	LineNumbering      string
	LineLength         int
	Preamble           string
}

func (m *motionSlideCommonData) templateData(additional map[string]any) map[string]any {
	data := map[string]any{
		"IsParagraphBasedAmendment": !m.Motion.LeadMotionID.Null(),
		"LineLength":                m.LineLength,
		"LineNumbering":             m.LineNumbering,
		"Mode":                      m.Mode,
		"Motion":                    m.Motion,
		"MotionText":                template.HTML(m.Motion.Text),
		"Preamble":                  m.Preamble,
		"Reason":                    template.HTML(m.Motion.Reason),
		"ShowSidebox":               m.ShowSidebox,
		"ShowText":                  m.ShowText,
		"Submitters":                m.Submitters,
	}
	maps.Copy(data, additional)
	return data
}

func MotionSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no motion id provided for slide")
	}

	mQ := req.Fetch.Motion(*req.ContentObjectID)
	motion, err := mQ.Preload(mQ.SubmitterList().MeetingUser().User()).
		Preload(mQ.SubmitterList().MeetingUser().StructureLevelList()).
		First(ctx)
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
		ProjectionReq: req,
		Mode:          string(options.Mode),
		Motion:        &motion,
		Submitters:    strings.Join(motionSubmitterList(&motion), ", "),
	}

	req.Fetch.Meeting_MotionsDefaultLineNumbering(motion.MeetingID).Lazy(&data.LineNumbering)
	req.Fetch.Meeting_MotionsEnableSideboxOnProjector(motion.MeetingID).Lazy(&data.ShowSidebox)
	req.Fetch.Meeting_MotionsEnableTextOnProjector(motion.MeetingID).Lazy(&data.ShowText)
	req.Fetch.Meeting_MotionsLineLength(motion.MeetingID).Lazy(&data.LineLength)
	req.Fetch.Meeting_MotionsPreamble(motion.MeetingID).Lazy(&data.Preamble)
	req.Fetch.Meeting_MotionsHideMetadataBackground(motion.MeetingID).Lazy(&data.HideMetaBackground)
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

	return data.templateData(map[string]any{}), nil
}

func motionTextChangedSlide(ctx context.Context, req *motionSlideCommonData) (map[string]any, error) {
	changeRecoData, err := motionChangeRecos(ctx, req)
	return req.templateData(changeRecoData), err
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

func motionSubmitterList(motion *dsmodels.Motion) []string {
	submitters := []string{}
	slices.SortFunc(motion.SubmitterList, func(a dsmodels.MotionSubmitter, b dsmodels.MotionSubmitter) int {
		return a.Weight - b.Weight
	})
	for _, submitter := range motion.SubmitterList {
		submitters = append(submitters, viewmodels.MeetingUser_FullName(submitter.MeetingUser))
	}
	if motion.AdditionalSubmitter != "" {
		submitters = append(submitters, motion.AdditionalSubmitter)
	}

	return submitters
}

func motionChangeRecos(ctx context.Context, req *motionSlideCommonData) (map[string]any, error) {
	fetch := req.ProjectionReq.Fetch
	crIDs := req.Motion.ChangeRecommendationIDs
	crs, err := fetch.MotionChangeRecommendation(crIDs...).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("could fetch change recommendations slide data: %w", err)
	}

	type changeReco struct {
		ID       int
		Type     string
		LineFrom int
		LineTo   int
		Text     template.HTML
	}

	changeRecos := []changeReco{}
	titleChanges := []changeReco{}
	for _, cr := range crs {
		if !cr.Rejected && !cr.Internal {
			newCr := changeReco{
				ID:       cr.ID,
				Type:     cr.Type,
				LineFrom: cr.LineFrom,
				LineTo:   cr.LineTo,
				Text:     template.HTML(cr.Text),
			}

			if cr.LineFrom > 0 {
				changeRecos = append(changeRecos, newCr)
			} else {
				titleChanges = append(titleChanges, newCr)
			}
		}
	}

	return req.templateData(map[string]any{
		"TitleChangeRecos":  titleChanges,
		"MotionChangeRecos": changeRecos,
	}), nil
}

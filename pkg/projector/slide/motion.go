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
	ProjectionReq         *projectionRequest
	Mode                  string
	Motion                *dsmodels.Motion
	AmendmentParagraphs   map[string]template.HTML
	ShowSidebox           bool
	ShowReason            bool
	ShowRecommendation    bool
	ShowText              bool
	HideMetaBackground    bool
	Recommender           string
	Recommendation        string
	ReferencedRecoMotions string
	Submitters            string
	LineNumbering         string
	LineLength            int
	Preamble              string
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
		"ReferencedRecoMotions":     m.ReferencedRecoMotions,
		"ShowSidebox":               m.ShowSidebox,
		"ShowText":                  m.ShowText,
		"Submitters":                m.Submitters,
		"HideMetadataBackground":    m.HideMetaBackground,
	}

	if m.ShowReason {
		data["Reason"] = template.HTML(m.Motion.Reason)
	}

	if m.ShowRecommendation && m.Recommendation != "" && m.Recommender != "" {
		data["Recommender"] = m.Recommender
		data["Recommendation"] = m.Recommendation
	}

	if m.AmendmentParagraphs != nil {
		data["AmendmentParagraphs"] = m.AmendmentParagraphs
		if lMotion, ok := m.Motion.LeadMotion.Value(); ok {
			data["LeadMotionText"] = template.HTML(lMotion.Text)
		}
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
		Preload(mQ.LeadMotion()).
		Preload(mQ.Recommendation()).
		Preload(mQ.ReferencedInMotionRecommendationExtensionList()).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load motion block %w", err)
	}

	options := motionSlideOptions{
		Mode: motionTextOriginal,
	}

	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse motion slide options: %w", err)
		}
	}

	data := motionSlideCommonData{
		ProjectionReq: req,
		Mode:          string(options.Mode),
		Motion:        &motion,
		Submitters:    strings.Join(motionSubmitterList(&motion), ", "),
	}

	req.Fetch.Meeting_MotionsDefaultLineNumbering(motion.MeetingID).Lazy(&data.LineNumbering)
	req.Fetch.Meeting_MotionsEnableReasonOnProjector(motion.MeetingID).Lazy(&data.ShowReason)
	req.Fetch.Meeting_MotionsEnableRecommendationOnProjector(motion.MeetingID).Lazy(&data.ShowRecommendation)
	req.Fetch.Meeting_MotionsRecommendationsBy(motion.MeetingID).Lazy(&data.Recommender)
	req.Fetch.Meeting_MotionsEnableSideboxOnProjector(motion.MeetingID).Lazy(&data.ShowSidebox)
	req.Fetch.Meeting_MotionsEnableTextOnProjector(motion.MeetingID).Lazy(&data.ShowText)
	req.Fetch.Meeting_MotionsLineLength(motion.MeetingID).Lazy(&data.LineLength)
	req.Fetch.Meeting_MotionsPreamble(motion.MeetingID).Lazy(&data.Preamble)
	req.Fetch.Meeting_MotionsHideMetadataBackground(motion.MeetingID).Lazy(&data.HideMetaBackground)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("could not fetch motion slide data: %w", err)
	}

	if data.ShowRecommendation {
		if val, ok := motion.Recommendation.Value(); ok {
			data.Recommendation = val.RecommendationLabel
			if motion.RecommendationExtension != "" && val.ShowRecommendationExtensionField {
				ext, err := viewmodels.Motion_RecommendationParsed(ctx, req.Fetch, &motion)
				if err != nil {
					return nil, fmt.Errorf("error parsing motion recommendation: %w", err)
				}
				data.Recommendation = fmt.Sprintf("%s %s", data.Recommendation, ext)
			}
		}
	}

	if !motion.LeadMotionID.Null() && len(motion.AmendmentParagraphs) > 0 {
		amendmentParagrapphs := map[string]template.HTML{}
		if err := json.Unmarshal(motion.AmendmentParagraphs, &amendmentParagrapphs); err != nil {
			return nil, fmt.Errorf("error parsing amendment paragraphs: %w", err)
		} else {
			data.AmendmentParagraphs = amendmentParagrapphs
		}
	}

	if len(data.Motion.ReferencedInMotionRecommendationExtensionList) > 0 {
		refMotionNames := []string{}
		for _, refMotion := range data.Motion.ReferencedInMotionRecommendationExtensionList {
			title := refMotion.Number
			if title == "" {
				title = refMotion.Title
			}

			refMotionNames = append(refMotionNames, title)
		}
		data.ReferencedRecoMotions = strings.Join(refMotionNames, ", ")
	}

	switch options.Mode {
	case motionTextChanged:
		return data.motionTextChangedSlide(ctx)
	case motionTextDiff:
		return data.motionTextDiffSlide(ctx)
	case motionTextFinal:
		return data.motionTextDiffSlide(ctx)
	case motionTextModifiedFinal:
		return data.motionTextModifiedFinalSlide(ctx)
	}

	return data.templateData(map[string]any{}), nil
}

func (m *motionSlideCommonData) motionTextChangedSlide(ctx context.Context) (map[string]any, error) {
	changeRecoData, err := m.motionChangeRecos(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not fetch motion change recos: %w", err)
	}

	return m.templateData(changeRecoData), err
}

func (m *motionSlideCommonData) motionTextDiffSlide(ctx context.Context) (map[string]any, error) {
	data, err := m.motionChangeRecos(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not fetch motion change recos: %w", err)
	}
	amendments, err := m.motionAmendments(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not fetch amendments: %w", err)
	}

	maps.Copy(data, amendments)
	return m.templateData(data), nil
}

func (m *motionSlideCommonData) motionTextModifiedFinalSlide(ctx context.Context) (map[string]any, error) {
	fetch := m.ProjectionReq.Fetch
	crIDs := m.Motion.ChangeRecommendationIDs
	crs, err := fetch.MotionChangeRecommendation(crIDs...).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not fetch change recommendations slide data: %w", err)
	}

	titleChanges := []motionChangeReco{}
	for _, cr := range crs {
		if !cr.Internal && !cr.Rejected {
			newCr := motionChangeReco{
				ID:   cr.ID,
				Type: cr.Type,
				Text: template.HTML(cr.Text),
			}

			if cr.LineFrom == 0 {
				titleChanges = append(titleChanges, newCr)
			}
		}
	}

	return m.templateData(map[string]any{
		"TitleChangeRecos": titleChanges,
		"MotionText":       template.HTML(m.Motion.ModifiedFinalVersion),
	}), nil
}

func motionSubmitterList(motion *dsmodels.Motion) []string {
	submitters := []string{}
	slices.SortFunc(motion.SubmitterList, func(a dsmodels.MotionSubmitter, b dsmodels.MotionSubmitter) int {
		return a.Weight - b.Weight
	})
	for _, submitter := range motion.SubmitterList {
		if meetingUser, ok := submitter.MeetingUser.Value(); ok {
			submitters = append(submitters, viewmodels.MeetingUser_FullName(&meetingUser))
		}
	}
	if motion.AdditionalSubmitter != "" {
		submitters = append(submitters, motion.AdditionalSubmitter)
	}

	return submitters
}

type motionChangeReco struct {
	ID       int
	Type     string
	LineFrom int
	LineTo   int
	Rejected bool
	Text     template.HTML
}

func (m *motionSlideCommonData) motionChangeRecos(ctx context.Context) (map[string]any, error) {
	fetch := m.ProjectionReq.Fetch
	crIDs := m.Motion.ChangeRecommendationIDs
	crs, err := fetch.MotionChangeRecommendation(crIDs...).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not fetch change recommendations slide data: %w", err)
	}

	changeRecos := []motionChangeReco{}
	titleChanges := []motionChangeReco{}
	for _, cr := range crs {
		if !cr.Internal {
			newCr := motionChangeReco{
				ID:       cr.ID,
				Type:     cr.Type,
				LineFrom: cr.LineFrom,
				LineTo:   cr.LineTo,
				Rejected: cr.Rejected,
				Text:     template.HTML(cr.Text),
			}

			if cr.LineFrom > 0 {
				changeRecos = append(changeRecos, newCr)
			} else {
				titleChanges = append(titleChanges, newCr)
			}
		}
	}

	return m.templateData(map[string]any{
		"HasTitleChanges":   len(titleChanges) > 0,
		"TitleChangeRecos":  titleChanges,
		"MotionChangeRecos": changeRecos,
	}), nil
}

type motionAmendment struct {
	ID          int
	Number      string
	Title       string
	ChangeTitle string
	Paragraphs  map[string]template.HTML
	ChangeRecos []motionChangeReco
}

func (m *motionSlideCommonData) motionAmendments(ctx context.Context) (map[string]any, error) {
	fetch := m.ProjectionReq.Fetch
	amendmentIDs := m.Motion.AmendmentIDs
	mQ := fetch.Motion(amendmentIDs...)
	amendments, err := mQ.Preload(mQ.ChangeRecommendationList()).Preload(mQ.State()).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not fetch change recommendations slide data: %w", err)
	}

	tmplAmendments := []motionAmendment{}
	for _, amendment := range amendments {
		if amendment.State.MergeAmendmentIntoFinal != "do_merge" {
			continue
		}

		changeRecos := []motionChangeReco{}
		for _, cr := range amendment.ChangeRecommendationList {
			if !cr.Rejected && !cr.Internal {
				newCr := motionChangeReco{
					ID:       cr.ID,
					Type:     cr.Type,
					LineFrom: cr.LineFrom,
					LineTo:   cr.LineTo,
					Text:     template.HTML(cr.Text),
				}

				if cr.LineFrom > 0 {
					changeRecos = append(changeRecos, newCr)
				}
			}
		}

		changeTitle := m.ProjectionReq.Locale.Get("Amendment")
		if amendment.Number != "" {
			changeTitle = amendment.Number
		}

		data := motionAmendment{
			ID:          amendment.ID,
			Number:      amendment.Number,
			Title:       amendment.Title,
			ChangeTitle: changeTitle,
			ChangeRecos: changeRecos,
		}

		if err := json.Unmarshal(amendment.AmendmentParagraphs, &data.Paragraphs); err != nil {
			return nil, fmt.Errorf("could not parse amendment paragraphs: %w", err)
		}

		tmplAmendments = append(tmplAmendments, data)
	}

	return m.templateData(map[string]any{
		"Amendments": tmplAmendments,
	}), nil
}

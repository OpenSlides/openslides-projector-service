package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MotionState struct {
	AllowCreatePoll                  *bool    `json:"allow_create_poll"`
	AllowMotionForwarding            *bool    `json:"allow_motion_forwarding"`
	AllowSubmitterEdit               *bool    `json:"allow_submitter_edit"`
	AllowSupport                     *bool    `json:"allow_support"`
	CssClass                         string   `json:"css_class"`
	FirstStateOfWorkflowID           *int     `json:"first_state_of_workflow_id"`
	ID                               int      `json:"id"`
	IsInternal                       *bool    `json:"is_internal"`
	MeetingID                        int      `json:"meeting_id"`
	MergeAmendmentIntoFinal          *string  `json:"merge_amendment_into_final"`
	MotionIDs                        []int    `json:"motion_ids"`
	MotionRecommendationIDs          []int    `json:"motion_recommendation_ids"`
	Name                             string   `json:"name"`
	NextStateIDs                     []int    `json:"next_state_ids"`
	PreviousStateIDs                 []int    `json:"previous_state_ids"`
	RecommendationLabel              *string  `json:"recommendation_label"`
	Restrictions                     []string `json:"restrictions"`
	SetNumber                        *bool    `json:"set_number"`
	SetWorkflowTimestamp             *bool    `json:"set_workflow_timestamp"`
	ShowRecommendationExtensionField *bool    `json:"show_recommendation_extension_field"`
	ShowStateExtensionField          *bool    `json:"show_state_extension_field"`
	SubmitterWithdrawBackIDs         []int    `json:"submitter_withdraw_back_ids"`
	SubmitterWithdrawStateID         *int     `json:"submitter_withdraw_state_id"`
	Weight                           int      `json:"weight"`
	WorkflowID                       int      `json:"workflow_id"`
	loadedRelations                  map[string]struct{}
	firstStateOfWorkflow             *MotionWorkflow
	meeting                          *Meeting
	motionRecommendations            []*Motion
	motions                          []*Motion
	nextStates                       []*MotionState
	previousStates                   []*MotionState
	submitterWithdrawBacks           []*MotionState
	submitterWithdrawState           *MotionState
	workflow                         *MotionWorkflow
}

func (m *MotionState) CollectionName() string {
	return "motion_state"
}

func (m *MotionState) FirstStateOfWorkflow() *MotionWorkflow {
	if _, ok := m.loadedRelations["first_state_of_workflow_id"]; !ok {
		log.Panic().Msg("Tried to access FirstStateOfWorkflow relation of MotionState which was not loaded.")
	}

	return m.firstStateOfWorkflow
}

func (m *MotionState) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionState which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionState) MotionRecommendations() []*Motion {
	if _, ok := m.loadedRelations["motion_recommendation_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionRecommendations relation of MotionState which was not loaded.")
	}

	return m.motionRecommendations
}

func (m *MotionState) Motions() []*Motion {
	if _, ok := m.loadedRelations["motion_ids"]; !ok {
		log.Panic().Msg("Tried to access Motions relation of MotionState which was not loaded.")
	}

	return m.motions
}

func (m *MotionState) NextStates() []*MotionState {
	if _, ok := m.loadedRelations["next_state_ids"]; !ok {
		log.Panic().Msg("Tried to access NextStates relation of MotionState which was not loaded.")
	}

	return m.nextStates
}

func (m *MotionState) PreviousStates() []*MotionState {
	if _, ok := m.loadedRelations["previous_state_ids"]; !ok {
		log.Panic().Msg("Tried to access PreviousStates relation of MotionState which was not loaded.")
	}

	return m.previousStates
}

func (m *MotionState) SubmitterWithdrawBacks() []*MotionState {
	if _, ok := m.loadedRelations["submitter_withdraw_back_ids"]; !ok {
		log.Panic().Msg("Tried to access SubmitterWithdrawBacks relation of MotionState which was not loaded.")
	}

	return m.submitterWithdrawBacks
}

func (m *MotionState) SubmitterWithdrawState() *MotionState {
	if _, ok := m.loadedRelations["submitter_withdraw_state_id"]; !ok {
		log.Panic().Msg("Tried to access SubmitterWithdrawState relation of MotionState which was not loaded.")
	}

	return m.submitterWithdrawState
}

func (m *MotionState) Workflow() MotionWorkflow {
	if _, ok := m.loadedRelations["workflow_id"]; !ok {
		log.Panic().Msg("Tried to access Workflow relation of MotionState which was not loaded.")
	}

	return *m.workflow
}

func (m *MotionState) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "first_state_of_workflow_id":
		return m.firstStateOfWorkflow.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "motion_recommendation_ids":
		for _, r := range m.motionRecommendations {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_ids":
		for _, r := range m.motions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "next_state_ids":
		for _, r := range m.nextStates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "previous_state_ids":
		for _, r := range m.previousStates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "submitter_withdraw_back_ids":
		for _, r := range m.submitterWithdrawBacks {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "submitter_withdraw_state_id":
		return m.submitterWithdrawState.GetRelatedModelsAccessor()
	case "workflow_id":
		return m.workflow.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *MotionState) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "first_state_of_workflow_id":
			m.firstStateOfWorkflow = content.(*MotionWorkflow)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "motion_recommendation_ids":
			m.motionRecommendations = content.([]*Motion)
		case "motion_ids":
			m.motions = content.([]*Motion)
		case "next_state_ids":
			m.nextStates = content.([]*MotionState)
		case "previous_state_ids":
			m.previousStates = content.([]*MotionState)
		case "submitter_withdraw_back_ids":
			m.submitterWithdrawBacks = content.([]*MotionState)
		case "submitter_withdraw_state_id":
			m.submitterWithdrawState = content.(*MotionState)
		case "workflow_id":
			m.workflow = content.(*MotionWorkflow)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionState) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "first_state_of_workflow_id":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.firstStateOfWorkflow = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_recommendation_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionRecommendations = append(m.motionRecommendations, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motions = append(m.motions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "next_state_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.nextStates = append(m.nextStates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "previous_state_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.previousStates = append(m.previousStates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "submitter_withdraw_back_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.submitterWithdrawBacks = append(m.submitterWithdrawBacks, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "submitter_withdraw_state_id":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.submitterWithdrawState = &entry

		result = entry.GetRelatedModelsAccessor()
	case "workflow_id":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.workflow = &entry

		result = entry.GetRelatedModelsAccessor()
	default:
		return nil, fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return result, nil
}

func (m *MotionState) Get(field string) interface{} {
	switch field {
	case "allow_create_poll":
		return m.AllowCreatePoll
	case "allow_motion_forwarding":
		return m.AllowMotionForwarding
	case "allow_submitter_edit":
		return m.AllowSubmitterEdit
	case "allow_support":
		return m.AllowSupport
	case "css_class":
		return m.CssClass
	case "first_state_of_workflow_id":
		return m.FirstStateOfWorkflowID
	case "id":
		return m.ID
	case "is_internal":
		return m.IsInternal
	case "meeting_id":
		return m.MeetingID
	case "merge_amendment_into_final":
		return m.MergeAmendmentIntoFinal
	case "motion_ids":
		return m.MotionIDs
	case "motion_recommendation_ids":
		return m.MotionRecommendationIDs
	case "name":
		return m.Name
	case "next_state_ids":
		return m.NextStateIDs
	case "previous_state_ids":
		return m.PreviousStateIDs
	case "recommendation_label":
		return m.RecommendationLabel
	case "restrictions":
		return m.Restrictions
	case "set_number":
		return m.SetNumber
	case "set_workflow_timestamp":
		return m.SetWorkflowTimestamp
	case "show_recommendation_extension_field":
		return m.ShowRecommendationExtensionField
	case "show_state_extension_field":
		return m.ShowStateExtensionField
	case "submitter_withdraw_back_ids":
		return m.SubmitterWithdrawBackIDs
	case "submitter_withdraw_state_id":
		return m.SubmitterWithdrawStateID
	case "weight":
		return m.Weight
	case "workflow_id":
		return m.WorkflowID
	}

	return nil
}

func (m *MotionState) GetFqids(field string) []string {
	switch field {
	case "first_state_of_workflow_id":
		if m.FirstStateOfWorkflowID != nil {
			return []string{"motion_workflow/" + strconv.Itoa(*m.FirstStateOfWorkflowID)}
		}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "motion_recommendation_ids":
		r := make([]string, len(m.MotionRecommendationIDs))
		for i, id := range m.MotionRecommendationIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "motion_ids":
		r := make([]string, len(m.MotionIDs))
		for i, id := range m.MotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "next_state_ids":
		r := make([]string, len(m.NextStateIDs))
		for i, id := range m.NextStateIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r

	case "previous_state_ids":
		r := make([]string, len(m.PreviousStateIDs))
		for i, id := range m.PreviousStateIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r

	case "submitter_withdraw_back_ids":
		r := make([]string, len(m.SubmitterWithdrawBackIDs))
		for i, id := range m.SubmitterWithdrawBackIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r

	case "submitter_withdraw_state_id":
		if m.SubmitterWithdrawStateID != nil {
			return []string{"motion_state/" + strconv.Itoa(*m.SubmitterWithdrawStateID)}
		}

	case "workflow_id":
		return []string{"motion_workflow/" + strconv.Itoa(m.WorkflowID)}
	}
	return []string{}
}

func (m *MotionState) Update(data map[string]string) error {
	if val, ok := data["allow_create_poll"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllowCreatePoll)
		if err != nil {
			return err
		}
	}

	if val, ok := data["allow_motion_forwarding"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllowMotionForwarding)
		if err != nil {
			return err
		}
	}

	if val, ok := data["allow_submitter_edit"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllowSubmitterEdit)
		if err != nil {
			return err
		}
	}

	if val, ok := data["allow_support"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllowSupport)
		if err != nil {
			return err
		}
	}

	if val, ok := data["css_class"]; ok {
		err := json.Unmarshal([]byte(val), &m.CssClass)
		if err != nil {
			return err
		}
	}

	if val, ok := data["first_state_of_workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FirstStateOfWorkflowID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsInternal)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["merge_amendment_into_final"]; ok {
		err := json.Unmarshal([]byte(val), &m.MergeAmendmentIntoFinal)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_ids"]; ok {
			m.motions = slices.DeleteFunc(m.motions, func(r *Motion) bool {
				return !slices.Contains(m.MotionIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_recommendation_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionRecommendationIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_recommendation_ids"]; ok {
			m.motionRecommendations = slices.DeleteFunc(m.motionRecommendations, func(r *Motion) bool {
				return !slices.Contains(m.MotionRecommendationIDs, r.ID)
			})
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["next_state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.NextStateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["next_state_ids"]; ok {
			m.nextStates = slices.DeleteFunc(m.nextStates, func(r *MotionState) bool {
				return !slices.Contains(m.NextStateIDs, r.ID)
			})
		}
	}

	if val, ok := data["previous_state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PreviousStateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["previous_state_ids"]; ok {
			m.previousStates = slices.DeleteFunc(m.previousStates, func(r *MotionState) bool {
				return !slices.Contains(m.PreviousStateIDs, r.ID)
			})
		}
	}

	if val, ok := data["recommendation_label"]; ok {
		err := json.Unmarshal([]byte(val), &m.RecommendationLabel)
		if err != nil {
			return err
		}
	}

	if val, ok := data["restrictions"]; ok {
		err := json.Unmarshal([]byte(val), &m.Restrictions)
		if err != nil {
			return err
		}
	}

	if val, ok := data["set_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SetNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["set_workflow_timestamp"]; ok {
		err := json.Unmarshal([]byte(val), &m.SetWorkflowTimestamp)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_recommendation_extension_field"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowRecommendationExtensionField)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_state_extension_field"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowStateExtensionField)
		if err != nil {
			return err
		}
	}

	if val, ok := data["submitter_withdraw_back_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SubmitterWithdrawBackIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["submitter_withdraw_back_ids"]; ok {
			m.submitterWithdrawBacks = slices.DeleteFunc(m.submitterWithdrawBacks, func(r *MotionState) bool {
				return !slices.Contains(m.SubmitterWithdrawBackIDs, r.ID)
			})
		}
	}

	if val, ok := data["submitter_withdraw_state_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.SubmitterWithdrawStateID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.WorkflowID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionState) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

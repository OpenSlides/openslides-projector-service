package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MotionWorkflow struct {
	DefaultAmendmentWorkflowMeetingID *int   `json:"default_amendment_workflow_meeting_id"`
	DefaultWorkflowMeetingID          *int   `json:"default_workflow_meeting_id"`
	FirstStateID                      int    `json:"first_state_id"`
	ID                                int    `json:"id"`
	MeetingID                         int    `json:"meeting_id"`
	Name                              string `json:"name"`
	SequentialNumber                  int    `json:"sequential_number"`
	StateIDs                          []int  `json:"state_ids"`
	loadedRelations                   map[string]struct{}
	defaultAmendmentWorkflowMeeting   *Meeting
	defaultWorkflowMeeting            *Meeting
	firstState                        *MotionState
	meeting                           *Meeting
	states                            []*MotionState
}

func (m *MotionWorkflow) CollectionName() string {
	return "motion_workflow"
}

func (m *MotionWorkflow) DefaultAmendmentWorkflowMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_amendment_workflow_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultAmendmentWorkflowMeeting relation of MotionWorkflow which was not loaded.")
	}

	return m.defaultAmendmentWorkflowMeeting
}

func (m *MotionWorkflow) DefaultWorkflowMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_workflow_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultWorkflowMeeting relation of MotionWorkflow which was not loaded.")
	}

	return m.defaultWorkflowMeeting
}

func (m *MotionWorkflow) FirstState() MotionState {
	if _, ok := m.loadedRelations["first_state_id"]; !ok {
		log.Panic().Msg("Tried to access FirstState relation of MotionWorkflow which was not loaded.")
	}

	return *m.firstState
}

func (m *MotionWorkflow) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionWorkflow which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionWorkflow) States() []*MotionState {
	if _, ok := m.loadedRelations["state_ids"]; !ok {
		log.Panic().Msg("Tried to access States relation of MotionWorkflow which was not loaded.")
	}

	return m.states
}

func (m *MotionWorkflow) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "default_amendment_workflow_meeting_id":
			m.defaultAmendmentWorkflowMeeting = content.(*Meeting)
		case "default_workflow_meeting_id":
			m.defaultWorkflowMeeting = content.(*Meeting)
		case "first_state_id":
			m.firstState = content.(*MotionState)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "state_ids":
			m.states = content.([]*MotionState)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionWorkflow) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "default_amendment_workflow_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultAmendmentWorkflowMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "default_workflow_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultWorkflowMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "first_state_id":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.firstState = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "state_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.states = append(m.states, &entry)

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

func (m *MotionWorkflow) Get(field string) interface{} {
	switch field {
	case "default_amendment_workflow_meeting_id":
		return m.DefaultAmendmentWorkflowMeetingID
	case "default_workflow_meeting_id":
		return m.DefaultWorkflowMeetingID
	case "first_state_id":
		return m.FirstStateID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "sequential_number":
		return m.SequentialNumber
	case "state_ids":
		return m.StateIDs
	}

	return nil
}

func (m *MotionWorkflow) GetFqids(field string) []string {
	switch field {
	case "default_amendment_workflow_meeting_id":
		if m.DefaultAmendmentWorkflowMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultAmendmentWorkflowMeetingID)}
		}

	case "default_workflow_meeting_id":
		if m.DefaultWorkflowMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultWorkflowMeetingID)}
		}

	case "first_state_id":
		return []string{"motion_state/" + strconv.Itoa(m.FirstStateID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "state_ids":
		r := make([]string, len(m.StateIDs))
		for i, id := range m.StateIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *MotionWorkflow) Update(data map[string]string) error {
	if val, ok := data["default_amendment_workflow_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultAmendmentWorkflowMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_workflow_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultWorkflowMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["first_state_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FirstStateID)
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

	if val, ok := data["meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["state_ids"]; ok {
			m.states = slices.DeleteFunc(m.states, func(r *MotionState) bool {
				return !slices.Contains(m.StateIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *MotionWorkflow) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

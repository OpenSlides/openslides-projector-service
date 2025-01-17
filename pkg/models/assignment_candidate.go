package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type AssignmentCandidate struct {
	AssignmentID    int  `json:"assignment_id"`
	ID              int  `json:"id"`
	MeetingID       int  `json:"meeting_id"`
	MeetingUserID   *int `json:"meeting_user_id"`
	Weight          *int `json:"weight"`
	loadedRelations map[string]struct{}
	meetingUser     *MeetingUser
	assignment      *Assignment
	meeting         *Meeting
}

func (m *AssignmentCandidate) CollectionName() string {
	return "assignment_candidate"
}

func (m *AssignmentCandidate) MeetingUser() *MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of AssignmentCandidate which was not loaded.")
	}

	return m.meetingUser
}

func (m *AssignmentCandidate) Assignment() Assignment {
	if _, ok := m.loadedRelations["assignment_id"]; !ok {
		log.Panic().Msg("Tried to access Assignment relation of AssignmentCandidate which was not loaded.")
	}

	return *m.assignment
}

func (m *AssignmentCandidate) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of AssignmentCandidate which was not loaded.")
	}

	return *m.meeting
}

func (m *AssignmentCandidate) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_user_id":
			m.meetingUser = content.(*MeetingUser)
		case "assignment_id":
			m.assignment = content.(*Assignment)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *AssignmentCandidate) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "meeting_user_id":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUser = &entry

		result = entry.GetRelatedModelsAccessor()
	case "assignment_id":
		var entry Assignment
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.assignment = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

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

func (m *AssignmentCandidate) Get(field string) interface{} {
	switch field {
	case "assignment_id":
		return m.AssignmentID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *AssignmentCandidate) GetFqids(field string) []string {
	switch field {
	case "meeting_user_id":
		if m.MeetingUserID != nil {
			return []string{"meeting_user/" + strconv.Itoa(*m.MeetingUserID)}
		}

	case "assignment_id":
		return []string{"assignment/" + strconv.Itoa(m.AssignmentID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}
	}
	return []string{}
}

func (m *AssignmentCandidate) Update(data map[string]string) error {
	if val, ok := data["assignment_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentID)
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

	if val, ok := data["meeting_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserID)
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

	return nil
}

func (m *AssignmentCandidate) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

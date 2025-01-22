package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type PollCandidate struct {
	ID                  int  `json:"id"`
	MeetingID           int  `json:"meeting_id"`
	PollCandidateListID int  `json:"poll_candidate_list_id"`
	UserID              *int `json:"user_id"`
	Weight              int  `json:"weight"`
	loadedRelations     map[string]struct{}
	meeting             *Meeting
	pollCandidateList   *PollCandidateList
	user                *User
}

func (m *PollCandidate) CollectionName() string {
	return "poll_candidate"
}

func (m *PollCandidate) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PollCandidate which was not loaded.")
	}

	return *m.meeting
}

func (m *PollCandidate) PollCandidateList() PollCandidateList {
	if _, ok := m.loadedRelations["poll_candidate_list_id"]; !ok {
		log.Panic().Msg("Tried to access PollCandidateList relation of PollCandidate which was not loaded.")
	}

	return *m.pollCandidateList
}

func (m *PollCandidate) User() *User {
	if _, ok := m.loadedRelations["user_id"]; !ok {
		log.Panic().Msg("Tried to access User relation of PollCandidate which was not loaded.")
	}

	return m.user
}

func (m *PollCandidate) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "poll_candidate_list_id":
		return m.pollCandidateList.GetRelatedModelsAccessor()
	case "user_id":
		return m.user.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *PollCandidate) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "poll_candidate_list_id":
			m.pollCandidateList = content.(*PollCandidateList)
		case "user_id":
			m.user = content.(*User)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *PollCandidate) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "poll_candidate_list_id":
		var entry PollCandidateList
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCandidateList = &entry

		result = entry.GetRelatedModelsAccessor()
	case "user_id":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.user = &entry

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

func (m *PollCandidate) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "poll_candidate_list_id":
		return m.PollCandidateListID
	case "user_id":
		return m.UserID
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *PollCandidate) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "poll_candidate_list_id":
		return []string{"poll_candidate_list/" + strconv.Itoa(m.PollCandidateListID)}

	case "user_id":
		if m.UserID != nil {
			return []string{"user/" + strconv.Itoa(*m.UserID)}
		}
	}
	return []string{}
}

func (m *PollCandidate) Update(data map[string]string) error {
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

	if val, ok := data["poll_candidate_list_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateListID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserID)
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

func (m *PollCandidate) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

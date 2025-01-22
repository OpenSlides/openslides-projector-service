package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

type PollCandidateList struct {
	ID               int   `json:"id"`
	MeetingID        int   `json:"meeting_id"`
	OptionID         int   `json:"option_id"`
	PollCandidateIDs []int `json:"poll_candidate_ids"`
	loadedRelations  map[string]struct{}
	meeting          *Meeting
	option           *Option
	pollCandidates   []*PollCandidate
}

func (m *PollCandidateList) CollectionName() string {
	return "poll_candidate_list"
}

func (m *PollCandidateList) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PollCandidateList which was not loaded.")
	}

	return *m.meeting
}

func (m *PollCandidateList) Option() Option {
	if _, ok := m.loadedRelations["option_id"]; !ok {
		log.Panic().Msg("Tried to access Option relation of PollCandidateList which was not loaded.")
	}

	return *m.option
}

func (m *PollCandidateList) PollCandidates() []*PollCandidate {
	if _, ok := m.loadedRelations["poll_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access PollCandidates relation of PollCandidateList which was not loaded.")
	}

	return m.pollCandidates
}

func (m *PollCandidateList) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "option_id":
		return m.option.GetRelatedModelsAccessor()
	case "poll_candidate_ids":
		for _, r := range m.pollCandidates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *PollCandidateList) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "option_id":
			m.option = content.(*Option)
		case "poll_candidate_ids":
			m.pollCandidates = content.([]*PollCandidate)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *PollCandidateList) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "option_id":
		var entry Option
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.option = &entry

		result = entry.GetRelatedModelsAccessor()
	case "poll_candidate_ids":
		var entry PollCandidate
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCandidates = append(m.pollCandidates, &entry)

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

func (m *PollCandidateList) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "option_id":
		return m.OptionID
	case "poll_candidate_ids":
		return m.PollCandidateIDs
	}

	return nil
}

func (m *PollCandidateList) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "option_id":
		return []string{"option/" + strconv.Itoa(m.OptionID)}

	case "poll_candidate_ids":
		r := make([]string, len(m.PollCandidateIDs))
		for i, id := range m.PollCandidateIDs {
			r[i] = "poll_candidate/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *PollCandidateList) Update(data map[string]string) error {
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

	if val, ok := data["option_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["poll_candidate_ids"]; ok {
			m.pollCandidates = slices.DeleteFunc(m.pollCandidates, func(r *PollCandidate) bool {
				return !slices.Contains(m.PollCandidateIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *PollCandidateList) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

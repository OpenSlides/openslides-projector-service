package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MotionStatuteParagraph struct {
	ID               int     `json:"id"`
	MeetingID        int     `json:"meeting_id"`
	MotionIDs        []int   `json:"motion_ids"`
	SequentialNumber int     `json:"sequential_number"`
	Text             *string `json:"text"`
	Title            string  `json:"title"`
	Weight           *int    `json:"weight"`
	loadedRelations  map[string]struct{}
	motions          []*Motion
	meeting          *Meeting
}

func (m *MotionStatuteParagraph) CollectionName() string {
	return "motion_statute_paragraph"
}

func (m *MotionStatuteParagraph) Motions() []*Motion {
	if _, ok := m.loadedRelations["motion_ids"]; !ok {
		log.Panic().Msg("Tried to access Motions relation of MotionStatuteParagraph which was not loaded.")
	}

	return m.motions
}

func (m *MotionStatuteParagraph) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionStatuteParagraph which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionStatuteParagraph) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "motion_ids":
			m.motions = content.([]*Motion)
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

func (m *MotionStatuteParagraph) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motions = append(m.motions, &entry)

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

func (m *MotionStatuteParagraph) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "motion_ids":
		return m.MotionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "text":
		return m.Text
	case "title":
		return m.Title
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *MotionStatuteParagraph) GetFqids(field string) []string {
	switch field {
	case "motion_ids":
		r := make([]string, len(m.MotionIDs))
		for i, id := range m.MotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}
	}
	return []string{}
}

func (m *MotionStatuteParagraph) Update(data map[string]string) error {
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

	if val, ok := data["motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionIDs)
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

	if val, ok := data["text"]; ok {
		err := json.Unmarshal([]byte(val), &m.Text)
		if err != nil {
			return err
		}
	}

	if val, ok := data["title"]; ok {
		err := json.Unmarshal([]byte(val), &m.Title)
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

func (m *MotionStatuteParagraph) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

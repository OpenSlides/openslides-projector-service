package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MotionChangeRecommendation struct {
	CreationTime     *int    `json:"creation_time"`
	ID               int     `json:"id"`
	Internal         *bool   `json:"internal"`
	LineFrom         *int    `json:"line_from"`
	LineTo           *int    `json:"line_to"`
	MeetingID        int     `json:"meeting_id"`
	MotionID         int     `json:"motion_id"`
	OtherDescription *string `json:"other_description"`
	Rejected         *bool   `json:"rejected"`
	Text             *string `json:"text"`
	Type             *string `json:"type"`
	loadedRelations  map[string]struct{}
	motion           *Motion
	meeting          *Meeting
}

func (m *MotionChangeRecommendation) CollectionName() string {
	return "motion_change_recommendation"
}

func (m *MotionChangeRecommendation) Motion() Motion {
	if _, ok := m.loadedRelations["motion_id"]; !ok {
		log.Panic().Msg("Tried to access Motion relation of MotionChangeRecommendation which was not loaded.")
	}

	return *m.motion
}

func (m *MotionChangeRecommendation) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionChangeRecommendation which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionChangeRecommendation) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "motion_id":
			m.motion = content.(*Motion)
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

func (m *MotionChangeRecommendation) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "motion_id":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motion = &entry

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

func (m *MotionChangeRecommendation) Get(field string) interface{} {
	switch field {
	case "creation_time":
		return m.CreationTime
	case "id":
		return m.ID
	case "internal":
		return m.Internal
	case "line_from":
		return m.LineFrom
	case "line_to":
		return m.LineTo
	case "meeting_id":
		return m.MeetingID
	case "motion_id":
		return m.MotionID
	case "other_description":
		return m.OtherDescription
	case "rejected":
		return m.Rejected
	case "text":
		return m.Text
	case "type":
		return m.Type
	}

	return nil
}

func (m *MotionChangeRecommendation) GetFqids(field string) []string {
	switch field {
	case "motion_id":
		return []string{"motion/" + strconv.Itoa(m.MotionID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}
	}
	return []string{}
}

func (m *MotionChangeRecommendation) Update(data map[string]string) error {
	if val, ok := data["creation_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.CreationTime)
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

	if val, ok := data["internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.Internal)
		if err != nil {
			return err
		}
	}

	if val, ok := data["line_from"]; ok {
		err := json.Unmarshal([]byte(val), &m.LineFrom)
		if err != nil {
			return err
		}
	}

	if val, ok := data["line_to"]; ok {
		err := json.Unmarshal([]byte(val), &m.LineTo)
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

	if val, ok := data["motion_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["other_description"]; ok {
		err := json.Unmarshal([]byte(val), &m.OtherDescription)
		if err != nil {
			return err
		}
	}

	if val, ok := data["rejected"]; ok {
		err := json.Unmarshal([]byte(val), &m.Rejected)
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

	if val, ok := data["type"]; ok {
		err := json.Unmarshal([]byte(val), &m.Type)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionChangeRecommendation) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

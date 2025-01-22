package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

type PointOfOrderCategory struct {
	ID              int    `json:"id"`
	MeetingID       int    `json:"meeting_id"`
	Rank            int    `json:"rank"`
	SpeakerIDs      []int  `json:"speaker_ids"`
	Text            string `json:"text"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	speakers        []*Speaker
}

func (m *PointOfOrderCategory) CollectionName() string {
	return "point_of_order_category"
}

func (m *PointOfOrderCategory) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PointOfOrderCategory which was not loaded.")
	}

	return *m.meeting
}

func (m *PointOfOrderCategory) Speakers() []*Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of PointOfOrderCategory which was not loaded.")
	}

	return m.speakers
}

func (m *PointOfOrderCategory) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "speaker_ids":
			m.speakers = content.([]*Speaker)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *PointOfOrderCategory) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "speaker_ids":
		var entry Speaker
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.speakers = append(m.speakers, &entry)

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

func (m *PointOfOrderCategory) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "rank":
		return m.Rank
	case "speaker_ids":
		return m.SpeakerIDs
	case "text":
		return m.Text
	}

	return nil
}

func (m *PointOfOrderCategory) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "speaker_ids":
		r := make([]string, len(m.SpeakerIDs))
		for i, id := range m.SpeakerIDs {
			r[i] = "speaker/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *PointOfOrderCategory) Update(data map[string]string) error {
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

	if val, ok := data["rank"]; ok {
		err := json.Unmarshal([]byte(val), &m.Rank)
		if err != nil {
			return err
		}
	}

	if val, ok := data["speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SpeakerIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["speaker_ids"]; ok {
			m.speakers = slices.DeleteFunc(m.speakers, func(r *Speaker) bool {
				return !slices.Contains(m.SpeakerIDs, r.ID)
			})
		}
	}

	if val, ok := data["text"]; ok {
		err := json.Unmarshal([]byte(val), &m.Text)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *PointOfOrderCategory) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

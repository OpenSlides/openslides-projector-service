package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

type ProjectorMessage struct {
	ID              int     `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	Message         *string `json:"message"`
	ProjectionIDs   []int   `json:"projection_ids"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	projections     []*Projection
}

func (m *ProjectorMessage) CollectionName() string {
	return "projector_message"
}

func (m *ProjectorMessage) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ProjectorMessage which was not loaded.")
	}

	return *m.meeting
}

func (m *ProjectorMessage) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of ProjectorMessage which was not loaded.")
	}

	return m.projections
}

func (m *ProjectorMessage) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "projection_ids":
			m.projections = content.([]*Projection)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *ProjectorMessage) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projections = append(m.projections, &entry)

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

func (m *ProjectorMessage) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "message":
		return m.Message
	case "projection_ids":
		return m.ProjectionIDs
	}

	return nil
}

func (m *ProjectorMessage) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *ProjectorMessage) Update(data map[string]string) error {
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

	if val, ok := data["message"]; ok {
		err := json.Unmarshal([]byte(val), &m.Message)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["projection_ids"]; ok {
			m.projections = slices.DeleteFunc(m.projections, func(r *Projection) bool {
				return !slices.Contains(m.ProjectionIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *ProjectorMessage) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MotionCategory struct {
	ChildIDs         []int   `json:"child_ids"`
	ID               int     `json:"id"`
	Level            *int    `json:"level"`
	MeetingID        int     `json:"meeting_id"`
	MotionIDs        []int   `json:"motion_ids"`
	Name             string  `json:"name"`
	ParentID         *int    `json:"parent_id"`
	Prefix           *string `json:"prefix"`
	SequentialNumber int     `json:"sequential_number"`
	Weight           *int    `json:"weight"`
	loadedRelations  map[string]struct{}
	childs           []*MotionCategory
	meeting          *Meeting
	motions          []*Motion
	parent           *MotionCategory
}

func (m *MotionCategory) CollectionName() string {
	return "motion_category"
}

func (m *MotionCategory) Childs() []*MotionCategory {
	if _, ok := m.loadedRelations["child_ids"]; !ok {
		log.Panic().Msg("Tried to access Childs relation of MotionCategory which was not loaded.")
	}

	return m.childs
}

func (m *MotionCategory) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionCategory which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionCategory) Motions() []*Motion {
	if _, ok := m.loadedRelations["motion_ids"]; !ok {
		log.Panic().Msg("Tried to access Motions relation of MotionCategory which was not loaded.")
	}

	return m.motions
}

func (m *MotionCategory) Parent() *MotionCategory {
	if _, ok := m.loadedRelations["parent_id"]; !ok {
		log.Panic().Msg("Tried to access Parent relation of MotionCategory which was not loaded.")
	}

	return m.parent
}

func (m *MotionCategory) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "child_ids":
			m.childs = content.([]*MotionCategory)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "motion_ids":
			m.motions = content.([]*Motion)
		case "parent_id":
			m.parent = content.(*MotionCategory)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionCategory) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "child_ids":
		var entry MotionCategory
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.childs = append(m.childs, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motions = append(m.motions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "parent_id":
		var entry MotionCategory
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.parent = &entry

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

func (m *MotionCategory) Get(field string) interface{} {
	switch field {
	case "child_ids":
		return m.ChildIDs
	case "id":
		return m.ID
	case "level":
		return m.Level
	case "meeting_id":
		return m.MeetingID
	case "motion_ids":
		return m.MotionIDs
	case "name":
		return m.Name
	case "parent_id":
		return m.ParentID
	case "prefix":
		return m.Prefix
	case "sequential_number":
		return m.SequentialNumber
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *MotionCategory) GetFqids(field string) []string {
	switch field {
	case "child_ids":
		r := make([]string, len(m.ChildIDs))
		for i, id := range m.ChildIDs {
			r[i] = "motion_category/" + strconv.Itoa(id)
		}
		return r

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "motion_ids":
		r := make([]string, len(m.MotionIDs))
		for i, id := range m.MotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "parent_id":
		if m.ParentID != nil {
			return []string{"motion_category/" + strconv.Itoa(*m.ParentID)}
		}
	}
	return []string{}
}

func (m *MotionCategory) Update(data map[string]string) error {
	if val, ok := data["child_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChildIDs)
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

	if val, ok := data["level"]; ok {
		err := json.Unmarshal([]byte(val), &m.Level)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["parent_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ParentID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["prefix"]; ok {
		err := json.Unmarshal([]byte(val), &m.Prefix)
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

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionCategory) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type MotionComment struct {
	Comment         *string `json:"comment"`
	ID              int     `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	MotionID        int     `json:"motion_id"`
	SectionID       int     `json:"section_id"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	motion          *Motion
	section         *MotionCommentSection
}

func (m *MotionComment) CollectionName() string {
	return "motion_comment"
}

func (m *MotionComment) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionComment which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionComment) Motion() Motion {
	if _, ok := m.loadedRelations["motion_id"]; !ok {
		log.Panic().Msg("Tried to access Motion relation of MotionComment which was not loaded.")
	}

	return *m.motion
}

func (m *MotionComment) Section() MotionCommentSection {
	if _, ok := m.loadedRelations["section_id"]; !ok {
		log.Panic().Msg("Tried to access Section relation of MotionComment which was not loaded.")
	}

	return *m.section
}

func (m *MotionComment) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "motion_id":
		return m.motion.GetRelatedModelsAccessor()
	case "section_id":
		return m.section.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *MotionComment) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "motion_id":
			m.motion = content.(*Motion)
		case "section_id":
			m.section = content.(*MotionCommentSection)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionComment) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "motion_id":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motion = &entry

		result = entry.GetRelatedModelsAccessor()
	case "section_id":
		var entry MotionCommentSection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.section = &entry

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

func (m *MotionComment) Get(field string) interface{} {
	switch field {
	case "comment":
		return m.Comment
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "motion_id":
		return m.MotionID
	case "section_id":
		return m.SectionID
	}

	return nil
}

func (m *MotionComment) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "motion_id":
		return []string{"motion/" + strconv.Itoa(m.MotionID)}

	case "section_id":
		return []string{"motion_comment_section/" + strconv.Itoa(m.SectionID)}
	}
	return []string{}
}

func (m *MotionComment) Update(data map[string]string) error {
	if val, ok := data["comment"]; ok {
		err := json.Unmarshal([]byte(val), &m.Comment)
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

	if val, ok := data["motion_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["section_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.SectionID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionComment) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

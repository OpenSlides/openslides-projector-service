package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

type StructureLevel struct {
	Color                           *string `json:"color"`
	DefaultTime                     *int    `json:"default_time"`
	ID                              int     `json:"id"`
	MeetingID                       int     `json:"meeting_id"`
	MeetingUserIDs                  []int   `json:"meeting_user_ids"`
	Name                            string  `json:"name"`
	StructureLevelListOfSpeakersIDs []int   `json:"structure_level_list_of_speakers_ids"`
	loadedRelations                 map[string]struct{}
	meeting                         *Meeting
	meetingUsers                    []*MeetingUser
	structureLevelListOfSpeakerss   []*StructureLevelListOfSpeakers
}

func (m *StructureLevel) CollectionName() string {
	return "structure_level"
}

func (m *StructureLevel) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of StructureLevel which was not loaded.")
	}

	return *m.meeting
}

func (m *StructureLevel) MeetingUsers() []*MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingUsers relation of StructureLevel which was not loaded.")
	}

	return m.meetingUsers
}

func (m *StructureLevel) StructureLevelListOfSpeakerss() []*StructureLevelListOfSpeakers {
	if _, ok := m.loadedRelations["structure_level_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevelListOfSpeakerss relation of StructureLevel which was not loaded.")
	}

	return m.structureLevelListOfSpeakerss
}

func (m *StructureLevel) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "meeting_user_ids":
		for _, r := range m.meetingUsers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "structure_level_list_of_speakers_ids":
		for _, r := range m.structureLevelListOfSpeakerss {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *StructureLevel) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_ids":
			m.meetingUsers = content.([]*MeetingUser)
		case "structure_level_list_of_speakers_ids":
			m.structureLevelListOfSpeakerss = content.([]*StructureLevelListOfSpeakers)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *StructureLevel) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "meeting_user_ids":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUsers = append(m.meetingUsers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "structure_level_list_of_speakers_ids":
		var entry StructureLevelListOfSpeakers
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.structureLevelListOfSpeakerss = append(m.structureLevelListOfSpeakerss, &entry)

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

func (m *StructureLevel) Get(field string) interface{} {
	switch field {
	case "color":
		return m.Color
	case "default_time":
		return m.DefaultTime
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "name":
		return m.Name
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	}

	return nil
}

func (m *StructureLevel) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_ids":
		r := make([]string, len(m.MeetingUserIDs))
		for i, id := range m.MeetingUserIDs {
			r[i] = "meeting_user/" + strconv.Itoa(id)
		}
		return r

	case "structure_level_list_of_speakers_ids":
		r := make([]string, len(m.StructureLevelListOfSpeakersIDs))
		for i, id := range m.StructureLevelListOfSpeakersIDs {
			r[i] = "structure_level_list_of_speakers/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *StructureLevel) Update(data map[string]string) error {
	if val, ok := data["color"]; ok {
		err := json.Unmarshal([]byte(val), &m.Color)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultTime)
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

	if val, ok := data["meeting_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["meeting_user_ids"]; ok {
			m.meetingUsers = slices.DeleteFunc(m.meetingUsers, func(r *MeetingUser) bool {
				return !slices.Contains(m.MeetingUserIDs, r.ID)
			})
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelListOfSpeakersIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["structure_level_list_of_speakers_ids"]; ok {
			m.structureLevelListOfSpeakerss = slices.DeleteFunc(m.structureLevelListOfSpeakerss, func(r *StructureLevelListOfSpeakers) bool {
				return !slices.Contains(m.StructureLevelListOfSpeakersIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *StructureLevel) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

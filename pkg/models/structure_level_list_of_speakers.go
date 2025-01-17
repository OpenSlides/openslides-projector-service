package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type StructureLevelListOfSpeakers struct {
	AdditionalTime   *float32 `json:"additional_time"`
	CurrentStartTime *int     `json:"current_start_time"`
	ID               int      `json:"id"`
	InitialTime      int      `json:"initial_time"`
	ListOfSpeakersID int      `json:"list_of_speakers_id"`
	MeetingID        int      `json:"meeting_id"`
	RemainingTime    float32  `json:"remaining_time"`
	SpeakerIDs       []int    `json:"speaker_ids"`
	StructureLevelID int      `json:"structure_level_id"`
	loadedRelations  map[string]struct{}
	structureLevel   *StructureLevel
	listOfSpeakers   *ListOfSpeakers
	meeting          *Meeting
	speakers         []*Speaker
}

func (m *StructureLevelListOfSpeakers) CollectionName() string {
	return "structure_level_list_of_speakers"
}

func (m *StructureLevelListOfSpeakers) StructureLevel() StructureLevel {
	if _, ok := m.loadedRelations["structure_level_id"]; !ok {
		log.Panic().Msg("Tried to access StructureLevel relation of StructureLevelListOfSpeakers which was not loaded.")
	}

	return *m.structureLevel
}

func (m *StructureLevelListOfSpeakers) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of StructureLevelListOfSpeakers which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *StructureLevelListOfSpeakers) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of StructureLevelListOfSpeakers which was not loaded.")
	}

	return *m.meeting
}

func (m *StructureLevelListOfSpeakers) Speakers() []*Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of StructureLevelListOfSpeakers which was not loaded.")
	}

	return m.speakers
}

func (m *StructureLevelListOfSpeakers) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "structure_level_id":
			m.structureLevel = content.(*StructureLevel)
		case "list_of_speakers_id":
			m.listOfSpeakers = content.(*ListOfSpeakers)
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

func (m *StructureLevelListOfSpeakers) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "structure_level_id":
		var entry StructureLevel
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.structureLevel = &entry

		result = entry.GetRelatedModelsAccessor()
	case "list_of_speakers_id":
		var entry ListOfSpeakers
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.listOfSpeakers = &entry

		result = entry.GetRelatedModelsAccessor()
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

func (m *StructureLevelListOfSpeakers) Get(field string) interface{} {
	switch field {
	case "additional_time":
		return m.AdditionalTime
	case "current_start_time":
		return m.CurrentStartTime
	case "id":
		return m.ID
	case "initial_time":
		return m.InitialTime
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "remaining_time":
		return m.RemainingTime
	case "speaker_ids":
		return m.SpeakerIDs
	case "structure_level_id":
		return m.StructureLevelID
	}

	return nil
}

func (m *StructureLevelListOfSpeakers) GetFqids(field string) []string {
	switch field {
	case "structure_level_id":
		return []string{"structure_level/" + strconv.Itoa(m.StructureLevelID)}

	case "list_of_speakers_id":
		return []string{"list_of_speakers/" + strconv.Itoa(m.ListOfSpeakersID)}

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

func (m *StructureLevelListOfSpeakers) Update(data map[string]string) error {
	if val, ok := data["additional_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.AdditionalTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["current_start_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.CurrentStartTime)
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

	if val, ok := data["initial_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.InitialTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersID)
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

	if val, ok := data["remaining_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.RemainingTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SpeakerIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *StructureLevelListOfSpeakers) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

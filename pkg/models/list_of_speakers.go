package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type ListOfSpeakers struct {
	Closed                          *bool   `json:"closed"`
	ContentObjectID                 string  `json:"content_object_id"`
	ID                              int     `json:"id"`
	MeetingID                       int     `json:"meeting_id"`
	ModeratorNotes                  *string `json:"moderator_notes"`
	ProjectionIDs                   []int   `json:"projection_ids"`
	SequentialNumber                int     `json:"sequential_number"`
	SpeakerIDs                      []int   `json:"speaker_ids"`
	StructureLevelListOfSpeakersIDs []int   `json:"structure_level_list_of_speakers_ids"`
	loadedRelations                 map[string]struct{}
	contentObject                   IBaseModel
	meeting                         *Meeting
	projections                     []*Projection
	speakers                        []*Speaker
	structureLevelListOfSpeakerss   []*StructureLevelListOfSpeakers
}

func (m *ListOfSpeakers) CollectionName() string {
	return "list_of_speakers"
}

func (m *ListOfSpeakers) ContentObject() IBaseModel {
	if _, ok := m.loadedRelations["content_object_id"]; !ok {
		log.Panic().Msg("Tried to access ContentObject relation of ListOfSpeakers which was not loaded.")
	}

	return m.contentObject
}

func (m *ListOfSpeakers) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ListOfSpeakers which was not loaded.")
	}

	return *m.meeting
}

func (m *ListOfSpeakers) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of ListOfSpeakers which was not loaded.")
	}

	return m.projections
}

func (m *ListOfSpeakers) Speakers() []*Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of ListOfSpeakers which was not loaded.")
	}

	return m.speakers
}

func (m *ListOfSpeakers) StructureLevelListOfSpeakerss() []*StructureLevelListOfSpeakers {
	if _, ok := m.loadedRelations["structure_level_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevelListOfSpeakerss relation of ListOfSpeakers which was not loaded.")
	}

	return m.structureLevelListOfSpeakerss
}

func (m *ListOfSpeakers) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "content_object_id":
			panic("not implemented")
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "projection_ids":
			m.projections = content.([]*Projection)
		case "speaker_ids":
			m.speakers = content.([]*Speaker)
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

func (m *ListOfSpeakers) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "content_object_id":
		parts := strings.Split(m.ContentObjectID, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf("could not parse id field")
		}

		switch parts[0] {
		case "assignment":
			var entry Assignment
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "meeting_mediafile":
			var entry MeetingMediafile
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "motion":
			var entry Motion
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "motion_block":
			var entry MotionBlock
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "topic":
			var entry Topic
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()
		}

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
	case "speaker_ids":
		var entry Speaker
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.speakers = append(m.speakers, &entry)

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

func (m *ListOfSpeakers) Get(field string) interface{} {
	switch field {
	case "closed":
		return m.Closed
	case "content_object_id":
		return m.ContentObjectID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "moderator_notes":
		return m.ModeratorNotes
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "speaker_ids":
		return m.SpeakerIDs
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	}

	return nil
}

func (m *ListOfSpeakers) GetFqids(field string) []string {
	switch field {
	case "content_object_id":
		return []string{m.ContentObjectID}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "speaker_ids":
		r := make([]string, len(m.SpeakerIDs))
		for i, id := range m.SpeakerIDs {
			r[i] = "speaker/" + strconv.Itoa(id)
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

func (m *ListOfSpeakers) Update(data map[string]string) error {
	if val, ok := data["closed"]; ok {
		err := json.Unmarshal([]byte(val), &m.Closed)
		if err != nil {
			return err
		}
	}

	if val, ok := data["content_object_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ContentObjectID)
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

	if val, ok := data["moderator_notes"]; ok {
		err := json.Unmarshal([]byte(val), &m.ModeratorNotes)
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

	if val, ok := data["sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SequentialNumber)
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

func (m *ListOfSpeakers) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

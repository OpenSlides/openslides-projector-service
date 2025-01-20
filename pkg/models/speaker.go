package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Speaker struct {
	BeginTime                      *int    `json:"begin_time"`
	EndTime                        *int    `json:"end_time"`
	ID                             int     `json:"id"`
	ListOfSpeakersID               int     `json:"list_of_speakers_id"`
	MeetingID                      int     `json:"meeting_id"`
	MeetingUserID                  *int    `json:"meeting_user_id"`
	Note                           *string `json:"note"`
	PauseTime                      *int    `json:"pause_time"`
	PointOfOrder                   *bool   `json:"point_of_order"`
	PointOfOrderCategoryID         *int    `json:"point_of_order_category_id"`
	SpeechState                    *string `json:"speech_state"`
	StructureLevelListOfSpeakersID *int    `json:"structure_level_list_of_speakers_id"`
	TotalPause                     *int    `json:"total_pause"`
	UnpauseTime                    *int    `json:"unpause_time"`
	Weight                         *int    `json:"weight"`
	loadedRelations                map[string]struct{}
	listOfSpeakers                 *ListOfSpeakers
	meeting                        *Meeting
	meetingUser                    *MeetingUser
	pointOfOrderCategory           *PointOfOrderCategory
	structureLevelListOfSpeakers   *StructureLevelListOfSpeakers
}

func (m *Speaker) CollectionName() string {
	return "speaker"
}

func (m *Speaker) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of Speaker which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *Speaker) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Speaker which was not loaded.")
	}

	return *m.meeting
}

func (m *Speaker) MeetingUser() *MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of Speaker which was not loaded.")
	}

	return m.meetingUser
}

func (m *Speaker) PointOfOrderCategory() *PointOfOrderCategory {
	if _, ok := m.loadedRelations["point_of_order_category_id"]; !ok {
		log.Panic().Msg("Tried to access PointOfOrderCategory relation of Speaker which was not loaded.")
	}

	return m.pointOfOrderCategory
}

func (m *Speaker) StructureLevelListOfSpeakers() *StructureLevelListOfSpeakers {
	if _, ok := m.loadedRelations["structure_level_list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access StructureLevelListOfSpeakers relation of Speaker which was not loaded.")
	}

	return m.structureLevelListOfSpeakers
}

func (m *Speaker) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "list_of_speakers_id":
			m.listOfSpeakers = content.(*ListOfSpeakers)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_id":
			m.meetingUser = content.(*MeetingUser)
		case "point_of_order_category_id":
			m.pointOfOrderCategory = content.(*PointOfOrderCategory)
		case "structure_level_list_of_speakers_id":
			m.structureLevelListOfSpeakers = content.(*StructureLevelListOfSpeakers)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Speaker) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
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
	case "meeting_user_id":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUser = &entry

		result = entry.GetRelatedModelsAccessor()
	case "point_of_order_category_id":
		var entry PointOfOrderCategory
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pointOfOrderCategory = &entry

		result = entry.GetRelatedModelsAccessor()
	case "structure_level_list_of_speakers_id":
		var entry StructureLevelListOfSpeakers
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.structureLevelListOfSpeakers = &entry

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

func (m *Speaker) Get(field string) interface{} {
	switch field {
	case "begin_time":
		return m.BeginTime
	case "end_time":
		return m.EndTime
	case "id":
		return m.ID
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "note":
		return m.Note
	case "pause_time":
		return m.PauseTime
	case "point_of_order":
		return m.PointOfOrder
	case "point_of_order_category_id":
		return m.PointOfOrderCategoryID
	case "speech_state":
		return m.SpeechState
	case "structure_level_list_of_speakers_id":
		return m.StructureLevelListOfSpeakersID
	case "total_pause":
		return m.TotalPause
	case "unpause_time":
		return m.UnpauseTime
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *Speaker) GetFqids(field string) []string {
	switch field {
	case "list_of_speakers_id":
		return []string{"list_of_speakers/" + strconv.Itoa(m.ListOfSpeakersID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_id":
		if m.MeetingUserID != nil {
			return []string{"meeting_user/" + strconv.Itoa(*m.MeetingUserID)}
		}

	case "point_of_order_category_id":
		if m.PointOfOrderCategoryID != nil {
			return []string{"point_of_order_category/" + strconv.Itoa(*m.PointOfOrderCategoryID)}
		}

	case "structure_level_list_of_speakers_id":
		if m.StructureLevelListOfSpeakersID != nil {
			return []string{"structure_level_list_of_speakers/" + strconv.Itoa(*m.StructureLevelListOfSpeakersID)}
		}
	}
	return []string{}
}

func (m *Speaker) Update(data map[string]string) error {
	if val, ok := data["begin_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.BeginTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["end_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.EndTime)
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

	if val, ok := data["meeting_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["note"]; ok {
		err := json.Unmarshal([]byte(val), &m.Note)
		if err != nil {
			return err
		}
	}

	if val, ok := data["pause_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.PauseTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["point_of_order"]; ok {
		err := json.Unmarshal([]byte(val), &m.PointOfOrder)
		if err != nil {
			return err
		}
	}

	if val, ok := data["point_of_order_category_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PointOfOrderCategoryID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["speech_state"]; ok {
		err := json.Unmarshal([]byte(val), &m.SpeechState)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_list_of_speakers_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelListOfSpeakersID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["total_pause"]; ok {
		err := json.Unmarshal([]byte(val), &m.TotalPause)
		if err != nil {
			return err
		}
	}

	if val, ok := data["unpause_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.UnpauseTime)
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

func (m *Speaker) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

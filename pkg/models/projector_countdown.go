package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

type ProjectorCountdown struct {
	CountdownTime                          *float32 `json:"countdown_time"`
	DefaultTime                            *int     `json:"default_time"`
	Description                            *string  `json:"description"`
	ID                                     int      `json:"id"`
	MeetingID                              int      `json:"meeting_id"`
	ProjectionIDs                          []int    `json:"projection_ids"`
	Running                                *bool    `json:"running"`
	Title                                  string   `json:"title"`
	UsedAsListOfSpeakersCountdownMeetingID *int     `json:"used_as_list_of_speakers_countdown_meeting_id"`
	UsedAsPollCountdownMeetingID           *int     `json:"used_as_poll_countdown_meeting_id"`
	loadedRelations                        map[string]struct{}
	meeting                                *Meeting
	projections                            []*Projection
	usedAsListOfSpeakersCountdownMeeting   *Meeting
	usedAsPollCountdownMeeting             *Meeting
}

func (m *ProjectorCountdown) CollectionName() string {
	return "projector_countdown"
}

func (m *ProjectorCountdown) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ProjectorCountdown which was not loaded.")
	}

	return *m.meeting
}

func (m *ProjectorCountdown) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of ProjectorCountdown which was not loaded.")
	}

	return m.projections
}

func (m *ProjectorCountdown) UsedAsListOfSpeakersCountdownMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_list_of_speakers_countdown_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsListOfSpeakersCountdownMeeting relation of ProjectorCountdown which was not loaded.")
	}

	return m.usedAsListOfSpeakersCountdownMeeting
}

func (m *ProjectorCountdown) UsedAsPollCountdownMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_poll_countdown_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsPollCountdownMeeting relation of ProjectorCountdown which was not loaded.")
	}

	return m.usedAsPollCountdownMeeting
}

func (m *ProjectorCountdown) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "projection_ids":
		for _, r := range m.projections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "used_as_list_of_speakers_countdown_meeting_id":
		return m.usedAsListOfSpeakersCountdownMeeting.GetRelatedModelsAccessor()
	case "used_as_poll_countdown_meeting_id":
		return m.usedAsPollCountdownMeeting.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *ProjectorCountdown) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "projection_ids":
			m.projections = content.([]*Projection)
		case "used_as_list_of_speakers_countdown_meeting_id":
			m.usedAsListOfSpeakersCountdownMeeting = content.(*Meeting)
		case "used_as_poll_countdown_meeting_id":
			m.usedAsPollCountdownMeeting = content.(*Meeting)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *ProjectorCountdown) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "used_as_list_of_speakers_countdown_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsListOfSpeakersCountdownMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_poll_countdown_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsPollCountdownMeeting = &entry

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

func (m *ProjectorCountdown) Get(field string) interface{} {
	switch field {
	case "countdown_time":
		return m.CountdownTime
	case "default_time":
		return m.DefaultTime
	case "description":
		return m.Description
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "projection_ids":
		return m.ProjectionIDs
	case "running":
		return m.Running
	case "title":
		return m.Title
	case "used_as_list_of_speakers_countdown_meeting_id":
		return m.UsedAsListOfSpeakersCountdownMeetingID
	case "used_as_poll_countdown_meeting_id":
		return m.UsedAsPollCountdownMeetingID
	}

	return nil
}

func (m *ProjectorCountdown) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "used_as_list_of_speakers_countdown_meeting_id":
		if m.UsedAsListOfSpeakersCountdownMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsListOfSpeakersCountdownMeetingID)}
		}

	case "used_as_poll_countdown_meeting_id":
		if m.UsedAsPollCountdownMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsPollCountdownMeetingID)}
		}
	}
	return []string{}
}

func (m *ProjectorCountdown) Update(data map[string]string) error {
	if val, ok := data["countdown_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.CountdownTime)
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

	if val, ok := data["description"]; ok {
		err := json.Unmarshal([]byte(val), &m.Description)
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

	if val, ok := data["running"]; ok {
		err := json.Unmarshal([]byte(val), &m.Running)
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

	if val, ok := data["used_as_list_of_speakers_countdown_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsListOfSpeakersCountdownMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_poll_countdown_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsPollCountdownMeetingID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ProjectorCountdown) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

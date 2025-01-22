package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Topic struct {
	AgendaItemID                  int     `json:"agenda_item_id"`
	AttachmentMeetingMediafileIDs []int   `json:"attachment_meeting_mediafile_ids"`
	ID                            int     `json:"id"`
	ListOfSpeakersID              int     `json:"list_of_speakers_id"`
	MeetingID                     int     `json:"meeting_id"`
	PollIDs                       []int   `json:"poll_ids"`
	ProjectionIDs                 []int   `json:"projection_ids"`
	SequentialNumber              int     `json:"sequential_number"`
	Text                          *string `json:"text"`
	Title                         string  `json:"title"`
	loadedRelations               map[string]struct{}
	agendaItem                    *AgendaItem
	attachmentMeetingMediafiles   []*MeetingMediafile
	listOfSpeakers                *ListOfSpeakers
	meeting                       *Meeting
	polls                         []*Poll
	projections                   []*Projection
}

func (m *Topic) CollectionName() string {
	return "topic"
}

func (m *Topic) AgendaItem() AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_id"]; !ok {
		log.Panic().Msg("Tried to access AgendaItem relation of Topic which was not loaded.")
	}

	return *m.agendaItem
}

func (m *Topic) AttachmentMeetingMediafiles() []*MeetingMediafile {
	if _, ok := m.loadedRelations["attachment_meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access AttachmentMeetingMediafiles relation of Topic which was not loaded.")
	}

	return m.attachmentMeetingMediafiles
}

func (m *Topic) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of Topic which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *Topic) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Topic which was not loaded.")
	}

	return *m.meeting
}

func (m *Topic) Polls() []*Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Topic which was not loaded.")
	}

	return m.polls
}

func (m *Topic) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of Topic which was not loaded.")
	}

	return m.projections
}

func (m *Topic) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "agenda_item_id":
		return m.agendaItem.GetRelatedModelsAccessor()
	case "attachment_meeting_mediafile_ids":
		for _, r := range m.attachmentMeetingMediafiles {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "list_of_speakers_id":
		return m.listOfSpeakers.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "poll_ids":
		for _, r := range m.polls {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "projection_ids":
		for _, r := range m.projections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Topic) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "agenda_item_id":
			m.agendaItem = content.(*AgendaItem)
		case "attachment_meeting_mediafile_ids":
			m.attachmentMeetingMediafiles = content.([]*MeetingMediafile)
		case "list_of_speakers_id":
			m.listOfSpeakers = content.(*ListOfSpeakers)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "poll_ids":
			m.polls = content.([]*Poll)
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

func (m *Topic) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "agenda_item_id":
		var entry AgendaItem
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.agendaItem = &entry

		result = entry.GetRelatedModelsAccessor()
	case "attachment_meeting_mediafile_ids":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.attachmentMeetingMediafiles = append(m.attachmentMeetingMediafiles, &entry)

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
	case "poll_ids":
		var entry Poll
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.polls = append(m.polls, &entry)

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

func (m *Topic) Get(field string) interface{} {
	switch field {
	case "agenda_item_id":
		return m.AgendaItemID
	case "attachment_meeting_mediafile_ids":
		return m.AttachmentMeetingMediafileIDs
	case "id":
		return m.ID
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "poll_ids":
		return m.PollIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "text":
		return m.Text
	case "title":
		return m.Title
	}

	return nil
}

func (m *Topic) GetFqids(field string) []string {
	switch field {
	case "agenda_item_id":
		return []string{"agenda_item/" + strconv.Itoa(m.AgendaItemID)}

	case "attachment_meeting_mediafile_ids":
		r := make([]string, len(m.AttachmentMeetingMediafileIDs))
		for i, id := range m.AttachmentMeetingMediafileIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "list_of_speakers_id":
		return []string{"list_of_speakers/" + strconv.Itoa(m.ListOfSpeakersID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "poll_ids":
		r := make([]string, len(m.PollIDs))
		for i, id := range m.PollIDs {
			r[i] = "poll/" + strconv.Itoa(id)
		}
		return r

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Topic) Update(data map[string]string) error {
	if val, ok := data["agenda_item_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["attachment_meeting_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AttachmentMeetingMediafileIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["attachment_meeting_mediafile_ids"]; ok {
			m.attachmentMeetingMediafiles = slices.DeleteFunc(m.attachmentMeetingMediafiles, func(r *MeetingMediafile) bool {
				return !slices.Contains(m.AttachmentMeetingMediafileIDs, r.ID)
			})
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

	if val, ok := data["poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["poll_ids"]; ok {
			m.polls = slices.DeleteFunc(m.polls, func(r *Poll) bool {
				return !slices.Contains(m.PollIDs, r.ID)
			})
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

	if val, ok := data["text"]; ok {
		err := json.Unmarshal([]byte(val), &m.Text)
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

	return nil
}

func (m *Topic) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

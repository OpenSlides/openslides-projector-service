package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type ActionWorker struct {
	Created   int             `json:"created"`
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Result    json.RawMessage `json:"result"`
	State     string          `json:"state"`
	Timestamp int             `json:"timestamp"`
	UserID    int             `json:"user_id"`
}

func (m *ActionWorker) CollectionName() string {
	return "action_worker"
}

func (m *ActionWorker) GetRelated(field string, id int) *RelatedModelsAccessor {
	return nil
}

func (m *ActionWorker) SetRelated(field string, content interface{}) {}

func (m *ActionWorker) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	return nil, nil
}

func (m *ActionWorker) Get(field string) interface{} {
	switch field {
	case "created":
		return m.Created
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "result":
		return m.Result
	case "state":
		return m.State
	case "timestamp":
		return m.Timestamp
	case "user_id":
		return m.UserID
	}

	return nil
}

func (m *ActionWorker) GetFqids(field string) []string {
	return []string{}
}

func (m *ActionWorker) Update(data map[string]string) error {
	if val, ok := data["created"]; ok {
		err := json.Unmarshal([]byte(val), &m.Created)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["result"]; ok {
		err := json.Unmarshal([]byte(val), &m.Result)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state"]; ok {
		err := json.Unmarshal([]byte(val), &m.State)
		if err != nil {
			return err
		}
	}

	if val, ok := data["timestamp"]; ok {
		err := json.Unmarshal([]byte(val), &m.Timestamp)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ActionWorker) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type AgendaItem struct {
	ChildIDs        []int   `json:"child_ids"`
	Closed          *bool   `json:"closed"`
	Comment         *string `json:"comment"`
	ContentObjectID string  `json:"content_object_id"`
	Duration        *int    `json:"duration"`
	ID              int     `json:"id"`
	IsHidden        *bool   `json:"is_hidden"`
	IsInternal      *bool   `json:"is_internal"`
	ItemNumber      *string `json:"item_number"`
	Level           *int    `json:"level"`
	MeetingID       int     `json:"meeting_id"`
	ParentID        *int    `json:"parent_id"`
	ProjectionIDs   []int   `json:"projection_ids"`
	TagIDs          []int   `json:"tag_ids"`
	Type            *string `json:"type"`
	Weight          *int    `json:"weight"`
	loadedRelations map[string]struct{}
	childs          []*AgendaItem
	contentObject   IBaseModel
	meeting         *Meeting
	parent          *AgendaItem
	projections     []*Projection
	tags            []*Tag
}

func (m *AgendaItem) CollectionName() string {
	return "agenda_item"
}

func (m *AgendaItem) Childs() []*AgendaItem {
	if _, ok := m.loadedRelations["child_ids"]; !ok {
		log.Panic().Msg("Tried to access Childs relation of AgendaItem which was not loaded.")
	}

	return m.childs
}

func (m *AgendaItem) ContentObject() IBaseModel {
	if _, ok := m.loadedRelations["content_object_id"]; !ok {
		log.Panic().Msg("Tried to access ContentObject relation of AgendaItem which was not loaded.")
	}

	return m.contentObject
}

func (m *AgendaItem) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of AgendaItem which was not loaded.")
	}

	return *m.meeting
}

func (m *AgendaItem) Parent() *AgendaItem {
	if _, ok := m.loadedRelations["parent_id"]; !ok {
		log.Panic().Msg("Tried to access Parent relation of AgendaItem which was not loaded.")
	}

	return m.parent
}

func (m *AgendaItem) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of AgendaItem which was not loaded.")
	}

	return m.projections
}

func (m *AgendaItem) Tags() []*Tag {
	if _, ok := m.loadedRelations["tag_ids"]; !ok {
		log.Panic().Msg("Tried to access Tags relation of AgendaItem which was not loaded.")
	}

	return m.tags
}

func (m *AgendaItem) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "child_ids":
		for _, r := range m.childs {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "content_object_id":
		return m.contentObject.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "parent_id":
		return m.parent.GetRelatedModelsAccessor()
	case "projection_ids":
		for _, r := range m.projections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "tag_ids":
		for _, r := range m.tags {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *AgendaItem) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "child_ids":
			m.childs = content.([]*AgendaItem)
		case "content_object_id":
			panic("not implemented")
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "parent_id":
			m.parent = content.(*AgendaItem)
		case "projection_ids":
			m.projections = content.([]*Projection)
		case "tag_ids":
			m.tags = content.([]*Tag)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *AgendaItem) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "child_ids":
		var entry AgendaItem
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.childs = append(m.childs, &entry)

		result = entry.GetRelatedModelsAccessor()
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
	case "parent_id":
		var entry AgendaItem
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.parent = &entry

		result = entry.GetRelatedModelsAccessor()
	case "projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projections = append(m.projections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "tag_ids":
		var entry Tag
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.tags = append(m.tags, &entry)

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

func (m *AgendaItem) Get(field string) interface{} {
	switch field {
	case "child_ids":
		return m.ChildIDs
	case "closed":
		return m.Closed
	case "comment":
		return m.Comment
	case "content_object_id":
		return m.ContentObjectID
	case "duration":
		return m.Duration
	case "id":
		return m.ID
	case "is_hidden":
		return m.IsHidden
	case "is_internal":
		return m.IsInternal
	case "item_number":
		return m.ItemNumber
	case "level":
		return m.Level
	case "meeting_id":
		return m.MeetingID
	case "parent_id":
		return m.ParentID
	case "projection_ids":
		return m.ProjectionIDs
	case "tag_ids":
		return m.TagIDs
	case "type":
		return m.Type
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *AgendaItem) GetFqids(field string) []string {
	switch field {
	case "child_ids":
		r := make([]string, len(m.ChildIDs))
		for i, id := range m.ChildIDs {
			r[i] = "agenda_item/" + strconv.Itoa(id)
		}
		return r

	case "content_object_id":
		return []string{m.ContentObjectID}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "parent_id":
		if m.ParentID != nil {
			return []string{"agenda_item/" + strconv.Itoa(*m.ParentID)}
		}

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "tag_ids":
		r := make([]string, len(m.TagIDs))
		for i, id := range m.TagIDs {
			r[i] = "tag/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *AgendaItem) Update(data map[string]string) error {
	if val, ok := data["child_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChildIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["child_ids"]; ok {
			m.childs = slices.DeleteFunc(m.childs, func(r *AgendaItem) bool {
				return !slices.Contains(m.ChildIDs, r.ID)
			})
		}
	}

	if val, ok := data["closed"]; ok {
		err := json.Unmarshal([]byte(val), &m.Closed)
		if err != nil {
			return err
		}
	}

	if val, ok := data["comment"]; ok {
		err := json.Unmarshal([]byte(val), &m.Comment)
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

	if val, ok := data["duration"]; ok {
		err := json.Unmarshal([]byte(val), &m.Duration)
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

	if val, ok := data["is_hidden"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsHidden)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsInternal)
		if err != nil {
			return err
		}
	}

	if val, ok := data["item_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.ItemNumber)
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

	if val, ok := data["parent_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ParentID)
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

	if val, ok := data["tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TagIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["tag_ids"]; ok {
			m.tags = slices.DeleteFunc(m.tags, func(r *Tag) bool {
				return !slices.Contains(m.TagIDs, r.ID)
			})
		}
	}

	if val, ok := data["type"]; ok {
		err := json.Unmarshal([]byte(val), &m.Type)
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

func (m *AgendaItem) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Assignment struct {
	AgendaItemID                  *int    `json:"agenda_item_id"`
	AttachmentMeetingMediafileIDs []int   `json:"attachment_meeting_mediafile_ids"`
	CandidateIDs                  []int   `json:"candidate_ids"`
	DefaultPollDescription        *string `json:"default_poll_description"`
	Description                   *string `json:"description"`
	ID                            int     `json:"id"`
	ListOfSpeakersID              int     `json:"list_of_speakers_id"`
	MeetingID                     int     `json:"meeting_id"`
	NumberPollCandidates          *bool   `json:"number_poll_candidates"`
	OpenPosts                     *int    `json:"open_posts"`
	Phase                         *string `json:"phase"`
	PollIDs                       []int   `json:"poll_ids"`
	ProjectionIDs                 []int   `json:"projection_ids"`
	SequentialNumber              int     `json:"sequential_number"`
	TagIDs                        []int   `json:"tag_ids"`
	Title                         string  `json:"title"`
	loadedRelations               map[string]struct{}
	agendaItem                    *AgendaItem
	attachmentMeetingMediafiles   []*MeetingMediafile
	candidates                    []*AssignmentCandidate
	listOfSpeakers                *ListOfSpeakers
	meeting                       *Meeting
	polls                         []*Poll
	projections                   []*Projection
	tags                          []*Tag
}

func (m *Assignment) CollectionName() string {
	return "assignment"
}

func (m *Assignment) AgendaItem() *AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_id"]; !ok {
		log.Panic().Msg("Tried to access AgendaItem relation of Assignment which was not loaded.")
	}

	return m.agendaItem
}

func (m *Assignment) AttachmentMeetingMediafiles() []*MeetingMediafile {
	if _, ok := m.loadedRelations["attachment_meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access AttachmentMeetingMediafiles relation of Assignment which was not loaded.")
	}

	return m.attachmentMeetingMediafiles
}

func (m *Assignment) Candidates() []*AssignmentCandidate {
	if _, ok := m.loadedRelations["candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access Candidates relation of Assignment which was not loaded.")
	}

	return m.candidates
}

func (m *Assignment) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of Assignment which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *Assignment) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Assignment which was not loaded.")
	}

	return *m.meeting
}

func (m *Assignment) Polls() []*Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Assignment which was not loaded.")
	}

	return m.polls
}

func (m *Assignment) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of Assignment which was not loaded.")
	}

	return m.projections
}

func (m *Assignment) Tags() []*Tag {
	if _, ok := m.loadedRelations["tag_ids"]; !ok {
		log.Panic().Msg("Tried to access Tags relation of Assignment which was not loaded.")
	}

	return m.tags
}

func (m *Assignment) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "agenda_item_id":
		return m.agendaItem.GetRelatedModelsAccessor()
	case "attachment_meeting_mediafile_ids":
		for _, r := range m.attachmentMeetingMediafiles {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "candidate_ids":
		for _, r := range m.candidates {
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
	case "tag_ids":
		for _, r := range m.tags {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Assignment) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "agenda_item_id":
			m.agendaItem = content.(*AgendaItem)
		case "attachment_meeting_mediafile_ids":
			m.attachmentMeetingMediafiles = content.([]*MeetingMediafile)
		case "candidate_ids":
			m.candidates = content.([]*AssignmentCandidate)
		case "list_of_speakers_id":
			m.listOfSpeakers = content.(*ListOfSpeakers)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "poll_ids":
			m.polls = content.([]*Poll)
		case "projection_ids":
			m.projections = content.([]*Projection)
		case "tag_ids":
			m.tags = content.([]*Tag)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Assignment) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "candidate_ids":
		var entry AssignmentCandidate
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.candidates = append(m.candidates, &entry)

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
	case "tag_ids":
		var entry Tag
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.tags = append(m.tags, &entry)

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

func (m *Assignment) Get(field string) interface{} {
	switch field {
	case "agenda_item_id":
		return m.AgendaItemID
	case "attachment_meeting_mediafile_ids":
		return m.AttachmentMeetingMediafileIDs
	case "candidate_ids":
		return m.CandidateIDs
	case "default_poll_description":
		return m.DefaultPollDescription
	case "description":
		return m.Description
	case "id":
		return m.ID
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "number_poll_candidates":
		return m.NumberPollCandidates
	case "open_posts":
		return m.OpenPosts
	case "phase":
		return m.Phase
	case "poll_ids":
		return m.PollIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "tag_ids":
		return m.TagIDs
	case "title":
		return m.Title
	}

	return nil
}

func (m *Assignment) GetFqids(field string) []string {
	switch field {
	case "agenda_item_id":
		if m.AgendaItemID != nil {
			return []string{"agenda_item/" + strconv.Itoa(*m.AgendaItemID)}
		}

	case "attachment_meeting_mediafile_ids":
		r := make([]string, len(m.AttachmentMeetingMediafileIDs))
		for i, id := range m.AttachmentMeetingMediafileIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "candidate_ids":
		r := make([]string, len(m.CandidateIDs))
		for i, id := range m.CandidateIDs {
			r[i] = "assignment_candidate/" + strconv.Itoa(id)
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

	case "tag_ids":
		r := make([]string, len(m.TagIDs))
		for i, id := range m.TagIDs {
			r[i] = "tag/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Assignment) Update(data map[string]string) error {
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

	if val, ok := data["candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CandidateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["candidate_ids"]; ok {
			m.candidates = slices.DeleteFunc(m.candidates, func(r *AssignmentCandidate) bool {
				return !slices.Contains(m.CandidateIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_poll_description"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultPollDescription)
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

	if val, ok := data["number_poll_candidates"]; ok {
		err := json.Unmarshal([]byte(val), &m.NumberPollCandidates)
		if err != nil {
			return err
		}
	}

	if val, ok := data["open_posts"]; ok {
		err := json.Unmarshal([]byte(val), &m.OpenPosts)
		if err != nil {
			return err
		}
	}

	if val, ok := data["phase"]; ok {
		err := json.Unmarshal([]byte(val), &m.Phase)
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

	if val, ok := data["tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TagIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["tag_ids"]; ok {
			m.tags = slices.DeleteFunc(m.tags, func(r *Tag) bool {
				return !slices.Contains(m.TagIDs, r.ID)
			})
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

func (m *Assignment) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type AssignmentCandidate struct {
	AssignmentID    int  `json:"assignment_id"`
	ID              int  `json:"id"`
	MeetingID       int  `json:"meeting_id"`
	MeetingUserID   *int `json:"meeting_user_id"`
	Weight          *int `json:"weight"`
	loadedRelations map[string]struct{}
	assignment      *Assignment
	meeting         *Meeting
	meetingUser     *MeetingUser
}

func (m *AssignmentCandidate) CollectionName() string {
	return "assignment_candidate"
}

func (m *AssignmentCandidate) Assignment() Assignment {
	if _, ok := m.loadedRelations["assignment_id"]; !ok {
		log.Panic().Msg("Tried to access Assignment relation of AssignmentCandidate which was not loaded.")
	}

	return *m.assignment
}

func (m *AssignmentCandidate) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of AssignmentCandidate which was not loaded.")
	}

	return *m.meeting
}

func (m *AssignmentCandidate) MeetingUser() *MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of AssignmentCandidate which was not loaded.")
	}

	return m.meetingUser
}

func (m *AssignmentCandidate) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "assignment_id":
		return m.assignment.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "meeting_user_id":
		return m.meetingUser.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *AssignmentCandidate) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "assignment_id":
			m.assignment = content.(*Assignment)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_id":
			m.meetingUser = content.(*MeetingUser)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *AssignmentCandidate) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "assignment_id":
		var entry Assignment
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.assignment = &entry

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
	default:
		return nil, fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return result, nil
}

func (m *AssignmentCandidate) Get(field string) interface{} {
	switch field {
	case "assignment_id":
		return m.AssignmentID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *AssignmentCandidate) GetFqids(field string) []string {
	switch field {
	case "assignment_id":
		return []string{"assignment/" + strconv.Itoa(m.AssignmentID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_id":
		if m.MeetingUserID != nil {
			return []string{"meeting_user/" + strconv.Itoa(*m.MeetingUserID)}
		}
	}
	return []string{}
}

func (m *AssignmentCandidate) Update(data map[string]string) error {
	if val, ok := data["assignment_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentID)
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

	if val, ok := data["meeting_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserID)
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

func (m *AssignmentCandidate) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type ChatGroup struct {
	ChatMessageIDs  []int  `json:"chat_message_ids"`
	ID              int    `json:"id"`
	MeetingID       int    `json:"meeting_id"`
	Name            string `json:"name"`
	ReadGroupIDs    []int  `json:"read_group_ids"`
	Weight          *int   `json:"weight"`
	WriteGroupIDs   []int  `json:"write_group_ids"`
	loadedRelations map[string]struct{}
	chatMessages    []*ChatMessage
	meeting         *Meeting
	readGroups      []*Group
	writeGroups     []*Group
}

func (m *ChatGroup) CollectionName() string {
	return "chat_group"
}

func (m *ChatGroup) ChatMessages() []*ChatMessage {
	if _, ok := m.loadedRelations["chat_message_ids"]; !ok {
		log.Panic().Msg("Tried to access ChatMessages relation of ChatGroup which was not loaded.")
	}

	return m.chatMessages
}

func (m *ChatGroup) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ChatGroup which was not loaded.")
	}

	return *m.meeting
}

func (m *ChatGroup) ReadGroups() []*Group {
	if _, ok := m.loadedRelations["read_group_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadGroups relation of ChatGroup which was not loaded.")
	}

	return m.readGroups
}

func (m *ChatGroup) WriteGroups() []*Group {
	if _, ok := m.loadedRelations["write_group_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteGroups relation of ChatGroup which was not loaded.")
	}

	return m.writeGroups
}

func (m *ChatGroup) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "chat_message_ids":
		for _, r := range m.chatMessages {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "read_group_ids":
		for _, r := range m.readGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "write_group_ids":
		for _, r := range m.writeGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *ChatGroup) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "chat_message_ids":
			m.chatMessages = content.([]*ChatMessage)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "read_group_ids":
			m.readGroups = content.([]*Group)
		case "write_group_ids":
			m.writeGroups = content.([]*Group)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *ChatGroup) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "chat_message_ids":
		var entry ChatMessage
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.chatMessages = append(m.chatMessages, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "read_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.readGroups = append(m.readGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "write_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.writeGroups = append(m.writeGroups, &entry)

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

func (m *ChatGroup) Get(field string) interface{} {
	switch field {
	case "chat_message_ids":
		return m.ChatMessageIDs
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "read_group_ids":
		return m.ReadGroupIDs
	case "weight":
		return m.Weight
	case "write_group_ids":
		return m.WriteGroupIDs
	}

	return nil
}

func (m *ChatGroup) GetFqids(field string) []string {
	switch field {
	case "chat_message_ids":
		r := make([]string, len(m.ChatMessageIDs))
		for i, id := range m.ChatMessageIDs {
			r[i] = "chat_message/" + strconv.Itoa(id)
		}
		return r

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "read_group_ids":
		r := make([]string, len(m.ReadGroupIDs))
		for i, id := range m.ReadGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "write_group_ids":
		r := make([]string, len(m.WriteGroupIDs))
		for i, id := range m.WriteGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *ChatGroup) Update(data map[string]string) error {
	if val, ok := data["chat_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatMessageIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["chat_message_ids"]; ok {
			m.chatMessages = slices.DeleteFunc(m.chatMessages, func(r *ChatMessage) bool {
				return !slices.Contains(m.ChatMessageIDs, r.ID)
			})
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["read_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReadGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["read_group_ids"]; ok {
			m.readGroups = slices.DeleteFunc(m.readGroups, func(r *Group) bool {
				return !slices.Contains(m.ReadGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["write_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WriteGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["write_group_ids"]; ok {
			m.writeGroups = slices.DeleteFunc(m.writeGroups, func(r *Group) bool {
				return !slices.Contains(m.WriteGroupIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *ChatGroup) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type ChatMessage struct {
	ChatGroupID     int    `json:"chat_group_id"`
	Content         string `json:"content"`
	Created         int    `json:"created"`
	ID              int    `json:"id"`
	MeetingID       int    `json:"meeting_id"`
	MeetingUserID   *int   `json:"meeting_user_id"`
	loadedRelations map[string]struct{}
	chatGroup       *ChatGroup
	meeting         *Meeting
	meetingUser     *MeetingUser
}

func (m *ChatMessage) CollectionName() string {
	return "chat_message"
}

func (m *ChatMessage) ChatGroup() ChatGroup {
	if _, ok := m.loadedRelations["chat_group_id"]; !ok {
		log.Panic().Msg("Tried to access ChatGroup relation of ChatMessage which was not loaded.")
	}

	return *m.chatGroup
}

func (m *ChatMessage) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of ChatMessage which was not loaded.")
	}

	return *m.meeting
}

func (m *ChatMessage) MeetingUser() *MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of ChatMessage which was not loaded.")
	}

	return m.meetingUser
}

func (m *ChatMessage) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "chat_group_id":
		return m.chatGroup.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "meeting_user_id":
		return m.meetingUser.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *ChatMessage) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "chat_group_id":
			m.chatGroup = content.(*ChatGroup)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_id":
			m.meetingUser = content.(*MeetingUser)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *ChatMessage) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "chat_group_id":
		var entry ChatGroup
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.chatGroup = &entry

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
	default:
		return nil, fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return result, nil
}

func (m *ChatMessage) Get(field string) interface{} {
	switch field {
	case "chat_group_id":
		return m.ChatGroupID
	case "content":
		return m.Content
	case "created":
		return m.Created
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	}

	return nil
}

func (m *ChatMessage) GetFqids(field string) []string {
	switch field {
	case "chat_group_id":
		return []string{"chat_group/" + strconv.Itoa(m.ChatGroupID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_id":
		if m.MeetingUserID != nil {
			return []string{"meeting_user/" + strconv.Itoa(*m.MeetingUserID)}
		}
	}
	return []string{}
}

func (m *ChatMessage) Update(data map[string]string) error {
	if val, ok := data["chat_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["content"]; ok {
		err := json.Unmarshal([]byte(val), &m.Content)
		if err != nil {
			return err
		}
	}

	if val, ok := data["created"]; ok {
		err := json.Unmarshal([]byte(val), &m.Created)
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

	if val, ok := data["meeting_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ChatMessage) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Committee struct {
	DefaultMeetingID                   *int    `json:"default_meeting_id"`
	Description                        *string `json:"description"`
	ExternalID                         *string `json:"external_id"`
	ForwardToCommitteeIDs              []int   `json:"forward_to_committee_ids"`
	ForwardingUserID                   *int    `json:"forwarding_user_id"`
	ID                                 int     `json:"id"`
	ManagerIDs                         []int   `json:"manager_ids"`
	MeetingIDs                         []int   `json:"meeting_ids"`
	Name                               string  `json:"name"`
	OrganizationID                     int     `json:"organization_id"`
	OrganizationTagIDs                 []int   `json:"organization_tag_ids"`
	ReceiveForwardingsFromCommitteeIDs []int   `json:"receive_forwardings_from_committee_ids"`
	UserIDs                            []int   `json:"user_ids"`
	loadedRelations                    map[string]struct{}
	defaultMeeting                     *Meeting
	forwardToCommittees                []*Committee
	forwardingUser                     *User
	managers                           []*User
	meetings                           []*Meeting
	organization                       *Organization
	organizationTags                   []*OrganizationTag
	receiveForwardingsFromCommittees   []*Committee
	users                              []*User
}

func (m *Committee) CollectionName() string {
	return "committee"
}

func (m *Committee) DefaultMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultMeeting relation of Committee which was not loaded.")
	}

	return m.defaultMeeting
}

func (m *Committee) ForwardToCommittees() []*Committee {
	if _, ok := m.loadedRelations["forward_to_committee_ids"]; !ok {
		log.Panic().Msg("Tried to access ForwardToCommittees relation of Committee which was not loaded.")
	}

	return m.forwardToCommittees
}

func (m *Committee) ForwardingUser() *User {
	if _, ok := m.loadedRelations["forwarding_user_id"]; !ok {
		log.Panic().Msg("Tried to access ForwardingUser relation of Committee which was not loaded.")
	}

	return m.forwardingUser
}

func (m *Committee) Managers() []*User {
	if _, ok := m.loadedRelations["manager_ids"]; !ok {
		log.Panic().Msg("Tried to access Managers relation of Committee which was not loaded.")
	}

	return m.managers
}

func (m *Committee) Meetings() []*Meeting {
	if _, ok := m.loadedRelations["meeting_ids"]; !ok {
		log.Panic().Msg("Tried to access Meetings relation of Committee which was not loaded.")
	}

	return m.meetings
}

func (m *Committee) Organization() Organization {
	if _, ok := m.loadedRelations["organization_id"]; !ok {
		log.Panic().Msg("Tried to access Organization relation of Committee which was not loaded.")
	}

	return *m.organization
}

func (m *Committee) OrganizationTags() []*OrganizationTag {
	if _, ok := m.loadedRelations["organization_tag_ids"]; !ok {
		log.Panic().Msg("Tried to access OrganizationTags relation of Committee which was not loaded.")
	}

	return m.organizationTags
}

func (m *Committee) ReceiveForwardingsFromCommittees() []*Committee {
	if _, ok := m.loadedRelations["receive_forwardings_from_committee_ids"]; !ok {
		log.Panic().Msg("Tried to access ReceiveForwardingsFromCommittees relation of Committee which was not loaded.")
	}

	return m.receiveForwardingsFromCommittees
}

func (m *Committee) Users() []*User {
	if _, ok := m.loadedRelations["user_ids"]; !ok {
		log.Panic().Msg("Tried to access Users relation of Committee which was not loaded.")
	}

	return m.users
}

func (m *Committee) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "default_meeting_id":
		return m.defaultMeeting.GetRelatedModelsAccessor()
	case "forward_to_committee_ids":
		for _, r := range m.forwardToCommittees {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "forwarding_user_id":
		return m.forwardingUser.GetRelatedModelsAccessor()
	case "manager_ids":
		for _, r := range m.managers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_ids":
		for _, r := range m.meetings {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "organization_id":
		return m.organization.GetRelatedModelsAccessor()
	case "organization_tag_ids":
		for _, r := range m.organizationTags {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "receive_forwardings_from_committee_ids":
		for _, r := range m.receiveForwardingsFromCommittees {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "user_ids":
		for _, r := range m.users {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Committee) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "default_meeting_id":
			m.defaultMeeting = content.(*Meeting)
		case "forward_to_committee_ids":
			m.forwardToCommittees = content.([]*Committee)
		case "forwarding_user_id":
			m.forwardingUser = content.(*User)
		case "manager_ids":
			m.managers = content.([]*User)
		case "meeting_ids":
			m.meetings = content.([]*Meeting)
		case "organization_id":
			m.organization = content.(*Organization)
		case "organization_tag_ids":
			m.organizationTags = content.([]*OrganizationTag)
		case "receive_forwardings_from_committee_ids":
			m.receiveForwardingsFromCommittees = content.([]*Committee)
		case "user_ids":
			m.users = content.([]*User)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Committee) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "default_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "forward_to_committee_ids":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.forwardToCommittees = append(m.forwardToCommittees, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "forwarding_user_id":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.forwardingUser = &entry

		result = entry.GetRelatedModelsAccessor()
	case "manager_ids":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.managers = append(m.managers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_ids":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetings = append(m.meetings, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.organization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "organization_tag_ids":
		var entry OrganizationTag
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.organizationTags = append(m.organizationTags, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "receive_forwardings_from_committee_ids":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.receiveForwardingsFromCommittees = append(m.receiveForwardingsFromCommittees, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "user_ids":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.users = append(m.users, &entry)

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

func (m *Committee) Get(field string) interface{} {
	switch field {
	case "default_meeting_id":
		return m.DefaultMeetingID
	case "description":
		return m.Description
	case "external_id":
		return m.ExternalID
	case "forward_to_committee_ids":
		return m.ForwardToCommitteeIDs
	case "forwarding_user_id":
		return m.ForwardingUserID
	case "id":
		return m.ID
	case "manager_ids":
		return m.ManagerIDs
	case "meeting_ids":
		return m.MeetingIDs
	case "name":
		return m.Name
	case "organization_id":
		return m.OrganizationID
	case "organization_tag_ids":
		return m.OrganizationTagIDs
	case "receive_forwardings_from_committee_ids":
		return m.ReceiveForwardingsFromCommitteeIDs
	case "user_ids":
		return m.UserIDs
	}

	return nil
}

func (m *Committee) GetFqids(field string) []string {
	switch field {
	case "default_meeting_id":
		if m.DefaultMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultMeetingID)}
		}

	case "forward_to_committee_ids":
		r := make([]string, len(m.ForwardToCommitteeIDs))
		for i, id := range m.ForwardToCommitteeIDs {
			r[i] = "committee/" + strconv.Itoa(id)
		}
		return r

	case "forwarding_user_id":
		if m.ForwardingUserID != nil {
			return []string{"user/" + strconv.Itoa(*m.ForwardingUserID)}
		}

	case "manager_ids":
		r := make([]string, len(m.ManagerIDs))
		for i, id := range m.ManagerIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r

	case "meeting_ids":
		r := make([]string, len(m.MeetingIDs))
		for i, id := range m.MeetingIDs {
			r[i] = "meeting/" + strconv.Itoa(id)
		}
		return r

	case "organization_id":
		return []string{"organization/" + strconv.Itoa(m.OrganizationID)}

	case "organization_tag_ids":
		r := make([]string, len(m.OrganizationTagIDs))
		for i, id := range m.OrganizationTagIDs {
			r[i] = "organization_tag/" + strconv.Itoa(id)
		}
		return r

	case "receive_forwardings_from_committee_ids":
		r := make([]string, len(m.ReceiveForwardingsFromCommitteeIDs))
		for i, id := range m.ReceiveForwardingsFromCommitteeIDs {
			r[i] = "committee/" + strconv.Itoa(id)
		}
		return r

	case "user_ids":
		r := make([]string, len(m.UserIDs))
		for i, id := range m.UserIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Committee) Update(data map[string]string) error {
	if val, ok := data["default_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultMeetingID)
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

	if val, ok := data["external_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExternalID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["forward_to_committee_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ForwardToCommitteeIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["forward_to_committee_ids"]; ok {
			m.forwardToCommittees = slices.DeleteFunc(m.forwardToCommittees, func(r *Committee) bool {
				return !slices.Contains(m.ForwardToCommitteeIDs, r.ID)
			})
		}
	}

	if val, ok := data["forwarding_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ForwardingUserID)
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

	if val, ok := data["manager_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ManagerIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["manager_ids"]; ok {
			m.managers = slices.DeleteFunc(m.managers, func(r *User) bool {
				return !slices.Contains(m.ManagerIDs, r.ID)
			})
		}
	}

	if val, ok := data["meeting_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["meeting_ids"]; ok {
			m.meetings = slices.DeleteFunc(m.meetings, func(r *Meeting) bool {
				return !slices.Contains(m.MeetingIDs, r.ID)
			})
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["organization_tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationTagIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["organization_tag_ids"]; ok {
			m.organizationTags = slices.DeleteFunc(m.organizationTags, func(r *OrganizationTag) bool {
				return !slices.Contains(m.OrganizationTagIDs, r.ID)
			})
		}
	}

	if val, ok := data["receive_forwardings_from_committee_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReceiveForwardingsFromCommitteeIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["receive_forwardings_from_committee_ids"]; ok {
			m.receiveForwardingsFromCommittees = slices.DeleteFunc(m.receiveForwardingsFromCommittees, func(r *Committee) bool {
				return !slices.Contains(m.ReceiveForwardingsFromCommitteeIDs, r.ID)
			})
		}
	}

	if val, ok := data["user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["user_ids"]; ok {
			m.users = slices.DeleteFunc(m.users, func(r *User) bool {
				return !slices.Contains(m.UserIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *Committee) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Gender struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	OrganizationID  int    `json:"organization_id"`
	UserIDs         []int  `json:"user_ids"`
	loadedRelations map[string]struct{}
	organization    *Organization
	users           []*User
}

func (m *Gender) CollectionName() string {
	return "gender"
}

func (m *Gender) Organization() Organization {
	if _, ok := m.loadedRelations["organization_id"]; !ok {
		log.Panic().Msg("Tried to access Organization relation of Gender which was not loaded.")
	}

	return *m.organization
}

func (m *Gender) Users() []*User {
	if _, ok := m.loadedRelations["user_ids"]; !ok {
		log.Panic().Msg("Tried to access Users relation of Gender which was not loaded.")
	}

	return m.users
}

func (m *Gender) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "organization_id":
		return m.organization.GetRelatedModelsAccessor()
	case "user_ids":
		for _, r := range m.users {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Gender) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "organization_id":
			m.organization = content.(*Organization)
		case "user_ids":
			m.users = content.([]*User)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Gender) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.organization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "user_ids":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.users = append(m.users, &entry)

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

func (m *Gender) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "organization_id":
		return m.OrganizationID
	case "user_ids":
		return m.UserIDs
	}

	return nil
}

func (m *Gender) GetFqids(field string) []string {
	switch field {
	case "organization_id":
		return []string{"organization/" + strconv.Itoa(m.OrganizationID)}

	case "user_ids":
		r := make([]string, len(m.UserIDs))
		for i, id := range m.UserIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Gender) Update(data map[string]string) error {
	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
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

	if val, ok := data["organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["user_ids"]; ok {
			m.users = slices.DeleteFunc(m.users, func(r *User) bool {
				return !slices.Contains(m.UserIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *Gender) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Group struct {
	AdminGroupForMeetingID                  *int     `json:"admin_group_for_meeting_id"`
	AnonymousGroupForMeetingID              *int     `json:"anonymous_group_for_meeting_id"`
	DefaultGroupForMeetingID                *int     `json:"default_group_for_meeting_id"`
	ExternalID                              *string  `json:"external_id"`
	ID                                      int      `json:"id"`
	MeetingID                               int      `json:"meeting_id"`
	MeetingMediafileAccessGroupIDs          []int    `json:"meeting_mediafile_access_group_ids"`
	MeetingMediafileInheritedAccessGroupIDs []int    `json:"meeting_mediafile_inherited_access_group_ids"`
	MeetingUserIDs                          []int    `json:"meeting_user_ids"`
	Name                                    string   `json:"name"`
	Permissions                             []string `json:"permissions"`
	PollIDs                                 []int    `json:"poll_ids"`
	ReadChatGroupIDs                        []int    `json:"read_chat_group_ids"`
	ReadCommentSectionIDs                   []int    `json:"read_comment_section_ids"`
	UsedAsAssignmentPollDefaultID           *int     `json:"used_as_assignment_poll_default_id"`
	UsedAsMotionPollDefaultID               *int     `json:"used_as_motion_poll_default_id"`
	UsedAsPollDefaultID                     *int     `json:"used_as_poll_default_id"`
	UsedAsTopicPollDefaultID                *int     `json:"used_as_topic_poll_default_id"`
	Weight                                  *int     `json:"weight"`
	WriteChatGroupIDs                       []int    `json:"write_chat_group_ids"`
	WriteCommentSectionIDs                  []int    `json:"write_comment_section_ids"`
	loadedRelations                         map[string]struct{}
	adminGroupForMeeting                    *Meeting
	anonymousGroupForMeeting                *Meeting
	defaultGroupForMeeting                  *Meeting
	meeting                                 *Meeting
	meetingMediafileAccessGroups            []*MeetingMediafile
	meetingMediafileInheritedAccessGroups   []*MeetingMediafile
	meetingUsers                            []*MeetingUser
	polls                                   []*Poll
	readChatGroups                          []*ChatGroup
	readCommentSections                     []*MotionCommentSection
	usedAsAssignmentPollDefault             *Meeting
	usedAsMotionPollDefault                 *Meeting
	usedAsPollDefault                       *Meeting
	usedAsTopicPollDefault                  *Meeting
	writeChatGroups                         []*ChatGroup
	writeCommentSections                    []*MotionCommentSection
}

func (m *Group) CollectionName() string {
	return "group"
}

func (m *Group) AdminGroupForMeeting() *Meeting {
	if _, ok := m.loadedRelations["admin_group_for_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access AdminGroupForMeeting relation of Group which was not loaded.")
	}

	return m.adminGroupForMeeting
}

func (m *Group) AnonymousGroupForMeeting() *Meeting {
	if _, ok := m.loadedRelations["anonymous_group_for_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access AnonymousGroupForMeeting relation of Group which was not loaded.")
	}

	return m.anonymousGroupForMeeting
}

func (m *Group) DefaultGroupForMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_group_for_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultGroupForMeeting relation of Group which was not loaded.")
	}

	return m.defaultGroupForMeeting
}

func (m *Group) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Group which was not loaded.")
	}

	return *m.meeting
}

func (m *Group) MeetingMediafileAccessGroups() []*MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafileAccessGroups relation of Group which was not loaded.")
	}

	return m.meetingMediafileAccessGroups
}

func (m *Group) MeetingMediafileInheritedAccessGroups() []*MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_inherited_access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafileInheritedAccessGroups relation of Group which was not loaded.")
	}

	return m.meetingMediafileInheritedAccessGroups
}

func (m *Group) MeetingUsers() []*MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingUsers relation of Group which was not loaded.")
	}

	return m.meetingUsers
}

func (m *Group) Polls() []*Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Group which was not loaded.")
	}

	return m.polls
}

func (m *Group) ReadChatGroups() []*ChatGroup {
	if _, ok := m.loadedRelations["read_chat_group_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadChatGroups relation of Group which was not loaded.")
	}

	return m.readChatGroups
}

func (m *Group) ReadCommentSections() []*MotionCommentSection {
	if _, ok := m.loadedRelations["read_comment_section_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadCommentSections relation of Group which was not loaded.")
	}

	return m.readCommentSections
}

func (m *Group) UsedAsAssignmentPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_assignment_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsAssignmentPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsAssignmentPollDefault
}

func (m *Group) UsedAsMotionPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_motion_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsMotionPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsMotionPollDefault
}

func (m *Group) UsedAsPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsPollDefault
}

func (m *Group) UsedAsTopicPollDefault() *Meeting {
	if _, ok := m.loadedRelations["used_as_topic_poll_default_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsTopicPollDefault relation of Group which was not loaded.")
	}

	return m.usedAsTopicPollDefault
}

func (m *Group) WriteChatGroups() []*ChatGroup {
	if _, ok := m.loadedRelations["write_chat_group_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteChatGroups relation of Group which was not loaded.")
	}

	return m.writeChatGroups
}

func (m *Group) WriteCommentSections() []*MotionCommentSection {
	if _, ok := m.loadedRelations["write_comment_section_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteCommentSections relation of Group which was not loaded.")
	}

	return m.writeCommentSections
}

func (m *Group) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "admin_group_for_meeting_id":
		return m.adminGroupForMeeting.GetRelatedModelsAccessor()
	case "anonymous_group_for_meeting_id":
		return m.anonymousGroupForMeeting.GetRelatedModelsAccessor()
	case "default_group_for_meeting_id":
		return m.defaultGroupForMeeting.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "meeting_mediafile_access_group_ids":
		for _, r := range m.meetingMediafileAccessGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_mediafile_inherited_access_group_ids":
		for _, r := range m.meetingMediafileInheritedAccessGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_user_ids":
		for _, r := range m.meetingUsers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "poll_ids":
		for _, r := range m.polls {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "read_chat_group_ids":
		for _, r := range m.readChatGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "read_comment_section_ids":
		for _, r := range m.readCommentSections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "used_as_assignment_poll_default_id":
		return m.usedAsAssignmentPollDefault.GetRelatedModelsAccessor()
	case "used_as_motion_poll_default_id":
		return m.usedAsMotionPollDefault.GetRelatedModelsAccessor()
	case "used_as_poll_default_id":
		return m.usedAsPollDefault.GetRelatedModelsAccessor()
	case "used_as_topic_poll_default_id":
		return m.usedAsTopicPollDefault.GetRelatedModelsAccessor()
	case "write_chat_group_ids":
		for _, r := range m.writeChatGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "write_comment_section_ids":
		for _, r := range m.writeCommentSections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Group) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "admin_group_for_meeting_id":
			m.adminGroupForMeeting = content.(*Meeting)
		case "anonymous_group_for_meeting_id":
			m.anonymousGroupForMeeting = content.(*Meeting)
		case "default_group_for_meeting_id":
			m.defaultGroupForMeeting = content.(*Meeting)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_mediafile_access_group_ids":
			m.meetingMediafileAccessGroups = content.([]*MeetingMediafile)
		case "meeting_mediafile_inherited_access_group_ids":
			m.meetingMediafileInheritedAccessGroups = content.([]*MeetingMediafile)
		case "meeting_user_ids":
			m.meetingUsers = content.([]*MeetingUser)
		case "poll_ids":
			m.polls = content.([]*Poll)
		case "read_chat_group_ids":
			m.readChatGroups = content.([]*ChatGroup)
		case "read_comment_section_ids":
			m.readCommentSections = content.([]*MotionCommentSection)
		case "used_as_assignment_poll_default_id":
			m.usedAsAssignmentPollDefault = content.(*Meeting)
		case "used_as_motion_poll_default_id":
			m.usedAsMotionPollDefault = content.(*Meeting)
		case "used_as_poll_default_id":
			m.usedAsPollDefault = content.(*Meeting)
		case "used_as_topic_poll_default_id":
			m.usedAsTopicPollDefault = content.(*Meeting)
		case "write_chat_group_ids":
			m.writeChatGroups = content.([]*ChatGroup)
		case "write_comment_section_ids":
			m.writeCommentSections = content.([]*MotionCommentSection)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Group) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "admin_group_for_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.adminGroupForMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "anonymous_group_for_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.anonymousGroupForMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "default_group_for_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultGroupForMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_mediafile_access_group_ids":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingMediafileAccessGroups = append(m.meetingMediafileAccessGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_mediafile_inherited_access_group_ids":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingMediafileInheritedAccessGroups = append(m.meetingMediafileInheritedAccessGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_user_ids":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUsers = append(m.meetingUsers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_ids":
		var entry Poll
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.polls = append(m.polls, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "read_chat_group_ids":
		var entry ChatGroup
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.readChatGroups = append(m.readChatGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "read_comment_section_ids":
		var entry MotionCommentSection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.readCommentSections = append(m.readCommentSections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "used_as_assignment_poll_default_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsAssignmentPollDefault = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_motion_poll_default_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsMotionPollDefault = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_poll_default_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsPollDefault = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_topic_poll_default_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsTopicPollDefault = &entry

		result = entry.GetRelatedModelsAccessor()
	case "write_chat_group_ids":
		var entry ChatGroup
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.writeChatGroups = append(m.writeChatGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "write_comment_section_ids":
		var entry MotionCommentSection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.writeCommentSections = append(m.writeCommentSections, &entry)

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

func (m *Group) Get(field string) interface{} {
	switch field {
	case "admin_group_for_meeting_id":
		return m.AdminGroupForMeetingID
	case "anonymous_group_for_meeting_id":
		return m.AnonymousGroupForMeetingID
	case "default_group_for_meeting_id":
		return m.DefaultGroupForMeetingID
	case "external_id":
		return m.ExternalID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_mediafile_access_group_ids":
		return m.MeetingMediafileAccessGroupIDs
	case "meeting_mediafile_inherited_access_group_ids":
		return m.MeetingMediafileInheritedAccessGroupIDs
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "name":
		return m.Name
	case "permissions":
		return m.Permissions
	case "poll_ids":
		return m.PollIDs
	case "read_chat_group_ids":
		return m.ReadChatGroupIDs
	case "read_comment_section_ids":
		return m.ReadCommentSectionIDs
	case "used_as_assignment_poll_default_id":
		return m.UsedAsAssignmentPollDefaultID
	case "used_as_motion_poll_default_id":
		return m.UsedAsMotionPollDefaultID
	case "used_as_poll_default_id":
		return m.UsedAsPollDefaultID
	case "used_as_topic_poll_default_id":
		return m.UsedAsTopicPollDefaultID
	case "weight":
		return m.Weight
	case "write_chat_group_ids":
		return m.WriteChatGroupIDs
	case "write_comment_section_ids":
		return m.WriteCommentSectionIDs
	}

	return nil
}

func (m *Group) GetFqids(field string) []string {
	switch field {
	case "admin_group_for_meeting_id":
		if m.AdminGroupForMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.AdminGroupForMeetingID)}
		}

	case "anonymous_group_for_meeting_id":
		if m.AnonymousGroupForMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.AnonymousGroupForMeetingID)}
		}

	case "default_group_for_meeting_id":
		if m.DefaultGroupForMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultGroupForMeetingID)}
		}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_mediafile_access_group_ids":
		r := make([]string, len(m.MeetingMediafileAccessGroupIDs))
		for i, id := range m.MeetingMediafileAccessGroupIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "meeting_mediafile_inherited_access_group_ids":
		r := make([]string, len(m.MeetingMediafileInheritedAccessGroupIDs))
		for i, id := range m.MeetingMediafileInheritedAccessGroupIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "meeting_user_ids":
		r := make([]string, len(m.MeetingUserIDs))
		for i, id := range m.MeetingUserIDs {
			r[i] = "meeting_user/" + strconv.Itoa(id)
		}
		return r

	case "poll_ids":
		r := make([]string, len(m.PollIDs))
		for i, id := range m.PollIDs {
			r[i] = "poll/" + strconv.Itoa(id)
		}
		return r

	case "read_chat_group_ids":
		r := make([]string, len(m.ReadChatGroupIDs))
		for i, id := range m.ReadChatGroupIDs {
			r[i] = "chat_group/" + strconv.Itoa(id)
		}
		return r

	case "read_comment_section_ids":
		r := make([]string, len(m.ReadCommentSectionIDs))
		for i, id := range m.ReadCommentSectionIDs {
			r[i] = "motion_comment_section/" + strconv.Itoa(id)
		}
		return r

	case "used_as_assignment_poll_default_id":
		if m.UsedAsAssignmentPollDefaultID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsAssignmentPollDefaultID)}
		}

	case "used_as_motion_poll_default_id":
		if m.UsedAsMotionPollDefaultID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsMotionPollDefaultID)}
		}

	case "used_as_poll_default_id":
		if m.UsedAsPollDefaultID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsPollDefaultID)}
		}

	case "used_as_topic_poll_default_id":
		if m.UsedAsTopicPollDefaultID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsTopicPollDefaultID)}
		}

	case "write_chat_group_ids":
		r := make([]string, len(m.WriteChatGroupIDs))
		for i, id := range m.WriteChatGroupIDs {
			r[i] = "chat_group/" + strconv.Itoa(id)
		}
		return r

	case "write_comment_section_ids":
		r := make([]string, len(m.WriteCommentSectionIDs))
		for i, id := range m.WriteCommentSectionIDs {
			r[i] = "motion_comment_section/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Group) Update(data map[string]string) error {
	if val, ok := data["admin_group_for_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AdminGroupForMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["anonymous_group_for_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AnonymousGroupForMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_group_for_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultGroupForMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["external_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExternalID)
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

	if val, ok := data["meeting_mediafile_access_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingMediafileAccessGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["meeting_mediafile_access_group_ids"]; ok {
			m.meetingMediafileAccessGroups = slices.DeleteFunc(m.meetingMediafileAccessGroups, func(r *MeetingMediafile) bool {
				return !slices.Contains(m.MeetingMediafileAccessGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["meeting_mediafile_inherited_access_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingMediafileInheritedAccessGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["meeting_mediafile_inherited_access_group_ids"]; ok {
			m.meetingMediafileInheritedAccessGroups = slices.DeleteFunc(m.meetingMediafileInheritedAccessGroups, func(r *MeetingMediafile) bool {
				return !slices.Contains(m.MeetingMediafileInheritedAccessGroupIDs, r.ID)
			})
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

	if val, ok := data["permissions"]; ok {
		err := json.Unmarshal([]byte(val), &m.Permissions)
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

	if val, ok := data["read_chat_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReadChatGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["read_chat_group_ids"]; ok {
			m.readChatGroups = slices.DeleteFunc(m.readChatGroups, func(r *ChatGroup) bool {
				return !slices.Contains(m.ReadChatGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["read_comment_section_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReadCommentSectionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["read_comment_section_ids"]; ok {
			m.readCommentSections = slices.DeleteFunc(m.readCommentSections, func(r *MotionCommentSection) bool {
				return !slices.Contains(m.ReadCommentSectionIDs, r.ID)
			})
		}
	}

	if val, ok := data["used_as_assignment_poll_default_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsAssignmentPollDefaultID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_motion_poll_default_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsMotionPollDefaultID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_poll_default_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsPollDefaultID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_topic_poll_default_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsTopicPollDefaultID)
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

	if val, ok := data["write_chat_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WriteChatGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["write_chat_group_ids"]; ok {
			m.writeChatGroups = slices.DeleteFunc(m.writeChatGroups, func(r *ChatGroup) bool {
				return !slices.Contains(m.WriteChatGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["write_comment_section_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WriteCommentSectionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["write_comment_section_ids"]; ok {
			m.writeCommentSections = slices.DeleteFunc(m.writeCommentSections, func(r *MotionCommentSection) bool {
				return !slices.Contains(m.WriteCommentSectionIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *Group) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type ImportPreview struct {
	Created int             `json:"created"`
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Result  json.RawMessage `json:"result"`
	State   string          `json:"state"`
}

func (m *ImportPreview) CollectionName() string {
	return "import_preview"
}

func (m *ImportPreview) GetRelated(field string, id int) *RelatedModelsAccessor {
	return nil
}

func (m *ImportPreview) SetRelated(field string, content interface{}) {}

func (m *ImportPreview) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	return nil, nil
}

func (m *ImportPreview) Get(field string) interface{} {
	switch field {
	case "created":
		return m.Created
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "result":
		return m.Result
	case "state":
		return m.State
	}

	return nil
}

func (m *ImportPreview) GetFqids(field string) []string {
	return []string{}
}

func (m *ImportPreview) Update(data map[string]string) error {
	if val, ok := data["created"]; ok {
		err := json.Unmarshal([]byte(val), &m.Created)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["result"]; ok {
		err := json.Unmarshal([]byte(val), &m.Result)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state"]; ok {
		err := json.Unmarshal([]byte(val), &m.State)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ImportPreview) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

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

func (m *ListOfSpeakers) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "content_object_id":
		return m.contentObject.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "projection_ids":
		for _, r := range m.projections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "speaker_ids":
		for _, r := range m.speakers {
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
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Mediafile struct {
	ChildIDs                            []int           `json:"child_ids"`
	CreateTimestamp                     *int            `json:"create_timestamp"`
	Filename                            *string         `json:"filename"`
	Filesize                            *int            `json:"filesize"`
	ID                                  int             `json:"id"`
	IsDirectory                         *bool           `json:"is_directory"`
	MeetingMediafileIDs                 []int           `json:"meeting_mediafile_ids"`
	Mimetype                            *string         `json:"mimetype"`
	OwnerID                             string          `json:"owner_id"`
	ParentID                            *int            `json:"parent_id"`
	PdfInformation                      json.RawMessage `json:"pdf_information"`
	PublishedToMeetingsInOrganizationID *int            `json:"published_to_meetings_in_organization_id"`
	Title                               *string         `json:"title"`
	Token                               *string         `json:"token"`
	loadedRelations                     map[string]struct{}
	childs                              []*Mediafile
	meetingMediafiles                   []*MeetingMediafile
	owner                               IBaseModel
	parent                              *Mediafile
	publishedToMeetingsInOrganization   *Organization
}

func (m *Mediafile) CollectionName() string {
	return "mediafile"
}

func (m *Mediafile) Childs() []*Mediafile {
	if _, ok := m.loadedRelations["child_ids"]; !ok {
		log.Panic().Msg("Tried to access Childs relation of Mediafile which was not loaded.")
	}

	return m.childs
}

func (m *Mediafile) MeetingMediafiles() []*MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafiles relation of Mediafile which was not loaded.")
	}

	return m.meetingMediafiles
}

func (m *Mediafile) Owner() IBaseModel {
	if _, ok := m.loadedRelations["owner_id"]; !ok {
		log.Panic().Msg("Tried to access Owner relation of Mediafile which was not loaded.")
	}

	return m.owner
}

func (m *Mediafile) Parent() *Mediafile {
	if _, ok := m.loadedRelations["parent_id"]; !ok {
		log.Panic().Msg("Tried to access Parent relation of Mediafile which was not loaded.")
	}

	return m.parent
}

func (m *Mediafile) PublishedToMeetingsInOrganization() *Organization {
	if _, ok := m.loadedRelations["published_to_meetings_in_organization_id"]; !ok {
		log.Panic().Msg("Tried to access PublishedToMeetingsInOrganization relation of Mediafile which was not loaded.")
	}

	return m.publishedToMeetingsInOrganization
}

func (m *Mediafile) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "child_ids":
		for _, r := range m.childs {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_mediafile_ids":
		for _, r := range m.meetingMediafiles {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "owner_id":
		return m.owner.GetRelatedModelsAccessor()
	case "parent_id":
		return m.parent.GetRelatedModelsAccessor()
	case "published_to_meetings_in_organization_id":
		return m.publishedToMeetingsInOrganization.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *Mediafile) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "child_ids":
			m.childs = content.([]*Mediafile)
		case "meeting_mediafile_ids":
			m.meetingMediafiles = content.([]*MeetingMediafile)
		case "owner_id":
			panic("not implemented")
		case "parent_id":
			m.parent = content.(*Mediafile)
		case "published_to_meetings_in_organization_id":
			m.publishedToMeetingsInOrganization = content.(*Organization)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Mediafile) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "child_ids":
		var entry Mediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.childs = append(m.childs, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_mediafile_ids":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingMediafiles = append(m.meetingMediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "owner_id":
		parts := strings.Split(m.OwnerID, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf("could not parse id field")
		}

		switch parts[0] {
		case "meeting":
			var entry Meeting
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.owner = &entry
			result = entry.GetRelatedModelsAccessor()

		case "organization":
			var entry Organization
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.owner = &entry
			result = entry.GetRelatedModelsAccessor()
		}

	case "parent_id":
		var entry Mediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.parent = &entry

		result = entry.GetRelatedModelsAccessor()
	case "published_to_meetings_in_organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.publishedToMeetingsInOrganization = &entry

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

func (m *Mediafile) Get(field string) interface{} {
	switch field {
	case "child_ids":
		return m.ChildIDs
	case "create_timestamp":
		return m.CreateTimestamp
	case "filename":
		return m.Filename
	case "filesize":
		return m.Filesize
	case "id":
		return m.ID
	case "is_directory":
		return m.IsDirectory
	case "meeting_mediafile_ids":
		return m.MeetingMediafileIDs
	case "mimetype":
		return m.Mimetype
	case "owner_id":
		return m.OwnerID
	case "parent_id":
		return m.ParentID
	case "pdf_information":
		return m.PdfInformation
	case "published_to_meetings_in_organization_id":
		return m.PublishedToMeetingsInOrganizationID
	case "title":
		return m.Title
	case "token":
		return m.Token
	}

	return nil
}

func (m *Mediafile) GetFqids(field string) []string {
	switch field {
	case "child_ids":
		r := make([]string, len(m.ChildIDs))
		for i, id := range m.ChildIDs {
			r[i] = "mediafile/" + strconv.Itoa(id)
		}
		return r

	case "meeting_mediafile_ids":
		r := make([]string, len(m.MeetingMediafileIDs))
		for i, id := range m.MeetingMediafileIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "owner_id":
		return []string{m.OwnerID}

	case "parent_id":
		if m.ParentID != nil {
			return []string{"mediafile/" + strconv.Itoa(*m.ParentID)}
		}

	case "published_to_meetings_in_organization_id":
		if m.PublishedToMeetingsInOrganizationID != nil {
			return []string{"organization/" + strconv.Itoa(*m.PublishedToMeetingsInOrganizationID)}
		}
	}
	return []string{}
}

func (m *Mediafile) Update(data map[string]string) error {
	if val, ok := data["child_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChildIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["child_ids"]; ok {
			m.childs = slices.DeleteFunc(m.childs, func(r *Mediafile) bool {
				return !slices.Contains(m.ChildIDs, r.ID)
			})
		}
	}

	if val, ok := data["create_timestamp"]; ok {
		err := json.Unmarshal([]byte(val), &m.CreateTimestamp)
		if err != nil {
			return err
		}
	}

	if val, ok := data["filename"]; ok {
		err := json.Unmarshal([]byte(val), &m.Filename)
		if err != nil {
			return err
		}
	}

	if val, ok := data["filesize"]; ok {
		err := json.Unmarshal([]byte(val), &m.Filesize)
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

	if val, ok := data["is_directory"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsDirectory)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingMediafileIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["meeting_mediafile_ids"]; ok {
			m.meetingMediafiles = slices.DeleteFunc(m.meetingMediafiles, func(r *MeetingMediafile) bool {
				return !slices.Contains(m.MeetingMediafileIDs, r.ID)
			})
		}
	}

	if val, ok := data["mimetype"]; ok {
		err := json.Unmarshal([]byte(val), &m.Mimetype)
		if err != nil {
			return err
		}
	}

	if val, ok := data["owner_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OwnerID)
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

	if val, ok := data["pdf_information"]; ok {
		err := json.Unmarshal([]byte(val), &m.PdfInformation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["published_to_meetings_in_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PublishedToMeetingsInOrganizationID)
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

	if val, ok := data["token"]; ok {
		err := json.Unmarshal([]byte(val), &m.Token)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Mediafile) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Meeting struct {
	AdminGroupID                                 *int            `json:"admin_group_id"`
	AgendaEnableNumbering                        *bool           `json:"agenda_enable_numbering"`
	AgendaItemCreation                           *string         `json:"agenda_item_creation"`
	AgendaItemIDs                                []int           `json:"agenda_item_ids"`
	AgendaNewItemsDefaultVisibility              *string         `json:"agenda_new_items_default_visibility"`
	AgendaNumberPrefix                           *string         `json:"agenda_number_prefix"`
	AgendaNumeralSystem                          *string         `json:"agenda_numeral_system"`
	AgendaShowInternalItemsOnProjector           *bool           `json:"agenda_show_internal_items_on_projector"`
	AgendaShowSubtitles                          *bool           `json:"agenda_show_subtitles"`
	AgendaShowTopicNavigationOnDetailView        *bool           `json:"agenda_show_topic_navigation_on_detail_view"`
	AllProjectionIDs                             []int           `json:"all_projection_ids"`
	AnonymousGroupID                             *int            `json:"anonymous_group_id"`
	ApplauseEnable                               *bool           `json:"applause_enable"`
	ApplauseMaxAmount                            *int            `json:"applause_max_amount"`
	ApplauseMinAmount                            *int            `json:"applause_min_amount"`
	ApplauseParticleImageUrl                     *string         `json:"applause_particle_image_url"`
	ApplauseShowLevel                            *bool           `json:"applause_show_level"`
	ApplauseTimeout                              *int            `json:"applause_timeout"`
	ApplauseType                                 *string         `json:"applause_type"`
	AssignmentCandidateIDs                       []int           `json:"assignment_candidate_ids"`
	AssignmentIDs                                []int           `json:"assignment_ids"`
	AssignmentPollAddCandidatesToListOfSpeakers  *bool           `json:"assignment_poll_add_candidates_to_list_of_speakers"`
	AssignmentPollBallotPaperNumber              *int            `json:"assignment_poll_ballot_paper_number"`
	AssignmentPollBallotPaperSelection           *string         `json:"assignment_poll_ballot_paper_selection"`
	AssignmentPollDefaultBackend                 *string         `json:"assignment_poll_default_backend"`
	AssignmentPollDefaultGroupIDs                []int           `json:"assignment_poll_default_group_ids"`
	AssignmentPollDefaultMethod                  *string         `json:"assignment_poll_default_method"`
	AssignmentPollDefaultOnehundredPercentBase   *string         `json:"assignment_poll_default_onehundred_percent_base"`
	AssignmentPollDefaultType                    *string         `json:"assignment_poll_default_type"`
	AssignmentPollEnableMaxVotesPerOption        *bool           `json:"assignment_poll_enable_max_votes_per_option"`
	AssignmentPollSortPollResultByVotes          *bool           `json:"assignment_poll_sort_poll_result_by_votes"`
	AssignmentsExportPreamble                    *string         `json:"assignments_export_preamble"`
	AssignmentsExportTitle                       *string         `json:"assignments_export_title"`
	ChatGroupIDs                                 []int           `json:"chat_group_ids"`
	ChatMessageIDs                               []int           `json:"chat_message_ids"`
	CommitteeID                                  int             `json:"committee_id"`
	ConferenceAutoConnect                        *bool           `json:"conference_auto_connect"`
	ConferenceAutoConnectNextSpeakers            *int            `json:"conference_auto_connect_next_speakers"`
	ConferenceEnableHelpdesk                     *bool           `json:"conference_enable_helpdesk"`
	ConferenceLosRestriction                     *bool           `json:"conference_los_restriction"`
	ConferenceOpenMicrophone                     *bool           `json:"conference_open_microphone"`
	ConferenceOpenVideo                          *bool           `json:"conference_open_video"`
	ConferenceShow                               *bool           `json:"conference_show"`
	ConferenceStreamPosterUrl                    *string         `json:"conference_stream_poster_url"`
	ConferenceStreamUrl                          *string         `json:"conference_stream_url"`
	CustomTranslations                           json.RawMessage `json:"custom_translations"`
	DefaultGroupID                               int             `json:"default_group_id"`
	DefaultMeetingForCommitteeID                 *int            `json:"default_meeting_for_committee_id"`
	DefaultProjectorAgendaItemListIDs            []int           `json:"default_projector_agenda_item_list_ids"`
	DefaultProjectorAmendmentIDs                 []int           `json:"default_projector_amendment_ids"`
	DefaultProjectorAssignmentIDs                []int           `json:"default_projector_assignment_ids"`
	DefaultProjectorAssignmentPollIDs            []int           `json:"default_projector_assignment_poll_ids"`
	DefaultProjectorCountdownIDs                 []int           `json:"default_projector_countdown_ids"`
	DefaultProjectorCurrentListOfSpeakersIDs     []int           `json:"default_projector_current_list_of_speakers_ids"`
	DefaultProjectorListOfSpeakersIDs            []int           `json:"default_projector_list_of_speakers_ids"`
	DefaultProjectorMediafileIDs                 []int           `json:"default_projector_mediafile_ids"`
	DefaultProjectorMessageIDs                   []int           `json:"default_projector_message_ids"`
	DefaultProjectorMotionBlockIDs               []int           `json:"default_projector_motion_block_ids"`
	DefaultProjectorMotionIDs                    []int           `json:"default_projector_motion_ids"`
	DefaultProjectorMotionPollIDs                []int           `json:"default_projector_motion_poll_ids"`
	DefaultProjectorPollIDs                      []int           `json:"default_projector_poll_ids"`
	DefaultProjectorTopicIDs                     []int           `json:"default_projector_topic_ids"`
	Description                                  *string         `json:"description"`
	EnableAnonymous                              *bool           `json:"enable_anonymous"`
	EndTime                                      *int            `json:"end_time"`
	ExportCsvEncoding                            *string         `json:"export_csv_encoding"`
	ExportCsvSeparator                           *string         `json:"export_csv_separator"`
	ExportPdfFontsize                            *int            `json:"export_pdf_fontsize"`
	ExportPdfLineHeight                          *float32        `json:"export_pdf_line_height"`
	ExportPdfPageMarginBottom                    *int            `json:"export_pdf_page_margin_bottom"`
	ExportPdfPageMarginLeft                      *int            `json:"export_pdf_page_margin_left"`
	ExportPdfPageMarginRight                     *int            `json:"export_pdf_page_margin_right"`
	ExportPdfPageMarginTop                       *int            `json:"export_pdf_page_margin_top"`
	ExportPdfPagenumberAlignment                 *string         `json:"export_pdf_pagenumber_alignment"`
	ExportPdfPagesize                            *string         `json:"export_pdf_pagesize"`
	ExternalID                                   *string         `json:"external_id"`
	FontBoldID                                   *int            `json:"font_bold_id"`
	FontBoldItalicID                             *int            `json:"font_bold_italic_id"`
	FontChyronSpeakerNameID                      *int            `json:"font_chyron_speaker_name_id"`
	FontItalicID                                 *int            `json:"font_italic_id"`
	FontMonospaceID                              *int            `json:"font_monospace_id"`
	FontProjectorH1ID                            *int            `json:"font_projector_h1_id"`
	FontProjectorH2ID                            *int            `json:"font_projector_h2_id"`
	FontRegularID                                *int            `json:"font_regular_id"`
	ForwardedMotionIDs                           []int           `json:"forwarded_motion_ids"`
	GroupIDs                                     []int           `json:"group_ids"`
	ID                                           int             `json:"id"`
	ImportedAt                                   *int            `json:"imported_at"`
	IsActiveInOrganizationID                     *int            `json:"is_active_in_organization_id"`
	IsArchivedInOrganizationID                   *int            `json:"is_archived_in_organization_id"`
	JitsiDomain                                  *string         `json:"jitsi_domain"`
	JitsiRoomName                                *string         `json:"jitsi_room_name"`
	JitsiRoomPassword                            *string         `json:"jitsi_room_password"`
	Language                                     string          `json:"language"`
	ListOfSpeakersAllowMultipleSpeakers          *bool           `json:"list_of_speakers_allow_multiple_speakers"`
	ListOfSpeakersAmountLastOnProjector          *int            `json:"list_of_speakers_amount_last_on_projector"`
	ListOfSpeakersAmountNextOnProjector          *int            `json:"list_of_speakers_amount_next_on_projector"`
	ListOfSpeakersCanCreatePointOfOrderForOthers *bool           `json:"list_of_speakers_can_create_point_of_order_for_others"`
	ListOfSpeakersCanSetContributionSelf         *bool           `json:"list_of_speakers_can_set_contribution_self"`
	ListOfSpeakersClosingDisablesPointOfOrder    *bool           `json:"list_of_speakers_closing_disables_point_of_order"`
	ListOfSpeakersCountdownID                    *int            `json:"list_of_speakers_countdown_id"`
	ListOfSpeakersCoupleCountdown                *bool           `json:"list_of_speakers_couple_countdown"`
	ListOfSpeakersDefaultStructureLevelTime      *int            `json:"list_of_speakers_default_structure_level_time"`
	ListOfSpeakersEnableInterposedQuestion       *bool           `json:"list_of_speakers_enable_interposed_question"`
	ListOfSpeakersEnablePointOfOrderCategories   *bool           `json:"list_of_speakers_enable_point_of_order_categories"`
	ListOfSpeakersEnablePointOfOrderSpeakers     *bool           `json:"list_of_speakers_enable_point_of_order_speakers"`
	ListOfSpeakersEnableProContraSpeech          *bool           `json:"list_of_speakers_enable_pro_contra_speech"`
	ListOfSpeakersHideContributionCount          *bool           `json:"list_of_speakers_hide_contribution_count"`
	ListOfSpeakersIDs                            []int           `json:"list_of_speakers_ids"`
	ListOfSpeakersInitiallyClosed                *bool           `json:"list_of_speakers_initially_closed"`
	ListOfSpeakersInterventionTime               *int            `json:"list_of_speakers_intervention_time"`
	ListOfSpeakersPresentUsersOnly               *bool           `json:"list_of_speakers_present_users_only"`
	ListOfSpeakersShowAmountOfSpeakersOnSlide    *bool           `json:"list_of_speakers_show_amount_of_speakers_on_slide"`
	ListOfSpeakersShowFirstContribution          *bool           `json:"list_of_speakers_show_first_contribution"`
	ListOfSpeakersSpeakerNoteForEveryone         *bool           `json:"list_of_speakers_speaker_note_for_everyone"`
	Location                                     *string         `json:"location"`
	LockedFromInside                             *bool           `json:"locked_from_inside"`
	LogoPdfBallotPaperID                         *int            `json:"logo_pdf_ballot_paper_id"`
	LogoPdfFooterLID                             *int            `json:"logo_pdf_footer_l_id"`
	LogoPdfFooterRID                             *int            `json:"logo_pdf_footer_r_id"`
	LogoPdfHeaderLID                             *int            `json:"logo_pdf_header_l_id"`
	LogoPdfHeaderRID                             *int            `json:"logo_pdf_header_r_id"`
	LogoProjectorHeaderID                        *int            `json:"logo_projector_header_id"`
	LogoProjectorMainID                          *int            `json:"logo_projector_main_id"`
	LogoWebHeaderID                              *int            `json:"logo_web_header_id"`
	MediafileIDs                                 []int           `json:"mediafile_ids"`
	MeetingMediafileIDs                          []int           `json:"meeting_mediafile_ids"`
	MeetingUserIDs                               []int           `json:"meeting_user_ids"`
	MotionBlockIDs                               []int           `json:"motion_block_ids"`
	MotionCategoryIDs                            []int           `json:"motion_category_ids"`
	MotionChangeRecommendationIDs                []int           `json:"motion_change_recommendation_ids"`
	MotionCommentIDs                             []int           `json:"motion_comment_ids"`
	MotionCommentSectionIDs                      []int           `json:"motion_comment_section_ids"`
	MotionEditorIDs                              []int           `json:"motion_editor_ids"`
	MotionIDs                                    []int           `json:"motion_ids"`
	MotionPollBallotPaperNumber                  *int            `json:"motion_poll_ballot_paper_number"`
	MotionPollBallotPaperSelection               *string         `json:"motion_poll_ballot_paper_selection"`
	MotionPollDefaultBackend                     *string         `json:"motion_poll_default_backend"`
	MotionPollDefaultGroupIDs                    []int           `json:"motion_poll_default_group_ids"`
	MotionPollDefaultMethod                      *string         `json:"motion_poll_default_method"`
	MotionPollDefaultOnehundredPercentBase       *string         `json:"motion_poll_default_onehundred_percent_base"`
	MotionPollDefaultType                        *string         `json:"motion_poll_default_type"`
	MotionStateIDs                               []int           `json:"motion_state_ids"`
	MotionSubmitterIDs                           []int           `json:"motion_submitter_ids"`
	MotionWorkflowIDs                            []int           `json:"motion_workflow_ids"`
	MotionWorkingGroupSpeakerIDs                 []int           `json:"motion_working_group_speaker_ids"`
	MotionsAmendmentsEnabled                     *bool           `json:"motions_amendments_enabled"`
	MotionsAmendmentsInMainList                  *bool           `json:"motions_amendments_in_main_list"`
	MotionsAmendmentsMultipleParagraphs          *bool           `json:"motions_amendments_multiple_paragraphs"`
	MotionsAmendmentsOfAmendments                *bool           `json:"motions_amendments_of_amendments"`
	MotionsAmendmentsPrefix                      *string         `json:"motions_amendments_prefix"`
	MotionsAmendmentsTextMode                    *string         `json:"motions_amendments_text_mode"`
	MotionsBlockSlideColumns                     *int            `json:"motions_block_slide_columns"`
	MotionsCreateEnableAdditionalSubmitterText   *bool           `json:"motions_create_enable_additional_submitter_text"`
	MotionsDefaultAmendmentWorkflowID            int             `json:"motions_default_amendment_workflow_id"`
	MotionsDefaultLineNumbering                  *string         `json:"motions_default_line_numbering"`
	MotionsDefaultSorting                        *string         `json:"motions_default_sorting"`
	MotionsDefaultWorkflowID                     int             `json:"motions_default_workflow_id"`
	MotionsEnableEditor                          *bool           `json:"motions_enable_editor"`
	MotionsEnableReasonOnProjector               *bool           `json:"motions_enable_reason_on_projector"`
	MotionsEnableRecommendationOnProjector       *bool           `json:"motions_enable_recommendation_on_projector"`
	MotionsEnableSideboxOnProjector              *bool           `json:"motions_enable_sidebox_on_projector"`
	MotionsEnableTextOnProjector                 *bool           `json:"motions_enable_text_on_projector"`
	MotionsEnableWorkingGroupSpeaker             *bool           `json:"motions_enable_working_group_speaker"`
	MotionsExportFollowRecommendation            *bool           `json:"motions_export_follow_recommendation"`
	MotionsExportPreamble                        *string         `json:"motions_export_preamble"`
	MotionsExportSubmitterRecommendation         *bool           `json:"motions_export_submitter_recommendation"`
	MotionsExportTitle                           *string         `json:"motions_export_title"`
	MotionsHideMetadataBackground                *bool           `json:"motions_hide_metadata_background"`
	MotionsLineLength                            *int            `json:"motions_line_length"`
	MotionsNumberMinDigits                       *int            `json:"motions_number_min_digits"`
	MotionsNumberType                            *string         `json:"motions_number_type"`
	MotionsNumberWithBlank                       *bool           `json:"motions_number_with_blank"`
	MotionsPreamble                              *string         `json:"motions_preamble"`
	MotionsReasonRequired                        *bool           `json:"motions_reason_required"`
	MotionsRecommendationTextMode                *string         `json:"motions_recommendation_text_mode"`
	MotionsRecommendationsBy                     *string         `json:"motions_recommendations_by"`
	MotionsShowReferringMotions                  *bool           `json:"motions_show_referring_motions"`
	MotionsShowSequentialNumber                  *bool           `json:"motions_show_sequential_number"`
	MotionsSupportersMinAmount                   *int            `json:"motions_supporters_min_amount"`
	Name                                         string          `json:"name"`
	OptionIDs                                    []int           `json:"option_ids"`
	OrganizationTagIDs                           []int           `json:"organization_tag_ids"`
	PersonalNoteIDs                              []int           `json:"personal_note_ids"`
	PointOfOrderCategoryIDs                      []int           `json:"point_of_order_category_ids"`
	PollBallotPaperNumber                        *int            `json:"poll_ballot_paper_number"`
	PollBallotPaperSelection                     *string         `json:"poll_ballot_paper_selection"`
	PollCandidateIDs                             []int           `json:"poll_candidate_ids"`
	PollCandidateListIDs                         []int           `json:"poll_candidate_list_ids"`
	PollCountdownID                              *int            `json:"poll_countdown_id"`
	PollCoupleCountdown                          *bool           `json:"poll_couple_countdown"`
	PollDefaultBackend                           *string         `json:"poll_default_backend"`
	PollDefaultGroupIDs                          []int           `json:"poll_default_group_ids"`
	PollDefaultMethod                            *string         `json:"poll_default_method"`
	PollDefaultOnehundredPercentBase             *string         `json:"poll_default_onehundred_percent_base"`
	PollDefaultType                              *string         `json:"poll_default_type"`
	PollIDs                                      []int           `json:"poll_ids"`
	PollSortPollResultByVotes                    *bool           `json:"poll_sort_poll_result_by_votes"`
	PresentUserIDs                               []int           `json:"present_user_ids"`
	ProjectionIDs                                []int           `json:"projection_ids"`
	ProjectorCountdownDefaultTime                int             `json:"projector_countdown_default_time"`
	ProjectorCountdownIDs                        []int           `json:"projector_countdown_ids"`
	ProjectorCountdownWarningTime                int             `json:"projector_countdown_warning_time"`
	ProjectorIDs                                 []int           `json:"projector_ids"`
	ProjectorMessageIDs                          []int           `json:"projector_message_ids"`
	ReferenceProjectorID                         int             `json:"reference_projector_id"`
	SpeakerIDs                                   []int           `json:"speaker_ids"`
	StartTime                                    *int            `json:"start_time"`
	StructureLevelIDs                            []int           `json:"structure_level_ids"`
	StructureLevelListOfSpeakersIDs              []int           `json:"structure_level_list_of_speakers_ids"`
	TagIDs                                       []int           `json:"tag_ids"`
	TemplateForOrganizationID                    *int            `json:"template_for_organization_id"`
	TopicIDs                                     []int           `json:"topic_ids"`
	TopicPollDefaultGroupIDs                     []int           `json:"topic_poll_default_group_ids"`
	UserIDs                                      []int           `json:"user_ids"`
	UsersAllowSelfSetPresent                     *bool           `json:"users_allow_self_set_present"`
	UsersEmailBody                               *string         `json:"users_email_body"`
	UsersEmailReplyto                            *string         `json:"users_email_replyto"`
	UsersEmailSender                             *string         `json:"users_email_sender"`
	UsersEmailSubject                            *string         `json:"users_email_subject"`
	UsersEnablePresenceView                      *bool           `json:"users_enable_presence_view"`
	UsersEnableVoteDelegations                   *bool           `json:"users_enable_vote_delegations"`
	UsersEnableVoteWeight                        *bool           `json:"users_enable_vote_weight"`
	UsersForbidDelegatorAsSubmitter              *bool           `json:"users_forbid_delegator_as_submitter"`
	UsersForbidDelegatorAsSupporter              *bool           `json:"users_forbid_delegator_as_supporter"`
	UsersForbidDelegatorInListOfSpeakers         *bool           `json:"users_forbid_delegator_in_list_of_speakers"`
	UsersForbidDelegatorToVote                   *bool           `json:"users_forbid_delegator_to_vote"`
	UsersPdfWelcometext                          *string         `json:"users_pdf_welcometext"`
	UsersPdfWelcometitle                         *string         `json:"users_pdf_welcometitle"`
	UsersPdfWlanEncryption                       *string         `json:"users_pdf_wlan_encryption"`
	UsersPdfWlanPassword                         *string         `json:"users_pdf_wlan_password"`
	UsersPdfWlanSsid                             *string         `json:"users_pdf_wlan_ssid"`
	VoteIDs                                      []int           `json:"vote_ids"`
	WelcomeText                                  *string         `json:"welcome_text"`
	WelcomeTitle                                 *string         `json:"welcome_title"`
	loadedRelations                              map[string]struct{}
	adminGroup                                   *Group
	agendaItems                                  []*AgendaItem
	allProjections                               []*Projection
	anonymousGroup                               *Group
	assignmentCandidates                         []*AssignmentCandidate
	assignmentPollDefaultGroups                  []*Group
	assignments                                  []*Assignment
	chatGroups                                   []*ChatGroup
	chatMessages                                 []*ChatMessage
	committee                                    *Committee
	defaultGroup                                 *Group
	defaultMeetingForCommittee                   *Committee
	defaultProjectorAgendaItemLists              []*Projector
	defaultProjectorAmendments                   []*Projector
	defaultProjectorAssignmentPolls              []*Projector
	defaultProjectorAssignments                  []*Projector
	defaultProjectorCountdowns                   []*Projector
	defaultProjectorCurrentListOfSpeakerss       []*Projector
	defaultProjectorListOfSpeakerss              []*Projector
	defaultProjectorMediafiles                   []*Projector
	defaultProjectorMessages                     []*Projector
	defaultProjectorMotionBlocks                 []*Projector
	defaultProjectorMotionPolls                  []*Projector
	defaultProjectorMotions                      []*Projector
	defaultProjectorPolls                        []*Projector
	defaultProjectorTopics                       []*Projector
	fontBold                                     *MeetingMediafile
	fontBoldItalic                               *MeetingMediafile
	fontChyronSpeakerName                        *MeetingMediafile
	fontItalic                                   *MeetingMediafile
	fontMonospace                                *MeetingMediafile
	fontProjectorH1                              *MeetingMediafile
	fontProjectorH2                              *MeetingMediafile
	fontRegular                                  *MeetingMediafile
	forwardedMotions                             []*Motion
	groups                                       []*Group
	isActiveInOrganization                       *Organization
	isArchivedInOrganization                     *Organization
	listOfSpeakersCountdown                      *ProjectorCountdown
	listOfSpeakerss                              []*ListOfSpeakers
	logoPdfBallotPaper                           *MeetingMediafile
	logoPdfFooterL                               *MeetingMediafile
	logoPdfFooterR                               *MeetingMediafile
	logoPdfHeaderL                               *MeetingMediafile
	logoPdfHeaderR                               *MeetingMediafile
	logoProjectorHeader                          *MeetingMediafile
	logoProjectorMain                            *MeetingMediafile
	logoWebHeader                                *MeetingMediafile
	mediafiles                                   []*Mediafile
	meetingMediafiles                            []*MeetingMediafile
	meetingUsers                                 []*MeetingUser
	motionBlocks                                 []*MotionBlock
	motionCategorys                              []*MotionCategory
	motionChangeRecommendations                  []*MotionChangeRecommendation
	motionCommentSections                        []*MotionCommentSection
	motionComments                               []*MotionComment
	motionEditors                                []*MotionEditor
	motionPollDefaultGroups                      []*Group
	motionStates                                 []*MotionState
	motionSubmitters                             []*MotionSubmitter
	motionWorkflows                              []*MotionWorkflow
	motionWorkingGroupSpeakers                   []*MotionWorkingGroupSpeaker
	motions                                      []*Motion
	motionsDefaultAmendmentWorkflow              *MotionWorkflow
	motionsDefaultWorkflow                       *MotionWorkflow
	options                                      []*Option
	organizationTags                             []*OrganizationTag
	personalNotes                                []*PersonalNote
	pointOfOrderCategorys                        []*PointOfOrderCategory
	pollCandidateLists                           []*PollCandidateList
	pollCandidates                               []*PollCandidate
	pollCountdown                                *ProjectorCountdown
	pollDefaultGroups                            []*Group
	polls                                        []*Poll
	presentUsers                                 []*User
	projections                                  []*Projection
	projectorCountdowns                          []*ProjectorCountdown
	projectorMessages                            []*ProjectorMessage
	projectors                                   []*Projector
	referenceProjector                           *Projector
	speakers                                     []*Speaker
	structureLevelListOfSpeakerss                []*StructureLevelListOfSpeakers
	structureLevels                              []*StructureLevel
	tags                                         []*Tag
	templateForOrganization                      *Organization
	topicPollDefaultGroups                       []*Group
	topics                                       []*Topic
	votes                                        []*Vote
}

func (m *Meeting) CollectionName() string {
	return "meeting"
}

func (m *Meeting) AdminGroup() *Group {
	if _, ok := m.loadedRelations["admin_group_id"]; !ok {
		log.Panic().Msg("Tried to access AdminGroup relation of Meeting which was not loaded.")
	}

	return m.adminGroup
}

func (m *Meeting) AgendaItems() []*AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_ids"]; !ok {
		log.Panic().Msg("Tried to access AgendaItems relation of Meeting which was not loaded.")
	}

	return m.agendaItems
}

func (m *Meeting) AllProjections() []*Projection {
	if _, ok := m.loadedRelations["all_projection_ids"]; !ok {
		log.Panic().Msg("Tried to access AllProjections relation of Meeting which was not loaded.")
	}

	return m.allProjections
}

func (m *Meeting) AnonymousGroup() *Group {
	if _, ok := m.loadedRelations["anonymous_group_id"]; !ok {
		log.Panic().Msg("Tried to access AnonymousGroup relation of Meeting which was not loaded.")
	}

	return m.anonymousGroup
}

func (m *Meeting) AssignmentCandidates() []*AssignmentCandidate {
	if _, ok := m.loadedRelations["assignment_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access AssignmentCandidates relation of Meeting which was not loaded.")
	}

	return m.assignmentCandidates
}

func (m *Meeting) AssignmentPollDefaultGroups() []*Group {
	if _, ok := m.loadedRelations["assignment_poll_default_group_ids"]; !ok {
		log.Panic().Msg("Tried to access AssignmentPollDefaultGroups relation of Meeting which was not loaded.")
	}

	return m.assignmentPollDefaultGroups
}

func (m *Meeting) Assignments() []*Assignment {
	if _, ok := m.loadedRelations["assignment_ids"]; !ok {
		log.Panic().Msg("Tried to access Assignments relation of Meeting which was not loaded.")
	}

	return m.assignments
}

func (m *Meeting) ChatGroups() []*ChatGroup {
	if _, ok := m.loadedRelations["chat_group_ids"]; !ok {
		log.Panic().Msg("Tried to access ChatGroups relation of Meeting which was not loaded.")
	}

	return m.chatGroups
}

func (m *Meeting) ChatMessages() []*ChatMessage {
	if _, ok := m.loadedRelations["chat_message_ids"]; !ok {
		log.Panic().Msg("Tried to access ChatMessages relation of Meeting which was not loaded.")
	}

	return m.chatMessages
}

func (m *Meeting) Committee() Committee {
	if _, ok := m.loadedRelations["committee_id"]; !ok {
		log.Panic().Msg("Tried to access Committee relation of Meeting which was not loaded.")
	}

	return *m.committee
}

func (m *Meeting) DefaultGroup() Group {
	if _, ok := m.loadedRelations["default_group_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultGroup relation of Meeting which was not loaded.")
	}

	return *m.defaultGroup
}

func (m *Meeting) DefaultMeetingForCommittee() *Committee {
	if _, ok := m.loadedRelations["default_meeting_for_committee_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultMeetingForCommittee relation of Meeting which was not loaded.")
	}

	return m.defaultMeetingForCommittee
}

func (m *Meeting) DefaultProjectorAgendaItemLists() []*Projector {
	if _, ok := m.loadedRelations["default_projector_agenda_item_list_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorAgendaItemLists relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorAgendaItemLists
}

func (m *Meeting) DefaultProjectorAmendments() []*Projector {
	if _, ok := m.loadedRelations["default_projector_amendment_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorAmendments relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorAmendments
}

func (m *Meeting) DefaultProjectorAssignmentPolls() []*Projector {
	if _, ok := m.loadedRelations["default_projector_assignment_poll_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorAssignmentPolls relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorAssignmentPolls
}

func (m *Meeting) DefaultProjectorAssignments() []*Projector {
	if _, ok := m.loadedRelations["default_projector_assignment_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorAssignments relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorAssignments
}

func (m *Meeting) DefaultProjectorCountdowns() []*Projector {
	if _, ok := m.loadedRelations["default_projector_countdown_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorCountdowns relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorCountdowns
}

func (m *Meeting) DefaultProjectorCurrentListOfSpeakerss() []*Projector {
	if _, ok := m.loadedRelations["default_projector_current_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorCurrentListOfSpeakerss relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorCurrentListOfSpeakerss
}

func (m *Meeting) DefaultProjectorListOfSpeakerss() []*Projector {
	if _, ok := m.loadedRelations["default_projector_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorListOfSpeakerss relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorListOfSpeakerss
}

func (m *Meeting) DefaultProjectorMediafiles() []*Projector {
	if _, ok := m.loadedRelations["default_projector_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMediafiles relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMediafiles
}

func (m *Meeting) DefaultProjectorMessages() []*Projector {
	if _, ok := m.loadedRelations["default_projector_message_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMessages relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMessages
}

func (m *Meeting) DefaultProjectorMotionBlocks() []*Projector {
	if _, ok := m.loadedRelations["default_projector_motion_block_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMotionBlocks relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMotionBlocks
}

func (m *Meeting) DefaultProjectorMotionPolls() []*Projector {
	if _, ok := m.loadedRelations["default_projector_motion_poll_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMotionPolls relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMotionPolls
}

func (m *Meeting) DefaultProjectorMotions() []*Projector {
	if _, ok := m.loadedRelations["default_projector_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMotions relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMotions
}

func (m *Meeting) DefaultProjectorPolls() []*Projector {
	if _, ok := m.loadedRelations["default_projector_poll_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorPolls relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorPolls
}

func (m *Meeting) DefaultProjectorTopics() []*Projector {
	if _, ok := m.loadedRelations["default_projector_topic_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorTopics relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorTopics
}

func (m *Meeting) FontBold() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_bold_id"]; !ok {
		log.Panic().Msg("Tried to access FontBold relation of Meeting which was not loaded.")
	}

	return m.fontBold
}

func (m *Meeting) FontBoldItalic() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_bold_italic_id"]; !ok {
		log.Panic().Msg("Tried to access FontBoldItalic relation of Meeting which was not loaded.")
	}

	return m.fontBoldItalic
}

func (m *Meeting) FontChyronSpeakerName() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_chyron_speaker_name_id"]; !ok {
		log.Panic().Msg("Tried to access FontChyronSpeakerName relation of Meeting which was not loaded.")
	}

	return m.fontChyronSpeakerName
}

func (m *Meeting) FontItalic() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_italic_id"]; !ok {
		log.Panic().Msg("Tried to access FontItalic relation of Meeting which was not loaded.")
	}

	return m.fontItalic
}

func (m *Meeting) FontMonospace() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_monospace_id"]; !ok {
		log.Panic().Msg("Tried to access FontMonospace relation of Meeting which was not loaded.")
	}

	return m.fontMonospace
}

func (m *Meeting) FontProjectorH1() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_projector_h1_id"]; !ok {
		log.Panic().Msg("Tried to access FontProjectorH1 relation of Meeting which was not loaded.")
	}

	return m.fontProjectorH1
}

func (m *Meeting) FontProjectorH2() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_projector_h2_id"]; !ok {
		log.Panic().Msg("Tried to access FontProjectorH2 relation of Meeting which was not loaded.")
	}

	return m.fontProjectorH2
}

func (m *Meeting) FontRegular() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_regular_id"]; !ok {
		log.Panic().Msg("Tried to access FontRegular relation of Meeting which was not loaded.")
	}

	return m.fontRegular
}

func (m *Meeting) ForwardedMotions() []*Motion {
	if _, ok := m.loadedRelations["forwarded_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access ForwardedMotions relation of Meeting which was not loaded.")
	}

	return m.forwardedMotions
}

func (m *Meeting) Groups() []*Group {
	if _, ok := m.loadedRelations["group_ids"]; !ok {
		log.Panic().Msg("Tried to access Groups relation of Meeting which was not loaded.")
	}

	return m.groups
}

func (m *Meeting) IsActiveInOrganization() *Organization {
	if _, ok := m.loadedRelations["is_active_in_organization_id"]; !ok {
		log.Panic().Msg("Tried to access IsActiveInOrganization relation of Meeting which was not loaded.")
	}

	return m.isActiveInOrganization
}

func (m *Meeting) IsArchivedInOrganization() *Organization {
	if _, ok := m.loadedRelations["is_archived_in_organization_id"]; !ok {
		log.Panic().Msg("Tried to access IsArchivedInOrganization relation of Meeting which was not loaded.")
	}

	return m.isArchivedInOrganization
}

func (m *Meeting) ListOfSpeakersCountdown() *ProjectorCountdown {
	if _, ok := m.loadedRelations["list_of_speakers_countdown_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakersCountdown relation of Meeting which was not loaded.")
	}

	return m.listOfSpeakersCountdown
}

func (m *Meeting) ListOfSpeakerss() []*ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakerss relation of Meeting which was not loaded.")
	}

	return m.listOfSpeakerss
}

func (m *Meeting) LogoPdfBallotPaper() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_ballot_paper_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfBallotPaper relation of Meeting which was not loaded.")
	}

	return m.logoPdfBallotPaper
}

func (m *Meeting) LogoPdfFooterL() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_footer_l_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfFooterL relation of Meeting which was not loaded.")
	}

	return m.logoPdfFooterL
}

func (m *Meeting) LogoPdfFooterR() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_footer_r_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfFooterR relation of Meeting which was not loaded.")
	}

	return m.logoPdfFooterR
}

func (m *Meeting) LogoPdfHeaderL() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_header_l_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfHeaderL relation of Meeting which was not loaded.")
	}

	return m.logoPdfHeaderL
}

func (m *Meeting) LogoPdfHeaderR() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_header_r_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfHeaderR relation of Meeting which was not loaded.")
	}

	return m.logoPdfHeaderR
}

func (m *Meeting) LogoProjectorHeader() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_projector_header_id"]; !ok {
		log.Panic().Msg("Tried to access LogoProjectorHeader relation of Meeting which was not loaded.")
	}

	return m.logoProjectorHeader
}

func (m *Meeting) LogoProjectorMain() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_projector_main_id"]; !ok {
		log.Panic().Msg("Tried to access LogoProjectorMain relation of Meeting which was not loaded.")
	}

	return m.logoProjectorMain
}

func (m *Meeting) LogoWebHeader() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_web_header_id"]; !ok {
		log.Panic().Msg("Tried to access LogoWebHeader relation of Meeting which was not loaded.")
	}

	return m.logoWebHeader
}

func (m *Meeting) Mediafiles() []*Mediafile {
	if _, ok := m.loadedRelations["mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access Mediafiles relation of Meeting which was not loaded.")
	}

	return m.mediafiles
}

func (m *Meeting) MeetingMediafiles() []*MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafiles relation of Meeting which was not loaded.")
	}

	return m.meetingMediafiles
}

func (m *Meeting) MeetingUsers() []*MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingUsers relation of Meeting which was not loaded.")
	}

	return m.meetingUsers
}

func (m *Meeting) MotionBlocks() []*MotionBlock {
	if _, ok := m.loadedRelations["motion_block_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionBlocks relation of Meeting which was not loaded.")
	}

	return m.motionBlocks
}

func (m *Meeting) MotionCategorys() []*MotionCategory {
	if _, ok := m.loadedRelations["motion_category_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionCategorys relation of Meeting which was not loaded.")
	}

	return m.motionCategorys
}

func (m *Meeting) MotionChangeRecommendations() []*MotionChangeRecommendation {
	if _, ok := m.loadedRelations["motion_change_recommendation_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionChangeRecommendations relation of Meeting which was not loaded.")
	}

	return m.motionChangeRecommendations
}

func (m *Meeting) MotionCommentSections() []*MotionCommentSection {
	if _, ok := m.loadedRelations["motion_comment_section_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionCommentSections relation of Meeting which was not loaded.")
	}

	return m.motionCommentSections
}

func (m *Meeting) MotionComments() []*MotionComment {
	if _, ok := m.loadedRelations["motion_comment_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionComments relation of Meeting which was not loaded.")
	}

	return m.motionComments
}

func (m *Meeting) MotionEditors() []*MotionEditor {
	if _, ok := m.loadedRelations["motion_editor_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionEditors relation of Meeting which was not loaded.")
	}

	return m.motionEditors
}

func (m *Meeting) MotionPollDefaultGroups() []*Group {
	if _, ok := m.loadedRelations["motion_poll_default_group_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionPollDefaultGroups relation of Meeting which was not loaded.")
	}

	return m.motionPollDefaultGroups
}

func (m *Meeting) MotionStates() []*MotionState {
	if _, ok := m.loadedRelations["motion_state_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionStates relation of Meeting which was not loaded.")
	}

	return m.motionStates
}

func (m *Meeting) MotionSubmitters() []*MotionSubmitter {
	if _, ok := m.loadedRelations["motion_submitter_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionSubmitters relation of Meeting which was not loaded.")
	}

	return m.motionSubmitters
}

func (m *Meeting) MotionWorkflows() []*MotionWorkflow {
	if _, ok := m.loadedRelations["motion_workflow_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionWorkflows relation of Meeting which was not loaded.")
	}

	return m.motionWorkflows
}

func (m *Meeting) MotionWorkingGroupSpeakers() []*MotionWorkingGroupSpeaker {
	if _, ok := m.loadedRelations["motion_working_group_speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionWorkingGroupSpeakers relation of Meeting which was not loaded.")
	}

	return m.motionWorkingGroupSpeakers
}

func (m *Meeting) Motions() []*Motion {
	if _, ok := m.loadedRelations["motion_ids"]; !ok {
		log.Panic().Msg("Tried to access Motions relation of Meeting which was not loaded.")
	}

	return m.motions
}

func (m *Meeting) MotionsDefaultAmendmentWorkflow() MotionWorkflow {
	if _, ok := m.loadedRelations["motions_default_amendment_workflow_id"]; !ok {
		log.Panic().Msg("Tried to access MotionsDefaultAmendmentWorkflow relation of Meeting which was not loaded.")
	}

	return *m.motionsDefaultAmendmentWorkflow
}

func (m *Meeting) MotionsDefaultWorkflow() MotionWorkflow {
	if _, ok := m.loadedRelations["motions_default_workflow_id"]; !ok {
		log.Panic().Msg("Tried to access MotionsDefaultWorkflow relation of Meeting which was not loaded.")
	}

	return *m.motionsDefaultWorkflow
}

func (m *Meeting) Options() []*Option {
	if _, ok := m.loadedRelations["option_ids"]; !ok {
		log.Panic().Msg("Tried to access Options relation of Meeting which was not loaded.")
	}

	return m.options
}

func (m *Meeting) OrganizationTags() []*OrganizationTag {
	if _, ok := m.loadedRelations["organization_tag_ids"]; !ok {
		log.Panic().Msg("Tried to access OrganizationTags relation of Meeting which was not loaded.")
	}

	return m.organizationTags
}

func (m *Meeting) PersonalNotes() []*PersonalNote {
	if _, ok := m.loadedRelations["personal_note_ids"]; !ok {
		log.Panic().Msg("Tried to access PersonalNotes relation of Meeting which was not loaded.")
	}

	return m.personalNotes
}

func (m *Meeting) PointOfOrderCategorys() []*PointOfOrderCategory {
	if _, ok := m.loadedRelations["point_of_order_category_ids"]; !ok {
		log.Panic().Msg("Tried to access PointOfOrderCategorys relation of Meeting which was not loaded.")
	}

	return m.pointOfOrderCategorys
}

func (m *Meeting) PollCandidateLists() []*PollCandidateList {
	if _, ok := m.loadedRelations["poll_candidate_list_ids"]; !ok {
		log.Panic().Msg("Tried to access PollCandidateLists relation of Meeting which was not loaded.")
	}

	return m.pollCandidateLists
}

func (m *Meeting) PollCandidates() []*PollCandidate {
	if _, ok := m.loadedRelations["poll_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access PollCandidates relation of Meeting which was not loaded.")
	}

	return m.pollCandidates
}

func (m *Meeting) PollCountdown() *ProjectorCountdown {
	if _, ok := m.loadedRelations["poll_countdown_id"]; !ok {
		log.Panic().Msg("Tried to access PollCountdown relation of Meeting which was not loaded.")
	}

	return m.pollCountdown
}

func (m *Meeting) PollDefaultGroups() []*Group {
	if _, ok := m.loadedRelations["poll_default_group_ids"]; !ok {
		log.Panic().Msg("Tried to access PollDefaultGroups relation of Meeting which was not loaded.")
	}

	return m.pollDefaultGroups
}

func (m *Meeting) Polls() []*Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Meeting which was not loaded.")
	}

	return m.polls
}

func (m *Meeting) PresentUsers() []*User {
	if _, ok := m.loadedRelations["present_user_ids"]; !ok {
		log.Panic().Msg("Tried to access PresentUsers relation of Meeting which was not loaded.")
	}

	return m.presentUsers
}

func (m *Meeting) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of Meeting which was not loaded.")
	}

	return m.projections
}

func (m *Meeting) ProjectorCountdowns() []*ProjectorCountdown {
	if _, ok := m.loadedRelations["projector_countdown_ids"]; !ok {
		log.Panic().Msg("Tried to access ProjectorCountdowns relation of Meeting which was not loaded.")
	}

	return m.projectorCountdowns
}

func (m *Meeting) ProjectorMessages() []*ProjectorMessage {
	if _, ok := m.loadedRelations["projector_message_ids"]; !ok {
		log.Panic().Msg("Tried to access ProjectorMessages relation of Meeting which was not loaded.")
	}

	return m.projectorMessages
}

func (m *Meeting) Projectors() []*Projector {
	if _, ok := m.loadedRelations["projector_ids"]; !ok {
		log.Panic().Msg("Tried to access Projectors relation of Meeting which was not loaded.")
	}

	return m.projectors
}

func (m *Meeting) ReferenceProjector() Projector {
	if _, ok := m.loadedRelations["reference_projector_id"]; !ok {
		log.Panic().Msg("Tried to access ReferenceProjector relation of Meeting which was not loaded.")
	}

	return *m.referenceProjector
}

func (m *Meeting) Speakers() []*Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of Meeting which was not loaded.")
	}

	return m.speakers
}

func (m *Meeting) StructureLevelListOfSpeakerss() []*StructureLevelListOfSpeakers {
	if _, ok := m.loadedRelations["structure_level_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevelListOfSpeakerss relation of Meeting which was not loaded.")
	}

	return m.structureLevelListOfSpeakerss
}

func (m *Meeting) StructureLevels() []*StructureLevel {
	if _, ok := m.loadedRelations["structure_level_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevels relation of Meeting which was not loaded.")
	}

	return m.structureLevels
}

func (m *Meeting) Tags() []*Tag {
	if _, ok := m.loadedRelations["tag_ids"]; !ok {
		log.Panic().Msg("Tried to access Tags relation of Meeting which was not loaded.")
	}

	return m.tags
}

func (m *Meeting) TemplateForOrganization() *Organization {
	if _, ok := m.loadedRelations["template_for_organization_id"]; !ok {
		log.Panic().Msg("Tried to access TemplateForOrganization relation of Meeting which was not loaded.")
	}

	return m.templateForOrganization
}

func (m *Meeting) TopicPollDefaultGroups() []*Group {
	if _, ok := m.loadedRelations["topic_poll_default_group_ids"]; !ok {
		log.Panic().Msg("Tried to access TopicPollDefaultGroups relation of Meeting which was not loaded.")
	}

	return m.topicPollDefaultGroups
}

func (m *Meeting) Topics() []*Topic {
	if _, ok := m.loadedRelations["topic_ids"]; !ok {
		log.Panic().Msg("Tried to access Topics relation of Meeting which was not loaded.")
	}

	return m.topics
}

func (m *Meeting) Votes() []*Vote {
	if _, ok := m.loadedRelations["vote_ids"]; !ok {
		log.Panic().Msg("Tried to access Votes relation of Meeting which was not loaded.")
	}

	return m.votes
}

func (m *Meeting) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "admin_group_id":
		return m.adminGroup.GetRelatedModelsAccessor()
	case "agenda_item_ids":
		for _, r := range m.agendaItems {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "all_projection_ids":
		for _, r := range m.allProjections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "anonymous_group_id":
		return m.anonymousGroup.GetRelatedModelsAccessor()
	case "assignment_candidate_ids":
		for _, r := range m.assignmentCandidates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "assignment_poll_default_group_ids":
		for _, r := range m.assignmentPollDefaultGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "assignment_ids":
		for _, r := range m.assignments {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "chat_group_ids":
		for _, r := range m.chatGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "chat_message_ids":
		for _, r := range m.chatMessages {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "committee_id":
		return m.committee.GetRelatedModelsAccessor()
	case "default_group_id":
		return m.defaultGroup.GetRelatedModelsAccessor()
	case "default_meeting_for_committee_id":
		return m.defaultMeetingForCommittee.GetRelatedModelsAccessor()
	case "default_projector_agenda_item_list_ids":
		for _, r := range m.defaultProjectorAgendaItemLists {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_amendment_ids":
		for _, r := range m.defaultProjectorAmendments {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_assignment_poll_ids":
		for _, r := range m.defaultProjectorAssignmentPolls {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_assignment_ids":
		for _, r := range m.defaultProjectorAssignments {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_countdown_ids":
		for _, r := range m.defaultProjectorCountdowns {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_current_list_of_speakers_ids":
		for _, r := range m.defaultProjectorCurrentListOfSpeakerss {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_list_of_speakers_ids":
		for _, r := range m.defaultProjectorListOfSpeakerss {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_mediafile_ids":
		for _, r := range m.defaultProjectorMediafiles {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_message_ids":
		for _, r := range m.defaultProjectorMessages {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_motion_block_ids":
		for _, r := range m.defaultProjectorMotionBlocks {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_motion_poll_ids":
		for _, r := range m.defaultProjectorMotionPolls {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_motion_ids":
		for _, r := range m.defaultProjectorMotions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_poll_ids":
		for _, r := range m.defaultProjectorPolls {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "default_projector_topic_ids":
		for _, r := range m.defaultProjectorTopics {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "font_bold_id":
		return m.fontBold.GetRelatedModelsAccessor()
	case "font_bold_italic_id":
		return m.fontBoldItalic.GetRelatedModelsAccessor()
	case "font_chyron_speaker_name_id":
		return m.fontChyronSpeakerName.GetRelatedModelsAccessor()
	case "font_italic_id":
		return m.fontItalic.GetRelatedModelsAccessor()
	case "font_monospace_id":
		return m.fontMonospace.GetRelatedModelsAccessor()
	case "font_projector_h1_id":
		return m.fontProjectorH1.GetRelatedModelsAccessor()
	case "font_projector_h2_id":
		return m.fontProjectorH2.GetRelatedModelsAccessor()
	case "font_regular_id":
		return m.fontRegular.GetRelatedModelsAccessor()
	case "forwarded_motion_ids":
		for _, r := range m.forwardedMotions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "group_ids":
		for _, r := range m.groups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "is_active_in_organization_id":
		return m.isActiveInOrganization.GetRelatedModelsAccessor()
	case "is_archived_in_organization_id":
		return m.isArchivedInOrganization.GetRelatedModelsAccessor()
	case "list_of_speakers_countdown_id":
		return m.listOfSpeakersCountdown.GetRelatedModelsAccessor()
	case "list_of_speakers_ids":
		for _, r := range m.listOfSpeakerss {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "logo_pdf_ballot_paper_id":
		return m.logoPdfBallotPaper.GetRelatedModelsAccessor()
	case "logo_pdf_footer_l_id":
		return m.logoPdfFooterL.GetRelatedModelsAccessor()
	case "logo_pdf_footer_r_id":
		return m.logoPdfFooterR.GetRelatedModelsAccessor()
	case "logo_pdf_header_l_id":
		return m.logoPdfHeaderL.GetRelatedModelsAccessor()
	case "logo_pdf_header_r_id":
		return m.logoPdfHeaderR.GetRelatedModelsAccessor()
	case "logo_projector_header_id":
		return m.logoProjectorHeader.GetRelatedModelsAccessor()
	case "logo_projector_main_id":
		return m.logoProjectorMain.GetRelatedModelsAccessor()
	case "logo_web_header_id":
		return m.logoWebHeader.GetRelatedModelsAccessor()
	case "mediafile_ids":
		for _, r := range m.mediafiles {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_mediafile_ids":
		for _, r := range m.meetingMediafiles {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_user_ids":
		for _, r := range m.meetingUsers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_block_ids":
		for _, r := range m.motionBlocks {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_category_ids":
		for _, r := range m.motionCategorys {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_change_recommendation_ids":
		for _, r := range m.motionChangeRecommendations {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_comment_section_ids":
		for _, r := range m.motionCommentSections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_comment_ids":
		for _, r := range m.motionComments {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_editor_ids":
		for _, r := range m.motionEditors {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_poll_default_group_ids":
		for _, r := range m.motionPollDefaultGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_state_ids":
		for _, r := range m.motionStates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_submitter_ids":
		for _, r := range m.motionSubmitters {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_workflow_ids":
		for _, r := range m.motionWorkflows {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_working_group_speaker_ids":
		for _, r := range m.motionWorkingGroupSpeakers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_ids":
		for _, r := range m.motions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motions_default_amendment_workflow_id":
		return m.motionsDefaultAmendmentWorkflow.GetRelatedModelsAccessor()
	case "motions_default_workflow_id":
		return m.motionsDefaultWorkflow.GetRelatedModelsAccessor()
	case "option_ids":
		for _, r := range m.options {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "organization_tag_ids":
		for _, r := range m.organizationTags {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "personal_note_ids":
		for _, r := range m.personalNotes {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "point_of_order_category_ids":
		for _, r := range m.pointOfOrderCategorys {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "poll_candidate_list_ids":
		for _, r := range m.pollCandidateLists {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "poll_candidate_ids":
		for _, r := range m.pollCandidates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "poll_countdown_id":
		return m.pollCountdown.GetRelatedModelsAccessor()
	case "poll_default_group_ids":
		for _, r := range m.pollDefaultGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "poll_ids":
		for _, r := range m.polls {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "present_user_ids":
		for _, r := range m.presentUsers {
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
	case "projector_countdown_ids":
		for _, r := range m.projectorCountdowns {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "projector_message_ids":
		for _, r := range m.projectorMessages {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "projector_ids":
		for _, r := range m.projectors {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "reference_projector_id":
		return m.referenceProjector.GetRelatedModelsAccessor()
	case "speaker_ids":
		for _, r := range m.speakers {
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
	case "structure_level_ids":
		for _, r := range m.structureLevels {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "tag_ids":
		for _, r := range m.tags {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "template_for_organization_id":
		return m.templateForOrganization.GetRelatedModelsAccessor()
	case "topic_poll_default_group_ids":
		for _, r := range m.topicPollDefaultGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "topic_ids":
		for _, r := range m.topics {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "vote_ids":
		for _, r := range m.votes {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Meeting) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "admin_group_id":
			m.adminGroup = content.(*Group)
		case "agenda_item_ids":
			m.agendaItems = content.([]*AgendaItem)
		case "all_projection_ids":
			m.allProjections = content.([]*Projection)
		case "anonymous_group_id":
			m.anonymousGroup = content.(*Group)
		case "assignment_candidate_ids":
			m.assignmentCandidates = content.([]*AssignmentCandidate)
		case "assignment_poll_default_group_ids":
			m.assignmentPollDefaultGroups = content.([]*Group)
		case "assignment_ids":
			m.assignments = content.([]*Assignment)
		case "chat_group_ids":
			m.chatGroups = content.([]*ChatGroup)
		case "chat_message_ids":
			m.chatMessages = content.([]*ChatMessage)
		case "committee_id":
			m.committee = content.(*Committee)
		case "default_group_id":
			m.defaultGroup = content.(*Group)
		case "default_meeting_for_committee_id":
			m.defaultMeetingForCommittee = content.(*Committee)
		case "default_projector_agenda_item_list_ids":
			m.defaultProjectorAgendaItemLists = content.([]*Projector)
		case "default_projector_amendment_ids":
			m.defaultProjectorAmendments = content.([]*Projector)
		case "default_projector_assignment_poll_ids":
			m.defaultProjectorAssignmentPolls = content.([]*Projector)
		case "default_projector_assignment_ids":
			m.defaultProjectorAssignments = content.([]*Projector)
		case "default_projector_countdown_ids":
			m.defaultProjectorCountdowns = content.([]*Projector)
		case "default_projector_current_list_of_speakers_ids":
			m.defaultProjectorCurrentListOfSpeakerss = content.([]*Projector)
		case "default_projector_list_of_speakers_ids":
			m.defaultProjectorListOfSpeakerss = content.([]*Projector)
		case "default_projector_mediafile_ids":
			m.defaultProjectorMediafiles = content.([]*Projector)
		case "default_projector_message_ids":
			m.defaultProjectorMessages = content.([]*Projector)
		case "default_projector_motion_block_ids":
			m.defaultProjectorMotionBlocks = content.([]*Projector)
		case "default_projector_motion_poll_ids":
			m.defaultProjectorMotionPolls = content.([]*Projector)
		case "default_projector_motion_ids":
			m.defaultProjectorMotions = content.([]*Projector)
		case "default_projector_poll_ids":
			m.defaultProjectorPolls = content.([]*Projector)
		case "default_projector_topic_ids":
			m.defaultProjectorTopics = content.([]*Projector)
		case "font_bold_id":
			m.fontBold = content.(*MeetingMediafile)
		case "font_bold_italic_id":
			m.fontBoldItalic = content.(*MeetingMediafile)
		case "font_chyron_speaker_name_id":
			m.fontChyronSpeakerName = content.(*MeetingMediafile)
		case "font_italic_id":
			m.fontItalic = content.(*MeetingMediafile)
		case "font_monospace_id":
			m.fontMonospace = content.(*MeetingMediafile)
		case "font_projector_h1_id":
			m.fontProjectorH1 = content.(*MeetingMediafile)
		case "font_projector_h2_id":
			m.fontProjectorH2 = content.(*MeetingMediafile)
		case "font_regular_id":
			m.fontRegular = content.(*MeetingMediafile)
		case "forwarded_motion_ids":
			m.forwardedMotions = content.([]*Motion)
		case "group_ids":
			m.groups = content.([]*Group)
		case "is_active_in_organization_id":
			m.isActiveInOrganization = content.(*Organization)
		case "is_archived_in_organization_id":
			m.isArchivedInOrganization = content.(*Organization)
		case "list_of_speakers_countdown_id":
			m.listOfSpeakersCountdown = content.(*ProjectorCountdown)
		case "list_of_speakers_ids":
			m.listOfSpeakerss = content.([]*ListOfSpeakers)
		case "logo_pdf_ballot_paper_id":
			m.logoPdfBallotPaper = content.(*MeetingMediafile)
		case "logo_pdf_footer_l_id":
			m.logoPdfFooterL = content.(*MeetingMediafile)
		case "logo_pdf_footer_r_id":
			m.logoPdfFooterR = content.(*MeetingMediafile)
		case "logo_pdf_header_l_id":
			m.logoPdfHeaderL = content.(*MeetingMediafile)
		case "logo_pdf_header_r_id":
			m.logoPdfHeaderR = content.(*MeetingMediafile)
		case "logo_projector_header_id":
			m.logoProjectorHeader = content.(*MeetingMediafile)
		case "logo_projector_main_id":
			m.logoProjectorMain = content.(*MeetingMediafile)
		case "logo_web_header_id":
			m.logoWebHeader = content.(*MeetingMediafile)
		case "mediafile_ids":
			m.mediafiles = content.([]*Mediafile)
		case "meeting_mediafile_ids":
			m.meetingMediafiles = content.([]*MeetingMediafile)
		case "meeting_user_ids":
			m.meetingUsers = content.([]*MeetingUser)
		case "motion_block_ids":
			m.motionBlocks = content.([]*MotionBlock)
		case "motion_category_ids":
			m.motionCategorys = content.([]*MotionCategory)
		case "motion_change_recommendation_ids":
			m.motionChangeRecommendations = content.([]*MotionChangeRecommendation)
		case "motion_comment_section_ids":
			m.motionCommentSections = content.([]*MotionCommentSection)
		case "motion_comment_ids":
			m.motionComments = content.([]*MotionComment)
		case "motion_editor_ids":
			m.motionEditors = content.([]*MotionEditor)
		case "motion_poll_default_group_ids":
			m.motionPollDefaultGroups = content.([]*Group)
		case "motion_state_ids":
			m.motionStates = content.([]*MotionState)
		case "motion_submitter_ids":
			m.motionSubmitters = content.([]*MotionSubmitter)
		case "motion_workflow_ids":
			m.motionWorkflows = content.([]*MotionWorkflow)
		case "motion_working_group_speaker_ids":
			m.motionWorkingGroupSpeakers = content.([]*MotionWorkingGroupSpeaker)
		case "motion_ids":
			m.motions = content.([]*Motion)
		case "motions_default_amendment_workflow_id":
			m.motionsDefaultAmendmentWorkflow = content.(*MotionWorkflow)
		case "motions_default_workflow_id":
			m.motionsDefaultWorkflow = content.(*MotionWorkflow)
		case "option_ids":
			m.options = content.([]*Option)
		case "organization_tag_ids":
			m.organizationTags = content.([]*OrganizationTag)
		case "personal_note_ids":
			m.personalNotes = content.([]*PersonalNote)
		case "point_of_order_category_ids":
			m.pointOfOrderCategorys = content.([]*PointOfOrderCategory)
		case "poll_candidate_list_ids":
			m.pollCandidateLists = content.([]*PollCandidateList)
		case "poll_candidate_ids":
			m.pollCandidates = content.([]*PollCandidate)
		case "poll_countdown_id":
			m.pollCountdown = content.(*ProjectorCountdown)
		case "poll_default_group_ids":
			m.pollDefaultGroups = content.([]*Group)
		case "poll_ids":
			m.polls = content.([]*Poll)
		case "present_user_ids":
			m.presentUsers = content.([]*User)
		case "projection_ids":
			m.projections = content.([]*Projection)
		case "projector_countdown_ids":
			m.projectorCountdowns = content.([]*ProjectorCountdown)
		case "projector_message_ids":
			m.projectorMessages = content.([]*ProjectorMessage)
		case "projector_ids":
			m.projectors = content.([]*Projector)
		case "reference_projector_id":
			m.referenceProjector = content.(*Projector)
		case "speaker_ids":
			m.speakers = content.([]*Speaker)
		case "structure_level_list_of_speakers_ids":
			m.structureLevelListOfSpeakerss = content.([]*StructureLevelListOfSpeakers)
		case "structure_level_ids":
			m.structureLevels = content.([]*StructureLevel)
		case "tag_ids":
			m.tags = content.([]*Tag)
		case "template_for_organization_id":
			m.templateForOrganization = content.(*Organization)
		case "topic_poll_default_group_ids":
			m.topicPollDefaultGroups = content.([]*Group)
		case "topic_ids":
			m.topics = content.([]*Topic)
		case "vote_ids":
			m.votes = content.([]*Vote)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Meeting) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "admin_group_id":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.adminGroup = &entry

		result = entry.GetRelatedModelsAccessor()
	case "agenda_item_ids":
		var entry AgendaItem
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.agendaItems = append(m.agendaItems, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "all_projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.allProjections = append(m.allProjections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "anonymous_group_id":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.anonymousGroup = &entry

		result = entry.GetRelatedModelsAccessor()
	case "assignment_candidate_ids":
		var entry AssignmentCandidate
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.assignmentCandidates = append(m.assignmentCandidates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "assignment_poll_default_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.assignmentPollDefaultGroups = append(m.assignmentPollDefaultGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "assignment_ids":
		var entry Assignment
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.assignments = append(m.assignments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "chat_group_ids":
		var entry ChatGroup
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.chatGroups = append(m.chatGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "chat_message_ids":
		var entry ChatMessage
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.chatMessages = append(m.chatMessages, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "committee_id":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.committee = &entry

		result = entry.GetRelatedModelsAccessor()
	case "default_group_id":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultGroup = &entry

		result = entry.GetRelatedModelsAccessor()
	case "default_meeting_for_committee_id":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultMeetingForCommittee = &entry

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_agenda_item_list_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorAgendaItemLists = append(m.defaultProjectorAgendaItemLists, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_amendment_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorAmendments = append(m.defaultProjectorAmendments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_assignment_poll_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorAssignmentPolls = append(m.defaultProjectorAssignmentPolls, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_assignment_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorAssignments = append(m.defaultProjectorAssignments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_countdown_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorCountdowns = append(m.defaultProjectorCountdowns, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_current_list_of_speakers_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorCurrentListOfSpeakerss = append(m.defaultProjectorCurrentListOfSpeakerss, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_list_of_speakers_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorListOfSpeakerss = append(m.defaultProjectorListOfSpeakerss, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_mediafile_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMediafiles = append(m.defaultProjectorMediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_message_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMessages = append(m.defaultProjectorMessages, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_motion_block_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMotionBlocks = append(m.defaultProjectorMotionBlocks, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_motion_poll_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMotionPolls = append(m.defaultProjectorMotionPolls, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_motion_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMotions = append(m.defaultProjectorMotions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_poll_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorPolls = append(m.defaultProjectorPolls, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_topic_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorTopics = append(m.defaultProjectorTopics, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "font_bold_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontBold = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_bold_italic_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontBoldItalic = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_chyron_speaker_name_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontChyronSpeakerName = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_italic_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontItalic = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_monospace_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontMonospace = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_projector_h1_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontProjectorH1 = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_projector_h2_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontProjectorH2 = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_regular_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontRegular = &entry

		result = entry.GetRelatedModelsAccessor()
	case "forwarded_motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.forwardedMotions = append(m.forwardedMotions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.groups = append(m.groups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "is_active_in_organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.isActiveInOrganization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "is_archived_in_organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.isArchivedInOrganization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "list_of_speakers_countdown_id":
		var entry ProjectorCountdown
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.listOfSpeakersCountdown = &entry

		result = entry.GetRelatedModelsAccessor()
	case "list_of_speakers_ids":
		var entry ListOfSpeakers
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.listOfSpeakerss = append(m.listOfSpeakerss, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_ballot_paper_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfBallotPaper = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_footer_l_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfFooterL = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_footer_r_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfFooterR = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_header_l_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfHeaderL = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_header_r_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfHeaderR = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_projector_header_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoProjectorHeader = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_projector_main_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoProjectorMain = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_web_header_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoWebHeader = &entry

		result = entry.GetRelatedModelsAccessor()
	case "mediafile_ids":
		var entry Mediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.mediafiles = append(m.mediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_mediafile_ids":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingMediafiles = append(m.meetingMediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_user_ids":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUsers = append(m.meetingUsers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_block_ids":
		var entry MotionBlock
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionBlocks = append(m.motionBlocks, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_category_ids":
		var entry MotionCategory
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionCategorys = append(m.motionCategorys, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_change_recommendation_ids":
		var entry MotionChangeRecommendation
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionChangeRecommendations = append(m.motionChangeRecommendations, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_comment_section_ids":
		var entry MotionCommentSection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionCommentSections = append(m.motionCommentSections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_comment_ids":
		var entry MotionComment
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionComments = append(m.motionComments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_editor_ids":
		var entry MotionEditor
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionEditors = append(m.motionEditors, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_poll_default_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionPollDefaultGroups = append(m.motionPollDefaultGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_state_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionStates = append(m.motionStates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_submitter_ids":
		var entry MotionSubmitter
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionSubmitters = append(m.motionSubmitters, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_workflow_ids":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionWorkflows = append(m.motionWorkflows, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_working_group_speaker_ids":
		var entry MotionWorkingGroupSpeaker
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionWorkingGroupSpeakers = append(m.motionWorkingGroupSpeakers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motions = append(m.motions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motions_default_amendment_workflow_id":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionsDefaultAmendmentWorkflow = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motions_default_workflow_id":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionsDefaultWorkflow = &entry

		result = entry.GetRelatedModelsAccessor()
	case "option_ids":
		var entry Option
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.options = append(m.options, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "organization_tag_ids":
		var entry OrganizationTag
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.organizationTags = append(m.organizationTags, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "personal_note_ids":
		var entry PersonalNote
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.personalNotes = append(m.personalNotes, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "point_of_order_category_ids":
		var entry PointOfOrderCategory
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pointOfOrderCategorys = append(m.pointOfOrderCategorys, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_candidate_list_ids":
		var entry PollCandidateList
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCandidateLists = append(m.pollCandidateLists, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_candidate_ids":
		var entry PollCandidate
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCandidates = append(m.pollCandidates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_countdown_id":
		var entry ProjectorCountdown
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCountdown = &entry

		result = entry.GetRelatedModelsAccessor()
	case "poll_default_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollDefaultGroups = append(m.pollDefaultGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_ids":
		var entry Poll
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.polls = append(m.polls, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "present_user_ids":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.presentUsers = append(m.presentUsers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projections = append(m.projections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "projector_countdown_ids":
		var entry ProjectorCountdown
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projectorCountdowns = append(m.projectorCountdowns, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "projector_message_ids":
		var entry ProjectorMessage
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projectorMessages = append(m.projectorMessages, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "projector_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projectors = append(m.projectors, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "reference_projector_id":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.referenceProjector = &entry

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
	case "structure_level_ids":
		var entry StructureLevel
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.structureLevels = append(m.structureLevels, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "tag_ids":
		var entry Tag
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.tags = append(m.tags, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "template_for_organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.templateForOrganization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "topic_poll_default_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.topicPollDefaultGroups = append(m.topicPollDefaultGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "topic_ids":
		var entry Topic
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.topics = append(m.topics, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "vote_ids":
		var entry Vote
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.votes = append(m.votes, &entry)

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

func (m *Meeting) Get(field string) interface{} {
	switch field {
	case "admin_group_id":
		return m.AdminGroupID
	case "agenda_enable_numbering":
		return m.AgendaEnableNumbering
	case "agenda_item_creation":
		return m.AgendaItemCreation
	case "agenda_item_ids":
		return m.AgendaItemIDs
	case "agenda_new_items_default_visibility":
		return m.AgendaNewItemsDefaultVisibility
	case "agenda_number_prefix":
		return m.AgendaNumberPrefix
	case "agenda_numeral_system":
		return m.AgendaNumeralSystem
	case "agenda_show_internal_items_on_projector":
		return m.AgendaShowInternalItemsOnProjector
	case "agenda_show_subtitles":
		return m.AgendaShowSubtitles
	case "agenda_show_topic_navigation_on_detail_view":
		return m.AgendaShowTopicNavigationOnDetailView
	case "all_projection_ids":
		return m.AllProjectionIDs
	case "anonymous_group_id":
		return m.AnonymousGroupID
	case "applause_enable":
		return m.ApplauseEnable
	case "applause_max_amount":
		return m.ApplauseMaxAmount
	case "applause_min_amount":
		return m.ApplauseMinAmount
	case "applause_particle_image_url":
		return m.ApplauseParticleImageUrl
	case "applause_show_level":
		return m.ApplauseShowLevel
	case "applause_timeout":
		return m.ApplauseTimeout
	case "applause_type":
		return m.ApplauseType
	case "assignment_candidate_ids":
		return m.AssignmentCandidateIDs
	case "assignment_ids":
		return m.AssignmentIDs
	case "assignment_poll_add_candidates_to_list_of_speakers":
		return m.AssignmentPollAddCandidatesToListOfSpeakers
	case "assignment_poll_ballot_paper_number":
		return m.AssignmentPollBallotPaperNumber
	case "assignment_poll_ballot_paper_selection":
		return m.AssignmentPollBallotPaperSelection
	case "assignment_poll_default_backend":
		return m.AssignmentPollDefaultBackend
	case "assignment_poll_default_group_ids":
		return m.AssignmentPollDefaultGroupIDs
	case "assignment_poll_default_method":
		return m.AssignmentPollDefaultMethod
	case "assignment_poll_default_onehundred_percent_base":
		return m.AssignmentPollDefaultOnehundredPercentBase
	case "assignment_poll_default_type":
		return m.AssignmentPollDefaultType
	case "assignment_poll_enable_max_votes_per_option":
		return m.AssignmentPollEnableMaxVotesPerOption
	case "assignment_poll_sort_poll_result_by_votes":
		return m.AssignmentPollSortPollResultByVotes
	case "assignments_export_preamble":
		return m.AssignmentsExportPreamble
	case "assignments_export_title":
		return m.AssignmentsExportTitle
	case "chat_group_ids":
		return m.ChatGroupIDs
	case "chat_message_ids":
		return m.ChatMessageIDs
	case "committee_id":
		return m.CommitteeID
	case "conference_auto_connect":
		return m.ConferenceAutoConnect
	case "conference_auto_connect_next_speakers":
		return m.ConferenceAutoConnectNextSpeakers
	case "conference_enable_helpdesk":
		return m.ConferenceEnableHelpdesk
	case "conference_los_restriction":
		return m.ConferenceLosRestriction
	case "conference_open_microphone":
		return m.ConferenceOpenMicrophone
	case "conference_open_video":
		return m.ConferenceOpenVideo
	case "conference_show":
		return m.ConferenceShow
	case "conference_stream_poster_url":
		return m.ConferenceStreamPosterUrl
	case "conference_stream_url":
		return m.ConferenceStreamUrl
	case "custom_translations":
		return m.CustomTranslations
	case "default_group_id":
		return m.DefaultGroupID
	case "default_meeting_for_committee_id":
		return m.DefaultMeetingForCommitteeID
	case "default_projector_agenda_item_list_ids":
		return m.DefaultProjectorAgendaItemListIDs
	case "default_projector_amendment_ids":
		return m.DefaultProjectorAmendmentIDs
	case "default_projector_assignment_ids":
		return m.DefaultProjectorAssignmentIDs
	case "default_projector_assignment_poll_ids":
		return m.DefaultProjectorAssignmentPollIDs
	case "default_projector_countdown_ids":
		return m.DefaultProjectorCountdownIDs
	case "default_projector_current_list_of_speakers_ids":
		return m.DefaultProjectorCurrentListOfSpeakersIDs
	case "default_projector_list_of_speakers_ids":
		return m.DefaultProjectorListOfSpeakersIDs
	case "default_projector_mediafile_ids":
		return m.DefaultProjectorMediafileIDs
	case "default_projector_message_ids":
		return m.DefaultProjectorMessageIDs
	case "default_projector_motion_block_ids":
		return m.DefaultProjectorMotionBlockIDs
	case "default_projector_motion_ids":
		return m.DefaultProjectorMotionIDs
	case "default_projector_motion_poll_ids":
		return m.DefaultProjectorMotionPollIDs
	case "default_projector_poll_ids":
		return m.DefaultProjectorPollIDs
	case "default_projector_topic_ids":
		return m.DefaultProjectorTopicIDs
	case "description":
		return m.Description
	case "enable_anonymous":
		return m.EnableAnonymous
	case "end_time":
		return m.EndTime
	case "export_csv_encoding":
		return m.ExportCsvEncoding
	case "export_csv_separator":
		return m.ExportCsvSeparator
	case "export_pdf_fontsize":
		return m.ExportPdfFontsize
	case "export_pdf_line_height":
		return m.ExportPdfLineHeight
	case "export_pdf_page_margin_bottom":
		return m.ExportPdfPageMarginBottom
	case "export_pdf_page_margin_left":
		return m.ExportPdfPageMarginLeft
	case "export_pdf_page_margin_right":
		return m.ExportPdfPageMarginRight
	case "export_pdf_page_margin_top":
		return m.ExportPdfPageMarginTop
	case "export_pdf_pagenumber_alignment":
		return m.ExportPdfPagenumberAlignment
	case "export_pdf_pagesize":
		return m.ExportPdfPagesize
	case "external_id":
		return m.ExternalID
	case "font_bold_id":
		return m.FontBoldID
	case "font_bold_italic_id":
		return m.FontBoldItalicID
	case "font_chyron_speaker_name_id":
		return m.FontChyronSpeakerNameID
	case "font_italic_id":
		return m.FontItalicID
	case "font_monospace_id":
		return m.FontMonospaceID
	case "font_projector_h1_id":
		return m.FontProjectorH1ID
	case "font_projector_h2_id":
		return m.FontProjectorH2ID
	case "font_regular_id":
		return m.FontRegularID
	case "forwarded_motion_ids":
		return m.ForwardedMotionIDs
	case "group_ids":
		return m.GroupIDs
	case "id":
		return m.ID
	case "imported_at":
		return m.ImportedAt
	case "is_active_in_organization_id":
		return m.IsActiveInOrganizationID
	case "is_archived_in_organization_id":
		return m.IsArchivedInOrganizationID
	case "jitsi_domain":
		return m.JitsiDomain
	case "jitsi_room_name":
		return m.JitsiRoomName
	case "jitsi_room_password":
		return m.JitsiRoomPassword
	case "language":
		return m.Language
	case "list_of_speakers_allow_multiple_speakers":
		return m.ListOfSpeakersAllowMultipleSpeakers
	case "list_of_speakers_amount_last_on_projector":
		return m.ListOfSpeakersAmountLastOnProjector
	case "list_of_speakers_amount_next_on_projector":
		return m.ListOfSpeakersAmountNextOnProjector
	case "list_of_speakers_can_create_point_of_order_for_others":
		return m.ListOfSpeakersCanCreatePointOfOrderForOthers
	case "list_of_speakers_can_set_contribution_self":
		return m.ListOfSpeakersCanSetContributionSelf
	case "list_of_speakers_closing_disables_point_of_order":
		return m.ListOfSpeakersClosingDisablesPointOfOrder
	case "list_of_speakers_countdown_id":
		return m.ListOfSpeakersCountdownID
	case "list_of_speakers_couple_countdown":
		return m.ListOfSpeakersCoupleCountdown
	case "list_of_speakers_default_structure_level_time":
		return m.ListOfSpeakersDefaultStructureLevelTime
	case "list_of_speakers_enable_interposed_question":
		return m.ListOfSpeakersEnableInterposedQuestion
	case "list_of_speakers_enable_point_of_order_categories":
		return m.ListOfSpeakersEnablePointOfOrderCategories
	case "list_of_speakers_enable_point_of_order_speakers":
		return m.ListOfSpeakersEnablePointOfOrderSpeakers
	case "list_of_speakers_enable_pro_contra_speech":
		return m.ListOfSpeakersEnableProContraSpeech
	case "list_of_speakers_hide_contribution_count":
		return m.ListOfSpeakersHideContributionCount
	case "list_of_speakers_ids":
		return m.ListOfSpeakersIDs
	case "list_of_speakers_initially_closed":
		return m.ListOfSpeakersInitiallyClosed
	case "list_of_speakers_intervention_time":
		return m.ListOfSpeakersInterventionTime
	case "list_of_speakers_present_users_only":
		return m.ListOfSpeakersPresentUsersOnly
	case "list_of_speakers_show_amount_of_speakers_on_slide":
		return m.ListOfSpeakersShowAmountOfSpeakersOnSlide
	case "list_of_speakers_show_first_contribution":
		return m.ListOfSpeakersShowFirstContribution
	case "list_of_speakers_speaker_note_for_everyone":
		return m.ListOfSpeakersSpeakerNoteForEveryone
	case "location":
		return m.Location
	case "locked_from_inside":
		return m.LockedFromInside
	case "logo_pdf_ballot_paper_id":
		return m.LogoPdfBallotPaperID
	case "logo_pdf_footer_l_id":
		return m.LogoPdfFooterLID
	case "logo_pdf_footer_r_id":
		return m.LogoPdfFooterRID
	case "logo_pdf_header_l_id":
		return m.LogoPdfHeaderLID
	case "logo_pdf_header_r_id":
		return m.LogoPdfHeaderRID
	case "logo_projector_header_id":
		return m.LogoProjectorHeaderID
	case "logo_projector_main_id":
		return m.LogoProjectorMainID
	case "logo_web_header_id":
		return m.LogoWebHeaderID
	case "mediafile_ids":
		return m.MediafileIDs
	case "meeting_mediafile_ids":
		return m.MeetingMediafileIDs
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "motion_block_ids":
		return m.MotionBlockIDs
	case "motion_category_ids":
		return m.MotionCategoryIDs
	case "motion_change_recommendation_ids":
		return m.MotionChangeRecommendationIDs
	case "motion_comment_ids":
		return m.MotionCommentIDs
	case "motion_comment_section_ids":
		return m.MotionCommentSectionIDs
	case "motion_editor_ids":
		return m.MotionEditorIDs
	case "motion_ids":
		return m.MotionIDs
	case "motion_poll_ballot_paper_number":
		return m.MotionPollBallotPaperNumber
	case "motion_poll_ballot_paper_selection":
		return m.MotionPollBallotPaperSelection
	case "motion_poll_default_backend":
		return m.MotionPollDefaultBackend
	case "motion_poll_default_group_ids":
		return m.MotionPollDefaultGroupIDs
	case "motion_poll_default_method":
		return m.MotionPollDefaultMethod
	case "motion_poll_default_onehundred_percent_base":
		return m.MotionPollDefaultOnehundredPercentBase
	case "motion_poll_default_type":
		return m.MotionPollDefaultType
	case "motion_state_ids":
		return m.MotionStateIDs
	case "motion_submitter_ids":
		return m.MotionSubmitterIDs
	case "motion_workflow_ids":
		return m.MotionWorkflowIDs
	case "motion_working_group_speaker_ids":
		return m.MotionWorkingGroupSpeakerIDs
	case "motions_amendments_enabled":
		return m.MotionsAmendmentsEnabled
	case "motions_amendments_in_main_list":
		return m.MotionsAmendmentsInMainList
	case "motions_amendments_multiple_paragraphs":
		return m.MotionsAmendmentsMultipleParagraphs
	case "motions_amendments_of_amendments":
		return m.MotionsAmendmentsOfAmendments
	case "motions_amendments_prefix":
		return m.MotionsAmendmentsPrefix
	case "motions_amendments_text_mode":
		return m.MotionsAmendmentsTextMode
	case "motions_block_slide_columns":
		return m.MotionsBlockSlideColumns
	case "motions_create_enable_additional_submitter_text":
		return m.MotionsCreateEnableAdditionalSubmitterText
	case "motions_default_amendment_workflow_id":
		return m.MotionsDefaultAmendmentWorkflowID
	case "motions_default_line_numbering":
		return m.MotionsDefaultLineNumbering
	case "motions_default_sorting":
		return m.MotionsDefaultSorting
	case "motions_default_workflow_id":
		return m.MotionsDefaultWorkflowID
	case "motions_enable_editor":
		return m.MotionsEnableEditor
	case "motions_enable_reason_on_projector":
		return m.MotionsEnableReasonOnProjector
	case "motions_enable_recommendation_on_projector":
		return m.MotionsEnableRecommendationOnProjector
	case "motions_enable_sidebox_on_projector":
		return m.MotionsEnableSideboxOnProjector
	case "motions_enable_text_on_projector":
		return m.MotionsEnableTextOnProjector
	case "motions_enable_working_group_speaker":
		return m.MotionsEnableWorkingGroupSpeaker
	case "motions_export_follow_recommendation":
		return m.MotionsExportFollowRecommendation
	case "motions_export_preamble":
		return m.MotionsExportPreamble
	case "motions_export_submitter_recommendation":
		return m.MotionsExportSubmitterRecommendation
	case "motions_export_title":
		return m.MotionsExportTitle
	case "motions_hide_metadata_background":
		return m.MotionsHideMetadataBackground
	case "motions_line_length":
		return m.MotionsLineLength
	case "motions_number_min_digits":
		return m.MotionsNumberMinDigits
	case "motions_number_type":
		return m.MotionsNumberType
	case "motions_number_with_blank":
		return m.MotionsNumberWithBlank
	case "motions_preamble":
		return m.MotionsPreamble
	case "motions_reason_required":
		return m.MotionsReasonRequired
	case "motions_recommendation_text_mode":
		return m.MotionsRecommendationTextMode
	case "motions_recommendations_by":
		return m.MotionsRecommendationsBy
	case "motions_show_referring_motions":
		return m.MotionsShowReferringMotions
	case "motions_show_sequential_number":
		return m.MotionsShowSequentialNumber
	case "motions_supporters_min_amount":
		return m.MotionsSupportersMinAmount
	case "name":
		return m.Name
	case "option_ids":
		return m.OptionIDs
	case "organization_tag_ids":
		return m.OrganizationTagIDs
	case "personal_note_ids":
		return m.PersonalNoteIDs
	case "point_of_order_category_ids":
		return m.PointOfOrderCategoryIDs
	case "poll_ballot_paper_number":
		return m.PollBallotPaperNumber
	case "poll_ballot_paper_selection":
		return m.PollBallotPaperSelection
	case "poll_candidate_ids":
		return m.PollCandidateIDs
	case "poll_candidate_list_ids":
		return m.PollCandidateListIDs
	case "poll_countdown_id":
		return m.PollCountdownID
	case "poll_couple_countdown":
		return m.PollCoupleCountdown
	case "poll_default_backend":
		return m.PollDefaultBackend
	case "poll_default_group_ids":
		return m.PollDefaultGroupIDs
	case "poll_default_method":
		return m.PollDefaultMethod
	case "poll_default_onehundred_percent_base":
		return m.PollDefaultOnehundredPercentBase
	case "poll_default_type":
		return m.PollDefaultType
	case "poll_ids":
		return m.PollIDs
	case "poll_sort_poll_result_by_votes":
		return m.PollSortPollResultByVotes
	case "present_user_ids":
		return m.PresentUserIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "projector_countdown_default_time":
		return m.ProjectorCountdownDefaultTime
	case "projector_countdown_ids":
		return m.ProjectorCountdownIDs
	case "projector_countdown_warning_time":
		return m.ProjectorCountdownWarningTime
	case "projector_ids":
		return m.ProjectorIDs
	case "projector_message_ids":
		return m.ProjectorMessageIDs
	case "reference_projector_id":
		return m.ReferenceProjectorID
	case "speaker_ids":
		return m.SpeakerIDs
	case "start_time":
		return m.StartTime
	case "structure_level_ids":
		return m.StructureLevelIDs
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	case "tag_ids":
		return m.TagIDs
	case "template_for_organization_id":
		return m.TemplateForOrganizationID
	case "topic_ids":
		return m.TopicIDs
	case "topic_poll_default_group_ids":
		return m.TopicPollDefaultGroupIDs
	case "user_ids":
		return m.UserIDs
	case "users_allow_self_set_present":
		return m.UsersAllowSelfSetPresent
	case "users_email_body":
		return m.UsersEmailBody
	case "users_email_replyto":
		return m.UsersEmailReplyto
	case "users_email_sender":
		return m.UsersEmailSender
	case "users_email_subject":
		return m.UsersEmailSubject
	case "users_enable_presence_view":
		return m.UsersEnablePresenceView
	case "users_enable_vote_delegations":
		return m.UsersEnableVoteDelegations
	case "users_enable_vote_weight":
		return m.UsersEnableVoteWeight
	case "users_forbid_delegator_as_submitter":
		return m.UsersForbidDelegatorAsSubmitter
	case "users_forbid_delegator_as_supporter":
		return m.UsersForbidDelegatorAsSupporter
	case "users_forbid_delegator_in_list_of_speakers":
		return m.UsersForbidDelegatorInListOfSpeakers
	case "users_forbid_delegator_to_vote":
		return m.UsersForbidDelegatorToVote
	case "users_pdf_welcometext":
		return m.UsersPdfWelcometext
	case "users_pdf_welcometitle":
		return m.UsersPdfWelcometitle
	case "users_pdf_wlan_encryption":
		return m.UsersPdfWlanEncryption
	case "users_pdf_wlan_password":
		return m.UsersPdfWlanPassword
	case "users_pdf_wlan_ssid":
		return m.UsersPdfWlanSsid
	case "vote_ids":
		return m.VoteIDs
	case "welcome_text":
		return m.WelcomeText
	case "welcome_title":
		return m.WelcomeTitle
	}

	return nil
}

func (m *Meeting) GetFqids(field string) []string {
	switch field {
	case "admin_group_id":
		if m.AdminGroupID != nil {
			return []string{"group/" + strconv.Itoa(*m.AdminGroupID)}
		}

	case "agenda_item_ids":
		r := make([]string, len(m.AgendaItemIDs))
		for i, id := range m.AgendaItemIDs {
			r[i] = "agenda_item/" + strconv.Itoa(id)
		}
		return r

	case "all_projection_ids":
		r := make([]string, len(m.AllProjectionIDs))
		for i, id := range m.AllProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "anonymous_group_id":
		if m.AnonymousGroupID != nil {
			return []string{"group/" + strconv.Itoa(*m.AnonymousGroupID)}
		}

	case "assignment_candidate_ids":
		r := make([]string, len(m.AssignmentCandidateIDs))
		for i, id := range m.AssignmentCandidateIDs {
			r[i] = "assignment_candidate/" + strconv.Itoa(id)
		}
		return r

	case "assignment_poll_default_group_ids":
		r := make([]string, len(m.AssignmentPollDefaultGroupIDs))
		for i, id := range m.AssignmentPollDefaultGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "assignment_ids":
		r := make([]string, len(m.AssignmentIDs))
		for i, id := range m.AssignmentIDs {
			r[i] = "assignment/" + strconv.Itoa(id)
		}
		return r

	case "chat_group_ids":
		r := make([]string, len(m.ChatGroupIDs))
		for i, id := range m.ChatGroupIDs {
			r[i] = "chat_group/" + strconv.Itoa(id)
		}
		return r

	case "chat_message_ids":
		r := make([]string, len(m.ChatMessageIDs))
		for i, id := range m.ChatMessageIDs {
			r[i] = "chat_message/" + strconv.Itoa(id)
		}
		return r

	case "committee_id":
		return []string{"committee/" + strconv.Itoa(m.CommitteeID)}

	case "default_group_id":
		return []string{"group/" + strconv.Itoa(m.DefaultGroupID)}

	case "default_meeting_for_committee_id":
		if m.DefaultMeetingForCommitteeID != nil {
			return []string{"committee/" + strconv.Itoa(*m.DefaultMeetingForCommitteeID)}
		}

	case "default_projector_agenda_item_list_ids":
		r := make([]string, len(m.DefaultProjectorAgendaItemListIDs))
		for i, id := range m.DefaultProjectorAgendaItemListIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_amendment_ids":
		r := make([]string, len(m.DefaultProjectorAmendmentIDs))
		for i, id := range m.DefaultProjectorAmendmentIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_assignment_poll_ids":
		r := make([]string, len(m.DefaultProjectorAssignmentPollIDs))
		for i, id := range m.DefaultProjectorAssignmentPollIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_assignment_ids":
		r := make([]string, len(m.DefaultProjectorAssignmentIDs))
		for i, id := range m.DefaultProjectorAssignmentIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_countdown_ids":
		r := make([]string, len(m.DefaultProjectorCountdownIDs))
		for i, id := range m.DefaultProjectorCountdownIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_current_list_of_speakers_ids":
		r := make([]string, len(m.DefaultProjectorCurrentListOfSpeakersIDs))
		for i, id := range m.DefaultProjectorCurrentListOfSpeakersIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_list_of_speakers_ids":
		r := make([]string, len(m.DefaultProjectorListOfSpeakersIDs))
		for i, id := range m.DefaultProjectorListOfSpeakersIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_mediafile_ids":
		r := make([]string, len(m.DefaultProjectorMediafileIDs))
		for i, id := range m.DefaultProjectorMediafileIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_message_ids":
		r := make([]string, len(m.DefaultProjectorMessageIDs))
		for i, id := range m.DefaultProjectorMessageIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_motion_block_ids":
		r := make([]string, len(m.DefaultProjectorMotionBlockIDs))
		for i, id := range m.DefaultProjectorMotionBlockIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_motion_poll_ids":
		r := make([]string, len(m.DefaultProjectorMotionPollIDs))
		for i, id := range m.DefaultProjectorMotionPollIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_motion_ids":
		r := make([]string, len(m.DefaultProjectorMotionIDs))
		for i, id := range m.DefaultProjectorMotionIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_poll_ids":
		r := make([]string, len(m.DefaultProjectorPollIDs))
		for i, id := range m.DefaultProjectorPollIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_topic_ids":
		r := make([]string, len(m.DefaultProjectorTopicIDs))
		for i, id := range m.DefaultProjectorTopicIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "font_bold_id":
		if m.FontBoldID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontBoldID)}
		}

	case "font_bold_italic_id":
		if m.FontBoldItalicID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontBoldItalicID)}
		}

	case "font_chyron_speaker_name_id":
		if m.FontChyronSpeakerNameID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontChyronSpeakerNameID)}
		}

	case "font_italic_id":
		if m.FontItalicID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontItalicID)}
		}

	case "font_monospace_id":
		if m.FontMonospaceID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontMonospaceID)}
		}

	case "font_projector_h1_id":
		if m.FontProjectorH1ID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontProjectorH1ID)}
		}

	case "font_projector_h2_id":
		if m.FontProjectorH2ID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontProjectorH2ID)}
		}

	case "font_regular_id":
		if m.FontRegularID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontRegularID)}
		}

	case "forwarded_motion_ids":
		r := make([]string, len(m.ForwardedMotionIDs))
		for i, id := range m.ForwardedMotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "group_ids":
		r := make([]string, len(m.GroupIDs))
		for i, id := range m.GroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "is_active_in_organization_id":
		if m.IsActiveInOrganizationID != nil {
			return []string{"organization/" + strconv.Itoa(*m.IsActiveInOrganizationID)}
		}

	case "is_archived_in_organization_id":
		if m.IsArchivedInOrganizationID != nil {
			return []string{"organization/" + strconv.Itoa(*m.IsArchivedInOrganizationID)}
		}

	case "list_of_speakers_countdown_id":
		if m.ListOfSpeakersCountdownID != nil {
			return []string{"projector_countdown/" + strconv.Itoa(*m.ListOfSpeakersCountdownID)}
		}

	case "list_of_speakers_ids":
		r := make([]string, len(m.ListOfSpeakersIDs))
		for i, id := range m.ListOfSpeakersIDs {
			r[i] = "list_of_speakers/" + strconv.Itoa(id)
		}
		return r

	case "logo_pdf_ballot_paper_id":
		if m.LogoPdfBallotPaperID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfBallotPaperID)}
		}

	case "logo_pdf_footer_l_id":
		if m.LogoPdfFooterLID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfFooterLID)}
		}

	case "logo_pdf_footer_r_id":
		if m.LogoPdfFooterRID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfFooterRID)}
		}

	case "logo_pdf_header_l_id":
		if m.LogoPdfHeaderLID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfHeaderLID)}
		}

	case "logo_pdf_header_r_id":
		if m.LogoPdfHeaderRID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfHeaderRID)}
		}

	case "logo_projector_header_id":
		if m.LogoProjectorHeaderID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoProjectorHeaderID)}
		}

	case "logo_projector_main_id":
		if m.LogoProjectorMainID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoProjectorMainID)}
		}

	case "logo_web_header_id":
		if m.LogoWebHeaderID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoWebHeaderID)}
		}

	case "mediafile_ids":
		r := make([]string, len(m.MediafileIDs))
		for i, id := range m.MediafileIDs {
			r[i] = "mediafile/" + strconv.Itoa(id)
		}
		return r

	case "meeting_mediafile_ids":
		r := make([]string, len(m.MeetingMediafileIDs))
		for i, id := range m.MeetingMediafileIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "meeting_user_ids":
		r := make([]string, len(m.MeetingUserIDs))
		for i, id := range m.MeetingUserIDs {
			r[i] = "meeting_user/" + strconv.Itoa(id)
		}
		return r

	case "motion_block_ids":
		r := make([]string, len(m.MotionBlockIDs))
		for i, id := range m.MotionBlockIDs {
			r[i] = "motion_block/" + strconv.Itoa(id)
		}
		return r

	case "motion_category_ids":
		r := make([]string, len(m.MotionCategoryIDs))
		for i, id := range m.MotionCategoryIDs {
			r[i] = "motion_category/" + strconv.Itoa(id)
		}
		return r

	case "motion_change_recommendation_ids":
		r := make([]string, len(m.MotionChangeRecommendationIDs))
		for i, id := range m.MotionChangeRecommendationIDs {
			r[i] = "motion_change_recommendation/" + strconv.Itoa(id)
		}
		return r

	case "motion_comment_section_ids":
		r := make([]string, len(m.MotionCommentSectionIDs))
		for i, id := range m.MotionCommentSectionIDs {
			r[i] = "motion_comment_section/" + strconv.Itoa(id)
		}
		return r

	case "motion_comment_ids":
		r := make([]string, len(m.MotionCommentIDs))
		for i, id := range m.MotionCommentIDs {
			r[i] = "motion_comment/" + strconv.Itoa(id)
		}
		return r

	case "motion_editor_ids":
		r := make([]string, len(m.MotionEditorIDs))
		for i, id := range m.MotionEditorIDs {
			r[i] = "motion_editor/" + strconv.Itoa(id)
		}
		return r

	case "motion_poll_default_group_ids":
		r := make([]string, len(m.MotionPollDefaultGroupIDs))
		for i, id := range m.MotionPollDefaultGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "motion_state_ids":
		r := make([]string, len(m.MotionStateIDs))
		for i, id := range m.MotionStateIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r

	case "motion_submitter_ids":
		r := make([]string, len(m.MotionSubmitterIDs))
		for i, id := range m.MotionSubmitterIDs {
			r[i] = "motion_submitter/" + strconv.Itoa(id)
		}
		return r

	case "motion_workflow_ids":
		r := make([]string, len(m.MotionWorkflowIDs))
		for i, id := range m.MotionWorkflowIDs {
			r[i] = "motion_workflow/" + strconv.Itoa(id)
		}
		return r

	case "motion_working_group_speaker_ids":
		r := make([]string, len(m.MotionWorkingGroupSpeakerIDs))
		for i, id := range m.MotionWorkingGroupSpeakerIDs {
			r[i] = "motion_working_group_speaker/" + strconv.Itoa(id)
		}
		return r

	case "motion_ids":
		r := make([]string, len(m.MotionIDs))
		for i, id := range m.MotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "motions_default_amendment_workflow_id":
		return []string{"motion_workflow/" + strconv.Itoa(m.MotionsDefaultAmendmentWorkflowID)}

	case "motions_default_workflow_id":
		return []string{"motion_workflow/" + strconv.Itoa(m.MotionsDefaultWorkflowID)}

	case "option_ids":
		r := make([]string, len(m.OptionIDs))
		for i, id := range m.OptionIDs {
			r[i] = "option/" + strconv.Itoa(id)
		}
		return r

	case "organization_tag_ids":
		r := make([]string, len(m.OrganizationTagIDs))
		for i, id := range m.OrganizationTagIDs {
			r[i] = "organization_tag/" + strconv.Itoa(id)
		}
		return r

	case "personal_note_ids":
		r := make([]string, len(m.PersonalNoteIDs))
		for i, id := range m.PersonalNoteIDs {
			r[i] = "personal_note/" + strconv.Itoa(id)
		}
		return r

	case "point_of_order_category_ids":
		r := make([]string, len(m.PointOfOrderCategoryIDs))
		for i, id := range m.PointOfOrderCategoryIDs {
			r[i] = "point_of_order_category/" + strconv.Itoa(id)
		}
		return r

	case "poll_candidate_list_ids":
		r := make([]string, len(m.PollCandidateListIDs))
		for i, id := range m.PollCandidateListIDs {
			r[i] = "poll_candidate_list/" + strconv.Itoa(id)
		}
		return r

	case "poll_candidate_ids":
		r := make([]string, len(m.PollCandidateIDs))
		for i, id := range m.PollCandidateIDs {
			r[i] = "poll_candidate/" + strconv.Itoa(id)
		}
		return r

	case "poll_countdown_id":
		if m.PollCountdownID != nil {
			return []string{"projector_countdown/" + strconv.Itoa(*m.PollCountdownID)}
		}

	case "poll_default_group_ids":
		r := make([]string, len(m.PollDefaultGroupIDs))
		for i, id := range m.PollDefaultGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "poll_ids":
		r := make([]string, len(m.PollIDs))
		for i, id := range m.PollIDs {
			r[i] = "poll/" + strconv.Itoa(id)
		}
		return r

	case "present_user_ids":
		r := make([]string, len(m.PresentUserIDs))
		for i, id := range m.PresentUserIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "projector_countdown_ids":
		r := make([]string, len(m.ProjectorCountdownIDs))
		for i, id := range m.ProjectorCountdownIDs {
			r[i] = "projector_countdown/" + strconv.Itoa(id)
		}
		return r

	case "projector_message_ids":
		r := make([]string, len(m.ProjectorMessageIDs))
		for i, id := range m.ProjectorMessageIDs {
			r[i] = "projector_message/" + strconv.Itoa(id)
		}
		return r

	case "projector_ids":
		r := make([]string, len(m.ProjectorIDs))
		for i, id := range m.ProjectorIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "reference_projector_id":
		return []string{"projector/" + strconv.Itoa(m.ReferenceProjectorID)}

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

	case "structure_level_ids":
		r := make([]string, len(m.StructureLevelIDs))
		for i, id := range m.StructureLevelIDs {
			r[i] = "structure_level/" + strconv.Itoa(id)
		}
		return r

	case "tag_ids":
		r := make([]string, len(m.TagIDs))
		for i, id := range m.TagIDs {
			r[i] = "tag/" + strconv.Itoa(id)
		}
		return r

	case "template_for_organization_id":
		if m.TemplateForOrganizationID != nil {
			return []string{"organization/" + strconv.Itoa(*m.TemplateForOrganizationID)}
		}

	case "topic_poll_default_group_ids":
		r := make([]string, len(m.TopicPollDefaultGroupIDs))
		for i, id := range m.TopicPollDefaultGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "topic_ids":
		r := make([]string, len(m.TopicIDs))
		for i, id := range m.TopicIDs {
			r[i] = "topic/" + strconv.Itoa(id)
		}
		return r

	case "vote_ids":
		r := make([]string, len(m.VoteIDs))
		for i, id := range m.VoteIDs {
			r[i] = "vote/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Meeting) Update(data map[string]string) error {
	if val, ok := data["admin_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AdminGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_enable_numbering"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaEnableNumbering)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_item_creation"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemCreation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_item_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["agenda_item_ids"]; ok {
			m.agendaItems = slices.DeleteFunc(m.agendaItems, func(r *AgendaItem) bool {
				return !slices.Contains(m.AgendaItemIDs, r.ID)
			})
		}
	}

	if val, ok := data["agenda_new_items_default_visibility"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaNewItemsDefaultVisibility)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_number_prefix"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaNumberPrefix)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_numeral_system"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaNumeralSystem)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_show_internal_items_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaShowInternalItemsOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_show_subtitles"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaShowSubtitles)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_show_topic_navigation_on_detail_view"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaShowTopicNavigationOnDetailView)
		if err != nil {
			return err
		}
	}

	if val, ok := data["all_projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllProjectionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["all_projection_ids"]; ok {
			m.allProjections = slices.DeleteFunc(m.allProjections, func(r *Projection) bool {
				return !slices.Contains(m.AllProjectionIDs, r.ID)
			})
		}
	}

	if val, ok := data["anonymous_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AnonymousGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_enable"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseEnable)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_max_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseMaxAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_min_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseMinAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_particle_image_url"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseParticleImageUrl)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_show_level"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseShowLevel)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_timeout"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseTimeout)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentCandidateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["assignment_candidate_ids"]; ok {
			m.assignmentCandidates = slices.DeleteFunc(m.assignmentCandidates, func(r *AssignmentCandidate) bool {
				return !slices.Contains(m.AssignmentCandidateIDs, r.ID)
			})
		}
	}

	if val, ok := data["assignment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["assignment_ids"]; ok {
			m.assignments = slices.DeleteFunc(m.assignments, func(r *Assignment) bool {
				return !slices.Contains(m.AssignmentIDs, r.ID)
			})
		}
	}

	if val, ok := data["assignment_poll_add_candidates_to_list_of_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollAddCandidatesToListOfSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_ballot_paper_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollBallotPaperNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_ballot_paper_selection"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollBallotPaperSelection)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultBackend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["assignment_poll_default_group_ids"]; ok {
			m.assignmentPollDefaultGroups = slices.DeleteFunc(m.assignmentPollDefaultGroups, func(r *Group) bool {
				return !slices.Contains(m.AssignmentPollDefaultGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["assignment_poll_default_method"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultMethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultOnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_enable_max_votes_per_option"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollEnableMaxVotesPerOption)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_sort_poll_result_by_votes"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollSortPollResultByVotes)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignments_export_preamble"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentsExportPreamble)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignments_export_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentsExportTitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chat_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["chat_group_ids"]; ok {
			m.chatGroups = slices.DeleteFunc(m.chatGroups, func(r *ChatGroup) bool {
				return !slices.Contains(m.ChatGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["chat_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatMessageIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["chat_message_ids"]; ok {
			m.chatMessages = slices.DeleteFunc(m.chatMessages, func(r *ChatMessage) bool {
				return !slices.Contains(m.ChatMessageIDs, r.ID)
			})
		}
	}

	if val, ok := data["committee_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommitteeID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_auto_connect"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceAutoConnect)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_auto_connect_next_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceAutoConnectNextSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_enable_helpdesk"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceEnableHelpdesk)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_los_restriction"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceLosRestriction)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_open_microphone"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceOpenMicrophone)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_open_video"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceOpenVideo)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_show"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceShow)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_stream_poster_url"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceStreamPosterUrl)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_stream_url"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceStreamUrl)
		if err != nil {
			return err
		}
	}

	if val, ok := data["custom_translations"]; ok {
		err := json.Unmarshal([]byte(val), &m.CustomTranslations)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_meeting_for_committee_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultMeetingForCommitteeID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_agenda_item_list_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAgendaItemListIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_agenda_item_list_ids"]; ok {
			m.defaultProjectorAgendaItemLists = slices.DeleteFunc(m.defaultProjectorAgendaItemLists, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorAgendaItemListIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_amendment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAmendmentIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_amendment_ids"]; ok {
			m.defaultProjectorAmendments = slices.DeleteFunc(m.defaultProjectorAmendments, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorAmendmentIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_assignment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAssignmentIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_assignment_ids"]; ok {
			m.defaultProjectorAssignments = slices.DeleteFunc(m.defaultProjectorAssignments, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorAssignmentIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_assignment_poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAssignmentPollIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_assignment_poll_ids"]; ok {
			m.defaultProjectorAssignmentPolls = slices.DeleteFunc(m.defaultProjectorAssignmentPolls, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorAssignmentPollIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_countdown_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorCountdownIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_countdown_ids"]; ok {
			m.defaultProjectorCountdowns = slices.DeleteFunc(m.defaultProjectorCountdowns, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorCountdownIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_current_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorCurrentListOfSpeakersIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_current_list_of_speakers_ids"]; ok {
			m.defaultProjectorCurrentListOfSpeakerss = slices.DeleteFunc(m.defaultProjectorCurrentListOfSpeakerss, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorCurrentListOfSpeakersIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorListOfSpeakersIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_list_of_speakers_ids"]; ok {
			m.defaultProjectorListOfSpeakerss = slices.DeleteFunc(m.defaultProjectorListOfSpeakerss, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorListOfSpeakersIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMediafileIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_mediafile_ids"]; ok {
			m.defaultProjectorMediafiles = slices.DeleteFunc(m.defaultProjectorMediafiles, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorMediafileIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMessageIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_message_ids"]; ok {
			m.defaultProjectorMessages = slices.DeleteFunc(m.defaultProjectorMessages, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorMessageIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_motion_block_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMotionBlockIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_motion_block_ids"]; ok {
			m.defaultProjectorMotionBlocks = slices.DeleteFunc(m.defaultProjectorMotionBlocks, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorMotionBlockIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_motion_ids"]; ok {
			m.defaultProjectorMotions = slices.DeleteFunc(m.defaultProjectorMotions, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorMotionIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_motion_poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMotionPollIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_motion_poll_ids"]; ok {
			m.defaultProjectorMotionPolls = slices.DeleteFunc(m.defaultProjectorMotionPolls, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorMotionPollIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorPollIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_poll_ids"]; ok {
			m.defaultProjectorPolls = slices.DeleteFunc(m.defaultProjectorPolls, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorPollIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_projector_topic_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorTopicIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["default_projector_topic_ids"]; ok {
			m.defaultProjectorTopics = slices.DeleteFunc(m.defaultProjectorTopics, func(r *Projector) bool {
				return !slices.Contains(m.DefaultProjectorTopicIDs, r.ID)
			})
		}
	}

	if val, ok := data["description"]; ok {
		err := json.Unmarshal([]byte(val), &m.Description)
		if err != nil {
			return err
		}
	}

	if val, ok := data["enable_anonymous"]; ok {
		err := json.Unmarshal([]byte(val), &m.EnableAnonymous)
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

	if val, ok := data["export_csv_encoding"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportCsvEncoding)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_csv_separator"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportCsvSeparator)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_fontsize"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfFontsize)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_line_height"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfLineHeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_bottom"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginBottom)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_left"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginLeft)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_right"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginRight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_top"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginTop)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_pagenumber_alignment"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPagenumberAlignment)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_pagesize"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPagesize)
		if err != nil {
			return err
		}
	}

	if val, ok := data["external_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExternalID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_bold_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontBoldID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_bold_italic_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontBoldItalicID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_chyron_speaker_name_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontChyronSpeakerNameID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_italic_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontItalicID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_monospace_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontMonospaceID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_projector_h1_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontProjectorH1ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_projector_h2_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontProjectorH2ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_regular_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontRegularID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["forwarded_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ForwardedMotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["forwarded_motion_ids"]; ok {
			m.forwardedMotions = slices.DeleteFunc(m.forwardedMotions, func(r *Motion) bool {
				return !slices.Contains(m.ForwardedMotionIDs, r.ID)
			})
		}
	}

	if val, ok := data["group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.GroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["group_ids"]; ok {
			m.groups = slices.DeleteFunc(m.groups, func(r *Group) bool {
				return !slices.Contains(m.GroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["imported_at"]; ok {
		err := json.Unmarshal([]byte(val), &m.ImportedAt)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_active_in_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsActiveInOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_archived_in_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsArchivedInOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["jitsi_domain"]; ok {
		err := json.Unmarshal([]byte(val), &m.JitsiDomain)
		if err != nil {
			return err
		}
	}

	if val, ok := data["jitsi_room_name"]; ok {
		err := json.Unmarshal([]byte(val), &m.JitsiRoomName)
		if err != nil {
			return err
		}
	}

	if val, ok := data["jitsi_room_password"]; ok {
		err := json.Unmarshal([]byte(val), &m.JitsiRoomPassword)
		if err != nil {
			return err
		}
	}

	if val, ok := data["language"]; ok {
		err := json.Unmarshal([]byte(val), &m.Language)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_allow_multiple_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersAllowMultipleSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_amount_last_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersAmountLastOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_amount_next_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersAmountNextOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_can_create_point_of_order_for_others"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCanCreatePointOfOrderForOthers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_can_set_contribution_self"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCanSetContributionSelf)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_closing_disables_point_of_order"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersClosingDisablesPointOfOrder)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_countdown_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCountdownID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_couple_countdown"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCoupleCountdown)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_default_structure_level_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersDefaultStructureLevelTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_interposed_question"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnableInterposedQuestion)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_point_of_order_categories"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnablePointOfOrderCategories)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_point_of_order_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnablePointOfOrderSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_pro_contra_speech"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnableProContraSpeech)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_hide_contribution_count"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersHideContributionCount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["list_of_speakers_ids"]; ok {
			m.listOfSpeakerss = slices.DeleteFunc(m.listOfSpeakerss, func(r *ListOfSpeakers) bool {
				return !slices.Contains(m.ListOfSpeakersIDs, r.ID)
			})
		}
	}

	if val, ok := data["list_of_speakers_initially_closed"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersInitiallyClosed)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_intervention_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersInterventionTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_present_users_only"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersPresentUsersOnly)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_show_amount_of_speakers_on_slide"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersShowAmountOfSpeakersOnSlide)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_show_first_contribution"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersShowFirstContribution)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_speaker_note_for_everyone"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersSpeakerNoteForEveryone)
		if err != nil {
			return err
		}
	}

	if val, ok := data["location"]; ok {
		err := json.Unmarshal([]byte(val), &m.Location)
		if err != nil {
			return err
		}
	}

	if val, ok := data["locked_from_inside"]; ok {
		err := json.Unmarshal([]byte(val), &m.LockedFromInside)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_ballot_paper_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfBallotPaperID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_footer_l_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfFooterLID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_footer_r_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfFooterRID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_header_l_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfHeaderLID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_header_r_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfHeaderRID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_projector_header_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoProjectorHeaderID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_projector_main_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoProjectorMainID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_web_header_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoWebHeaderID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MediafileIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["mediafile_ids"]; ok {
			m.mediafiles = slices.DeleteFunc(m.mediafiles, func(r *Mediafile) bool {
				return !slices.Contains(m.MediafileIDs, r.ID)
			})
		}
	}

	if val, ok := data["meeting_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingMediafileIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["meeting_mediafile_ids"]; ok {
			m.meetingMediafiles = slices.DeleteFunc(m.meetingMediafiles, func(r *MeetingMediafile) bool {
				return !slices.Contains(m.MeetingMediafileIDs, r.ID)
			})
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

	if val, ok := data["motion_block_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionBlockIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_block_ids"]; ok {
			m.motionBlocks = slices.DeleteFunc(m.motionBlocks, func(r *MotionBlock) bool {
				return !slices.Contains(m.MotionBlockIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_category_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionCategoryIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_category_ids"]; ok {
			m.motionCategorys = slices.DeleteFunc(m.motionCategorys, func(r *MotionCategory) bool {
				return !slices.Contains(m.MotionCategoryIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_change_recommendation_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionChangeRecommendationIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_change_recommendation_ids"]; ok {
			m.motionChangeRecommendations = slices.DeleteFunc(m.motionChangeRecommendations, func(r *MotionChangeRecommendation) bool {
				return !slices.Contains(m.MotionChangeRecommendationIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_comment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionCommentIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_comment_ids"]; ok {
			m.motionComments = slices.DeleteFunc(m.motionComments, func(r *MotionComment) bool {
				return !slices.Contains(m.MotionCommentIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_comment_section_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionCommentSectionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_comment_section_ids"]; ok {
			m.motionCommentSections = slices.DeleteFunc(m.motionCommentSections, func(r *MotionCommentSection) bool {
				return !slices.Contains(m.MotionCommentSectionIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_editor_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionEditorIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_editor_ids"]; ok {
			m.motionEditors = slices.DeleteFunc(m.motionEditors, func(r *MotionEditor) bool {
				return !slices.Contains(m.MotionEditorIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_ids"]; ok {
			m.motions = slices.DeleteFunc(m.motions, func(r *Motion) bool {
				return !slices.Contains(m.MotionIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_poll_ballot_paper_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollBallotPaperNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_ballot_paper_selection"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollBallotPaperSelection)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultBackend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_poll_default_group_ids"]; ok {
			m.motionPollDefaultGroups = slices.DeleteFunc(m.motionPollDefaultGroups, func(r *Group) bool {
				return !slices.Contains(m.MotionPollDefaultGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_poll_default_method"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultMethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultOnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionStateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_state_ids"]; ok {
			m.motionStates = slices.DeleteFunc(m.motionStates, func(r *MotionState) bool {
				return !slices.Contains(m.MotionStateIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_submitter_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionSubmitterIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_submitter_ids"]; ok {
			m.motionSubmitters = slices.DeleteFunc(m.motionSubmitters, func(r *MotionSubmitter) bool {
				return !slices.Contains(m.MotionSubmitterIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_workflow_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionWorkflowIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_workflow_ids"]; ok {
			m.motionWorkflows = slices.DeleteFunc(m.motionWorkflows, func(r *MotionWorkflow) bool {
				return !slices.Contains(m.MotionWorkflowIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_working_group_speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionWorkingGroupSpeakerIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_working_group_speaker_ids"]; ok {
			m.motionWorkingGroupSpeakers = slices.DeleteFunc(m.motionWorkingGroupSpeakers, func(r *MotionWorkingGroupSpeaker) bool {
				return !slices.Contains(m.MotionWorkingGroupSpeakerIDs, r.ID)
			})
		}
	}

	if val, ok := data["motions_amendments_enabled"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsEnabled)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_in_main_list"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsInMainList)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_multiple_paragraphs"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsMultipleParagraphs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_of_amendments"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsOfAmendments)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_prefix"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsPrefix)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_text_mode"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsTextMode)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_block_slide_columns"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsBlockSlideColumns)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_create_enable_additional_submitter_text"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsCreateEnableAdditionalSubmitterText)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_amendment_workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultAmendmentWorkflowID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_line_numbering"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultLineNumbering)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_sorting"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultSorting)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultWorkflowID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_editor"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableEditor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_reason_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableReasonOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_recommendation_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableRecommendationOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_sidebox_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableSideboxOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_text_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableTextOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_working_group_speaker"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableWorkingGroupSpeaker)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_follow_recommendation"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportFollowRecommendation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_preamble"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportPreamble)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_submitter_recommendation"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportSubmitterRecommendation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportTitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_hide_metadata_background"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsHideMetadataBackground)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_line_length"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsLineLength)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_number_min_digits"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsNumberMinDigits)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_number_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsNumberType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_number_with_blank"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsNumberWithBlank)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_preamble"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsPreamble)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_reason_required"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsReasonRequired)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_recommendation_text_mode"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsRecommendationTextMode)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_recommendations_by"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsRecommendationsBy)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_show_referring_motions"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsShowReferringMotions)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_show_sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsShowSequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_supporters_min_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsSupportersMinAmount)
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

	if val, ok := data["option_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["option_ids"]; ok {
			m.options = slices.DeleteFunc(m.options, func(r *Option) bool {
				return !slices.Contains(m.OptionIDs, r.ID)
			})
		}
	}

	if val, ok := data["organization_tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationTagIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["organization_tag_ids"]; ok {
			m.organizationTags = slices.DeleteFunc(m.organizationTags, func(r *OrganizationTag) bool {
				return !slices.Contains(m.OrganizationTagIDs, r.ID)
			})
		}
	}

	if val, ok := data["personal_note_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PersonalNoteIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["personal_note_ids"]; ok {
			m.personalNotes = slices.DeleteFunc(m.personalNotes, func(r *PersonalNote) bool {
				return !slices.Contains(m.PersonalNoteIDs, r.ID)
			})
		}
	}

	if val, ok := data["point_of_order_category_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PointOfOrderCategoryIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["point_of_order_category_ids"]; ok {
			m.pointOfOrderCategorys = slices.DeleteFunc(m.pointOfOrderCategorys, func(r *PointOfOrderCategory) bool {
				return !slices.Contains(m.PointOfOrderCategoryIDs, r.ID)
			})
		}
	}

	if val, ok := data["poll_ballot_paper_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollBallotPaperNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_ballot_paper_selection"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollBallotPaperSelection)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["poll_candidate_ids"]; ok {
			m.pollCandidates = slices.DeleteFunc(m.pollCandidates, func(r *PollCandidate) bool {
				return !slices.Contains(m.PollCandidateIDs, r.ID)
			})
		}
	}

	if val, ok := data["poll_candidate_list_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateListIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["poll_candidate_list_ids"]; ok {
			m.pollCandidateLists = slices.DeleteFunc(m.pollCandidateLists, func(r *PollCandidateList) bool {
				return !slices.Contains(m.PollCandidateListIDs, r.ID)
			})
		}
	}

	if val, ok := data["poll_countdown_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCountdownID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_couple_countdown"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCoupleCountdown)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultBackend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["poll_default_group_ids"]; ok {
			m.pollDefaultGroups = slices.DeleteFunc(m.pollDefaultGroups, func(r *Group) bool {
				return !slices.Contains(m.PollDefaultGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["poll_default_method"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultMethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultOnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultType)
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

	if val, ok := data["poll_sort_poll_result_by_votes"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollSortPollResultByVotes)
		if err != nil {
			return err
		}
	}

	if val, ok := data["present_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PresentUserIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["present_user_ids"]; ok {
			m.presentUsers = slices.DeleteFunc(m.presentUsers, func(r *User) bool {
				return !slices.Contains(m.PresentUserIDs, r.ID)
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

	if val, ok := data["projector_countdown_default_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorCountdownDefaultTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_countdown_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorCountdownIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["projector_countdown_ids"]; ok {
			m.projectorCountdowns = slices.DeleteFunc(m.projectorCountdowns, func(r *ProjectorCountdown) bool {
				return !slices.Contains(m.ProjectorCountdownIDs, r.ID)
			})
		}
	}

	if val, ok := data["projector_countdown_warning_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorCountdownWarningTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["projector_ids"]; ok {
			m.projectors = slices.DeleteFunc(m.projectors, func(r *Projector) bool {
				return !slices.Contains(m.ProjectorIDs, r.ID)
			})
		}
	}

	if val, ok := data["projector_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorMessageIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["projector_message_ids"]; ok {
			m.projectorMessages = slices.DeleteFunc(m.projectorMessages, func(r *ProjectorMessage) bool {
				return !slices.Contains(m.ProjectorMessageIDs, r.ID)
			})
		}
	}

	if val, ok := data["reference_projector_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReferenceProjectorID)
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

	if val, ok := data["start_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.StartTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["structure_level_ids"]; ok {
			m.structureLevels = slices.DeleteFunc(m.structureLevels, func(r *StructureLevel) bool {
				return !slices.Contains(m.StructureLevelIDs, r.ID)
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

	if val, ok := data["tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TagIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["tag_ids"]; ok {
			m.tags = slices.DeleteFunc(m.tags, func(r *Tag) bool {
				return !slices.Contains(m.TagIDs, r.ID)
			})
		}
	}

	if val, ok := data["template_for_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.TemplateForOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["topic_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TopicIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["topic_ids"]; ok {
			m.topics = slices.DeleteFunc(m.topics, func(r *Topic) bool {
				return !slices.Contains(m.TopicIDs, r.ID)
			})
		}
	}

	if val, ok := data["topic_poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TopicPollDefaultGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["topic_poll_default_group_ids"]; ok {
			m.topicPollDefaultGroups = slices.DeleteFunc(m.topicPollDefaultGroups, func(r *Group) bool {
				return !slices.Contains(m.TopicPollDefaultGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_allow_self_set_present"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersAllowSelfSetPresent)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_body"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailBody)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_replyto"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailReplyto)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_sender"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailSender)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_subject"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailSubject)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_enable_presence_view"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEnablePresenceView)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_enable_vote_delegations"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEnableVoteDelegations)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_enable_vote_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEnableVoteWeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_as_submitter"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorAsSubmitter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_as_supporter"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorAsSupporter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_in_list_of_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorInListOfSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_to_vote"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorToVote)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_welcometext"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWelcometext)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_welcometitle"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWelcometitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_wlan_encryption"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWlanEncryption)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_wlan_password"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWlanPassword)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_wlan_ssid"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWlanSsid)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["vote_ids"]; ok {
			m.votes = slices.DeleteFunc(m.votes, func(r *Vote) bool {
				return !slices.Contains(m.VoteIDs, r.ID)
			})
		}
	}

	if val, ok := data["welcome_text"]; ok {
		err := json.Unmarshal([]byte(val), &m.WelcomeText)
		if err != nil {
			return err
		}
	}

	if val, ok := data["welcome_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.WelcomeTitle)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Meeting) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type MeetingMediafile struct {
	AccessGroupIDs                         []int    `json:"access_group_ids"`
	AttachmentIDs                          []string `json:"attachment_ids"`
	ID                                     int      `json:"id"`
	InheritedAccessGroupIDs                []int    `json:"inherited_access_group_ids"`
	IsPublic                               bool     `json:"is_public"`
	ListOfSpeakersID                       *int     `json:"list_of_speakers_id"`
	MediafileID                            int      `json:"mediafile_id"`
	MeetingID                              int      `json:"meeting_id"`
	ProjectionIDs                          []int    `json:"projection_ids"`
	UsedAsFontBoldInMeetingID              *int     `json:"used_as_font_bold_in_meeting_id"`
	UsedAsFontBoldItalicInMeetingID        *int     `json:"used_as_font_bold_italic_in_meeting_id"`
	UsedAsFontChyronSpeakerNameInMeetingID *int     `json:"used_as_font_chyron_speaker_name_in_meeting_id"`
	UsedAsFontItalicInMeetingID            *int     `json:"used_as_font_italic_in_meeting_id"`
	UsedAsFontMonospaceInMeetingID         *int     `json:"used_as_font_monospace_in_meeting_id"`
	UsedAsFontProjectorH1InMeetingID       *int     `json:"used_as_font_projector_h1_in_meeting_id"`
	UsedAsFontProjectorH2InMeetingID       *int     `json:"used_as_font_projector_h2_in_meeting_id"`
	UsedAsFontRegularInMeetingID           *int     `json:"used_as_font_regular_in_meeting_id"`
	UsedAsLogoPdfBallotPaperInMeetingID    *int     `json:"used_as_logo_pdf_ballot_paper_in_meeting_id"`
	UsedAsLogoPdfFooterLInMeetingID        *int     `json:"used_as_logo_pdf_footer_l_in_meeting_id"`
	UsedAsLogoPdfFooterRInMeetingID        *int     `json:"used_as_logo_pdf_footer_r_in_meeting_id"`
	UsedAsLogoPdfHeaderLInMeetingID        *int     `json:"used_as_logo_pdf_header_l_in_meeting_id"`
	UsedAsLogoPdfHeaderRInMeetingID        *int     `json:"used_as_logo_pdf_header_r_in_meeting_id"`
	UsedAsLogoProjectorHeaderInMeetingID   *int     `json:"used_as_logo_projector_header_in_meeting_id"`
	UsedAsLogoProjectorMainInMeetingID     *int     `json:"used_as_logo_projector_main_in_meeting_id"`
	UsedAsLogoWebHeaderInMeetingID         *int     `json:"used_as_logo_web_header_in_meeting_id"`
	loadedRelations                        map[string]struct{}
	accessGroups                           []*Group
	inheritedAccessGroups                  []*Group
	listOfSpeakers                         *ListOfSpeakers
	mediafile                              *Mediafile
	meeting                                *Meeting
	projections                            []*Projection
	usedAsFontBoldInMeeting                *Meeting
	usedAsFontBoldItalicInMeeting          *Meeting
	usedAsFontChyronSpeakerNameInMeeting   *Meeting
	usedAsFontItalicInMeeting              *Meeting
	usedAsFontMonospaceInMeeting           *Meeting
	usedAsFontProjectorH1InMeeting         *Meeting
	usedAsFontProjectorH2InMeeting         *Meeting
	usedAsFontRegularInMeeting             *Meeting
	usedAsLogoPdfBallotPaperInMeeting      *Meeting
	usedAsLogoPdfFooterLInMeeting          *Meeting
	usedAsLogoPdfFooterRInMeeting          *Meeting
	usedAsLogoPdfHeaderLInMeeting          *Meeting
	usedAsLogoPdfHeaderRInMeeting          *Meeting
	usedAsLogoProjectorHeaderInMeeting     *Meeting
	usedAsLogoProjectorMainInMeeting       *Meeting
	usedAsLogoWebHeaderInMeeting           *Meeting
}

func (m *MeetingMediafile) CollectionName() string {
	return "meeting_mediafile"
}

func (m *MeetingMediafile) AccessGroups() []*Group {
	if _, ok := m.loadedRelations["access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access AccessGroups relation of MeetingMediafile which was not loaded.")
	}

	return m.accessGroups
}

func (m *MeetingMediafile) InheritedAccessGroups() []*Group {
	if _, ok := m.loadedRelations["inherited_access_group_ids"]; !ok {
		log.Panic().Msg("Tried to access InheritedAccessGroups relation of MeetingMediafile which was not loaded.")
	}

	return m.inheritedAccessGroups
}

func (m *MeetingMediafile) ListOfSpeakers() *ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of MeetingMediafile which was not loaded.")
	}

	return m.listOfSpeakers
}

func (m *MeetingMediafile) Mediafile() Mediafile {
	if _, ok := m.loadedRelations["mediafile_id"]; !ok {
		log.Panic().Msg("Tried to access Mediafile relation of MeetingMediafile which was not loaded.")
	}

	return *m.mediafile
}

func (m *MeetingMediafile) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MeetingMediafile which was not loaded.")
	}

	return *m.meeting
}

func (m *MeetingMediafile) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of MeetingMediafile which was not loaded.")
	}

	return m.projections
}

func (m *MeetingMediafile) UsedAsFontBoldInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_bold_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontBoldInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontBoldInMeeting
}

func (m *MeetingMediafile) UsedAsFontBoldItalicInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_bold_italic_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontBoldItalicInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontBoldItalicInMeeting
}

func (m *MeetingMediafile) UsedAsFontChyronSpeakerNameInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_chyron_speaker_name_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontChyronSpeakerNameInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontChyronSpeakerNameInMeeting
}

func (m *MeetingMediafile) UsedAsFontItalicInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_italic_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontItalicInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontItalicInMeeting
}

func (m *MeetingMediafile) UsedAsFontMonospaceInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_monospace_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontMonospaceInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontMonospaceInMeeting
}

func (m *MeetingMediafile) UsedAsFontProjectorH1InMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_projector_h1_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontProjectorH1InMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontProjectorH1InMeeting
}

func (m *MeetingMediafile) UsedAsFontProjectorH2InMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_projector_h2_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontProjectorH2InMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontProjectorH2InMeeting
}

func (m *MeetingMediafile) UsedAsFontRegularInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_font_regular_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsFontRegularInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsFontRegularInMeeting
}

func (m *MeetingMediafile) UsedAsLogoPdfBallotPaperInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_ballot_paper_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfBallotPaperInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfBallotPaperInMeeting
}

func (m *MeetingMediafile) UsedAsLogoPdfFooterLInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_footer_l_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfFooterLInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfFooterLInMeeting
}

func (m *MeetingMediafile) UsedAsLogoPdfFooterRInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_footer_r_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfFooterRInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfFooterRInMeeting
}

func (m *MeetingMediafile) UsedAsLogoPdfHeaderLInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_header_l_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfHeaderLInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfHeaderLInMeeting
}

func (m *MeetingMediafile) UsedAsLogoPdfHeaderRInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_pdf_header_r_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoPdfHeaderRInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoPdfHeaderRInMeeting
}

func (m *MeetingMediafile) UsedAsLogoProjectorHeaderInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_projector_header_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoProjectorHeaderInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoProjectorHeaderInMeeting
}

func (m *MeetingMediafile) UsedAsLogoProjectorMainInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_projector_main_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoProjectorMainInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoProjectorMainInMeeting
}

func (m *MeetingMediafile) UsedAsLogoWebHeaderInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_logo_web_header_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsLogoWebHeaderInMeeting relation of MeetingMediafile which was not loaded.")
	}

	return m.usedAsLogoWebHeaderInMeeting
}

func (m *MeetingMediafile) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "access_group_ids":
		for _, r := range m.accessGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "inherited_access_group_ids":
		for _, r := range m.inheritedAccessGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "list_of_speakers_id":
		return m.listOfSpeakers.GetRelatedModelsAccessor()
	case "mediafile_id":
		return m.mediafile.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "projection_ids":
		for _, r := range m.projections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "used_as_font_bold_in_meeting_id":
		return m.usedAsFontBoldInMeeting.GetRelatedModelsAccessor()
	case "used_as_font_bold_italic_in_meeting_id":
		return m.usedAsFontBoldItalicInMeeting.GetRelatedModelsAccessor()
	case "used_as_font_chyron_speaker_name_in_meeting_id":
		return m.usedAsFontChyronSpeakerNameInMeeting.GetRelatedModelsAccessor()
	case "used_as_font_italic_in_meeting_id":
		return m.usedAsFontItalicInMeeting.GetRelatedModelsAccessor()
	case "used_as_font_monospace_in_meeting_id":
		return m.usedAsFontMonospaceInMeeting.GetRelatedModelsAccessor()
	case "used_as_font_projector_h1_in_meeting_id":
		return m.usedAsFontProjectorH1InMeeting.GetRelatedModelsAccessor()
	case "used_as_font_projector_h2_in_meeting_id":
		return m.usedAsFontProjectorH2InMeeting.GetRelatedModelsAccessor()
	case "used_as_font_regular_in_meeting_id":
		return m.usedAsFontRegularInMeeting.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_ballot_paper_in_meeting_id":
		return m.usedAsLogoPdfBallotPaperInMeeting.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_footer_l_in_meeting_id":
		return m.usedAsLogoPdfFooterLInMeeting.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_footer_r_in_meeting_id":
		return m.usedAsLogoPdfFooterRInMeeting.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_header_l_in_meeting_id":
		return m.usedAsLogoPdfHeaderLInMeeting.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_header_r_in_meeting_id":
		return m.usedAsLogoPdfHeaderRInMeeting.GetRelatedModelsAccessor()
	case "used_as_logo_projector_header_in_meeting_id":
		return m.usedAsLogoProjectorHeaderInMeeting.GetRelatedModelsAccessor()
	case "used_as_logo_projector_main_in_meeting_id":
		return m.usedAsLogoProjectorMainInMeeting.GetRelatedModelsAccessor()
	case "used_as_logo_web_header_in_meeting_id":
		return m.usedAsLogoWebHeaderInMeeting.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *MeetingMediafile) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "access_group_ids":
			m.accessGroups = content.([]*Group)
		case "inherited_access_group_ids":
			m.inheritedAccessGroups = content.([]*Group)
		case "list_of_speakers_id":
			m.listOfSpeakers = content.(*ListOfSpeakers)
		case "mediafile_id":
			m.mediafile = content.(*Mediafile)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "projection_ids":
			m.projections = content.([]*Projection)
		case "used_as_font_bold_in_meeting_id":
			m.usedAsFontBoldInMeeting = content.(*Meeting)
		case "used_as_font_bold_italic_in_meeting_id":
			m.usedAsFontBoldItalicInMeeting = content.(*Meeting)
		case "used_as_font_chyron_speaker_name_in_meeting_id":
			m.usedAsFontChyronSpeakerNameInMeeting = content.(*Meeting)
		case "used_as_font_italic_in_meeting_id":
			m.usedAsFontItalicInMeeting = content.(*Meeting)
		case "used_as_font_monospace_in_meeting_id":
			m.usedAsFontMonospaceInMeeting = content.(*Meeting)
		case "used_as_font_projector_h1_in_meeting_id":
			m.usedAsFontProjectorH1InMeeting = content.(*Meeting)
		case "used_as_font_projector_h2_in_meeting_id":
			m.usedAsFontProjectorH2InMeeting = content.(*Meeting)
		case "used_as_font_regular_in_meeting_id":
			m.usedAsFontRegularInMeeting = content.(*Meeting)
		case "used_as_logo_pdf_ballot_paper_in_meeting_id":
			m.usedAsLogoPdfBallotPaperInMeeting = content.(*Meeting)
		case "used_as_logo_pdf_footer_l_in_meeting_id":
			m.usedAsLogoPdfFooterLInMeeting = content.(*Meeting)
		case "used_as_logo_pdf_footer_r_in_meeting_id":
			m.usedAsLogoPdfFooterRInMeeting = content.(*Meeting)
		case "used_as_logo_pdf_header_l_in_meeting_id":
			m.usedAsLogoPdfHeaderLInMeeting = content.(*Meeting)
		case "used_as_logo_pdf_header_r_in_meeting_id":
			m.usedAsLogoPdfHeaderRInMeeting = content.(*Meeting)
		case "used_as_logo_projector_header_in_meeting_id":
			m.usedAsLogoProjectorHeaderInMeeting = content.(*Meeting)
		case "used_as_logo_projector_main_in_meeting_id":
			m.usedAsLogoProjectorMainInMeeting = content.(*Meeting)
		case "used_as_logo_web_header_in_meeting_id":
			m.usedAsLogoWebHeaderInMeeting = content.(*Meeting)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MeetingMediafile) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "access_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.accessGroups = append(m.accessGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "inherited_access_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.inheritedAccessGroups = append(m.inheritedAccessGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "list_of_speakers_id":
		var entry ListOfSpeakers
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.listOfSpeakers = &entry

		result = entry.GetRelatedModelsAccessor()
	case "mediafile_id":
		var entry Mediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.mediafile = &entry

		result = entry.GetRelatedModelsAccessor()
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
	case "used_as_font_bold_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsFontBoldInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_font_bold_italic_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsFontBoldItalicInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_font_chyron_speaker_name_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsFontChyronSpeakerNameInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_font_italic_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsFontItalicInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_font_monospace_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsFontMonospaceInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_font_projector_h1_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsFontProjectorH1InMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_font_projector_h2_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsFontProjectorH2InMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_font_regular_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsFontRegularInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_ballot_paper_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsLogoPdfBallotPaperInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_footer_l_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsLogoPdfFooterLInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_footer_r_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsLogoPdfFooterRInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_header_l_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsLogoPdfHeaderLInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_logo_pdf_header_r_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsLogoPdfHeaderRInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_logo_projector_header_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsLogoProjectorHeaderInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_logo_projector_main_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsLogoProjectorMainInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_logo_web_header_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsLogoWebHeaderInMeeting = &entry

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

func (m *MeetingMediafile) Get(field string) interface{} {
	switch field {
	case "access_group_ids":
		return m.AccessGroupIDs
	case "attachment_ids":
		return m.AttachmentIDs
	case "id":
		return m.ID
	case "inherited_access_group_ids":
		return m.InheritedAccessGroupIDs
	case "is_public":
		return m.IsPublic
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "mediafile_id":
		return m.MediafileID
	case "meeting_id":
		return m.MeetingID
	case "projection_ids":
		return m.ProjectionIDs
	case "used_as_font_bold_in_meeting_id":
		return m.UsedAsFontBoldInMeetingID
	case "used_as_font_bold_italic_in_meeting_id":
		return m.UsedAsFontBoldItalicInMeetingID
	case "used_as_font_chyron_speaker_name_in_meeting_id":
		return m.UsedAsFontChyronSpeakerNameInMeetingID
	case "used_as_font_italic_in_meeting_id":
		return m.UsedAsFontItalicInMeetingID
	case "used_as_font_monospace_in_meeting_id":
		return m.UsedAsFontMonospaceInMeetingID
	case "used_as_font_projector_h1_in_meeting_id":
		return m.UsedAsFontProjectorH1InMeetingID
	case "used_as_font_projector_h2_in_meeting_id":
		return m.UsedAsFontProjectorH2InMeetingID
	case "used_as_font_regular_in_meeting_id":
		return m.UsedAsFontRegularInMeetingID
	case "used_as_logo_pdf_ballot_paper_in_meeting_id":
		return m.UsedAsLogoPdfBallotPaperInMeetingID
	case "used_as_logo_pdf_footer_l_in_meeting_id":
		return m.UsedAsLogoPdfFooterLInMeetingID
	case "used_as_logo_pdf_footer_r_in_meeting_id":
		return m.UsedAsLogoPdfFooterRInMeetingID
	case "used_as_logo_pdf_header_l_in_meeting_id":
		return m.UsedAsLogoPdfHeaderLInMeetingID
	case "used_as_logo_pdf_header_r_in_meeting_id":
		return m.UsedAsLogoPdfHeaderRInMeetingID
	case "used_as_logo_projector_header_in_meeting_id":
		return m.UsedAsLogoProjectorHeaderInMeetingID
	case "used_as_logo_projector_main_in_meeting_id":
		return m.UsedAsLogoProjectorMainInMeetingID
	case "used_as_logo_web_header_in_meeting_id":
		return m.UsedAsLogoWebHeaderInMeetingID
	}

	return nil
}

func (m *MeetingMediafile) GetFqids(field string) []string {
	switch field {
	case "access_group_ids":
		r := make([]string, len(m.AccessGroupIDs))
		for i, id := range m.AccessGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "inherited_access_group_ids":
		r := make([]string, len(m.InheritedAccessGroupIDs))
		for i, id := range m.InheritedAccessGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "list_of_speakers_id":
		if m.ListOfSpeakersID != nil {
			return []string{"list_of_speakers/" + strconv.Itoa(*m.ListOfSpeakersID)}
		}

	case "mediafile_id":
		return []string{"mediafile/" + strconv.Itoa(m.MediafileID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "used_as_font_bold_in_meeting_id":
		if m.UsedAsFontBoldInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsFontBoldInMeetingID)}
		}

	case "used_as_font_bold_italic_in_meeting_id":
		if m.UsedAsFontBoldItalicInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsFontBoldItalicInMeetingID)}
		}

	case "used_as_font_chyron_speaker_name_in_meeting_id":
		if m.UsedAsFontChyronSpeakerNameInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsFontChyronSpeakerNameInMeetingID)}
		}

	case "used_as_font_italic_in_meeting_id":
		if m.UsedAsFontItalicInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsFontItalicInMeetingID)}
		}

	case "used_as_font_monospace_in_meeting_id":
		if m.UsedAsFontMonospaceInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsFontMonospaceInMeetingID)}
		}

	case "used_as_font_projector_h1_in_meeting_id":
		if m.UsedAsFontProjectorH1InMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsFontProjectorH1InMeetingID)}
		}

	case "used_as_font_projector_h2_in_meeting_id":
		if m.UsedAsFontProjectorH2InMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsFontProjectorH2InMeetingID)}
		}

	case "used_as_font_regular_in_meeting_id":
		if m.UsedAsFontRegularInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsFontRegularInMeetingID)}
		}

	case "used_as_logo_pdf_ballot_paper_in_meeting_id":
		if m.UsedAsLogoPdfBallotPaperInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsLogoPdfBallotPaperInMeetingID)}
		}

	case "used_as_logo_pdf_footer_l_in_meeting_id":
		if m.UsedAsLogoPdfFooterLInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsLogoPdfFooterLInMeetingID)}
		}

	case "used_as_logo_pdf_footer_r_in_meeting_id":
		if m.UsedAsLogoPdfFooterRInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsLogoPdfFooterRInMeetingID)}
		}

	case "used_as_logo_pdf_header_l_in_meeting_id":
		if m.UsedAsLogoPdfHeaderLInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsLogoPdfHeaderLInMeetingID)}
		}

	case "used_as_logo_pdf_header_r_in_meeting_id":
		if m.UsedAsLogoPdfHeaderRInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsLogoPdfHeaderRInMeetingID)}
		}

	case "used_as_logo_projector_header_in_meeting_id":
		if m.UsedAsLogoProjectorHeaderInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsLogoProjectorHeaderInMeetingID)}
		}

	case "used_as_logo_projector_main_in_meeting_id":
		if m.UsedAsLogoProjectorMainInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsLogoProjectorMainInMeetingID)}
		}

	case "used_as_logo_web_header_in_meeting_id":
		if m.UsedAsLogoWebHeaderInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsLogoWebHeaderInMeetingID)}
		}
	}
	return []string{}
}

func (m *MeetingMediafile) Update(data map[string]string) error {
	if val, ok := data["access_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AccessGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["access_group_ids"]; ok {
			m.accessGroups = slices.DeleteFunc(m.accessGroups, func(r *Group) bool {
				return !slices.Contains(m.AccessGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["attachment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AttachmentIDs)
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

	if val, ok := data["inherited_access_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.InheritedAccessGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["inherited_access_group_ids"]; ok {
			m.inheritedAccessGroups = slices.DeleteFunc(m.inheritedAccessGroups, func(r *Group) bool {
				return !slices.Contains(m.InheritedAccessGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["is_public"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsPublic)
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

	if val, ok := data["mediafile_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MediafileID)
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

	if val, ok := data["used_as_font_bold_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontBoldInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_bold_italic_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontBoldItalicInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_chyron_speaker_name_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontChyronSpeakerNameInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_italic_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontItalicInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_monospace_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontMonospaceInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_projector_h1_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontProjectorH1InMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_projector_h2_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontProjectorH2InMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_font_regular_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsFontRegularInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_ballot_paper_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfBallotPaperInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_footer_l_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfFooterLInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_footer_r_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfFooterRInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_header_l_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfHeaderLInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_pdf_header_r_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoPdfHeaderRInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_projector_header_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoProjectorHeaderInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_projector_main_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoProjectorMainInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_logo_web_header_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsLogoWebHeaderInMeetingID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MeetingMediafile) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type MeetingUser struct {
	AboutMe                      *string `json:"about_me"`
	AssignmentCandidateIDs       []int   `json:"assignment_candidate_ids"`
	ChatMessageIDs               []int   `json:"chat_message_ids"`
	Comment                      *string `json:"comment"`
	GroupIDs                     []int   `json:"group_ids"`
	ID                           int     `json:"id"`
	LockedOut                    *bool   `json:"locked_out"`
	MeetingID                    int     `json:"meeting_id"`
	MotionEditorIDs              []int   `json:"motion_editor_ids"`
	MotionSubmitterIDs           []int   `json:"motion_submitter_ids"`
	MotionWorkingGroupSpeakerIDs []int   `json:"motion_working_group_speaker_ids"`
	Number                       *string `json:"number"`
	PersonalNoteIDs              []int   `json:"personal_note_ids"`
	SpeakerIDs                   []int   `json:"speaker_ids"`
	StructureLevelIDs            []int   `json:"structure_level_ids"`
	SupportedMotionIDs           []int   `json:"supported_motion_ids"`
	UserID                       int     `json:"user_id"`
	VoteDelegatedToID            *int    `json:"vote_delegated_to_id"`
	VoteDelegationsFromIDs       []int   `json:"vote_delegations_from_ids"`
	VoteWeight                   *string `json:"vote_weight"`
	loadedRelations              map[string]struct{}
	assignmentCandidates         []*AssignmentCandidate
	chatMessages                 []*ChatMessage
	groups                       []*Group
	meeting                      *Meeting
	motionEditors                []*MotionEditor
	motionSubmitters             []*MotionSubmitter
	motionWorkingGroupSpeakers   []*MotionWorkingGroupSpeaker
	personalNotes                []*PersonalNote
	speakers                     []*Speaker
	structureLevels              []*StructureLevel
	supportedMotions             []*Motion
	user                         *User
	voteDelegatedTo              *MeetingUser
	voteDelegationsFroms         []*MeetingUser
}

func (m *MeetingUser) CollectionName() string {
	return "meeting_user"
}

func (m *MeetingUser) AssignmentCandidates() []*AssignmentCandidate {
	if _, ok := m.loadedRelations["assignment_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access AssignmentCandidates relation of MeetingUser which was not loaded.")
	}

	return m.assignmentCandidates
}

func (m *MeetingUser) ChatMessages() []*ChatMessage {
	if _, ok := m.loadedRelations["chat_message_ids"]; !ok {
		log.Panic().Msg("Tried to access ChatMessages relation of MeetingUser which was not loaded.")
	}

	return m.chatMessages
}

func (m *MeetingUser) Groups() []*Group {
	if _, ok := m.loadedRelations["group_ids"]; !ok {
		log.Panic().Msg("Tried to access Groups relation of MeetingUser which was not loaded.")
	}

	return m.groups
}

func (m *MeetingUser) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MeetingUser which was not loaded.")
	}

	return *m.meeting
}

func (m *MeetingUser) MotionEditors() []*MotionEditor {
	if _, ok := m.loadedRelations["motion_editor_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionEditors relation of MeetingUser which was not loaded.")
	}

	return m.motionEditors
}

func (m *MeetingUser) MotionSubmitters() []*MotionSubmitter {
	if _, ok := m.loadedRelations["motion_submitter_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionSubmitters relation of MeetingUser which was not loaded.")
	}

	return m.motionSubmitters
}

func (m *MeetingUser) MotionWorkingGroupSpeakers() []*MotionWorkingGroupSpeaker {
	if _, ok := m.loadedRelations["motion_working_group_speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionWorkingGroupSpeakers relation of MeetingUser which was not loaded.")
	}

	return m.motionWorkingGroupSpeakers
}

func (m *MeetingUser) PersonalNotes() []*PersonalNote {
	if _, ok := m.loadedRelations["personal_note_ids"]; !ok {
		log.Panic().Msg("Tried to access PersonalNotes relation of MeetingUser which was not loaded.")
	}

	return m.personalNotes
}

func (m *MeetingUser) Speakers() []*Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of MeetingUser which was not loaded.")
	}

	return m.speakers
}

func (m *MeetingUser) StructureLevels() []*StructureLevel {
	if _, ok := m.loadedRelations["structure_level_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevels relation of MeetingUser which was not loaded.")
	}

	return m.structureLevels
}

func (m *MeetingUser) SupportedMotions() []*Motion {
	if _, ok := m.loadedRelations["supported_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access SupportedMotions relation of MeetingUser which was not loaded.")
	}

	return m.supportedMotions
}

func (m *MeetingUser) User() User {
	if _, ok := m.loadedRelations["user_id"]; !ok {
		log.Panic().Msg("Tried to access User relation of MeetingUser which was not loaded.")
	}

	return *m.user
}

func (m *MeetingUser) VoteDelegatedTo() *MeetingUser {
	if _, ok := m.loadedRelations["vote_delegated_to_id"]; !ok {
		log.Panic().Msg("Tried to access VoteDelegatedTo relation of MeetingUser which was not loaded.")
	}

	return m.voteDelegatedTo
}

func (m *MeetingUser) VoteDelegationsFroms() []*MeetingUser {
	if _, ok := m.loadedRelations["vote_delegations_from_ids"]; !ok {
		log.Panic().Msg("Tried to access VoteDelegationsFroms relation of MeetingUser which was not loaded.")
	}

	return m.voteDelegationsFroms
}

func (m *MeetingUser) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "assignment_candidate_ids":
		for _, r := range m.assignmentCandidates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "chat_message_ids":
		for _, r := range m.chatMessages {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "group_ids":
		for _, r := range m.groups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "motion_editor_ids":
		for _, r := range m.motionEditors {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_submitter_ids":
		for _, r := range m.motionSubmitters {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_working_group_speaker_ids":
		for _, r := range m.motionWorkingGroupSpeakers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "personal_note_ids":
		for _, r := range m.personalNotes {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "speaker_ids":
		for _, r := range m.speakers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "structure_level_ids":
		for _, r := range m.structureLevels {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "supported_motion_ids":
		for _, r := range m.supportedMotions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "user_id":
		return m.user.GetRelatedModelsAccessor()
	case "vote_delegated_to_id":
		return m.voteDelegatedTo.GetRelatedModelsAccessor()
	case "vote_delegations_from_ids":
		for _, r := range m.voteDelegationsFroms {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *MeetingUser) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "assignment_candidate_ids":
			m.assignmentCandidates = content.([]*AssignmentCandidate)
		case "chat_message_ids":
			m.chatMessages = content.([]*ChatMessage)
		case "group_ids":
			m.groups = content.([]*Group)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "motion_editor_ids":
			m.motionEditors = content.([]*MotionEditor)
		case "motion_submitter_ids":
			m.motionSubmitters = content.([]*MotionSubmitter)
		case "motion_working_group_speaker_ids":
			m.motionWorkingGroupSpeakers = content.([]*MotionWorkingGroupSpeaker)
		case "personal_note_ids":
			m.personalNotes = content.([]*PersonalNote)
		case "speaker_ids":
			m.speakers = content.([]*Speaker)
		case "structure_level_ids":
			m.structureLevels = content.([]*StructureLevel)
		case "supported_motion_ids":
			m.supportedMotions = content.([]*Motion)
		case "user_id":
			m.user = content.(*User)
		case "vote_delegated_to_id":
			m.voteDelegatedTo = content.(*MeetingUser)
		case "vote_delegations_from_ids":
			m.voteDelegationsFroms = content.([]*MeetingUser)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MeetingUser) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "assignment_candidate_ids":
		var entry AssignmentCandidate
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.assignmentCandidates = append(m.assignmentCandidates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "chat_message_ids":
		var entry ChatMessage
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.chatMessages = append(m.chatMessages, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.groups = append(m.groups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_editor_ids":
		var entry MotionEditor
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionEditors = append(m.motionEditors, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_submitter_ids":
		var entry MotionSubmitter
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionSubmitters = append(m.motionSubmitters, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_working_group_speaker_ids":
		var entry MotionWorkingGroupSpeaker
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionWorkingGroupSpeakers = append(m.motionWorkingGroupSpeakers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "personal_note_ids":
		var entry PersonalNote
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.personalNotes = append(m.personalNotes, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "speaker_ids":
		var entry Speaker
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.speakers = append(m.speakers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "structure_level_ids":
		var entry StructureLevel
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.structureLevels = append(m.structureLevels, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "supported_motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.supportedMotions = append(m.supportedMotions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "user_id":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.user = &entry

		result = entry.GetRelatedModelsAccessor()
	case "vote_delegated_to_id":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.voteDelegatedTo = &entry

		result = entry.GetRelatedModelsAccessor()
	case "vote_delegations_from_ids":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.voteDelegationsFroms = append(m.voteDelegationsFroms, &entry)

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

func (m *MeetingUser) Get(field string) interface{} {
	switch field {
	case "about_me":
		return m.AboutMe
	case "assignment_candidate_ids":
		return m.AssignmentCandidateIDs
	case "chat_message_ids":
		return m.ChatMessageIDs
	case "comment":
		return m.Comment
	case "group_ids":
		return m.GroupIDs
	case "id":
		return m.ID
	case "locked_out":
		return m.LockedOut
	case "meeting_id":
		return m.MeetingID
	case "motion_editor_ids":
		return m.MotionEditorIDs
	case "motion_submitter_ids":
		return m.MotionSubmitterIDs
	case "motion_working_group_speaker_ids":
		return m.MotionWorkingGroupSpeakerIDs
	case "number":
		return m.Number
	case "personal_note_ids":
		return m.PersonalNoteIDs
	case "speaker_ids":
		return m.SpeakerIDs
	case "structure_level_ids":
		return m.StructureLevelIDs
	case "supported_motion_ids":
		return m.SupportedMotionIDs
	case "user_id":
		return m.UserID
	case "vote_delegated_to_id":
		return m.VoteDelegatedToID
	case "vote_delegations_from_ids":
		return m.VoteDelegationsFromIDs
	case "vote_weight":
		return m.VoteWeight
	}

	return nil
}

func (m *MeetingUser) GetFqids(field string) []string {
	switch field {
	case "assignment_candidate_ids":
		r := make([]string, len(m.AssignmentCandidateIDs))
		for i, id := range m.AssignmentCandidateIDs {
			r[i] = "assignment_candidate/" + strconv.Itoa(id)
		}
		return r

	case "chat_message_ids":
		r := make([]string, len(m.ChatMessageIDs))
		for i, id := range m.ChatMessageIDs {
			r[i] = "chat_message/" + strconv.Itoa(id)
		}
		return r

	case "group_ids":
		r := make([]string, len(m.GroupIDs))
		for i, id := range m.GroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "motion_editor_ids":
		r := make([]string, len(m.MotionEditorIDs))
		for i, id := range m.MotionEditorIDs {
			r[i] = "motion_editor/" + strconv.Itoa(id)
		}
		return r

	case "motion_submitter_ids":
		r := make([]string, len(m.MotionSubmitterIDs))
		for i, id := range m.MotionSubmitterIDs {
			r[i] = "motion_submitter/" + strconv.Itoa(id)
		}
		return r

	case "motion_working_group_speaker_ids":
		r := make([]string, len(m.MotionWorkingGroupSpeakerIDs))
		for i, id := range m.MotionWorkingGroupSpeakerIDs {
			r[i] = "motion_working_group_speaker/" + strconv.Itoa(id)
		}
		return r

	case "personal_note_ids":
		r := make([]string, len(m.PersonalNoteIDs))
		for i, id := range m.PersonalNoteIDs {
			r[i] = "personal_note/" + strconv.Itoa(id)
		}
		return r

	case "speaker_ids":
		r := make([]string, len(m.SpeakerIDs))
		for i, id := range m.SpeakerIDs {
			r[i] = "speaker/" + strconv.Itoa(id)
		}
		return r

	case "structure_level_ids":
		r := make([]string, len(m.StructureLevelIDs))
		for i, id := range m.StructureLevelIDs {
			r[i] = "structure_level/" + strconv.Itoa(id)
		}
		return r

	case "supported_motion_ids":
		r := make([]string, len(m.SupportedMotionIDs))
		for i, id := range m.SupportedMotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "user_id":
		return []string{"user/" + strconv.Itoa(m.UserID)}

	case "vote_delegated_to_id":
		if m.VoteDelegatedToID != nil {
			return []string{"meeting_user/" + strconv.Itoa(*m.VoteDelegatedToID)}
		}

	case "vote_delegations_from_ids":
		r := make([]string, len(m.VoteDelegationsFromIDs))
		for i, id := range m.VoteDelegationsFromIDs {
			r[i] = "meeting_user/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *MeetingUser) Update(data map[string]string) error {
	if val, ok := data["about_me"]; ok {
		err := json.Unmarshal([]byte(val), &m.AboutMe)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentCandidateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["assignment_candidate_ids"]; ok {
			m.assignmentCandidates = slices.DeleteFunc(m.assignmentCandidates, func(r *AssignmentCandidate) bool {
				return !slices.Contains(m.AssignmentCandidateIDs, r.ID)
			})
		}
	}

	if val, ok := data["chat_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatMessageIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["chat_message_ids"]; ok {
			m.chatMessages = slices.DeleteFunc(m.chatMessages, func(r *ChatMessage) bool {
				return !slices.Contains(m.ChatMessageIDs, r.ID)
			})
		}
	}

	if val, ok := data["comment"]; ok {
		err := json.Unmarshal([]byte(val), &m.Comment)
		if err != nil {
			return err
		}
	}

	if val, ok := data["group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.GroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["group_ids"]; ok {
			m.groups = slices.DeleteFunc(m.groups, func(r *Group) bool {
				return !slices.Contains(m.GroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["locked_out"]; ok {
		err := json.Unmarshal([]byte(val), &m.LockedOut)
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

	if val, ok := data["motion_editor_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionEditorIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_editor_ids"]; ok {
			m.motionEditors = slices.DeleteFunc(m.motionEditors, func(r *MotionEditor) bool {
				return !slices.Contains(m.MotionEditorIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_submitter_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionSubmitterIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_submitter_ids"]; ok {
			m.motionSubmitters = slices.DeleteFunc(m.motionSubmitters, func(r *MotionSubmitter) bool {
				return !slices.Contains(m.MotionSubmitterIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_working_group_speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionWorkingGroupSpeakerIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_working_group_speaker_ids"]; ok {
			m.motionWorkingGroupSpeakers = slices.DeleteFunc(m.motionWorkingGroupSpeakers, func(r *MotionWorkingGroupSpeaker) bool {
				return !slices.Contains(m.MotionWorkingGroupSpeakerIDs, r.ID)
			})
		}
	}

	if val, ok := data["number"]; ok {
		err := json.Unmarshal([]byte(val), &m.Number)
		if err != nil {
			return err
		}
	}

	if val, ok := data["personal_note_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PersonalNoteIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["personal_note_ids"]; ok {
			m.personalNotes = slices.DeleteFunc(m.personalNotes, func(r *PersonalNote) bool {
				return !slices.Contains(m.PersonalNoteIDs, r.ID)
			})
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

	if val, ok := data["structure_level_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["structure_level_ids"]; ok {
			m.structureLevels = slices.DeleteFunc(m.structureLevels, func(r *StructureLevel) bool {
				return !slices.Contains(m.StructureLevelIDs, r.ID)
			})
		}
	}

	if val, ok := data["supported_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SupportedMotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["supported_motion_ids"]; ok {
			m.supportedMotions = slices.DeleteFunc(m.supportedMotions, func(r *Motion) bool {
				return !slices.Contains(m.SupportedMotionIDs, r.ID)
			})
		}
	}

	if val, ok := data["user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_delegated_to_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteDelegatedToID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_delegations_from_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteDelegationsFromIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["vote_delegations_from_ids"]; ok {
			m.voteDelegationsFroms = slices.DeleteFunc(m.voteDelegationsFroms, func(r *MeetingUser) bool {
				return !slices.Contains(m.VoteDelegationsFromIDs, r.ID)
			})
		}
	}

	if val, ok := data["vote_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteWeight)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MeetingUser) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Motion struct {
	AdditionalSubmitter                          *string         `json:"additional_submitter"`
	AgendaItemID                                 *int            `json:"agenda_item_id"`
	AllDerivedMotionIDs                          []int           `json:"all_derived_motion_ids"`
	AllOriginIDs                                 []int           `json:"all_origin_ids"`
	AmendmentIDs                                 []int           `json:"amendment_ids"`
	AmendmentParagraphs                          json.RawMessage `json:"amendment_paragraphs"`
	AttachmentMeetingMediafileIDs                []int           `json:"attachment_meeting_mediafile_ids"`
	BlockID                                      *int            `json:"block_id"`
	CategoryID                                   *int            `json:"category_id"`
	CategoryWeight                               *int            `json:"category_weight"`
	ChangeRecommendationIDs                      []int           `json:"change_recommendation_ids"`
	CommentIDs                                   []int           `json:"comment_ids"`
	Created                                      *int            `json:"created"`
	DerivedMotionIDs                             []int           `json:"derived_motion_ids"`
	EditorIDs                                    []int           `json:"editor_ids"`
	Forwarded                                    *int            `json:"forwarded"`
	ID                                           int             `json:"id"`
	IDenticalMotionIDs                           []int           `json:"identical_motion_ids"`
	LastModified                                 *int            `json:"last_modified"`
	LeadMotionID                                 *int            `json:"lead_motion_id"`
	ListOfSpeakersID                             int             `json:"list_of_speakers_id"`
	MeetingID                                    int             `json:"meeting_id"`
	ModifiedFinalVersion                         *string         `json:"modified_final_version"`
	Number                                       *string         `json:"number"`
	NumberValue                                  *int            `json:"number_value"`
	OptionIDs                                    []int           `json:"option_ids"`
	OriginID                                     *int            `json:"origin_id"`
	OriginMeetingID                              *int            `json:"origin_meeting_id"`
	PersonalNoteIDs                              []int           `json:"personal_note_ids"`
	PollIDs                                      []int           `json:"poll_ids"`
	ProjectionIDs                                []int           `json:"projection_ids"`
	Reason                                       *string         `json:"reason"`
	RecommendationExtension                      *string         `json:"recommendation_extension"`
	RecommendationExtensionReferenceIDs          []string        `json:"recommendation_extension_reference_ids"`
	RecommendationID                             *int            `json:"recommendation_id"`
	ReferencedInMotionRecommendationExtensionIDs []int           `json:"referenced_in_motion_recommendation_extension_ids"`
	ReferencedInMotionStateExtensionIDs          []int           `json:"referenced_in_motion_state_extension_ids"`
	SequentialNumber                             int             `json:"sequential_number"`
	SortChildIDs                                 []int           `json:"sort_child_ids"`
	SortParentID                                 *int            `json:"sort_parent_id"`
	SortWeight                                   *int            `json:"sort_weight"`
	StartLineNumber                              *int            `json:"start_line_number"`
	StateExtension                               *string         `json:"state_extension"`
	StateExtensionReferenceIDs                   []string        `json:"state_extension_reference_ids"`
	StateID                                      int             `json:"state_id"`
	SubmitterIDs                                 []int           `json:"submitter_ids"`
	SupporterMeetingUserIDs                      []int           `json:"supporter_meeting_user_ids"`
	TagIDs                                       []int           `json:"tag_ids"`
	Text                                         *string         `json:"text"`
	TextHash                                     *string         `json:"text_hash"`
	Title                                        string          `json:"title"`
	WorkflowTimestamp                            *int            `json:"workflow_timestamp"`
	WorkingGroupSpeakerIDs                       []int           `json:"working_group_speaker_ids"`
	loadedRelations                              map[string]struct{}
	agendaItem                                   *AgendaItem
	allDerivedMotions                            []*Motion
	allOrigins                                   []*Motion
	amendments                                   []*Motion
	attachmentMeetingMediafiles                  []*MeetingMediafile
	block                                        *MotionBlock
	category                                     *MotionCategory
	changeRecommendations                        []*MotionChangeRecommendation
	comments                                     []*MotionComment
	derivedMotions                               []*Motion
	editors                                      []*MotionEditor
	iDenticalMotions                             []*Motion
	leadMotion                                   *Motion
	listOfSpeakers                               *ListOfSpeakers
	meeting                                      *Meeting
	options                                      []*Option
	origin                                       *Motion
	originMeeting                                *Meeting
	personalNotes                                []*PersonalNote
	polls                                        []*Poll
	projections                                  []*Projection
	recommendation                               *MotionState
	referencedInMotionRecommendationExtensions   []*Motion
	referencedInMotionStateExtensions            []*Motion
	sortChilds                                   []*Motion
	sortParent                                   *Motion
	state                                        *MotionState
	submitters                                   []*MotionSubmitter
	supporterMeetingUsers                        []*MeetingUser
	tags                                         []*Tag
	workingGroupSpeakers                         []*MotionWorkingGroupSpeaker
}

func (m *Motion) CollectionName() string {
	return "motion"
}

func (m *Motion) AgendaItem() *AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_id"]; !ok {
		log.Panic().Msg("Tried to access AgendaItem relation of Motion which was not loaded.")
	}

	return m.agendaItem
}

func (m *Motion) AllDerivedMotions() []*Motion {
	if _, ok := m.loadedRelations["all_derived_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access AllDerivedMotions relation of Motion which was not loaded.")
	}

	return m.allDerivedMotions
}

func (m *Motion) AllOrigins() []*Motion {
	if _, ok := m.loadedRelations["all_origin_ids"]; !ok {
		log.Panic().Msg("Tried to access AllOrigins relation of Motion which was not loaded.")
	}

	return m.allOrigins
}

func (m *Motion) Amendments() []*Motion {
	if _, ok := m.loadedRelations["amendment_ids"]; !ok {
		log.Panic().Msg("Tried to access Amendments relation of Motion which was not loaded.")
	}

	return m.amendments
}

func (m *Motion) AttachmentMeetingMediafiles() []*MeetingMediafile {
	if _, ok := m.loadedRelations["attachment_meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access AttachmentMeetingMediafiles relation of Motion which was not loaded.")
	}

	return m.attachmentMeetingMediafiles
}

func (m *Motion) Block() *MotionBlock {
	if _, ok := m.loadedRelations["block_id"]; !ok {
		log.Panic().Msg("Tried to access Block relation of Motion which was not loaded.")
	}

	return m.block
}

func (m *Motion) Category() *MotionCategory {
	if _, ok := m.loadedRelations["category_id"]; !ok {
		log.Panic().Msg("Tried to access Category relation of Motion which was not loaded.")
	}

	return m.category
}

func (m *Motion) ChangeRecommendations() []*MotionChangeRecommendation {
	if _, ok := m.loadedRelations["change_recommendation_ids"]; !ok {
		log.Panic().Msg("Tried to access ChangeRecommendations relation of Motion which was not loaded.")
	}

	return m.changeRecommendations
}

func (m *Motion) Comments() []*MotionComment {
	if _, ok := m.loadedRelations["comment_ids"]; !ok {
		log.Panic().Msg("Tried to access Comments relation of Motion which was not loaded.")
	}

	return m.comments
}

func (m *Motion) DerivedMotions() []*Motion {
	if _, ok := m.loadedRelations["derived_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access DerivedMotions relation of Motion which was not loaded.")
	}

	return m.derivedMotions
}

func (m *Motion) Editors() []*MotionEditor {
	if _, ok := m.loadedRelations["editor_ids"]; !ok {
		log.Panic().Msg("Tried to access Editors relation of Motion which was not loaded.")
	}

	return m.editors
}

func (m *Motion) IDenticalMotions() []*Motion {
	if _, ok := m.loadedRelations["identical_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access IDenticalMotions relation of Motion which was not loaded.")
	}

	return m.iDenticalMotions
}

func (m *Motion) LeadMotion() *Motion {
	if _, ok := m.loadedRelations["lead_motion_id"]; !ok {
		log.Panic().Msg("Tried to access LeadMotion relation of Motion which was not loaded.")
	}

	return m.leadMotion
}

func (m *Motion) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of Motion which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *Motion) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Motion which was not loaded.")
	}

	return *m.meeting
}

func (m *Motion) Options() []*Option {
	if _, ok := m.loadedRelations["option_ids"]; !ok {
		log.Panic().Msg("Tried to access Options relation of Motion which was not loaded.")
	}

	return m.options
}

func (m *Motion) Origin() *Motion {
	if _, ok := m.loadedRelations["origin_id"]; !ok {
		log.Panic().Msg("Tried to access Origin relation of Motion which was not loaded.")
	}

	return m.origin
}

func (m *Motion) OriginMeeting() *Meeting {
	if _, ok := m.loadedRelations["origin_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access OriginMeeting relation of Motion which was not loaded.")
	}

	return m.originMeeting
}

func (m *Motion) PersonalNotes() []*PersonalNote {
	if _, ok := m.loadedRelations["personal_note_ids"]; !ok {
		log.Panic().Msg("Tried to access PersonalNotes relation of Motion which was not loaded.")
	}

	return m.personalNotes
}

func (m *Motion) Polls() []*Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Motion which was not loaded.")
	}

	return m.polls
}

func (m *Motion) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of Motion which was not loaded.")
	}

	return m.projections
}

func (m *Motion) Recommendation() *MotionState {
	if _, ok := m.loadedRelations["recommendation_id"]; !ok {
		log.Panic().Msg("Tried to access Recommendation relation of Motion which was not loaded.")
	}

	return m.recommendation
}

func (m *Motion) ReferencedInMotionRecommendationExtensions() []*Motion {
	if _, ok := m.loadedRelations["referenced_in_motion_recommendation_extension_ids"]; !ok {
		log.Panic().Msg("Tried to access ReferencedInMotionRecommendationExtensions relation of Motion which was not loaded.")
	}

	return m.referencedInMotionRecommendationExtensions
}

func (m *Motion) ReferencedInMotionStateExtensions() []*Motion {
	if _, ok := m.loadedRelations["referenced_in_motion_state_extension_ids"]; !ok {
		log.Panic().Msg("Tried to access ReferencedInMotionStateExtensions relation of Motion which was not loaded.")
	}

	return m.referencedInMotionStateExtensions
}

func (m *Motion) SortChilds() []*Motion {
	if _, ok := m.loadedRelations["sort_child_ids"]; !ok {
		log.Panic().Msg("Tried to access SortChilds relation of Motion which was not loaded.")
	}

	return m.sortChilds
}

func (m *Motion) SortParent() *Motion {
	if _, ok := m.loadedRelations["sort_parent_id"]; !ok {
		log.Panic().Msg("Tried to access SortParent relation of Motion which was not loaded.")
	}

	return m.sortParent
}

func (m *Motion) State() MotionState {
	if _, ok := m.loadedRelations["state_id"]; !ok {
		log.Panic().Msg("Tried to access State relation of Motion which was not loaded.")
	}

	return *m.state
}

func (m *Motion) Submitters() []*MotionSubmitter {
	if _, ok := m.loadedRelations["submitter_ids"]; !ok {
		log.Panic().Msg("Tried to access Submitters relation of Motion which was not loaded.")
	}

	return m.submitters
}

func (m *Motion) SupporterMeetingUsers() []*MeetingUser {
	if _, ok := m.loadedRelations["supporter_meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access SupporterMeetingUsers relation of Motion which was not loaded.")
	}

	return m.supporterMeetingUsers
}

func (m *Motion) Tags() []*Tag {
	if _, ok := m.loadedRelations["tag_ids"]; !ok {
		log.Panic().Msg("Tried to access Tags relation of Motion which was not loaded.")
	}

	return m.tags
}

func (m *Motion) WorkingGroupSpeakers() []*MotionWorkingGroupSpeaker {
	if _, ok := m.loadedRelations["working_group_speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access WorkingGroupSpeakers relation of Motion which was not loaded.")
	}

	return m.workingGroupSpeakers
}

func (m *Motion) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "agenda_item_id":
		return m.agendaItem.GetRelatedModelsAccessor()
	case "all_derived_motion_ids":
		for _, r := range m.allDerivedMotions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "all_origin_ids":
		for _, r := range m.allOrigins {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "amendment_ids":
		for _, r := range m.amendments {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "attachment_meeting_mediafile_ids":
		for _, r := range m.attachmentMeetingMediafiles {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "block_id":
		return m.block.GetRelatedModelsAccessor()
	case "category_id":
		return m.category.GetRelatedModelsAccessor()
	case "change_recommendation_ids":
		for _, r := range m.changeRecommendations {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "comment_ids":
		for _, r := range m.comments {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "derived_motion_ids":
		for _, r := range m.derivedMotions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "editor_ids":
		for _, r := range m.editors {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "identical_motion_ids":
		for _, r := range m.iDenticalMotions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "lead_motion_id":
		return m.leadMotion.GetRelatedModelsAccessor()
	case "list_of_speakers_id":
		return m.listOfSpeakers.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "option_ids":
		for _, r := range m.options {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "origin_id":
		return m.origin.GetRelatedModelsAccessor()
	case "origin_meeting_id":
		return m.originMeeting.GetRelatedModelsAccessor()
	case "personal_note_ids":
		for _, r := range m.personalNotes {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
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
	case "recommendation_id":
		return m.recommendation.GetRelatedModelsAccessor()
	case "referenced_in_motion_recommendation_extension_ids":
		for _, r := range m.referencedInMotionRecommendationExtensions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "referenced_in_motion_state_extension_ids":
		for _, r := range m.referencedInMotionStateExtensions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "sort_child_ids":
		for _, r := range m.sortChilds {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "sort_parent_id":
		return m.sortParent.GetRelatedModelsAccessor()
	case "state_id":
		return m.state.GetRelatedModelsAccessor()
	case "submitter_ids":
		for _, r := range m.submitters {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "supporter_meeting_user_ids":
		for _, r := range m.supporterMeetingUsers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "tag_ids":
		for _, r := range m.tags {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "working_group_speaker_ids":
		for _, r := range m.workingGroupSpeakers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Motion) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "agenda_item_id":
			m.agendaItem = content.(*AgendaItem)
		case "all_derived_motion_ids":
			m.allDerivedMotions = content.([]*Motion)
		case "all_origin_ids":
			m.allOrigins = content.([]*Motion)
		case "amendment_ids":
			m.amendments = content.([]*Motion)
		case "attachment_meeting_mediafile_ids":
			m.attachmentMeetingMediafiles = content.([]*MeetingMediafile)
		case "block_id":
			m.block = content.(*MotionBlock)
		case "category_id":
			m.category = content.(*MotionCategory)
		case "change_recommendation_ids":
			m.changeRecommendations = content.([]*MotionChangeRecommendation)
		case "comment_ids":
			m.comments = content.([]*MotionComment)
		case "derived_motion_ids":
			m.derivedMotions = content.([]*Motion)
		case "editor_ids":
			m.editors = content.([]*MotionEditor)
		case "identical_motion_ids":
			m.iDenticalMotions = content.([]*Motion)
		case "lead_motion_id":
			m.leadMotion = content.(*Motion)
		case "list_of_speakers_id":
			m.listOfSpeakers = content.(*ListOfSpeakers)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "option_ids":
			m.options = content.([]*Option)
		case "origin_id":
			m.origin = content.(*Motion)
		case "origin_meeting_id":
			m.originMeeting = content.(*Meeting)
		case "personal_note_ids":
			m.personalNotes = content.([]*PersonalNote)
		case "poll_ids":
			m.polls = content.([]*Poll)
		case "projection_ids":
			m.projections = content.([]*Projection)
		case "recommendation_id":
			m.recommendation = content.(*MotionState)
		case "referenced_in_motion_recommendation_extension_ids":
			m.referencedInMotionRecommendationExtensions = content.([]*Motion)
		case "referenced_in_motion_state_extension_ids":
			m.referencedInMotionStateExtensions = content.([]*Motion)
		case "sort_child_ids":
			m.sortChilds = content.([]*Motion)
		case "sort_parent_id":
			m.sortParent = content.(*Motion)
		case "state_id":
			m.state = content.(*MotionState)
		case "submitter_ids":
			m.submitters = content.([]*MotionSubmitter)
		case "supporter_meeting_user_ids":
			m.supporterMeetingUsers = content.([]*MeetingUser)
		case "tag_ids":
			m.tags = content.([]*Tag)
		case "working_group_speaker_ids":
			m.workingGroupSpeakers = content.([]*MotionWorkingGroupSpeaker)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Motion) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "all_derived_motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.allDerivedMotions = append(m.allDerivedMotions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "all_origin_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.allOrigins = append(m.allOrigins, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "amendment_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.amendments = append(m.amendments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "attachment_meeting_mediafile_ids":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.attachmentMeetingMediafiles = append(m.attachmentMeetingMediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "block_id":
		var entry MotionBlock
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.block = &entry

		result = entry.GetRelatedModelsAccessor()
	case "category_id":
		var entry MotionCategory
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.category = &entry

		result = entry.GetRelatedModelsAccessor()
	case "change_recommendation_ids":
		var entry MotionChangeRecommendation
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.changeRecommendations = append(m.changeRecommendations, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "comment_ids":
		var entry MotionComment
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.comments = append(m.comments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "derived_motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.derivedMotions = append(m.derivedMotions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "editor_ids":
		var entry MotionEditor
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.editors = append(m.editors, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "identical_motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.iDenticalMotions = append(m.iDenticalMotions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "lead_motion_id":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.leadMotion = &entry

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
	case "option_ids":
		var entry Option
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.options = append(m.options, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "origin_id":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.origin = &entry

		result = entry.GetRelatedModelsAccessor()
	case "origin_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.originMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "personal_note_ids":
		var entry PersonalNote
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.personalNotes = append(m.personalNotes, &entry)

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
	case "recommendation_id":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.recommendation = &entry

		result = entry.GetRelatedModelsAccessor()
	case "referenced_in_motion_recommendation_extension_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.referencedInMotionRecommendationExtensions = append(m.referencedInMotionRecommendationExtensions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "referenced_in_motion_state_extension_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.referencedInMotionStateExtensions = append(m.referencedInMotionStateExtensions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "sort_child_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.sortChilds = append(m.sortChilds, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "sort_parent_id":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.sortParent = &entry

		result = entry.GetRelatedModelsAccessor()
	case "state_id":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.state = &entry

		result = entry.GetRelatedModelsAccessor()
	case "submitter_ids":
		var entry MotionSubmitter
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.submitters = append(m.submitters, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "supporter_meeting_user_ids":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.supporterMeetingUsers = append(m.supporterMeetingUsers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "tag_ids":
		var entry Tag
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.tags = append(m.tags, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "working_group_speaker_ids":
		var entry MotionWorkingGroupSpeaker
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.workingGroupSpeakers = append(m.workingGroupSpeakers, &entry)

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

func (m *Motion) Get(field string) interface{} {
	switch field {
	case "additional_submitter":
		return m.AdditionalSubmitter
	case "agenda_item_id":
		return m.AgendaItemID
	case "all_derived_motion_ids":
		return m.AllDerivedMotionIDs
	case "all_origin_ids":
		return m.AllOriginIDs
	case "amendment_ids":
		return m.AmendmentIDs
	case "amendment_paragraphs":
		return m.AmendmentParagraphs
	case "attachment_meeting_mediafile_ids":
		return m.AttachmentMeetingMediafileIDs
	case "block_id":
		return m.BlockID
	case "category_id":
		return m.CategoryID
	case "category_weight":
		return m.CategoryWeight
	case "change_recommendation_ids":
		return m.ChangeRecommendationIDs
	case "comment_ids":
		return m.CommentIDs
	case "created":
		return m.Created
	case "derived_motion_ids":
		return m.DerivedMotionIDs
	case "editor_ids":
		return m.EditorIDs
	case "forwarded":
		return m.Forwarded
	case "id":
		return m.ID
	case "identical_motion_ids":
		return m.IDenticalMotionIDs
	case "last_modified":
		return m.LastModified
	case "lead_motion_id":
		return m.LeadMotionID
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "modified_final_version":
		return m.ModifiedFinalVersion
	case "number":
		return m.Number
	case "number_value":
		return m.NumberValue
	case "option_ids":
		return m.OptionIDs
	case "origin_id":
		return m.OriginID
	case "origin_meeting_id":
		return m.OriginMeetingID
	case "personal_note_ids":
		return m.PersonalNoteIDs
	case "poll_ids":
		return m.PollIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "reason":
		return m.Reason
	case "recommendation_extension":
		return m.RecommendationExtension
	case "recommendation_extension_reference_ids":
		return m.RecommendationExtensionReferenceIDs
	case "recommendation_id":
		return m.RecommendationID
	case "referenced_in_motion_recommendation_extension_ids":
		return m.ReferencedInMotionRecommendationExtensionIDs
	case "referenced_in_motion_state_extension_ids":
		return m.ReferencedInMotionStateExtensionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "sort_child_ids":
		return m.SortChildIDs
	case "sort_parent_id":
		return m.SortParentID
	case "sort_weight":
		return m.SortWeight
	case "start_line_number":
		return m.StartLineNumber
	case "state_extension":
		return m.StateExtension
	case "state_extension_reference_ids":
		return m.StateExtensionReferenceIDs
	case "state_id":
		return m.StateID
	case "submitter_ids":
		return m.SubmitterIDs
	case "supporter_meeting_user_ids":
		return m.SupporterMeetingUserIDs
	case "tag_ids":
		return m.TagIDs
	case "text":
		return m.Text
	case "text_hash":
		return m.TextHash
	case "title":
		return m.Title
	case "workflow_timestamp":
		return m.WorkflowTimestamp
	case "working_group_speaker_ids":
		return m.WorkingGroupSpeakerIDs
	}

	return nil
}

func (m *Motion) GetFqids(field string) []string {
	switch field {
	case "agenda_item_id":
		if m.AgendaItemID != nil {
			return []string{"agenda_item/" + strconv.Itoa(*m.AgendaItemID)}
		}

	case "all_derived_motion_ids":
		r := make([]string, len(m.AllDerivedMotionIDs))
		for i, id := range m.AllDerivedMotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "all_origin_ids":
		r := make([]string, len(m.AllOriginIDs))
		for i, id := range m.AllOriginIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "amendment_ids":
		r := make([]string, len(m.AmendmentIDs))
		for i, id := range m.AmendmentIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "attachment_meeting_mediafile_ids":
		r := make([]string, len(m.AttachmentMeetingMediafileIDs))
		for i, id := range m.AttachmentMeetingMediafileIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "block_id":
		if m.BlockID != nil {
			return []string{"motion_block/" + strconv.Itoa(*m.BlockID)}
		}

	case "category_id":
		if m.CategoryID != nil {
			return []string{"motion_category/" + strconv.Itoa(*m.CategoryID)}
		}

	case "change_recommendation_ids":
		r := make([]string, len(m.ChangeRecommendationIDs))
		for i, id := range m.ChangeRecommendationIDs {
			r[i] = "motion_change_recommendation/" + strconv.Itoa(id)
		}
		return r

	case "comment_ids":
		r := make([]string, len(m.CommentIDs))
		for i, id := range m.CommentIDs {
			r[i] = "motion_comment/" + strconv.Itoa(id)
		}
		return r

	case "derived_motion_ids":
		r := make([]string, len(m.DerivedMotionIDs))
		for i, id := range m.DerivedMotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "editor_ids":
		r := make([]string, len(m.EditorIDs))
		for i, id := range m.EditorIDs {
			r[i] = "motion_editor/" + strconv.Itoa(id)
		}
		return r

	case "identical_motion_ids":
		r := make([]string, len(m.IDenticalMotionIDs))
		for i, id := range m.IDenticalMotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "lead_motion_id":
		if m.LeadMotionID != nil {
			return []string{"motion/" + strconv.Itoa(*m.LeadMotionID)}
		}

	case "list_of_speakers_id":
		return []string{"list_of_speakers/" + strconv.Itoa(m.ListOfSpeakersID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "option_ids":
		r := make([]string, len(m.OptionIDs))
		for i, id := range m.OptionIDs {
			r[i] = "option/" + strconv.Itoa(id)
		}
		return r

	case "origin_id":
		if m.OriginID != nil {
			return []string{"motion/" + strconv.Itoa(*m.OriginID)}
		}

	case "origin_meeting_id":
		if m.OriginMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.OriginMeetingID)}
		}

	case "personal_note_ids":
		r := make([]string, len(m.PersonalNoteIDs))
		for i, id := range m.PersonalNoteIDs {
			r[i] = "personal_note/" + strconv.Itoa(id)
		}
		return r

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

	case "recommendation_id":
		if m.RecommendationID != nil {
			return []string{"motion_state/" + strconv.Itoa(*m.RecommendationID)}
		}

	case "referenced_in_motion_recommendation_extension_ids":
		r := make([]string, len(m.ReferencedInMotionRecommendationExtensionIDs))
		for i, id := range m.ReferencedInMotionRecommendationExtensionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "referenced_in_motion_state_extension_ids":
		r := make([]string, len(m.ReferencedInMotionStateExtensionIDs))
		for i, id := range m.ReferencedInMotionStateExtensionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "sort_child_ids":
		r := make([]string, len(m.SortChildIDs))
		for i, id := range m.SortChildIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "sort_parent_id":
		if m.SortParentID != nil {
			return []string{"motion/" + strconv.Itoa(*m.SortParentID)}
		}

	case "state_id":
		return []string{"motion_state/" + strconv.Itoa(m.StateID)}

	case "submitter_ids":
		r := make([]string, len(m.SubmitterIDs))
		for i, id := range m.SubmitterIDs {
			r[i] = "motion_submitter/" + strconv.Itoa(id)
		}
		return r

	case "supporter_meeting_user_ids":
		r := make([]string, len(m.SupporterMeetingUserIDs))
		for i, id := range m.SupporterMeetingUserIDs {
			r[i] = "meeting_user/" + strconv.Itoa(id)
		}
		return r

	case "tag_ids":
		r := make([]string, len(m.TagIDs))
		for i, id := range m.TagIDs {
			r[i] = "tag/" + strconv.Itoa(id)
		}
		return r

	case "working_group_speaker_ids":
		r := make([]string, len(m.WorkingGroupSpeakerIDs))
		for i, id := range m.WorkingGroupSpeakerIDs {
			r[i] = "motion_working_group_speaker/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Motion) Update(data map[string]string) error {
	if val, ok := data["additional_submitter"]; ok {
		err := json.Unmarshal([]byte(val), &m.AdditionalSubmitter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_item_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["all_derived_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllDerivedMotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["all_derived_motion_ids"]; ok {
			m.allDerivedMotions = slices.DeleteFunc(m.allDerivedMotions, func(r *Motion) bool {
				return !slices.Contains(m.AllDerivedMotionIDs, r.ID)
			})
		}
	}

	if val, ok := data["all_origin_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllOriginIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["all_origin_ids"]; ok {
			m.allOrigins = slices.DeleteFunc(m.allOrigins, func(r *Motion) bool {
				return !slices.Contains(m.AllOriginIDs, r.ID)
			})
		}
	}

	if val, ok := data["amendment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AmendmentIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["amendment_ids"]; ok {
			m.amendments = slices.DeleteFunc(m.amendments, func(r *Motion) bool {
				return !slices.Contains(m.AmendmentIDs, r.ID)
			})
		}
	}

	if val, ok := data["amendment_paragraphs"]; ok {
		err := json.Unmarshal([]byte(val), &m.AmendmentParagraphs)
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

	if val, ok := data["block_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.BlockID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["category_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.CategoryID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["category_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.CategoryWeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["change_recommendation_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChangeRecommendationIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["change_recommendation_ids"]; ok {
			m.changeRecommendations = slices.DeleteFunc(m.changeRecommendations, func(r *MotionChangeRecommendation) bool {
				return !slices.Contains(m.ChangeRecommendationIDs, r.ID)
			})
		}
	}

	if val, ok := data["comment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommentIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["comment_ids"]; ok {
			m.comments = slices.DeleteFunc(m.comments, func(r *MotionComment) bool {
				return !slices.Contains(m.CommentIDs, r.ID)
			})
		}
	}

	if val, ok := data["created"]; ok {
		err := json.Unmarshal([]byte(val), &m.Created)
		if err != nil {
			return err
		}
	}

	if val, ok := data["derived_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DerivedMotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["derived_motion_ids"]; ok {
			m.derivedMotions = slices.DeleteFunc(m.derivedMotions, func(r *Motion) bool {
				return !slices.Contains(m.DerivedMotionIDs, r.ID)
			})
		}
	}

	if val, ok := data["editor_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.EditorIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["editor_ids"]; ok {
			m.editors = slices.DeleteFunc(m.editors, func(r *MotionEditor) bool {
				return !slices.Contains(m.EditorIDs, r.ID)
			})
		}
	}

	if val, ok := data["forwarded"]; ok {
		err := json.Unmarshal([]byte(val), &m.Forwarded)
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

	if val, ok := data["identical_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.IDenticalMotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["identical_motion_ids"]; ok {
			m.iDenticalMotions = slices.DeleteFunc(m.iDenticalMotions, func(r *Motion) bool {
				return !slices.Contains(m.IDenticalMotionIDs, r.ID)
			})
		}
	}

	if val, ok := data["last_modified"]; ok {
		err := json.Unmarshal([]byte(val), &m.LastModified)
		if err != nil {
			return err
		}
	}

	if val, ok := data["lead_motion_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LeadMotionID)
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

	if val, ok := data["modified_final_version"]; ok {
		err := json.Unmarshal([]byte(val), &m.ModifiedFinalVersion)
		if err != nil {
			return err
		}
	}

	if val, ok := data["number"]; ok {
		err := json.Unmarshal([]byte(val), &m.Number)
		if err != nil {
			return err
		}
	}

	if val, ok := data["number_value"]; ok {
		err := json.Unmarshal([]byte(val), &m.NumberValue)
		if err != nil {
			return err
		}
	}

	if val, ok := data["option_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["option_ids"]; ok {
			m.options = slices.DeleteFunc(m.options, func(r *Option) bool {
				return !slices.Contains(m.OptionIDs, r.ID)
			})
		}
	}

	if val, ok := data["origin_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OriginID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["origin_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OriginMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["personal_note_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PersonalNoteIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["personal_note_ids"]; ok {
			m.personalNotes = slices.DeleteFunc(m.personalNotes, func(r *PersonalNote) bool {
				return !slices.Contains(m.PersonalNoteIDs, r.ID)
			})
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

	if val, ok := data["reason"]; ok {
		err := json.Unmarshal([]byte(val), &m.Reason)
		if err != nil {
			return err
		}
	}

	if val, ok := data["recommendation_extension"]; ok {
		err := json.Unmarshal([]byte(val), &m.RecommendationExtension)
		if err != nil {
			return err
		}
	}

	if val, ok := data["recommendation_extension_reference_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.RecommendationExtensionReferenceIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["recommendation_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.RecommendationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["referenced_in_motion_recommendation_extension_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReferencedInMotionRecommendationExtensionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["referenced_in_motion_recommendation_extension_ids"]; ok {
			m.referencedInMotionRecommendationExtensions = slices.DeleteFunc(m.referencedInMotionRecommendationExtensions, func(r *Motion) bool {
				return !slices.Contains(m.ReferencedInMotionRecommendationExtensionIDs, r.ID)
			})
		}
	}

	if val, ok := data["referenced_in_motion_state_extension_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReferencedInMotionStateExtensionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["referenced_in_motion_state_extension_ids"]; ok {
			m.referencedInMotionStateExtensions = slices.DeleteFunc(m.referencedInMotionStateExtensions, func(r *Motion) bool {
				return !slices.Contains(m.ReferencedInMotionStateExtensionIDs, r.ID)
			})
		}
	}

	if val, ok := data["sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sort_child_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SortChildIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["sort_child_ids"]; ok {
			m.sortChilds = slices.DeleteFunc(m.sortChilds, func(r *Motion) bool {
				return !slices.Contains(m.SortChildIDs, r.ID)
			})
		}
	}

	if val, ok := data["sort_parent_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.SortParentID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["sort_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.SortWeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["start_line_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.StartLineNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state_extension"]; ok {
		err := json.Unmarshal([]byte(val), &m.StateExtension)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state_extension_reference_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StateExtensionReferenceIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["state_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.StateID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["submitter_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SubmitterIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["submitter_ids"]; ok {
			m.submitters = slices.DeleteFunc(m.submitters, func(r *MotionSubmitter) bool {
				return !slices.Contains(m.SubmitterIDs, r.ID)
			})
		}
	}

	if val, ok := data["supporter_meeting_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SupporterMeetingUserIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["supporter_meeting_user_ids"]; ok {
			m.supporterMeetingUsers = slices.DeleteFunc(m.supporterMeetingUsers, func(r *MeetingUser) bool {
				return !slices.Contains(m.SupporterMeetingUserIDs, r.ID)
			})
		}
	}

	if val, ok := data["tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TagIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["tag_ids"]; ok {
			m.tags = slices.DeleteFunc(m.tags, func(r *Tag) bool {
				return !slices.Contains(m.TagIDs, r.ID)
			})
		}
	}

	if val, ok := data["text"]; ok {
		err := json.Unmarshal([]byte(val), &m.Text)
		if err != nil {
			return err
		}
	}

	if val, ok := data["text_hash"]; ok {
		err := json.Unmarshal([]byte(val), &m.TextHash)
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

	if val, ok := data["workflow_timestamp"]; ok {
		err := json.Unmarshal([]byte(val), &m.WorkflowTimestamp)
		if err != nil {
			return err
		}
	}

	if val, ok := data["working_group_speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WorkingGroupSpeakerIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["working_group_speaker_ids"]; ok {
			m.workingGroupSpeakers = slices.DeleteFunc(m.workingGroupSpeakers, func(r *MotionWorkingGroupSpeaker) bool {
				return !slices.Contains(m.WorkingGroupSpeakerIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *Motion) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type MotionBlock struct {
	AgendaItemID     *int   `json:"agenda_item_id"`
	ID               int    `json:"id"`
	Internal         *bool  `json:"internal"`
	ListOfSpeakersID int    `json:"list_of_speakers_id"`
	MeetingID        int    `json:"meeting_id"`
	MotionIDs        []int  `json:"motion_ids"`
	ProjectionIDs    []int  `json:"projection_ids"`
	SequentialNumber int    `json:"sequential_number"`
	Title            string `json:"title"`
	loadedRelations  map[string]struct{}
	agendaItem       *AgendaItem
	listOfSpeakers   *ListOfSpeakers
	meeting          *Meeting
	motions          []*Motion
	projections      []*Projection
}

func (m *MotionBlock) CollectionName() string {
	return "motion_block"
}

func (m *MotionBlock) AgendaItem() *AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_id"]; !ok {
		log.Panic().Msg("Tried to access AgendaItem relation of MotionBlock which was not loaded.")
	}

	return m.agendaItem
}

func (m *MotionBlock) ListOfSpeakers() ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakers relation of MotionBlock which was not loaded.")
	}

	return *m.listOfSpeakers
}

func (m *MotionBlock) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionBlock which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionBlock) Motions() []*Motion {
	if _, ok := m.loadedRelations["motion_ids"]; !ok {
		log.Panic().Msg("Tried to access Motions relation of MotionBlock which was not loaded.")
	}

	return m.motions
}

func (m *MotionBlock) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of MotionBlock which was not loaded.")
	}

	return m.projections
}

func (m *MotionBlock) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "agenda_item_id":
		return m.agendaItem.GetRelatedModelsAccessor()
	case "list_of_speakers_id":
		return m.listOfSpeakers.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "motion_ids":
		for _, r := range m.motions {
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

func (m *MotionBlock) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "agenda_item_id":
			m.agendaItem = content.(*AgendaItem)
		case "list_of_speakers_id":
			m.listOfSpeakers = content.(*ListOfSpeakers)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "motion_ids":
			m.motions = content.([]*Motion)
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

func (m *MotionBlock) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motions = append(m.motions, &entry)

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

func (m *MotionBlock) Get(field string) interface{} {
	switch field {
	case "agenda_item_id":
		return m.AgendaItemID
	case "id":
		return m.ID
	case "internal":
		return m.Internal
	case "list_of_speakers_id":
		return m.ListOfSpeakersID
	case "meeting_id":
		return m.MeetingID
	case "motion_ids":
		return m.MotionIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "title":
		return m.Title
	}

	return nil
}

func (m *MotionBlock) GetFqids(field string) []string {
	switch field {
	case "agenda_item_id":
		if m.AgendaItemID != nil {
			return []string{"agenda_item/" + strconv.Itoa(*m.AgendaItemID)}
		}

	case "list_of_speakers_id":
		return []string{"list_of_speakers/" + strconv.Itoa(m.ListOfSpeakersID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "motion_ids":
		r := make([]string, len(m.MotionIDs))
		for i, id := range m.MotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
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

func (m *MotionBlock) Update(data map[string]string) error {
	if val, ok := data["agenda_item_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemID)
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

	if val, ok := data["internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.Internal)
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

	if val, ok := data["motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_ids"]; ok {
			m.motions = slices.DeleteFunc(m.motions, func(r *Motion) bool {
				return !slices.Contains(m.MotionIDs, r.ID)
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

	if val, ok := data["title"]; ok {
		err := json.Unmarshal([]byte(val), &m.Title)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionBlock) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

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

func (m *MotionCategory) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "child_ids":
		for _, r := range m.childs {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "motion_ids":
		for _, r := range m.motions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "parent_id":
		return m.parent.GetRelatedModelsAccessor()
	}

	return nil
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

		if _, ok := m.loadedRelations["child_ids"]; ok {
			m.childs = slices.DeleteFunc(m.childs, func(r *MotionCategory) bool {
				return !slices.Contains(m.ChildIDs, r.ID)
			})
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

		if _, ok := m.loadedRelations["motion_ids"]; ok {
			m.motions = slices.DeleteFunc(m.motions, func(r *Motion) bool {
				return !slices.Contains(m.MotionIDs, r.ID)
			})
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
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type MotionChangeRecommendation struct {
	CreationTime     *int    `json:"creation_time"`
	ID               int     `json:"id"`
	Internal         *bool   `json:"internal"`
	LineFrom         *int    `json:"line_from"`
	LineTo           *int    `json:"line_to"`
	MeetingID        int     `json:"meeting_id"`
	MotionID         int     `json:"motion_id"`
	OtherDescription *string `json:"other_description"`
	Rejected         *bool   `json:"rejected"`
	Text             *string `json:"text"`
	Type             *string `json:"type"`
	loadedRelations  map[string]struct{}
	meeting          *Meeting
	motion           *Motion
}

func (m *MotionChangeRecommendation) CollectionName() string {
	return "motion_change_recommendation"
}

func (m *MotionChangeRecommendation) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionChangeRecommendation which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionChangeRecommendation) Motion() Motion {
	if _, ok := m.loadedRelations["motion_id"]; !ok {
		log.Panic().Msg("Tried to access Motion relation of MotionChangeRecommendation which was not loaded.")
	}

	return *m.motion
}

func (m *MotionChangeRecommendation) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "motion_id":
		return m.motion.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *MotionChangeRecommendation) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "motion_id":
			m.motion = content.(*Motion)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionChangeRecommendation) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	default:
		return nil, fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return result, nil
}

func (m *MotionChangeRecommendation) Get(field string) interface{} {
	switch field {
	case "creation_time":
		return m.CreationTime
	case "id":
		return m.ID
	case "internal":
		return m.Internal
	case "line_from":
		return m.LineFrom
	case "line_to":
		return m.LineTo
	case "meeting_id":
		return m.MeetingID
	case "motion_id":
		return m.MotionID
	case "other_description":
		return m.OtherDescription
	case "rejected":
		return m.Rejected
	case "text":
		return m.Text
	case "type":
		return m.Type
	}

	return nil
}

func (m *MotionChangeRecommendation) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "motion_id":
		return []string{"motion/" + strconv.Itoa(m.MotionID)}
	}
	return []string{}
}

func (m *MotionChangeRecommendation) Update(data map[string]string) error {
	if val, ok := data["creation_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.CreationTime)
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

	if val, ok := data["internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.Internal)
		if err != nil {
			return err
		}
	}

	if val, ok := data["line_from"]; ok {
		err := json.Unmarshal([]byte(val), &m.LineFrom)
		if err != nil {
			return err
		}
	}

	if val, ok := data["line_to"]; ok {
		err := json.Unmarshal([]byte(val), &m.LineTo)
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

	if val, ok := data["other_description"]; ok {
		err := json.Unmarshal([]byte(val), &m.OtherDescription)
		if err != nil {
			return err
		}
	}

	if val, ok := data["rejected"]; ok {
		err := json.Unmarshal([]byte(val), &m.Rejected)
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

	if val, ok := data["type"]; ok {
		err := json.Unmarshal([]byte(val), &m.Type)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionChangeRecommendation) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

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

type MotionCommentSection struct {
	CommentIDs        []int  `json:"comment_ids"`
	ID                int    `json:"id"`
	MeetingID         int    `json:"meeting_id"`
	Name              string `json:"name"`
	ReadGroupIDs      []int  `json:"read_group_ids"`
	SequentialNumber  int    `json:"sequential_number"`
	SubmitterCanWrite *bool  `json:"submitter_can_write"`
	Weight            *int   `json:"weight"`
	WriteGroupIDs     []int  `json:"write_group_ids"`
	loadedRelations   map[string]struct{}
	comments          []*MotionComment
	meeting           *Meeting
	readGroups        []*Group
	writeGroups       []*Group
}

func (m *MotionCommentSection) CollectionName() string {
	return "motion_comment_section"
}

func (m *MotionCommentSection) Comments() []*MotionComment {
	if _, ok := m.loadedRelations["comment_ids"]; !ok {
		log.Panic().Msg("Tried to access Comments relation of MotionCommentSection which was not loaded.")
	}

	return m.comments
}

func (m *MotionCommentSection) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionCommentSection which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionCommentSection) ReadGroups() []*Group {
	if _, ok := m.loadedRelations["read_group_ids"]; !ok {
		log.Panic().Msg("Tried to access ReadGroups relation of MotionCommentSection which was not loaded.")
	}

	return m.readGroups
}

func (m *MotionCommentSection) WriteGroups() []*Group {
	if _, ok := m.loadedRelations["write_group_ids"]; !ok {
		log.Panic().Msg("Tried to access WriteGroups relation of MotionCommentSection which was not loaded.")
	}

	return m.writeGroups
}

func (m *MotionCommentSection) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "comment_ids":
		for _, r := range m.comments {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "read_group_ids":
		for _, r := range m.readGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "write_group_ids":
		for _, r := range m.writeGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *MotionCommentSection) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "comment_ids":
			m.comments = content.([]*MotionComment)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "read_group_ids":
			m.readGroups = content.([]*Group)
		case "write_group_ids":
			m.writeGroups = content.([]*Group)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionCommentSection) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "comment_ids":
		var entry MotionComment
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.comments = append(m.comments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "read_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.readGroups = append(m.readGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "write_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.writeGroups = append(m.writeGroups, &entry)

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

func (m *MotionCommentSection) Get(field string) interface{} {
	switch field {
	case "comment_ids":
		return m.CommentIDs
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "read_group_ids":
		return m.ReadGroupIDs
	case "sequential_number":
		return m.SequentialNumber
	case "submitter_can_write":
		return m.SubmitterCanWrite
	case "weight":
		return m.Weight
	case "write_group_ids":
		return m.WriteGroupIDs
	}

	return nil
}

func (m *MotionCommentSection) GetFqids(field string) []string {
	switch field {
	case "comment_ids":
		r := make([]string, len(m.CommentIDs))
		for i, id := range m.CommentIDs {
			r[i] = "motion_comment/" + strconv.Itoa(id)
		}
		return r

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "read_group_ids":
		r := make([]string, len(m.ReadGroupIDs))
		for i, id := range m.ReadGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "write_group_ids":
		r := make([]string, len(m.WriteGroupIDs))
		for i, id := range m.WriteGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *MotionCommentSection) Update(data map[string]string) error {
	if val, ok := data["comment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommentIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["comment_ids"]; ok {
			m.comments = slices.DeleteFunc(m.comments, func(r *MotionComment) bool {
				return !slices.Contains(m.CommentIDs, r.ID)
			})
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["read_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReadGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["read_group_ids"]; ok {
			m.readGroups = slices.DeleteFunc(m.readGroups, func(r *Group) bool {
				return !slices.Contains(m.ReadGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["submitter_can_write"]; ok {
		err := json.Unmarshal([]byte(val), &m.SubmitterCanWrite)
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

	if val, ok := data["write_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.WriteGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["write_group_ids"]; ok {
			m.writeGroups = slices.DeleteFunc(m.writeGroups, func(r *Group) bool {
				return !slices.Contains(m.WriteGroupIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *MotionCommentSection) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type MotionEditor struct {
	ID              int  `json:"id"`
	MeetingID       int  `json:"meeting_id"`
	MeetingUserID   int  `json:"meeting_user_id"`
	MotionID        int  `json:"motion_id"`
	Weight          *int `json:"weight"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	meetingUser     *MeetingUser
	motion          *Motion
}

func (m *MotionEditor) CollectionName() string {
	return "motion_editor"
}

func (m *MotionEditor) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionEditor which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionEditor) MeetingUser() MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of MotionEditor which was not loaded.")
	}

	return *m.meetingUser
}

func (m *MotionEditor) Motion() Motion {
	if _, ok := m.loadedRelations["motion_id"]; !ok {
		log.Panic().Msg("Tried to access Motion relation of MotionEditor which was not loaded.")
	}

	return *m.motion
}

func (m *MotionEditor) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "meeting_user_id":
		return m.meetingUser.GetRelatedModelsAccessor()
	case "motion_id":
		return m.motion.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *MotionEditor) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_id":
			m.meetingUser = content.(*MeetingUser)
		case "motion_id":
			m.motion = content.(*Motion)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionEditor) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "meeting_user_id":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUser = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_id":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motion = &entry

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

func (m *MotionEditor) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "motion_id":
		return m.MotionID
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *MotionEditor) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_id":
		return []string{"meeting_user/" + strconv.Itoa(m.MeetingUserID)}

	case "motion_id":
		return []string{"motion/" + strconv.Itoa(m.MotionID)}
	}
	return []string{}
}

func (m *MotionEditor) Update(data map[string]string) error {
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

	if val, ok := data["meeting_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserID)
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

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionEditor) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type MotionState struct {
	AllowCreatePoll                  *bool    `json:"allow_create_poll"`
	AllowMotionForwarding            *bool    `json:"allow_motion_forwarding"`
	AllowSubmitterEdit               *bool    `json:"allow_submitter_edit"`
	AllowSupport                     *bool    `json:"allow_support"`
	CssClass                         string   `json:"css_class"`
	FirstStateOfWorkflowID           *int     `json:"first_state_of_workflow_id"`
	ID                               int      `json:"id"`
	IsInternal                       *bool    `json:"is_internal"`
	MeetingID                        int      `json:"meeting_id"`
	MergeAmendmentIntoFinal          *string  `json:"merge_amendment_into_final"`
	MotionIDs                        []int    `json:"motion_ids"`
	MotionRecommendationIDs          []int    `json:"motion_recommendation_ids"`
	Name                             string   `json:"name"`
	NextStateIDs                     []int    `json:"next_state_ids"`
	PreviousStateIDs                 []int    `json:"previous_state_ids"`
	RecommendationLabel              *string  `json:"recommendation_label"`
	Restrictions                     []string `json:"restrictions"`
	SetNumber                        *bool    `json:"set_number"`
	SetWorkflowTimestamp             *bool    `json:"set_workflow_timestamp"`
	ShowRecommendationExtensionField *bool    `json:"show_recommendation_extension_field"`
	ShowStateExtensionField          *bool    `json:"show_state_extension_field"`
	SubmitterWithdrawBackIDs         []int    `json:"submitter_withdraw_back_ids"`
	SubmitterWithdrawStateID         *int     `json:"submitter_withdraw_state_id"`
	Weight                           int      `json:"weight"`
	WorkflowID                       int      `json:"workflow_id"`
	loadedRelations                  map[string]struct{}
	firstStateOfWorkflow             *MotionWorkflow
	meeting                          *Meeting
	motionRecommendations            []*Motion
	motions                          []*Motion
	nextStates                       []*MotionState
	previousStates                   []*MotionState
	submitterWithdrawBacks           []*MotionState
	submitterWithdrawState           *MotionState
	workflow                         *MotionWorkflow
}

func (m *MotionState) CollectionName() string {
	return "motion_state"
}

func (m *MotionState) FirstStateOfWorkflow() *MotionWorkflow {
	if _, ok := m.loadedRelations["first_state_of_workflow_id"]; !ok {
		log.Panic().Msg("Tried to access FirstStateOfWorkflow relation of MotionState which was not loaded.")
	}

	return m.firstStateOfWorkflow
}

func (m *MotionState) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionState which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionState) MotionRecommendations() []*Motion {
	if _, ok := m.loadedRelations["motion_recommendation_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionRecommendations relation of MotionState which was not loaded.")
	}

	return m.motionRecommendations
}

func (m *MotionState) Motions() []*Motion {
	if _, ok := m.loadedRelations["motion_ids"]; !ok {
		log.Panic().Msg("Tried to access Motions relation of MotionState which was not loaded.")
	}

	return m.motions
}

func (m *MotionState) NextStates() []*MotionState {
	if _, ok := m.loadedRelations["next_state_ids"]; !ok {
		log.Panic().Msg("Tried to access NextStates relation of MotionState which was not loaded.")
	}

	return m.nextStates
}

func (m *MotionState) PreviousStates() []*MotionState {
	if _, ok := m.loadedRelations["previous_state_ids"]; !ok {
		log.Panic().Msg("Tried to access PreviousStates relation of MotionState which was not loaded.")
	}

	return m.previousStates
}

func (m *MotionState) SubmitterWithdrawBacks() []*MotionState {
	if _, ok := m.loadedRelations["submitter_withdraw_back_ids"]; !ok {
		log.Panic().Msg("Tried to access SubmitterWithdrawBacks relation of MotionState which was not loaded.")
	}

	return m.submitterWithdrawBacks
}

func (m *MotionState) SubmitterWithdrawState() *MotionState {
	if _, ok := m.loadedRelations["submitter_withdraw_state_id"]; !ok {
		log.Panic().Msg("Tried to access SubmitterWithdrawState relation of MotionState which was not loaded.")
	}

	return m.submitterWithdrawState
}

func (m *MotionState) Workflow() MotionWorkflow {
	if _, ok := m.loadedRelations["workflow_id"]; !ok {
		log.Panic().Msg("Tried to access Workflow relation of MotionState which was not loaded.")
	}

	return *m.workflow
}

func (m *MotionState) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "first_state_of_workflow_id":
		return m.firstStateOfWorkflow.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "motion_recommendation_ids":
		for _, r := range m.motionRecommendations {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "motion_ids":
		for _, r := range m.motions {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "next_state_ids":
		for _, r := range m.nextStates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "previous_state_ids":
		for _, r := range m.previousStates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "submitter_withdraw_back_ids":
		for _, r := range m.submitterWithdrawBacks {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "submitter_withdraw_state_id":
		return m.submitterWithdrawState.GetRelatedModelsAccessor()
	case "workflow_id":
		return m.workflow.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *MotionState) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "first_state_of_workflow_id":
			m.firstStateOfWorkflow = content.(*MotionWorkflow)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "motion_recommendation_ids":
			m.motionRecommendations = content.([]*Motion)
		case "motion_ids":
			m.motions = content.([]*Motion)
		case "next_state_ids":
			m.nextStates = content.([]*MotionState)
		case "previous_state_ids":
			m.previousStates = content.([]*MotionState)
		case "submitter_withdraw_back_ids":
			m.submitterWithdrawBacks = content.([]*MotionState)
		case "submitter_withdraw_state_id":
			m.submitterWithdrawState = content.(*MotionState)
		case "workflow_id":
			m.workflow = content.(*MotionWorkflow)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionState) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "first_state_of_workflow_id":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.firstStateOfWorkflow = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_recommendation_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionRecommendations = append(m.motionRecommendations, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motions = append(m.motions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "next_state_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.nextStates = append(m.nextStates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "previous_state_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.previousStates = append(m.previousStates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "submitter_withdraw_back_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.submitterWithdrawBacks = append(m.submitterWithdrawBacks, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "submitter_withdraw_state_id":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.submitterWithdrawState = &entry

		result = entry.GetRelatedModelsAccessor()
	case "workflow_id":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.workflow = &entry

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

func (m *MotionState) Get(field string) interface{} {
	switch field {
	case "allow_create_poll":
		return m.AllowCreatePoll
	case "allow_motion_forwarding":
		return m.AllowMotionForwarding
	case "allow_submitter_edit":
		return m.AllowSubmitterEdit
	case "allow_support":
		return m.AllowSupport
	case "css_class":
		return m.CssClass
	case "first_state_of_workflow_id":
		return m.FirstStateOfWorkflowID
	case "id":
		return m.ID
	case "is_internal":
		return m.IsInternal
	case "meeting_id":
		return m.MeetingID
	case "merge_amendment_into_final":
		return m.MergeAmendmentIntoFinal
	case "motion_ids":
		return m.MotionIDs
	case "motion_recommendation_ids":
		return m.MotionRecommendationIDs
	case "name":
		return m.Name
	case "next_state_ids":
		return m.NextStateIDs
	case "previous_state_ids":
		return m.PreviousStateIDs
	case "recommendation_label":
		return m.RecommendationLabel
	case "restrictions":
		return m.Restrictions
	case "set_number":
		return m.SetNumber
	case "set_workflow_timestamp":
		return m.SetWorkflowTimestamp
	case "show_recommendation_extension_field":
		return m.ShowRecommendationExtensionField
	case "show_state_extension_field":
		return m.ShowStateExtensionField
	case "submitter_withdraw_back_ids":
		return m.SubmitterWithdrawBackIDs
	case "submitter_withdraw_state_id":
		return m.SubmitterWithdrawStateID
	case "weight":
		return m.Weight
	case "workflow_id":
		return m.WorkflowID
	}

	return nil
}

func (m *MotionState) GetFqids(field string) []string {
	switch field {
	case "first_state_of_workflow_id":
		if m.FirstStateOfWorkflowID != nil {
			return []string{"motion_workflow/" + strconv.Itoa(*m.FirstStateOfWorkflowID)}
		}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "motion_recommendation_ids":
		r := make([]string, len(m.MotionRecommendationIDs))
		for i, id := range m.MotionRecommendationIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "motion_ids":
		r := make([]string, len(m.MotionIDs))
		for i, id := range m.MotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "next_state_ids":
		r := make([]string, len(m.NextStateIDs))
		for i, id := range m.NextStateIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r

	case "previous_state_ids":
		r := make([]string, len(m.PreviousStateIDs))
		for i, id := range m.PreviousStateIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r

	case "submitter_withdraw_back_ids":
		r := make([]string, len(m.SubmitterWithdrawBackIDs))
		for i, id := range m.SubmitterWithdrawBackIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r

	case "submitter_withdraw_state_id":
		if m.SubmitterWithdrawStateID != nil {
			return []string{"motion_state/" + strconv.Itoa(*m.SubmitterWithdrawStateID)}
		}

	case "workflow_id":
		return []string{"motion_workflow/" + strconv.Itoa(m.WorkflowID)}
	}
	return []string{}
}

func (m *MotionState) Update(data map[string]string) error {
	if val, ok := data["allow_create_poll"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllowCreatePoll)
		if err != nil {
			return err
		}
	}

	if val, ok := data["allow_motion_forwarding"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllowMotionForwarding)
		if err != nil {
			return err
		}
	}

	if val, ok := data["allow_submitter_edit"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllowSubmitterEdit)
		if err != nil {
			return err
		}
	}

	if val, ok := data["allow_support"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllowSupport)
		if err != nil {
			return err
		}
	}

	if val, ok := data["css_class"]; ok {
		err := json.Unmarshal([]byte(val), &m.CssClass)
		if err != nil {
			return err
		}
	}

	if val, ok := data["first_state_of_workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FirstStateOfWorkflowID)
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

	if val, ok := data["is_internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsInternal)
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

	if val, ok := data["merge_amendment_into_final"]; ok {
		err := json.Unmarshal([]byte(val), &m.MergeAmendmentIntoFinal)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_ids"]; ok {
			m.motions = slices.DeleteFunc(m.motions, func(r *Motion) bool {
				return !slices.Contains(m.MotionIDs, r.ID)
			})
		}
	}

	if val, ok := data["motion_recommendation_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionRecommendationIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["motion_recommendation_ids"]; ok {
			m.motionRecommendations = slices.DeleteFunc(m.motionRecommendations, func(r *Motion) bool {
				return !slices.Contains(m.MotionRecommendationIDs, r.ID)
			})
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["next_state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.NextStateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["next_state_ids"]; ok {
			m.nextStates = slices.DeleteFunc(m.nextStates, func(r *MotionState) bool {
				return !slices.Contains(m.NextStateIDs, r.ID)
			})
		}
	}

	if val, ok := data["previous_state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PreviousStateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["previous_state_ids"]; ok {
			m.previousStates = slices.DeleteFunc(m.previousStates, func(r *MotionState) bool {
				return !slices.Contains(m.PreviousStateIDs, r.ID)
			})
		}
	}

	if val, ok := data["recommendation_label"]; ok {
		err := json.Unmarshal([]byte(val), &m.RecommendationLabel)
		if err != nil {
			return err
		}
	}

	if val, ok := data["restrictions"]; ok {
		err := json.Unmarshal([]byte(val), &m.Restrictions)
		if err != nil {
			return err
		}
	}

	if val, ok := data["set_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.SetNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["set_workflow_timestamp"]; ok {
		err := json.Unmarshal([]byte(val), &m.SetWorkflowTimestamp)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_recommendation_extension_field"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowRecommendationExtensionField)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_state_extension_field"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowStateExtensionField)
		if err != nil {
			return err
		}
	}

	if val, ok := data["submitter_withdraw_back_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SubmitterWithdrawBackIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["submitter_withdraw_back_ids"]; ok {
			m.submitterWithdrawBacks = slices.DeleteFunc(m.submitterWithdrawBacks, func(r *MotionState) bool {
				return !slices.Contains(m.SubmitterWithdrawBackIDs, r.ID)
			})
		}
	}

	if val, ok := data["submitter_withdraw_state_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.SubmitterWithdrawStateID)
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

	if val, ok := data["workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.WorkflowID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionState) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type MotionSubmitter struct {
	ID              int  `json:"id"`
	MeetingID       int  `json:"meeting_id"`
	MeetingUserID   int  `json:"meeting_user_id"`
	MotionID        int  `json:"motion_id"`
	Weight          *int `json:"weight"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	meetingUser     *MeetingUser
	motion          *Motion
}

func (m *MotionSubmitter) CollectionName() string {
	return "motion_submitter"
}

func (m *MotionSubmitter) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionSubmitter which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionSubmitter) MeetingUser() MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of MotionSubmitter which was not loaded.")
	}

	return *m.meetingUser
}

func (m *MotionSubmitter) Motion() Motion {
	if _, ok := m.loadedRelations["motion_id"]; !ok {
		log.Panic().Msg("Tried to access Motion relation of MotionSubmitter which was not loaded.")
	}

	return *m.motion
}

func (m *MotionSubmitter) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "meeting_user_id":
		return m.meetingUser.GetRelatedModelsAccessor()
	case "motion_id":
		return m.motion.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *MotionSubmitter) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_id":
			m.meetingUser = content.(*MeetingUser)
		case "motion_id":
			m.motion = content.(*Motion)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionSubmitter) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "meeting_user_id":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUser = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_id":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motion = &entry

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

func (m *MotionSubmitter) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "motion_id":
		return m.MotionID
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *MotionSubmitter) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_id":
		return []string{"meeting_user/" + strconv.Itoa(m.MeetingUserID)}

	case "motion_id":
		return []string{"motion/" + strconv.Itoa(m.MotionID)}
	}
	return []string{}
}

func (m *MotionSubmitter) Update(data map[string]string) error {
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

	if val, ok := data["meeting_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserID)
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

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionSubmitter) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type MotionWorkflow struct {
	DefaultAmendmentWorkflowMeetingID *int   `json:"default_amendment_workflow_meeting_id"`
	DefaultWorkflowMeetingID          *int   `json:"default_workflow_meeting_id"`
	FirstStateID                      int    `json:"first_state_id"`
	ID                                int    `json:"id"`
	MeetingID                         int    `json:"meeting_id"`
	Name                              string `json:"name"`
	SequentialNumber                  int    `json:"sequential_number"`
	StateIDs                          []int  `json:"state_ids"`
	loadedRelations                   map[string]struct{}
	defaultAmendmentWorkflowMeeting   *Meeting
	defaultWorkflowMeeting            *Meeting
	firstState                        *MotionState
	meeting                           *Meeting
	states                            []*MotionState
}

func (m *MotionWorkflow) CollectionName() string {
	return "motion_workflow"
}

func (m *MotionWorkflow) DefaultAmendmentWorkflowMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_amendment_workflow_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultAmendmentWorkflowMeeting relation of MotionWorkflow which was not loaded.")
	}

	return m.defaultAmendmentWorkflowMeeting
}

func (m *MotionWorkflow) DefaultWorkflowMeeting() *Meeting {
	if _, ok := m.loadedRelations["default_workflow_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultWorkflowMeeting relation of MotionWorkflow which was not loaded.")
	}

	return m.defaultWorkflowMeeting
}

func (m *MotionWorkflow) FirstState() MotionState {
	if _, ok := m.loadedRelations["first_state_id"]; !ok {
		log.Panic().Msg("Tried to access FirstState relation of MotionWorkflow which was not loaded.")
	}

	return *m.firstState
}

func (m *MotionWorkflow) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionWorkflow which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionWorkflow) States() []*MotionState {
	if _, ok := m.loadedRelations["state_ids"]; !ok {
		log.Panic().Msg("Tried to access States relation of MotionWorkflow which was not loaded.")
	}

	return m.states
}

func (m *MotionWorkflow) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "default_amendment_workflow_meeting_id":
		return m.defaultAmendmentWorkflowMeeting.GetRelatedModelsAccessor()
	case "default_workflow_meeting_id":
		return m.defaultWorkflowMeeting.GetRelatedModelsAccessor()
	case "first_state_id":
		return m.firstState.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "state_ids":
		for _, r := range m.states {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *MotionWorkflow) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "default_amendment_workflow_meeting_id":
			m.defaultAmendmentWorkflowMeeting = content.(*Meeting)
		case "default_workflow_meeting_id":
			m.defaultWorkflowMeeting = content.(*Meeting)
		case "first_state_id":
			m.firstState = content.(*MotionState)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "state_ids":
			m.states = content.([]*MotionState)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionWorkflow) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "default_amendment_workflow_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultAmendmentWorkflowMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "default_workflow_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultWorkflowMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "first_state_id":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.firstState = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "state_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.states = append(m.states, &entry)

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

func (m *MotionWorkflow) Get(field string) interface{} {
	switch field {
	case "default_amendment_workflow_meeting_id":
		return m.DefaultAmendmentWorkflowMeetingID
	case "default_workflow_meeting_id":
		return m.DefaultWorkflowMeetingID
	case "first_state_id":
		return m.FirstStateID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "sequential_number":
		return m.SequentialNumber
	case "state_ids":
		return m.StateIDs
	}

	return nil
}

func (m *MotionWorkflow) GetFqids(field string) []string {
	switch field {
	case "default_amendment_workflow_meeting_id":
		if m.DefaultAmendmentWorkflowMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultAmendmentWorkflowMeetingID)}
		}

	case "default_workflow_meeting_id":
		if m.DefaultWorkflowMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.DefaultWorkflowMeetingID)}
		}

	case "first_state_id":
		return []string{"motion_state/" + strconv.Itoa(m.FirstStateID)}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "state_ids":
		r := make([]string, len(m.StateIDs))
		for i, id := range m.StateIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *MotionWorkflow) Update(data map[string]string) error {
	if val, ok := data["default_amendment_workflow_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultAmendmentWorkflowMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_workflow_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultWorkflowMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["first_state_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FirstStateID)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
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

	if val, ok := data["state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["state_ids"]; ok {
			m.states = slices.DeleteFunc(m.states, func(r *MotionState) bool {
				return !slices.Contains(m.StateIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *MotionWorkflow) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type MotionWorkingGroupSpeaker struct {
	ID              int  `json:"id"`
	MeetingID       int  `json:"meeting_id"`
	MeetingUserID   int  `json:"meeting_user_id"`
	MotionID        int  `json:"motion_id"`
	Weight          *int `json:"weight"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
	meetingUser     *MeetingUser
	motion          *Motion
}

func (m *MotionWorkingGroupSpeaker) CollectionName() string {
	return "motion_working_group_speaker"
}

func (m *MotionWorkingGroupSpeaker) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of MotionWorkingGroupSpeaker which was not loaded.")
	}

	return *m.meeting
}

func (m *MotionWorkingGroupSpeaker) MeetingUser() MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of MotionWorkingGroupSpeaker which was not loaded.")
	}

	return *m.meetingUser
}

func (m *MotionWorkingGroupSpeaker) Motion() Motion {
	if _, ok := m.loadedRelations["motion_id"]; !ok {
		log.Panic().Msg("Tried to access Motion relation of MotionWorkingGroupSpeaker which was not loaded.")
	}

	return *m.motion
}

func (m *MotionWorkingGroupSpeaker) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "meeting_user_id":
		return m.meetingUser.GetRelatedModelsAccessor()
	case "motion_id":
		return m.motion.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *MotionWorkingGroupSpeaker) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_id":
			m.meetingUser = content.(*MeetingUser)
		case "motion_id":
			m.motion = content.(*Motion)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *MotionWorkingGroupSpeaker) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "meeting_user_id":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUser = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_id":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motion = &entry

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

func (m *MotionWorkingGroupSpeaker) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "motion_id":
		return m.MotionID
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *MotionWorkingGroupSpeaker) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_id":
		return []string{"meeting_user/" + strconv.Itoa(m.MeetingUserID)}

	case "motion_id":
		return []string{"motion/" + strconv.Itoa(m.MotionID)}
	}
	return []string{}
}

func (m *MotionWorkingGroupSpeaker) Update(data map[string]string) error {
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

	if val, ok := data["meeting_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserID)
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

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MotionWorkingGroupSpeaker) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Option struct {
	Abstain                    *string `json:"abstain"`
	ContentObjectID            *string `json:"content_object_id"`
	ID                         int     `json:"id"`
	MeetingID                  int     `json:"meeting_id"`
	No                         *string `json:"no"`
	PollID                     *int    `json:"poll_id"`
	Text                       *string `json:"text"`
	UsedAsGlobalOptionInPollID *int    `json:"used_as_global_option_in_poll_id"`
	VoteIDs                    []int   `json:"vote_ids"`
	Weight                     *int    `json:"weight"`
	Yes                        *string `json:"yes"`
	loadedRelations            map[string]struct{}
	contentObject              IBaseModel
	meeting                    *Meeting
	poll                       *Poll
	usedAsGlobalOptionInPoll   *Poll
	votes                      []*Vote
}

func (m *Option) CollectionName() string {
	return "option"
}

func (m *Option) ContentObject() IBaseModel {
	if _, ok := m.loadedRelations["content_object_id"]; !ok {
		log.Panic().Msg("Tried to access ContentObject relation of Option which was not loaded.")
	}

	return m.contentObject
}

func (m *Option) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Option which was not loaded.")
	}

	return *m.meeting
}

func (m *Option) Poll() *Poll {
	if _, ok := m.loadedRelations["poll_id"]; !ok {
		log.Panic().Msg("Tried to access Poll relation of Option which was not loaded.")
	}

	return m.poll
}

func (m *Option) UsedAsGlobalOptionInPoll() *Poll {
	if _, ok := m.loadedRelations["used_as_global_option_in_poll_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsGlobalOptionInPoll relation of Option which was not loaded.")
	}

	return m.usedAsGlobalOptionInPoll
}

func (m *Option) Votes() []*Vote {
	if _, ok := m.loadedRelations["vote_ids"]; !ok {
		log.Panic().Msg("Tried to access Votes relation of Option which was not loaded.")
	}

	return m.votes
}

func (m *Option) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "content_object_id":
		return m.contentObject.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "poll_id":
		return m.poll.GetRelatedModelsAccessor()
	case "used_as_global_option_in_poll_id":
		return m.usedAsGlobalOptionInPoll.GetRelatedModelsAccessor()
	case "vote_ids":
		for _, r := range m.votes {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Option) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "content_object_id":
			panic("not implemented")
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "poll_id":
			m.poll = content.(*Poll)
		case "used_as_global_option_in_poll_id":
			m.usedAsGlobalOptionInPoll = content.(*Poll)
		case "vote_ids":
			m.votes = content.([]*Vote)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Option) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "content_object_id":
		if m.ContentObjectID == nil {
			return nil, fmt.Errorf("cannot fill relation for ContentObjectID while id field is empty")
		}
		parts := strings.Split(*m.ContentObjectID, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf("could not parse id field")
		}

		switch parts[0] {
		case "motion":
			var entry Motion
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "poll_candidate_list":
			var entry PollCandidateList
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "user":
			var entry User
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
	case "poll_id":
		var entry Poll
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.poll = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_global_option_in_poll_id":
		var entry Poll
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsGlobalOptionInPoll = &entry

		result = entry.GetRelatedModelsAccessor()
	case "vote_ids":
		var entry Vote
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.votes = append(m.votes, &entry)

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

func (m *Option) Get(field string) interface{} {
	switch field {
	case "abstain":
		return m.Abstain
	case "content_object_id":
		return m.ContentObjectID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "no":
		return m.No
	case "poll_id":
		return m.PollID
	case "text":
		return m.Text
	case "used_as_global_option_in_poll_id":
		return m.UsedAsGlobalOptionInPollID
	case "vote_ids":
		return m.VoteIDs
	case "weight":
		return m.Weight
	case "yes":
		return m.Yes
	}

	return nil
}

func (m *Option) GetFqids(field string) []string {
	switch field {
	case "content_object_id":
		if m.ContentObjectID != nil {
			return []string{*m.ContentObjectID}
		}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "poll_id":
		if m.PollID != nil {
			return []string{"poll/" + strconv.Itoa(*m.PollID)}
		}

	case "used_as_global_option_in_poll_id":
		if m.UsedAsGlobalOptionInPollID != nil {
			return []string{"poll/" + strconv.Itoa(*m.UsedAsGlobalOptionInPollID)}
		}

	case "vote_ids":
		r := make([]string, len(m.VoteIDs))
		for i, id := range m.VoteIDs {
			r[i] = "vote/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Option) Update(data map[string]string) error {
	if val, ok := data["abstain"]; ok {
		err := json.Unmarshal([]byte(val), &m.Abstain)
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

	if val, ok := data["no"]; ok {
		err := json.Unmarshal([]byte(val), &m.No)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollID)
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

	if val, ok := data["used_as_global_option_in_poll_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsGlobalOptionInPollID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["vote_ids"]; ok {
			m.votes = slices.DeleteFunc(m.votes, func(r *Vote) bool {
				return !slices.Contains(m.VoteIDs, r.ID)
			})
		}
	}

	if val, ok := data["weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.Weight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["yes"]; ok {
		err := json.Unmarshal([]byte(val), &m.Yes)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Option) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Organization struct {
	ActiveMeetingIDs           []int           `json:"active_meeting_ids"`
	ArchivedMeetingIDs         []int           `json:"archived_meeting_ids"`
	CommitteeIDs               []int           `json:"committee_ids"`
	DefaultLanguage            string          `json:"default_language"`
	Description                *string         `json:"description"`
	EnableAnonymous            *bool           `json:"enable_anonymous"`
	EnableChat                 *bool           `json:"enable_chat"`
	EnableElectronicVoting     *bool           `json:"enable_electronic_voting"`
	GenderIDs                  []int           `json:"gender_ids"`
	ID                         int             `json:"id"`
	LegalNotice                *string         `json:"legal_notice"`
	LimitOfMeetings            *int            `json:"limit_of_meetings"`
	LimitOfUsers               *int            `json:"limit_of_users"`
	LoginText                  *string         `json:"login_text"`
	MediafileIDs               []int           `json:"mediafile_ids"`
	Name                       *string         `json:"name"`
	OrganizationTagIDs         []int           `json:"organization_tag_ids"`
	PrivacyPolicy              *string         `json:"privacy_policy"`
	PublishedMediafileIDs      []int           `json:"published_mediafile_ids"`
	RequireDuplicateFrom       *bool           `json:"require_duplicate_from"`
	ResetPasswordVerboseErrors *bool           `json:"reset_password_verbose_errors"`
	SamlAttrMapping            json.RawMessage `json:"saml_attr_mapping"`
	SamlEnabled                *bool           `json:"saml_enabled"`
	SamlLoginButtonText        *string         `json:"saml_login_button_text"`
	SamlMetadataIDp            *string         `json:"saml_metadata_idp"`
	SamlMetadataSp             *string         `json:"saml_metadata_sp"`
	SamlPrivateKey             *string         `json:"saml_private_key"`
	TemplateMeetingIDs         []int           `json:"template_meeting_ids"`
	ThemeID                    int             `json:"theme_id"`
	ThemeIDs                   []int           `json:"theme_ids"`
	Url                        *string         `json:"url"`
	UserIDs                    []int           `json:"user_ids"`
	UsersEmailBody             *string         `json:"users_email_body"`
	UsersEmailReplyto          *string         `json:"users_email_replyto"`
	UsersEmailSender           *string         `json:"users_email_sender"`
	UsersEmailSubject          *string         `json:"users_email_subject"`
	VoteDecryptPublicMainKey   *string         `json:"vote_decrypt_public_main_key"`
	loadedRelations            map[string]struct{}
	activeMeetings             []*Meeting
	archivedMeetings           []*Meeting
	committees                 []*Committee
	genders                    []*Gender
	mediafiles                 []*Mediafile
	organizationTags           []*OrganizationTag
	publishedMediafiles        []*Mediafile
	templateMeetings           []*Meeting
	theme                      *Theme
	themes                     []*Theme
	users                      []*User
}

func (m *Organization) CollectionName() string {
	return "organization"
}

func (m *Organization) ActiveMeetings() []*Meeting {
	if _, ok := m.loadedRelations["active_meeting_ids"]; !ok {
		log.Panic().Msg("Tried to access ActiveMeetings relation of Organization which was not loaded.")
	}

	return m.activeMeetings
}

func (m *Organization) ArchivedMeetings() []*Meeting {
	if _, ok := m.loadedRelations["archived_meeting_ids"]; !ok {
		log.Panic().Msg("Tried to access ArchivedMeetings relation of Organization which was not loaded.")
	}

	return m.archivedMeetings
}

func (m *Organization) Committees() []*Committee {
	if _, ok := m.loadedRelations["committee_ids"]; !ok {
		log.Panic().Msg("Tried to access Committees relation of Organization which was not loaded.")
	}

	return m.committees
}

func (m *Organization) Genders() []*Gender {
	if _, ok := m.loadedRelations["gender_ids"]; !ok {
		log.Panic().Msg("Tried to access Genders relation of Organization which was not loaded.")
	}

	return m.genders
}

func (m *Organization) Mediafiles() []*Mediafile {
	if _, ok := m.loadedRelations["mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access Mediafiles relation of Organization which was not loaded.")
	}

	return m.mediafiles
}

func (m *Organization) OrganizationTags() []*OrganizationTag {
	if _, ok := m.loadedRelations["organization_tag_ids"]; !ok {
		log.Panic().Msg("Tried to access OrganizationTags relation of Organization which was not loaded.")
	}

	return m.organizationTags
}

func (m *Organization) PublishedMediafiles() []*Mediafile {
	if _, ok := m.loadedRelations["published_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access PublishedMediafiles relation of Organization which was not loaded.")
	}

	return m.publishedMediafiles
}

func (m *Organization) TemplateMeetings() []*Meeting {
	if _, ok := m.loadedRelations["template_meeting_ids"]; !ok {
		log.Panic().Msg("Tried to access TemplateMeetings relation of Organization which was not loaded.")
	}

	return m.templateMeetings
}

func (m *Organization) Theme() Theme {
	if _, ok := m.loadedRelations["theme_id"]; !ok {
		log.Panic().Msg("Tried to access Theme relation of Organization which was not loaded.")
	}

	return *m.theme
}

func (m *Organization) Themes() []*Theme {
	if _, ok := m.loadedRelations["theme_ids"]; !ok {
		log.Panic().Msg("Tried to access Themes relation of Organization which was not loaded.")
	}

	return m.themes
}

func (m *Organization) Users() []*User {
	if _, ok := m.loadedRelations["user_ids"]; !ok {
		log.Panic().Msg("Tried to access Users relation of Organization which was not loaded.")
	}

	return m.users
}

func (m *Organization) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "active_meeting_ids":
		for _, r := range m.activeMeetings {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "archived_meeting_ids":
		for _, r := range m.archivedMeetings {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "committee_ids":
		for _, r := range m.committees {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "gender_ids":
		for _, r := range m.genders {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "mediafile_ids":
		for _, r := range m.mediafiles {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "organization_tag_ids":
		for _, r := range m.organizationTags {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "published_mediafile_ids":
		for _, r := range m.publishedMediafiles {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "template_meeting_ids":
		for _, r := range m.templateMeetings {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "theme_id":
		return m.theme.GetRelatedModelsAccessor()
	case "theme_ids":
		for _, r := range m.themes {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "user_ids":
		for _, r := range m.users {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Organization) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "active_meeting_ids":
			m.activeMeetings = content.([]*Meeting)
		case "archived_meeting_ids":
			m.archivedMeetings = content.([]*Meeting)
		case "committee_ids":
			m.committees = content.([]*Committee)
		case "gender_ids":
			m.genders = content.([]*Gender)
		case "mediafile_ids":
			m.mediafiles = content.([]*Mediafile)
		case "organization_tag_ids":
			m.organizationTags = content.([]*OrganizationTag)
		case "published_mediafile_ids":
			m.publishedMediafiles = content.([]*Mediafile)
		case "template_meeting_ids":
			m.templateMeetings = content.([]*Meeting)
		case "theme_id":
			m.theme = content.(*Theme)
		case "theme_ids":
			m.themes = content.([]*Theme)
		case "user_ids":
			m.users = content.([]*User)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Organization) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "active_meeting_ids":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.activeMeetings = append(m.activeMeetings, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "archived_meeting_ids":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.archivedMeetings = append(m.archivedMeetings, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "committee_ids":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.committees = append(m.committees, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "gender_ids":
		var entry Gender
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.genders = append(m.genders, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "mediafile_ids":
		var entry Mediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.mediafiles = append(m.mediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "organization_tag_ids":
		var entry OrganizationTag
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.organizationTags = append(m.organizationTags, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "published_mediafile_ids":
		var entry Mediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.publishedMediafiles = append(m.publishedMediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "template_meeting_ids":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.templateMeetings = append(m.templateMeetings, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "theme_id":
		var entry Theme
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.theme = &entry

		result = entry.GetRelatedModelsAccessor()
	case "theme_ids":
		var entry Theme
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.themes = append(m.themes, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "user_ids":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.users = append(m.users, &entry)

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

func (m *Organization) Get(field string) interface{} {
	switch field {
	case "active_meeting_ids":
		return m.ActiveMeetingIDs
	case "archived_meeting_ids":
		return m.ArchivedMeetingIDs
	case "committee_ids":
		return m.CommitteeIDs
	case "default_language":
		return m.DefaultLanguage
	case "description":
		return m.Description
	case "enable_anonymous":
		return m.EnableAnonymous
	case "enable_chat":
		return m.EnableChat
	case "enable_electronic_voting":
		return m.EnableElectronicVoting
	case "gender_ids":
		return m.GenderIDs
	case "id":
		return m.ID
	case "legal_notice":
		return m.LegalNotice
	case "limit_of_meetings":
		return m.LimitOfMeetings
	case "limit_of_users":
		return m.LimitOfUsers
	case "login_text":
		return m.LoginText
	case "mediafile_ids":
		return m.MediafileIDs
	case "name":
		return m.Name
	case "organization_tag_ids":
		return m.OrganizationTagIDs
	case "privacy_policy":
		return m.PrivacyPolicy
	case "published_mediafile_ids":
		return m.PublishedMediafileIDs
	case "require_duplicate_from":
		return m.RequireDuplicateFrom
	case "reset_password_verbose_errors":
		return m.ResetPasswordVerboseErrors
	case "saml_attr_mapping":
		return m.SamlAttrMapping
	case "saml_enabled":
		return m.SamlEnabled
	case "saml_login_button_text":
		return m.SamlLoginButtonText
	case "saml_metadata_idp":
		return m.SamlMetadataIDp
	case "saml_metadata_sp":
		return m.SamlMetadataSp
	case "saml_private_key":
		return m.SamlPrivateKey
	case "template_meeting_ids":
		return m.TemplateMeetingIDs
	case "theme_id":
		return m.ThemeID
	case "theme_ids":
		return m.ThemeIDs
	case "url":
		return m.Url
	case "user_ids":
		return m.UserIDs
	case "users_email_body":
		return m.UsersEmailBody
	case "users_email_replyto":
		return m.UsersEmailReplyto
	case "users_email_sender":
		return m.UsersEmailSender
	case "users_email_subject":
		return m.UsersEmailSubject
	case "vote_decrypt_public_main_key":
		return m.VoteDecryptPublicMainKey
	}

	return nil
}

func (m *Organization) GetFqids(field string) []string {
	switch field {
	case "active_meeting_ids":
		r := make([]string, len(m.ActiveMeetingIDs))
		for i, id := range m.ActiveMeetingIDs {
			r[i] = "meeting/" + strconv.Itoa(id)
		}
		return r

	case "archived_meeting_ids":
		r := make([]string, len(m.ArchivedMeetingIDs))
		for i, id := range m.ArchivedMeetingIDs {
			r[i] = "meeting/" + strconv.Itoa(id)
		}
		return r

	case "committee_ids":
		r := make([]string, len(m.CommitteeIDs))
		for i, id := range m.CommitteeIDs {
			r[i] = "committee/" + strconv.Itoa(id)
		}
		return r

	case "gender_ids":
		r := make([]string, len(m.GenderIDs))
		for i, id := range m.GenderIDs {
			r[i] = "gender/" + strconv.Itoa(id)
		}
		return r

	case "mediafile_ids":
		r := make([]string, len(m.MediafileIDs))
		for i, id := range m.MediafileIDs {
			r[i] = "mediafile/" + strconv.Itoa(id)
		}
		return r

	case "organization_tag_ids":
		r := make([]string, len(m.OrganizationTagIDs))
		for i, id := range m.OrganizationTagIDs {
			r[i] = "organization_tag/" + strconv.Itoa(id)
		}
		return r

	case "published_mediafile_ids":
		r := make([]string, len(m.PublishedMediafileIDs))
		for i, id := range m.PublishedMediafileIDs {
			r[i] = "mediafile/" + strconv.Itoa(id)
		}
		return r

	case "template_meeting_ids":
		r := make([]string, len(m.TemplateMeetingIDs))
		for i, id := range m.TemplateMeetingIDs {
			r[i] = "meeting/" + strconv.Itoa(id)
		}
		return r

	case "theme_id":
		return []string{"theme/" + strconv.Itoa(m.ThemeID)}

	case "theme_ids":
		r := make([]string, len(m.ThemeIDs))
		for i, id := range m.ThemeIDs {
			r[i] = "theme/" + strconv.Itoa(id)
		}
		return r

	case "user_ids":
		r := make([]string, len(m.UserIDs))
		for i, id := range m.UserIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Organization) Update(data map[string]string) error {
	if val, ok := data["active_meeting_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ActiveMeetingIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["active_meeting_ids"]; ok {
			m.activeMeetings = slices.DeleteFunc(m.activeMeetings, func(r *Meeting) bool {
				return !slices.Contains(m.ActiveMeetingIDs, r.ID)
			})
		}
	}

	if val, ok := data["archived_meeting_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ArchivedMeetingIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["archived_meeting_ids"]; ok {
			m.archivedMeetings = slices.DeleteFunc(m.archivedMeetings, func(r *Meeting) bool {
				return !slices.Contains(m.ArchivedMeetingIDs, r.ID)
			})
		}
	}

	if val, ok := data["committee_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommitteeIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["committee_ids"]; ok {
			m.committees = slices.DeleteFunc(m.committees, func(r *Committee) bool {
				return !slices.Contains(m.CommitteeIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_language"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultLanguage)
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

	if val, ok := data["enable_anonymous"]; ok {
		err := json.Unmarshal([]byte(val), &m.EnableAnonymous)
		if err != nil {
			return err
		}
	}

	if val, ok := data["enable_chat"]; ok {
		err := json.Unmarshal([]byte(val), &m.EnableChat)
		if err != nil {
			return err
		}
	}

	if val, ok := data["enable_electronic_voting"]; ok {
		err := json.Unmarshal([]byte(val), &m.EnableElectronicVoting)
		if err != nil {
			return err
		}
	}

	if val, ok := data["gender_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.GenderIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["gender_ids"]; ok {
			m.genders = slices.DeleteFunc(m.genders, func(r *Gender) bool {
				return !slices.Contains(m.GenderIDs, r.ID)
			})
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["legal_notice"]; ok {
		err := json.Unmarshal([]byte(val), &m.LegalNotice)
		if err != nil {
			return err
		}
	}

	if val, ok := data["limit_of_meetings"]; ok {
		err := json.Unmarshal([]byte(val), &m.LimitOfMeetings)
		if err != nil {
			return err
		}
	}

	if val, ok := data["limit_of_users"]; ok {
		err := json.Unmarshal([]byte(val), &m.LimitOfUsers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["login_text"]; ok {
		err := json.Unmarshal([]byte(val), &m.LoginText)
		if err != nil {
			return err
		}
	}

	if val, ok := data["mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MediafileIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["mediafile_ids"]; ok {
			m.mediafiles = slices.DeleteFunc(m.mediafiles, func(r *Mediafile) bool {
				return !slices.Contains(m.MediafileIDs, r.ID)
			})
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["organization_tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationTagIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["organization_tag_ids"]; ok {
			m.organizationTags = slices.DeleteFunc(m.organizationTags, func(r *OrganizationTag) bool {
				return !slices.Contains(m.OrganizationTagIDs, r.ID)
			})
		}
	}

	if val, ok := data["privacy_policy"]; ok {
		err := json.Unmarshal([]byte(val), &m.PrivacyPolicy)
		if err != nil {
			return err
		}
	}

	if val, ok := data["published_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PublishedMediafileIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["published_mediafile_ids"]; ok {
			m.publishedMediafiles = slices.DeleteFunc(m.publishedMediafiles, func(r *Mediafile) bool {
				return !slices.Contains(m.PublishedMediafileIDs, r.ID)
			})
		}
	}

	if val, ok := data["require_duplicate_from"]; ok {
		err := json.Unmarshal([]byte(val), &m.RequireDuplicateFrom)
		if err != nil {
			return err
		}
	}

	if val, ok := data["reset_password_verbose_errors"]; ok {
		err := json.Unmarshal([]byte(val), &m.ResetPasswordVerboseErrors)
		if err != nil {
			return err
		}
	}

	if val, ok := data["saml_attr_mapping"]; ok {
		err := json.Unmarshal([]byte(val), &m.SamlAttrMapping)
		if err != nil {
			return err
		}
	}

	if val, ok := data["saml_enabled"]; ok {
		err := json.Unmarshal([]byte(val), &m.SamlEnabled)
		if err != nil {
			return err
		}
	}

	if val, ok := data["saml_login_button_text"]; ok {
		err := json.Unmarshal([]byte(val), &m.SamlLoginButtonText)
		if err != nil {
			return err
		}
	}

	if val, ok := data["saml_metadata_idp"]; ok {
		err := json.Unmarshal([]byte(val), &m.SamlMetadataIDp)
		if err != nil {
			return err
		}
	}

	if val, ok := data["saml_metadata_sp"]; ok {
		err := json.Unmarshal([]byte(val), &m.SamlMetadataSp)
		if err != nil {
			return err
		}
	}

	if val, ok := data["saml_private_key"]; ok {
		err := json.Unmarshal([]byte(val), &m.SamlPrivateKey)
		if err != nil {
			return err
		}
	}

	if val, ok := data["template_meeting_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TemplateMeetingIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["template_meeting_ids"]; ok {
			m.templateMeetings = slices.DeleteFunc(m.templateMeetings, func(r *Meeting) bool {
				return !slices.Contains(m.TemplateMeetingIDs, r.ID)
			})
		}
	}

	if val, ok := data["theme_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ThemeID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["theme_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ThemeIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["theme_ids"]; ok {
			m.themes = slices.DeleteFunc(m.themes, func(r *Theme) bool {
				return !slices.Contains(m.ThemeIDs, r.ID)
			})
		}
	}

	if val, ok := data["url"]; ok {
		err := json.Unmarshal([]byte(val), &m.Url)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["user_ids"]; ok {
			m.users = slices.DeleteFunc(m.users, func(r *User) bool {
				return !slices.Contains(m.UserIDs, r.ID)
			})
		}
	}

	if val, ok := data["users_email_body"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailBody)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_replyto"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailReplyto)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_sender"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailSender)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_subject"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailSubject)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_decrypt_public_main_key"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteDecryptPublicMainKey)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Organization) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type OrganizationTag struct {
	Color           string   `json:"color"`
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	OrganizationID  int      `json:"organization_id"`
	TaggedIDs       []string `json:"tagged_ids"`
	loadedRelations map[string]struct{}
	organization    *Organization
}

func (m *OrganizationTag) CollectionName() string {
	return "organization_tag"
}

func (m *OrganizationTag) Organization() Organization {
	if _, ok := m.loadedRelations["organization_id"]; !ok {
		log.Panic().Msg("Tried to access Organization relation of OrganizationTag which was not loaded.")
	}

	return *m.organization
}

func (m *OrganizationTag) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "organization_id":
		return m.organization.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *OrganizationTag) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "organization_id":
			m.organization = content.(*Organization)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *OrganizationTag) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.organization = &entry

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

func (m *OrganizationTag) Get(field string) interface{} {
	switch field {
	case "color":
		return m.Color
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "organization_id":
		return m.OrganizationID
	case "tagged_ids":
		return m.TaggedIDs
	}

	return nil
}

func (m *OrganizationTag) GetFqids(field string) []string {
	switch field {
	case "organization_id":
		return []string{"organization/" + strconv.Itoa(m.OrganizationID)}
	}
	return []string{}
}

func (m *OrganizationTag) Update(data map[string]string) error {
	if val, ok := data["color"]; ok {
		err := json.Unmarshal([]byte(val), &m.Color)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["tagged_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TaggedIDs)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *OrganizationTag) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type PersonalNote struct {
	ContentObjectID *string `json:"content_object_id"`
	ID              int     `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	MeetingUserID   int     `json:"meeting_user_id"`
	Note            *string `json:"note"`
	Star            *bool   `json:"star"`
	loadedRelations map[string]struct{}
	contentObject   IBaseModel
	meeting         *Meeting
	meetingUser     *MeetingUser
}

func (m *PersonalNote) CollectionName() string {
	return "personal_note"
}

func (m *PersonalNote) ContentObject() IBaseModel {
	if _, ok := m.loadedRelations["content_object_id"]; !ok {
		log.Panic().Msg("Tried to access ContentObject relation of PersonalNote which was not loaded.")
	}

	return m.contentObject
}

func (m *PersonalNote) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PersonalNote which was not loaded.")
	}

	return *m.meeting
}

func (m *PersonalNote) MeetingUser() MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_id"]; !ok {
		log.Panic().Msg("Tried to access MeetingUser relation of PersonalNote which was not loaded.")
	}

	return *m.meetingUser
}

func (m *PersonalNote) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "content_object_id":
		return m.contentObject.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "meeting_user_id":
		return m.meetingUser.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *PersonalNote) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "content_object_id":
			panic("not implemented")
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "meeting_user_id":
			m.meetingUser = content.(*MeetingUser)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *PersonalNote) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "content_object_id":
		if m.ContentObjectID == nil {
			return nil, fmt.Errorf("cannot fill relation for ContentObjectID while id field is empty")
		}
		parts := strings.Split(*m.ContentObjectID, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf("could not parse id field")
		}

		switch parts[0] {
		case "motion":
			var entry Motion
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
	case "meeting_user_id":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUser = &entry

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

func (m *PersonalNote) Get(field string) interface{} {
	switch field {
	case "content_object_id":
		return m.ContentObjectID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "meeting_user_id":
		return m.MeetingUserID
	case "note":
		return m.Note
	case "star":
		return m.Star
	}

	return nil
}

func (m *PersonalNote) GetFqids(field string) []string {
	switch field {
	case "content_object_id":
		if m.ContentObjectID != nil {
			return []string{*m.ContentObjectID}
		}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "meeting_user_id":
		return []string{"meeting_user/" + strconv.Itoa(m.MeetingUserID)}
	}
	return []string{}
}

func (m *PersonalNote) Update(data map[string]string) error {
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

	if val, ok := data["star"]; ok {
		err := json.Unmarshal([]byte(val), &m.Star)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *PersonalNote) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

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

func (m *PointOfOrderCategory) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "speaker_ids":
		for _, r := range m.speakers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
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
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Poll struct {
	Backend               string          `json:"backend"`
	ContentObjectID       string          `json:"content_object_id"`
	CryptKey              *string         `json:"crypt_key"`
	CryptSignature        *string         `json:"crypt_signature"`
	Description           *string         `json:"description"`
	EntitledGroupIDs      []int           `json:"entitled_group_ids"`
	EntitledUsersAtStop   json.RawMessage `json:"entitled_users_at_stop"`
	GlobalAbstain         *bool           `json:"global_abstain"`
	GlobalNo              *bool           `json:"global_no"`
	GlobalOptionID        *int            `json:"global_option_id"`
	GlobalYes             *bool           `json:"global_yes"`
	ID                    int             `json:"id"`
	IsPseudoanonymized    *bool           `json:"is_pseudoanonymized"`
	MaxVotesAmount        *int            `json:"max_votes_amount"`
	MaxVotesPerOption     *int            `json:"max_votes_per_option"`
	MeetingID             int             `json:"meeting_id"`
	MinVotesAmount        *int            `json:"min_votes_amount"`
	OnehundredPercentBase string          `json:"onehundred_percent_base"`
	OptionIDs             []int           `json:"option_ids"`
	Pollmethod            string          `json:"pollmethod"`
	ProjectionIDs         []int           `json:"projection_ids"`
	SequentialNumber      int             `json:"sequential_number"`
	State                 *string         `json:"state"`
	Title                 string          `json:"title"`
	Type                  string          `json:"type"`
	VoteCount             *int            `json:"vote_count"`
	VotedIDs              []int           `json:"voted_ids"`
	VotesRaw              *string         `json:"votes_raw"`
	VotesSignature        *string         `json:"votes_signature"`
	Votescast             *string         `json:"votescast"`
	Votesinvalid          *string         `json:"votesinvalid"`
	Votesvalid            *string         `json:"votesvalid"`
	loadedRelations       map[string]struct{}
	contentObject         IBaseModel
	entitledGroups        []*Group
	globalOption          *Option
	meeting               *Meeting
	options               []*Option
	projections           []*Projection
	voteds                []*User
}

func (m *Poll) CollectionName() string {
	return "poll"
}

func (m *Poll) ContentObject() IBaseModel {
	if _, ok := m.loadedRelations["content_object_id"]; !ok {
		log.Panic().Msg("Tried to access ContentObject relation of Poll which was not loaded.")
	}

	return m.contentObject
}

func (m *Poll) EntitledGroups() []*Group {
	if _, ok := m.loadedRelations["entitled_group_ids"]; !ok {
		log.Panic().Msg("Tried to access EntitledGroups relation of Poll which was not loaded.")
	}

	return m.entitledGroups
}

func (m *Poll) GlobalOption() *Option {
	if _, ok := m.loadedRelations["global_option_id"]; !ok {
		log.Panic().Msg("Tried to access GlobalOption relation of Poll which was not loaded.")
	}

	return m.globalOption
}

func (m *Poll) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Poll which was not loaded.")
	}

	return *m.meeting
}

func (m *Poll) Options() []*Option {
	if _, ok := m.loadedRelations["option_ids"]; !ok {
		log.Panic().Msg("Tried to access Options relation of Poll which was not loaded.")
	}

	return m.options
}

func (m *Poll) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of Poll which was not loaded.")
	}

	return m.projections
}

func (m *Poll) Voteds() []*User {
	if _, ok := m.loadedRelations["voted_ids"]; !ok {
		log.Panic().Msg("Tried to access Voteds relation of Poll which was not loaded.")
	}

	return m.voteds
}

func (m *Poll) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "content_object_id":
		return m.contentObject.GetRelatedModelsAccessor()
	case "entitled_group_ids":
		for _, r := range m.entitledGroups {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "global_option_id":
		return m.globalOption.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "option_ids":
		for _, r := range m.options {
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
	case "voted_ids":
		for _, r := range m.voteds {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *Poll) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "content_object_id":
			panic("not implemented")
		case "entitled_group_ids":
			m.entitledGroups = content.([]*Group)
		case "global_option_id":
			m.globalOption = content.(*Option)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "option_ids":
			m.options = content.([]*Option)
		case "projection_ids":
			m.projections = content.([]*Projection)
		case "voted_ids":
			m.voteds = content.([]*User)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Poll) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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

		case "motion":
			var entry Motion
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

	case "entitled_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.entitledGroups = append(m.entitledGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "global_option_id":
		var entry Option
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.globalOption = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "option_ids":
		var entry Option
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.options = append(m.options, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projections = append(m.projections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "voted_ids":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.voteds = append(m.voteds, &entry)

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

func (m *Poll) Get(field string) interface{} {
	switch field {
	case "backend":
		return m.Backend
	case "content_object_id":
		return m.ContentObjectID
	case "crypt_key":
		return m.CryptKey
	case "crypt_signature":
		return m.CryptSignature
	case "description":
		return m.Description
	case "entitled_group_ids":
		return m.EntitledGroupIDs
	case "entitled_users_at_stop":
		return m.EntitledUsersAtStop
	case "global_abstain":
		return m.GlobalAbstain
	case "global_no":
		return m.GlobalNo
	case "global_option_id":
		return m.GlobalOptionID
	case "global_yes":
		return m.GlobalYes
	case "id":
		return m.ID
	case "is_pseudoanonymized":
		return m.IsPseudoanonymized
	case "max_votes_amount":
		return m.MaxVotesAmount
	case "max_votes_per_option":
		return m.MaxVotesPerOption
	case "meeting_id":
		return m.MeetingID
	case "min_votes_amount":
		return m.MinVotesAmount
	case "onehundred_percent_base":
		return m.OnehundredPercentBase
	case "option_ids":
		return m.OptionIDs
	case "pollmethod":
		return m.Pollmethod
	case "projection_ids":
		return m.ProjectionIDs
	case "sequential_number":
		return m.SequentialNumber
	case "state":
		return m.State
	case "title":
		return m.Title
	case "type":
		return m.Type
	case "vote_count":
		return m.VoteCount
	case "voted_ids":
		return m.VotedIDs
	case "votes_raw":
		return m.VotesRaw
	case "votes_signature":
		return m.VotesSignature
	case "votescast":
		return m.Votescast
	case "votesinvalid":
		return m.Votesinvalid
	case "votesvalid":
		return m.Votesvalid
	}

	return nil
}

func (m *Poll) GetFqids(field string) []string {
	switch field {
	case "content_object_id":
		return []string{m.ContentObjectID}

	case "entitled_group_ids":
		r := make([]string, len(m.EntitledGroupIDs))
		for i, id := range m.EntitledGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "global_option_id":
		if m.GlobalOptionID != nil {
			return []string{"option/" + strconv.Itoa(*m.GlobalOptionID)}
		}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "option_ids":
		r := make([]string, len(m.OptionIDs))
		for i, id := range m.OptionIDs {
			r[i] = "option/" + strconv.Itoa(id)
		}
		return r

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "voted_ids":
		r := make([]string, len(m.VotedIDs))
		for i, id := range m.VotedIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *Poll) Update(data map[string]string) error {
	if val, ok := data["backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.Backend)
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

	if val, ok := data["crypt_key"]; ok {
		err := json.Unmarshal([]byte(val), &m.CryptKey)
		if err != nil {
			return err
		}
	}

	if val, ok := data["crypt_signature"]; ok {
		err := json.Unmarshal([]byte(val), &m.CryptSignature)
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

	if val, ok := data["entitled_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.EntitledGroupIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["entitled_group_ids"]; ok {
			m.entitledGroups = slices.DeleteFunc(m.entitledGroups, func(r *Group) bool {
				return !slices.Contains(m.EntitledGroupIDs, r.ID)
			})
		}
	}

	if val, ok := data["entitled_users_at_stop"]; ok {
		err := json.Unmarshal([]byte(val), &m.EntitledUsersAtStop)
		if err != nil {
			return err
		}
	}

	if val, ok := data["global_abstain"]; ok {
		err := json.Unmarshal([]byte(val), &m.GlobalAbstain)
		if err != nil {
			return err
		}
	}

	if val, ok := data["global_no"]; ok {
		err := json.Unmarshal([]byte(val), &m.GlobalNo)
		if err != nil {
			return err
		}
	}

	if val, ok := data["global_option_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.GlobalOptionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["global_yes"]; ok {
		err := json.Unmarshal([]byte(val), &m.GlobalYes)
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

	if val, ok := data["is_pseudoanonymized"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsPseudoanonymized)
		if err != nil {
			return err
		}
	}

	if val, ok := data["max_votes_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.MaxVotesAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["max_votes_per_option"]; ok {
		err := json.Unmarshal([]byte(val), &m.MaxVotesPerOption)
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

	if val, ok := data["min_votes_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.MinVotesAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.OnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["option_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["option_ids"]; ok {
			m.options = slices.DeleteFunc(m.options, func(r *Option) bool {
				return !slices.Contains(m.OptionIDs, r.ID)
			})
		}
	}

	if val, ok := data["pollmethod"]; ok {
		err := json.Unmarshal([]byte(val), &m.Pollmethod)
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

	if val, ok := data["state"]; ok {
		err := json.Unmarshal([]byte(val), &m.State)
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

	if val, ok := data["type"]; ok {
		err := json.Unmarshal([]byte(val), &m.Type)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_count"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteCount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["voted_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VotedIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["voted_ids"]; ok {
			m.voteds = slices.DeleteFunc(m.voteds, func(r *User) bool {
				return !slices.Contains(m.VotedIDs, r.ID)
			})
		}
	}

	if val, ok := data["votes_raw"]; ok {
		err := json.Unmarshal([]byte(val), &m.VotesRaw)
		if err != nil {
			return err
		}
	}

	if val, ok := data["votes_signature"]; ok {
		err := json.Unmarshal([]byte(val), &m.VotesSignature)
		if err != nil {
			return err
		}
	}

	if val, ok := data["votescast"]; ok {
		err := json.Unmarshal([]byte(val), &m.Votescast)
		if err != nil {
			return err
		}
	}

	if val, ok := data["votesinvalid"]; ok {
		err := json.Unmarshal([]byte(val), &m.Votesinvalid)
		if err != nil {
			return err
		}
	}

	if val, ok := data["votesvalid"]; ok {
		err := json.Unmarshal([]byte(val), &m.Votesvalid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Poll) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type PollCandidate struct {
	ID                  int  `json:"id"`
	MeetingID           int  `json:"meeting_id"`
	PollCandidateListID int  `json:"poll_candidate_list_id"`
	UserID              *int `json:"user_id"`
	Weight              int  `json:"weight"`
	loadedRelations     map[string]struct{}
	meeting             *Meeting
	pollCandidateList   *PollCandidateList
	user                *User
}

func (m *PollCandidate) CollectionName() string {
	return "poll_candidate"
}

func (m *PollCandidate) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PollCandidate which was not loaded.")
	}

	return *m.meeting
}

func (m *PollCandidate) PollCandidateList() PollCandidateList {
	if _, ok := m.loadedRelations["poll_candidate_list_id"]; !ok {
		log.Panic().Msg("Tried to access PollCandidateList relation of PollCandidate which was not loaded.")
	}

	return *m.pollCandidateList
}

func (m *PollCandidate) User() *User {
	if _, ok := m.loadedRelations["user_id"]; !ok {
		log.Panic().Msg("Tried to access User relation of PollCandidate which was not loaded.")
	}

	return m.user
}

func (m *PollCandidate) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "poll_candidate_list_id":
		return m.pollCandidateList.GetRelatedModelsAccessor()
	case "user_id":
		return m.user.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *PollCandidate) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "poll_candidate_list_id":
			m.pollCandidateList = content.(*PollCandidateList)
		case "user_id":
			m.user = content.(*User)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *PollCandidate) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "poll_candidate_list_id":
		var entry PollCandidateList
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCandidateList = &entry

		result = entry.GetRelatedModelsAccessor()
	case "user_id":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.user = &entry

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

func (m *PollCandidate) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "poll_candidate_list_id":
		return m.PollCandidateListID
	case "user_id":
		return m.UserID
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *PollCandidate) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "poll_candidate_list_id":
		return []string{"poll_candidate_list/" + strconv.Itoa(m.PollCandidateListID)}

	case "user_id":
		if m.UserID != nil {
			return []string{"user/" + strconv.Itoa(*m.UserID)}
		}
	}
	return []string{}
}

func (m *PollCandidate) Update(data map[string]string) error {
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

	if val, ok := data["poll_candidate_list_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateListID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserID)
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

func (m *PollCandidate) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type PollCandidateList struct {
	ID               int   `json:"id"`
	MeetingID        int   `json:"meeting_id"`
	OptionID         int   `json:"option_id"`
	PollCandidateIDs []int `json:"poll_candidate_ids"`
	loadedRelations  map[string]struct{}
	meeting          *Meeting
	option           *Option
	pollCandidates   []*PollCandidate
}

func (m *PollCandidateList) CollectionName() string {
	return "poll_candidate_list"
}

func (m *PollCandidateList) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of PollCandidateList which was not loaded.")
	}

	return *m.meeting
}

func (m *PollCandidateList) Option() Option {
	if _, ok := m.loadedRelations["option_id"]; !ok {
		log.Panic().Msg("Tried to access Option relation of PollCandidateList which was not loaded.")
	}

	return *m.option
}

func (m *PollCandidateList) PollCandidates() []*PollCandidate {
	if _, ok := m.loadedRelations["poll_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access PollCandidates relation of PollCandidateList which was not loaded.")
	}

	return m.pollCandidates
}

func (m *PollCandidateList) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "option_id":
		return m.option.GetRelatedModelsAccessor()
	case "poll_candidate_ids":
		for _, r := range m.pollCandidates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *PollCandidateList) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "option_id":
			m.option = content.(*Option)
		case "poll_candidate_ids":
			m.pollCandidates = content.([]*PollCandidate)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *PollCandidateList) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	case "option_id":
		var entry Option
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.option = &entry

		result = entry.GetRelatedModelsAccessor()
	case "poll_candidate_ids":
		var entry PollCandidate
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCandidates = append(m.pollCandidates, &entry)

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

func (m *PollCandidateList) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "option_id":
		return m.OptionID
	case "poll_candidate_ids":
		return m.PollCandidateIDs
	}

	return nil
}

func (m *PollCandidateList) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "option_id":
		return []string{"option/" + strconv.Itoa(m.OptionID)}

	case "poll_candidate_ids":
		r := make([]string, len(m.PollCandidateIDs))
		for i, id := range m.PollCandidateIDs {
			r[i] = "poll_candidate/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *PollCandidateList) Update(data map[string]string) error {
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

	if val, ok := data["option_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["poll_candidate_ids"]; ok {
			m.pollCandidates = slices.DeleteFunc(m.pollCandidates, func(r *PollCandidate) bool {
				return !slices.Contains(m.PollCandidateIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *PollCandidateList) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Projection struct {
	Content            json.RawMessage `json:"content"`
	ContentObjectID    string          `json:"content_object_id"`
	CurrentProjectorID *int            `json:"current_projector_id"`
	HistoryProjectorID *int            `json:"history_projector_id"`
	ID                 int             `json:"id"`
	MeetingID          int             `json:"meeting_id"`
	Options            json.RawMessage `json:"options"`
	PreviewProjectorID *int            `json:"preview_projector_id"`
	Stable             *bool           `json:"stable"`
	Type               *string         `json:"type"`
	Weight             *int            `json:"weight"`
	loadedRelations    map[string]struct{}
	contentObject      IBaseModel
	currentProjector   *Projector
	historyProjector   *Projector
	meeting            *Meeting
	previewProjector   *Projector
}

func (m *Projection) CollectionName() string {
	return "projection"
}

func (m *Projection) ContentObject() IBaseModel {
	if _, ok := m.loadedRelations["content_object_id"]; !ok {
		log.Panic().Msg("Tried to access ContentObject relation of Projection which was not loaded.")
	}

	return m.contentObject
}

func (m *Projection) CurrentProjector() *Projector {
	if _, ok := m.loadedRelations["current_projector_id"]; !ok {
		log.Panic().Msg("Tried to access CurrentProjector relation of Projection which was not loaded.")
	}

	return m.currentProjector
}

func (m *Projection) HistoryProjector() *Projector {
	if _, ok := m.loadedRelations["history_projector_id"]; !ok {
		log.Panic().Msg("Tried to access HistoryProjector relation of Projection which was not loaded.")
	}

	return m.historyProjector
}

func (m *Projection) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Projection which was not loaded.")
	}

	return *m.meeting
}

func (m *Projection) PreviewProjector() *Projector {
	if _, ok := m.loadedRelations["preview_projector_id"]; !ok {
		log.Panic().Msg("Tried to access PreviewProjector relation of Projection which was not loaded.")
	}

	return m.previewProjector
}

func (m *Projection) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "content_object_id":
		return m.contentObject.GetRelatedModelsAccessor()
	case "current_projector_id":
		return m.currentProjector.GetRelatedModelsAccessor()
	case "history_projector_id":
		return m.historyProjector.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "preview_projector_id":
		return m.previewProjector.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *Projection) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "content_object_id":
			panic("not implemented")
		case "current_projector_id":
			m.currentProjector = content.(*Projector)
		case "history_projector_id":
			m.historyProjector = content.(*Projector)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "preview_projector_id":
			m.previewProjector = content.(*Projector)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Projection) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "content_object_id":
		parts := strings.Split(m.ContentObjectID, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf("could not parse id field")
		}

		switch parts[0] {
		case "agenda_item":
			var entry AgendaItem
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "assignment":
			var entry Assignment
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "list_of_speakers":
			var entry ListOfSpeakers
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "meeting":
			var entry Meeting
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

		case "poll":
			var entry Poll
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "projector_countdown":
			var entry ProjectorCountdown
			err := json.Unmarshal(content, &entry)
			if err != nil {
				return nil, err
			}
			m.contentObject = &entry
			result = entry.GetRelatedModelsAccessor()

		case "projector_message":
			var entry ProjectorMessage
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

	case "current_projector_id":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.currentProjector = &entry

		result = entry.GetRelatedModelsAccessor()
	case "history_projector_id":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.historyProjector = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "preview_projector_id":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.previewProjector = &entry

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

func (m *Projection) Get(field string) interface{} {
	switch field {
	case "content":
		return m.Content
	case "content_object_id":
		return m.ContentObjectID
	case "current_projector_id":
		return m.CurrentProjectorID
	case "history_projector_id":
		return m.HistoryProjectorID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "options":
		return m.Options
	case "preview_projector_id":
		return m.PreviewProjectorID
	case "stable":
		return m.Stable
	case "type":
		return m.Type
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *Projection) GetFqids(field string) []string {
	switch field {
	case "content_object_id":
		return []string{m.ContentObjectID}

	case "current_projector_id":
		if m.CurrentProjectorID != nil {
			return []string{"projector/" + strconv.Itoa(*m.CurrentProjectorID)}
		}

	case "history_projector_id":
		if m.HistoryProjectorID != nil {
			return []string{"projector/" + strconv.Itoa(*m.HistoryProjectorID)}
		}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "preview_projector_id":
		if m.PreviewProjectorID != nil {
			return []string{"projector/" + strconv.Itoa(*m.PreviewProjectorID)}
		}
	}
	return []string{}
}

func (m *Projection) Update(data map[string]string) error {
	if val, ok := data["content"]; ok {
		err := json.Unmarshal([]byte(val), &m.Content)
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

	if val, ok := data["current_projector_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.CurrentProjectorID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["history_projector_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.HistoryProjectorID)
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

	if val, ok := data["options"]; ok {
		err := json.Unmarshal([]byte(val), &m.Options)
		if err != nil {
			return err
		}
	}

	if val, ok := data["preview_projector_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PreviewProjectorID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["stable"]; ok {
		err := json.Unmarshal([]byte(val), &m.Stable)
		if err != nil {
			return err
		}
	}

	if val, ok := data["type"]; ok {
		err := json.Unmarshal([]byte(val), &m.Type)
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

func (m *Projection) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Projector struct {
	AspectRatioDenominator                                    *int    `json:"aspect_ratio_denominator"`
	AspectRatioNumerator                                      *int    `json:"aspect_ratio_numerator"`
	BackgroundColor                                           *string `json:"background_color"`
	ChyronBackgroundColor                                     *string `json:"chyron_background_color"`
	ChyronBackgroundColor2                                    *string `json:"chyron_background_color_2"`
	ChyronFontColor                                           *string `json:"chyron_font_color"`
	ChyronFontColor2                                          *string `json:"chyron_font_color_2"`
	Color                                                     *string `json:"color"`
	CurrentProjectionIDs                                      []int   `json:"current_projection_ids"`
	HeaderBackgroundColor                                     *string `json:"header_background_color"`
	HeaderFontColor                                           *string `json:"header_font_color"`
	HeaderH1Color                                             *string `json:"header_h1_color"`
	HistoryProjectionIDs                                      []int   `json:"history_projection_ids"`
	ID                                                        int     `json:"id"`
	IsInternal                                                *bool   `json:"is_internal"`
	MeetingID                                                 int     `json:"meeting_id"`
	Name                                                      *string `json:"name"`
	PreviewProjectionIDs                                      []int   `json:"preview_projection_ids"`
	Scale                                                     *int    `json:"scale"`
	Scroll                                                    *int    `json:"scroll"`
	SequentialNumber                                          int     `json:"sequential_number"`
	ShowClock                                                 *bool   `json:"show_clock"`
	ShowHeaderFooter                                          *bool   `json:"show_header_footer"`
	ShowLogo                                                  *bool   `json:"show_logo"`
	ShowTitle                                                 *bool   `json:"show_title"`
	UsedAsDefaultProjectorForAgendaItemListInMeetingID        *int    `json:"used_as_default_projector_for_agenda_item_list_in_meeting_id"`
	UsedAsDefaultProjectorForAmendmentInMeetingID             *int    `json:"used_as_default_projector_for_amendment_in_meeting_id"`
	UsedAsDefaultProjectorForAssignmentInMeetingID            *int    `json:"used_as_default_projector_for_assignment_in_meeting_id"`
	UsedAsDefaultProjectorForAssignmentPollInMeetingID        *int    `json:"used_as_default_projector_for_assignment_poll_in_meeting_id"`
	UsedAsDefaultProjectorForCountdownInMeetingID             *int    `json:"used_as_default_projector_for_countdown_in_meeting_id"`
	UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID *int    `json:"used_as_default_projector_for_current_list_of_speakers_in_meeting_id"`
	UsedAsDefaultProjectorForListOfSpeakersInMeetingID        *int    `json:"used_as_default_projector_for_list_of_speakers_in_meeting_id"`
	UsedAsDefaultProjectorForMediafileInMeetingID             *int    `json:"used_as_default_projector_for_mediafile_in_meeting_id"`
	UsedAsDefaultProjectorForMessageInMeetingID               *int    `json:"used_as_default_projector_for_message_in_meeting_id"`
	UsedAsDefaultProjectorForMotionBlockInMeetingID           *int    `json:"used_as_default_projector_for_motion_block_in_meeting_id"`
	UsedAsDefaultProjectorForMotionInMeetingID                *int    `json:"used_as_default_projector_for_motion_in_meeting_id"`
	UsedAsDefaultProjectorForMotionPollInMeetingID            *int    `json:"used_as_default_projector_for_motion_poll_in_meeting_id"`
	UsedAsDefaultProjectorForPollInMeetingID                  *int    `json:"used_as_default_projector_for_poll_in_meeting_id"`
	UsedAsDefaultProjectorForTopicInMeetingID                 *int    `json:"used_as_default_projector_for_topic_in_meeting_id"`
	UsedAsReferenceProjectorMeetingID                         *int    `json:"used_as_reference_projector_meeting_id"`
	Width                                                     *int    `json:"width"`
	loadedRelations                                           map[string]struct{}
	currentProjections                                        []*Projection
	historyProjections                                        []*Projection
	meeting                                                   *Meeting
	previewProjections                                        []*Projection
	usedAsDefaultProjectorForAgendaItemListInMeeting          *Meeting
	usedAsDefaultProjectorForAmendmentInMeeting               *Meeting
	usedAsDefaultProjectorForAssignmentInMeeting              *Meeting
	usedAsDefaultProjectorForAssignmentPollInMeeting          *Meeting
	usedAsDefaultProjectorForCountdownInMeeting               *Meeting
	usedAsDefaultProjectorForCurrentListOfSpeakersInMeeting   *Meeting
	usedAsDefaultProjectorForListOfSpeakersInMeeting          *Meeting
	usedAsDefaultProjectorForMediafileInMeeting               *Meeting
	usedAsDefaultProjectorForMessageInMeeting                 *Meeting
	usedAsDefaultProjectorForMotionBlockInMeeting             *Meeting
	usedAsDefaultProjectorForMotionInMeeting                  *Meeting
	usedAsDefaultProjectorForMotionPollInMeeting              *Meeting
	usedAsDefaultProjectorForPollInMeeting                    *Meeting
	usedAsDefaultProjectorForTopicInMeeting                   *Meeting
	usedAsReferenceProjectorMeeting                           *Meeting
}

func (m *Projector) CollectionName() string {
	return "projector"
}

func (m *Projector) CurrentProjections() []*Projection {
	if _, ok := m.loadedRelations["current_projection_ids"]; !ok {
		log.Panic().Msg("Tried to access CurrentProjections relation of Projector which was not loaded.")
	}

	return m.currentProjections
}

func (m *Projector) HistoryProjections() []*Projection {
	if _, ok := m.loadedRelations["history_projection_ids"]; !ok {
		log.Panic().Msg("Tried to access HistoryProjections relation of Projector which was not loaded.")
	}

	return m.historyProjections
}

func (m *Projector) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Projector which was not loaded.")
	}

	return *m.meeting
}

func (m *Projector) PreviewProjections() []*Projection {
	if _, ok := m.loadedRelations["preview_projection_ids"]; !ok {
		log.Panic().Msg("Tried to access PreviewProjections relation of Projector which was not loaded.")
	}

	return m.previewProjections
}

func (m *Projector) UsedAsDefaultProjectorForAgendaItemListInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_agenda_item_list_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForAgendaItemListInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForAgendaItemListInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForAmendmentInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_amendment_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForAmendmentInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForAmendmentInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForAssignmentInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_assignment_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForAssignmentInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForAssignmentInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForAssignmentPollInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_assignment_poll_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForAssignmentPollInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForAssignmentPollInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForCountdownInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_countdown_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForCountdownInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForCountdownInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForCurrentListOfSpeakersInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_current_list_of_speakers_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForCurrentListOfSpeakersInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForCurrentListOfSpeakersInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForListOfSpeakersInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_list_of_speakers_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForListOfSpeakersInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForListOfSpeakersInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForMediafileInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_mediafile_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMediafileInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMediafileInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForMessageInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_message_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMessageInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMessageInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForMotionBlockInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_motion_block_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMotionBlockInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMotionBlockInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForMotionInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_motion_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMotionInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMotionInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForMotionPollInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_motion_poll_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForMotionPollInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForMotionPollInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForPollInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_poll_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForPollInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForPollInMeeting
}

func (m *Projector) UsedAsDefaultProjectorForTopicInMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_default_projector_for_topic_in_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsDefaultProjectorForTopicInMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsDefaultProjectorForTopicInMeeting
}

func (m *Projector) UsedAsReferenceProjectorMeeting() *Meeting {
	if _, ok := m.loadedRelations["used_as_reference_projector_meeting_id"]; !ok {
		log.Panic().Msg("Tried to access UsedAsReferenceProjectorMeeting relation of Projector which was not loaded.")
	}

	return m.usedAsReferenceProjectorMeeting
}

func (m *Projector) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "current_projection_ids":
		for _, r := range m.currentProjections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "history_projection_ids":
		for _, r := range m.historyProjections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "preview_projection_ids":
		for _, r := range m.previewProjections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "used_as_default_projector_for_agenda_item_list_in_meeting_id":
		return m.usedAsDefaultProjectorForAgendaItemListInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_amendment_in_meeting_id":
		return m.usedAsDefaultProjectorForAmendmentInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_assignment_in_meeting_id":
		return m.usedAsDefaultProjectorForAssignmentInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_assignment_poll_in_meeting_id":
		return m.usedAsDefaultProjectorForAssignmentPollInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_countdown_in_meeting_id":
		return m.usedAsDefaultProjectorForCountdownInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_current_list_of_speakers_in_meeting_id":
		return m.usedAsDefaultProjectorForCurrentListOfSpeakersInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_list_of_speakers_in_meeting_id":
		return m.usedAsDefaultProjectorForListOfSpeakersInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_mediafile_in_meeting_id":
		return m.usedAsDefaultProjectorForMediafileInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_message_in_meeting_id":
		return m.usedAsDefaultProjectorForMessageInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_motion_block_in_meeting_id":
		return m.usedAsDefaultProjectorForMotionBlockInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_motion_in_meeting_id":
		return m.usedAsDefaultProjectorForMotionInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_motion_poll_in_meeting_id":
		return m.usedAsDefaultProjectorForMotionPollInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_poll_in_meeting_id":
		return m.usedAsDefaultProjectorForPollInMeeting.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_topic_in_meeting_id":
		return m.usedAsDefaultProjectorForTopicInMeeting.GetRelatedModelsAccessor()
	case "used_as_reference_projector_meeting_id":
		return m.usedAsReferenceProjectorMeeting.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *Projector) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "current_projection_ids":
			m.currentProjections = content.([]*Projection)
		case "history_projection_ids":
			m.historyProjections = content.([]*Projection)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "preview_projection_ids":
			m.previewProjections = content.([]*Projection)
		case "used_as_default_projector_for_agenda_item_list_in_meeting_id":
			m.usedAsDefaultProjectorForAgendaItemListInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_amendment_in_meeting_id":
			m.usedAsDefaultProjectorForAmendmentInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_assignment_in_meeting_id":
			m.usedAsDefaultProjectorForAssignmentInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_assignment_poll_in_meeting_id":
			m.usedAsDefaultProjectorForAssignmentPollInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_countdown_in_meeting_id":
			m.usedAsDefaultProjectorForCountdownInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_current_list_of_speakers_in_meeting_id":
			m.usedAsDefaultProjectorForCurrentListOfSpeakersInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_list_of_speakers_in_meeting_id":
			m.usedAsDefaultProjectorForListOfSpeakersInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_mediafile_in_meeting_id":
			m.usedAsDefaultProjectorForMediafileInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_message_in_meeting_id":
			m.usedAsDefaultProjectorForMessageInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_motion_block_in_meeting_id":
			m.usedAsDefaultProjectorForMotionBlockInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_motion_in_meeting_id":
			m.usedAsDefaultProjectorForMotionInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_motion_poll_in_meeting_id":
			m.usedAsDefaultProjectorForMotionPollInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_poll_in_meeting_id":
			m.usedAsDefaultProjectorForPollInMeeting = content.(*Meeting)
		case "used_as_default_projector_for_topic_in_meeting_id":
			m.usedAsDefaultProjectorForTopicInMeeting = content.(*Meeting)
		case "used_as_reference_projector_meeting_id":
			m.usedAsReferenceProjectorMeeting = content.(*Meeting)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Projector) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "current_projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.currentProjections = append(m.currentProjections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "history_projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.historyProjections = append(m.historyProjections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "preview_projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.previewProjections = append(m.previewProjections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_agenda_item_list_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForAgendaItemListInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_amendment_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForAmendmentInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_assignment_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForAssignmentInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_assignment_poll_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForAssignmentPollInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_countdown_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForCountdownInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_current_list_of_speakers_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForCurrentListOfSpeakersInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_list_of_speakers_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForListOfSpeakersInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_mediafile_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForMediafileInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_message_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForMessageInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_motion_block_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForMotionBlockInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_motion_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForMotionInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_motion_poll_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForMotionPollInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_poll_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForPollInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_default_projector_for_topic_in_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsDefaultProjectorForTopicInMeeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "used_as_reference_projector_meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.usedAsReferenceProjectorMeeting = &entry

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

func (m *Projector) Get(field string) interface{} {
	switch field {
	case "aspect_ratio_denominator":
		return m.AspectRatioDenominator
	case "aspect_ratio_numerator":
		return m.AspectRatioNumerator
	case "background_color":
		return m.BackgroundColor
	case "chyron_background_color":
		return m.ChyronBackgroundColor
	case "chyron_background_color_2":
		return m.ChyronBackgroundColor2
	case "chyron_font_color":
		return m.ChyronFontColor
	case "chyron_font_color_2":
		return m.ChyronFontColor2
	case "color":
		return m.Color
	case "current_projection_ids":
		return m.CurrentProjectionIDs
	case "header_background_color":
		return m.HeaderBackgroundColor
	case "header_font_color":
		return m.HeaderFontColor
	case "header_h1_color":
		return m.HeaderH1Color
	case "history_projection_ids":
		return m.HistoryProjectionIDs
	case "id":
		return m.ID
	case "is_internal":
		return m.IsInternal
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "preview_projection_ids":
		return m.PreviewProjectionIDs
	case "scale":
		return m.Scale
	case "scroll":
		return m.Scroll
	case "sequential_number":
		return m.SequentialNumber
	case "show_clock":
		return m.ShowClock
	case "show_header_footer":
		return m.ShowHeaderFooter
	case "show_logo":
		return m.ShowLogo
	case "show_title":
		return m.ShowTitle
	case "used_as_default_projector_for_agenda_item_list_in_meeting_id":
		return m.UsedAsDefaultProjectorForAgendaItemListInMeetingID
	case "used_as_default_projector_for_amendment_in_meeting_id":
		return m.UsedAsDefaultProjectorForAmendmentInMeetingID
	case "used_as_default_projector_for_assignment_in_meeting_id":
		return m.UsedAsDefaultProjectorForAssignmentInMeetingID
	case "used_as_default_projector_for_assignment_poll_in_meeting_id":
		return m.UsedAsDefaultProjectorForAssignmentPollInMeetingID
	case "used_as_default_projector_for_countdown_in_meeting_id":
		return m.UsedAsDefaultProjectorForCountdownInMeetingID
	case "used_as_default_projector_for_current_list_of_speakers_in_meeting_id":
		return m.UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID
	case "used_as_default_projector_for_list_of_speakers_in_meeting_id":
		return m.UsedAsDefaultProjectorForListOfSpeakersInMeetingID
	case "used_as_default_projector_for_mediafile_in_meeting_id":
		return m.UsedAsDefaultProjectorForMediafileInMeetingID
	case "used_as_default_projector_for_message_in_meeting_id":
		return m.UsedAsDefaultProjectorForMessageInMeetingID
	case "used_as_default_projector_for_motion_block_in_meeting_id":
		return m.UsedAsDefaultProjectorForMotionBlockInMeetingID
	case "used_as_default_projector_for_motion_in_meeting_id":
		return m.UsedAsDefaultProjectorForMotionInMeetingID
	case "used_as_default_projector_for_motion_poll_in_meeting_id":
		return m.UsedAsDefaultProjectorForMotionPollInMeetingID
	case "used_as_default_projector_for_poll_in_meeting_id":
		return m.UsedAsDefaultProjectorForPollInMeetingID
	case "used_as_default_projector_for_topic_in_meeting_id":
		return m.UsedAsDefaultProjectorForTopicInMeetingID
	case "used_as_reference_projector_meeting_id":
		return m.UsedAsReferenceProjectorMeetingID
	case "width":
		return m.Width
	}

	return nil
}

func (m *Projector) GetFqids(field string) []string {
	switch field {
	case "current_projection_ids":
		r := make([]string, len(m.CurrentProjectionIDs))
		for i, id := range m.CurrentProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "history_projection_ids":
		r := make([]string, len(m.HistoryProjectionIDs))
		for i, id := range m.HistoryProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "preview_projection_ids":
		r := make([]string, len(m.PreviewProjectionIDs))
		for i, id := range m.PreviewProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "used_as_default_projector_for_agenda_item_list_in_meeting_id":
		if m.UsedAsDefaultProjectorForAgendaItemListInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForAgendaItemListInMeetingID)}
		}

	case "used_as_default_projector_for_amendment_in_meeting_id":
		if m.UsedAsDefaultProjectorForAmendmentInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForAmendmentInMeetingID)}
		}

	case "used_as_default_projector_for_assignment_in_meeting_id":
		if m.UsedAsDefaultProjectorForAssignmentInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForAssignmentInMeetingID)}
		}

	case "used_as_default_projector_for_assignment_poll_in_meeting_id":
		if m.UsedAsDefaultProjectorForAssignmentPollInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForAssignmentPollInMeetingID)}
		}

	case "used_as_default_projector_for_countdown_in_meeting_id":
		if m.UsedAsDefaultProjectorForCountdownInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForCountdownInMeetingID)}
		}

	case "used_as_default_projector_for_current_list_of_speakers_in_meeting_id":
		if m.UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID)}
		}

	case "used_as_default_projector_for_list_of_speakers_in_meeting_id":
		if m.UsedAsDefaultProjectorForListOfSpeakersInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForListOfSpeakersInMeetingID)}
		}

	case "used_as_default_projector_for_mediafile_in_meeting_id":
		if m.UsedAsDefaultProjectorForMediafileInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForMediafileInMeetingID)}
		}

	case "used_as_default_projector_for_message_in_meeting_id":
		if m.UsedAsDefaultProjectorForMessageInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForMessageInMeetingID)}
		}

	case "used_as_default_projector_for_motion_block_in_meeting_id":
		if m.UsedAsDefaultProjectorForMotionBlockInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForMotionBlockInMeetingID)}
		}

	case "used_as_default_projector_for_motion_in_meeting_id":
		if m.UsedAsDefaultProjectorForMotionInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForMotionInMeetingID)}
		}

	case "used_as_default_projector_for_motion_poll_in_meeting_id":
		if m.UsedAsDefaultProjectorForMotionPollInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForMotionPollInMeetingID)}
		}

	case "used_as_default_projector_for_poll_in_meeting_id":
		if m.UsedAsDefaultProjectorForPollInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForPollInMeetingID)}
		}

	case "used_as_default_projector_for_topic_in_meeting_id":
		if m.UsedAsDefaultProjectorForTopicInMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsDefaultProjectorForTopicInMeetingID)}
		}

	case "used_as_reference_projector_meeting_id":
		if m.UsedAsReferenceProjectorMeetingID != nil {
			return []string{"meeting/" + strconv.Itoa(*m.UsedAsReferenceProjectorMeetingID)}
		}
	}
	return []string{}
}

func (m *Projector) Update(data map[string]string) error {
	if val, ok := data["aspect_ratio_denominator"]; ok {
		err := json.Unmarshal([]byte(val), &m.AspectRatioDenominator)
		if err != nil {
			return err
		}
	}

	if val, ok := data["aspect_ratio_numerator"]; ok {
		err := json.Unmarshal([]byte(val), &m.AspectRatioNumerator)
		if err != nil {
			return err
		}
	}

	if val, ok := data["background_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.BackgroundColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chyron_background_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChyronBackgroundColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chyron_background_color_2"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChyronBackgroundColor2)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chyron_font_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChyronFontColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chyron_font_color_2"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChyronFontColor2)
		if err != nil {
			return err
		}
	}

	if val, ok := data["color"]; ok {
		err := json.Unmarshal([]byte(val), &m.Color)
		if err != nil {
			return err
		}
	}

	if val, ok := data["current_projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CurrentProjectionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["current_projection_ids"]; ok {
			m.currentProjections = slices.DeleteFunc(m.currentProjections, func(r *Projection) bool {
				return !slices.Contains(m.CurrentProjectionIDs, r.ID)
			})
		}
	}

	if val, ok := data["header_background_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.HeaderBackgroundColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["header_font_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.HeaderFontColor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["header_h1_color"]; ok {
		err := json.Unmarshal([]byte(val), &m.HeaderH1Color)
		if err != nil {
			return err
		}
	}

	if val, ok := data["history_projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.HistoryProjectionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["history_projection_ids"]; ok {
			m.historyProjections = slices.DeleteFunc(m.historyProjections, func(r *Projection) bool {
				return !slices.Contains(m.HistoryProjectionIDs, r.ID)
			})
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_internal"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsInternal)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["preview_projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PreviewProjectionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["preview_projection_ids"]; ok {
			m.previewProjections = slices.DeleteFunc(m.previewProjections, func(r *Projection) bool {
				return !slices.Contains(m.PreviewProjectionIDs, r.ID)
			})
		}
	}

	if val, ok := data["scale"]; ok {
		err := json.Unmarshal([]byte(val), &m.Scale)
		if err != nil {
			return err
		}
	}

	if val, ok := data["scroll"]; ok {
		err := json.Unmarshal([]byte(val), &m.Scroll)
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

	if val, ok := data["show_clock"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowClock)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_header_footer"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowHeaderFooter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_logo"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowLogo)
		if err != nil {
			return err
		}
	}

	if val, ok := data["show_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.ShowTitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_agenda_item_list_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForAgendaItemListInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_amendment_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForAmendmentInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_assignment_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForAssignmentInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_assignment_poll_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForAssignmentPollInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_countdown_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForCountdownInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_current_list_of_speakers_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_list_of_speakers_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForListOfSpeakersInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_mediafile_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMediafileInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_message_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMessageInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_motion_block_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMotionBlockInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_motion_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMotionInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_motion_poll_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForMotionPollInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_poll_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForPollInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_default_projector_for_topic_in_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsDefaultProjectorForTopicInMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["used_as_reference_projector_meeting_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsedAsReferenceProjectorMeetingID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["width"]; ok {
		err := json.Unmarshal([]byte(val), &m.Width)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Projector) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

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

func (m *ProjectorMessage) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "projection_ids":
		for _, r := range m.projections {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
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
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

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

func (m *Speaker) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "list_of_speakers_id":
		return m.listOfSpeakers.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "meeting_user_id":
		return m.meetingUser.GetRelatedModelsAccessor()
	case "point_of_order_category_id":
		return m.pointOfOrderCategory.GetRelatedModelsAccessor()
	case "structure_level_list_of_speakers_id":
		return m.structureLevelListOfSpeakers.GetRelatedModelsAccessor()
	}

	return nil
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
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

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
	listOfSpeakers   *ListOfSpeakers
	meeting          *Meeting
	speakers         []*Speaker
	structureLevel   *StructureLevel
}

func (m *StructureLevelListOfSpeakers) CollectionName() string {
	return "structure_level_list_of_speakers"
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

func (m *StructureLevelListOfSpeakers) StructureLevel() StructureLevel {
	if _, ok := m.loadedRelations["structure_level_id"]; !ok {
		log.Panic().Msg("Tried to access StructureLevel relation of StructureLevelListOfSpeakers which was not loaded.")
	}

	return *m.structureLevel
}

func (m *StructureLevelListOfSpeakers) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "list_of_speakers_id":
		return m.listOfSpeakers.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "speaker_ids":
		for _, r := range m.speakers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "structure_level_id":
		return m.structureLevel.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *StructureLevelListOfSpeakers) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "list_of_speakers_id":
			m.listOfSpeakers = content.(*ListOfSpeakers)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "speaker_ids":
			m.speakers = content.([]*Speaker)
		case "structure_level_id":
			m.structureLevel = content.(*StructureLevel)
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
	case "structure_level_id":
		var entry StructureLevel
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.structureLevel = &entry

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

	case "structure_level_id":
		return []string{"structure_level/" + strconv.Itoa(m.StructureLevelID)}
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

		if _, ok := m.loadedRelations["speaker_ids"]; ok {
			m.speakers = slices.DeleteFunc(m.speakers, func(r *Speaker) bool {
				return !slices.Contains(m.SpeakerIDs, r.ID)
			})
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
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Tag struct {
	ID              int      `json:"id"`
	MeetingID       int      `json:"meeting_id"`
	Name            string   `json:"name"`
	TaggedIDs       []string `json:"tagged_ids"`
	loadedRelations map[string]struct{}
	meeting         *Meeting
}

func (m *Tag) CollectionName() string {
	return "tag"
}

func (m *Tag) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Tag which was not loaded.")
	}

	return *m.meeting
}

func (m *Tag) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *Tag) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "meeting_id":
			m.meeting = content.(*Meeting)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Tag) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
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
	default:
		return nil, fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return result, nil
}

func (m *Tag) Get(field string) interface{} {
	switch field {
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "name":
		return m.Name
	case "tagged_ids":
		return m.TaggedIDs
	}

	return nil
}

func (m *Tag) GetFqids(field string) []string {
	switch field {
	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}
	}
	return []string{}
}

func (m *Tag) Update(data map[string]string) error {
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["tagged_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TaggedIDs)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Tag) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Theme struct {
	Abstain                *string `json:"abstain"`
	Accent100              *string `json:"accent_100"`
	Accent200              *string `json:"accent_200"`
	Accent300              *string `json:"accent_300"`
	Accent400              *string `json:"accent_400"`
	Accent50               *string `json:"accent_50"`
	Accent500              string  `json:"accent_500"`
	Accent600              *string `json:"accent_600"`
	Accent700              *string `json:"accent_700"`
	Accent800              *string `json:"accent_800"`
	Accent900              *string `json:"accent_900"`
	AccentA100             *string `json:"accent_a100"`
	AccentA200             *string `json:"accent_a200"`
	AccentA400             *string `json:"accent_a400"`
	AccentA700             *string `json:"accent_a700"`
	Headbar                *string `json:"headbar"`
	ID                     int     `json:"id"`
	Name                   string  `json:"name"`
	No                     *string `json:"no"`
	OrganizationID         int     `json:"organization_id"`
	Primary100             *string `json:"primary_100"`
	Primary200             *string `json:"primary_200"`
	Primary300             *string `json:"primary_300"`
	Primary400             *string `json:"primary_400"`
	Primary50              *string `json:"primary_50"`
	Primary500             string  `json:"primary_500"`
	Primary600             *string `json:"primary_600"`
	Primary700             *string `json:"primary_700"`
	Primary800             *string `json:"primary_800"`
	Primary900             *string `json:"primary_900"`
	PrimaryA100            *string `json:"primary_a100"`
	PrimaryA200            *string `json:"primary_a200"`
	PrimaryA400            *string `json:"primary_a400"`
	PrimaryA700            *string `json:"primary_a700"`
	ThemeForOrganizationID *int    `json:"theme_for_organization_id"`
	Warn100                *string `json:"warn_100"`
	Warn200                *string `json:"warn_200"`
	Warn300                *string `json:"warn_300"`
	Warn400                *string `json:"warn_400"`
	Warn50                 *string `json:"warn_50"`
	Warn500                string  `json:"warn_500"`
	Warn600                *string `json:"warn_600"`
	Warn700                *string `json:"warn_700"`
	Warn800                *string `json:"warn_800"`
	Warn900                *string `json:"warn_900"`
	WarnA100               *string `json:"warn_a100"`
	WarnA200               *string `json:"warn_a200"`
	WarnA400               *string `json:"warn_a400"`
	WarnA700               *string `json:"warn_a700"`
	Yes                    *string `json:"yes"`
	loadedRelations        map[string]struct{}
	organization           *Organization
	themeForOrganization   *Organization
}

func (m *Theme) CollectionName() string {
	return "theme"
}

func (m *Theme) Organization() Organization {
	if _, ok := m.loadedRelations["organization_id"]; !ok {
		log.Panic().Msg("Tried to access Organization relation of Theme which was not loaded.")
	}

	return *m.organization
}

func (m *Theme) ThemeForOrganization() *Organization {
	if _, ok := m.loadedRelations["theme_for_organization_id"]; !ok {
		log.Panic().Msg("Tried to access ThemeForOrganization relation of Theme which was not loaded.")
	}

	return m.themeForOrganization
}

func (m *Theme) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "organization_id":
		return m.organization.GetRelatedModelsAccessor()
	case "theme_for_organization_id":
		return m.themeForOrganization.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *Theme) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "organization_id":
			m.organization = content.(*Organization)
		case "theme_for_organization_id":
			m.themeForOrganization = content.(*Organization)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Theme) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.organization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "theme_for_organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.themeForOrganization = &entry

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

func (m *Theme) Get(field string) interface{} {
	switch field {
	case "abstain":
		return m.Abstain
	case "accent_100":
		return m.Accent100
	case "accent_200":
		return m.Accent200
	case "accent_300":
		return m.Accent300
	case "accent_400":
		return m.Accent400
	case "accent_50":
		return m.Accent50
	case "accent_500":
		return m.Accent500
	case "accent_600":
		return m.Accent600
	case "accent_700":
		return m.Accent700
	case "accent_800":
		return m.Accent800
	case "accent_900":
		return m.Accent900
	case "accent_a100":
		return m.AccentA100
	case "accent_a200":
		return m.AccentA200
	case "accent_a400":
		return m.AccentA400
	case "accent_a700":
		return m.AccentA700
	case "headbar":
		return m.Headbar
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "no":
		return m.No
	case "organization_id":
		return m.OrganizationID
	case "primary_100":
		return m.Primary100
	case "primary_200":
		return m.Primary200
	case "primary_300":
		return m.Primary300
	case "primary_400":
		return m.Primary400
	case "primary_50":
		return m.Primary50
	case "primary_500":
		return m.Primary500
	case "primary_600":
		return m.Primary600
	case "primary_700":
		return m.Primary700
	case "primary_800":
		return m.Primary800
	case "primary_900":
		return m.Primary900
	case "primary_a100":
		return m.PrimaryA100
	case "primary_a200":
		return m.PrimaryA200
	case "primary_a400":
		return m.PrimaryA400
	case "primary_a700":
		return m.PrimaryA700
	case "theme_for_organization_id":
		return m.ThemeForOrganizationID
	case "warn_100":
		return m.Warn100
	case "warn_200":
		return m.Warn200
	case "warn_300":
		return m.Warn300
	case "warn_400":
		return m.Warn400
	case "warn_50":
		return m.Warn50
	case "warn_500":
		return m.Warn500
	case "warn_600":
		return m.Warn600
	case "warn_700":
		return m.Warn700
	case "warn_800":
		return m.Warn800
	case "warn_900":
		return m.Warn900
	case "warn_a100":
		return m.WarnA100
	case "warn_a200":
		return m.WarnA200
	case "warn_a400":
		return m.WarnA400
	case "warn_a700":
		return m.WarnA700
	case "yes":
		return m.Yes
	}

	return nil
}

func (m *Theme) GetFqids(field string) []string {
	switch field {
	case "organization_id":
		return []string{"organization/" + strconv.Itoa(m.OrganizationID)}

	case "theme_for_organization_id":
		if m.ThemeForOrganizationID != nil {
			return []string{"organization/" + strconv.Itoa(*m.ThemeForOrganizationID)}
		}
	}
	return []string{}
}

func (m *Theme) Update(data map[string]string) error {
	if val, ok := data["abstain"]; ok {
		err := json.Unmarshal([]byte(val), &m.Abstain)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_100"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent100)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_200"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent200)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_300"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent300)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_400"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent400)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_50"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent50)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_500"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent500)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_600"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent600)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_700"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent700)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_800"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent800)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_900"]; ok {
		err := json.Unmarshal([]byte(val), &m.Accent900)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_a100"]; ok {
		err := json.Unmarshal([]byte(val), &m.AccentA100)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_a200"]; ok {
		err := json.Unmarshal([]byte(val), &m.AccentA200)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_a400"]; ok {
		err := json.Unmarshal([]byte(val), &m.AccentA400)
		if err != nil {
			return err
		}
	}

	if val, ok := data["accent_a700"]; ok {
		err := json.Unmarshal([]byte(val), &m.AccentA700)
		if err != nil {
			return err
		}
	}

	if val, ok := data["headbar"]; ok {
		err := json.Unmarshal([]byte(val), &m.Headbar)
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["no"]; ok {
		err := json.Unmarshal([]byte(val), &m.No)
		if err != nil {
			return err
		}
	}

	if val, ok := data["organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_100"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary100)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_200"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary200)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_300"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary300)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_400"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary400)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_50"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary50)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_500"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary500)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_600"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary600)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_700"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary700)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_800"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary800)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_900"]; ok {
		err := json.Unmarshal([]byte(val), &m.Primary900)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_a100"]; ok {
		err := json.Unmarshal([]byte(val), &m.PrimaryA100)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_a200"]; ok {
		err := json.Unmarshal([]byte(val), &m.PrimaryA200)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_a400"]; ok {
		err := json.Unmarshal([]byte(val), &m.PrimaryA400)
		if err != nil {
			return err
		}
	}

	if val, ok := data["primary_a700"]; ok {
		err := json.Unmarshal([]byte(val), &m.PrimaryA700)
		if err != nil {
			return err
		}
	}

	if val, ok := data["theme_for_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ThemeForOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_100"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn100)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_200"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn200)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_300"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn300)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_400"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn400)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_50"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn50)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_500"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn500)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_600"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn600)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_700"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn700)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_800"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn800)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_900"]; ok {
		err := json.Unmarshal([]byte(val), &m.Warn900)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_a100"]; ok {
		err := json.Unmarshal([]byte(val), &m.WarnA100)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_a200"]; ok {
		err := json.Unmarshal([]byte(val), &m.WarnA200)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_a400"]; ok {
		err := json.Unmarshal([]byte(val), &m.WarnA400)
		if err != nil {
			return err
		}
	}

	if val, ok := data["warn_a700"]; ok {
		err := json.Unmarshal([]byte(val), &m.WarnA700)
		if err != nil {
			return err
		}
	}

	if val, ok := data["yes"]; ok {
		err := json.Unmarshal([]byte(val), &m.Yes)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Theme) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

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

type User struct {
	CanChangeOwnPassword        *bool   `json:"can_change_own_password"`
	CommitteeIDs                []int   `json:"committee_ids"`
	CommitteeManagementIDs      []int   `json:"committee_management_ids"`
	DefaultPassword             *string `json:"default_password"`
	DefaultVoteWeight           *string `json:"default_vote_weight"`
	DelegatedVoteIDs            []int   `json:"delegated_vote_ids"`
	Email                       *string `json:"email"`
	FirstName                   *string `json:"first_name"`
	ForwardingCommitteeIDs      []int   `json:"forwarding_committee_ids"`
	GenderID                    *int    `json:"gender_id"`
	ID                          int     `json:"id"`
	IsActive                    *bool   `json:"is_active"`
	IsDemoUser                  *bool   `json:"is_demo_user"`
	IsPhysicalPerson            *bool   `json:"is_physical_person"`
	IsPresentInMeetingIDs       []int   `json:"is_present_in_meeting_ids"`
	LastEmailSent               *int    `json:"last_email_sent"`
	LastLogin                   *int    `json:"last_login"`
	LastName                    *string `json:"last_name"`
	MeetingIDs                  []int   `json:"meeting_ids"`
	MeetingUserIDs              []int   `json:"meeting_user_ids"`
	MemberNumber                *string `json:"member_number"`
	OptionIDs                   []int   `json:"option_ids"`
	OrganizationID              int     `json:"organization_id"`
	OrganizationManagementLevel *string `json:"organization_management_level"`
	Password                    *string `json:"password"`
	PollCandidateIDs            []int   `json:"poll_candidate_ids"`
	PollVotedIDs                []int   `json:"poll_voted_ids"`
	Pronoun                     *string `json:"pronoun"`
	SamlID                      *string `json:"saml_id"`
	Title                       *string `json:"title"`
	Username                    string  `json:"username"`
	VoteIDs                     []int   `json:"vote_ids"`
	loadedRelations             map[string]struct{}
	committeeManagements        []*Committee
	committees                  []*Committee
	delegatedVotes              []*Vote
	forwardingCommittees        []*Committee
	gender                      *Gender
	isPresentInMeetings         []*Meeting
	meetingUsers                []*MeetingUser
	options                     []*Option
	organization                *Organization
	pollCandidates              []*PollCandidate
	pollVoteds                  []*Poll
	votes                       []*Vote
}

func (m *User) CollectionName() string {
	return "user"
}

func (m *User) CommitteeManagements() []*Committee {
	if _, ok := m.loadedRelations["committee_management_ids"]; !ok {
		log.Panic().Msg("Tried to access CommitteeManagements relation of User which was not loaded.")
	}

	return m.committeeManagements
}

func (m *User) Committees() []*Committee {
	if _, ok := m.loadedRelations["committee_ids"]; !ok {
		log.Panic().Msg("Tried to access Committees relation of User which was not loaded.")
	}

	return m.committees
}

func (m *User) DelegatedVotes() []*Vote {
	if _, ok := m.loadedRelations["delegated_vote_ids"]; !ok {
		log.Panic().Msg("Tried to access DelegatedVotes relation of User which was not loaded.")
	}

	return m.delegatedVotes
}

func (m *User) ForwardingCommittees() []*Committee {
	if _, ok := m.loadedRelations["forwarding_committee_ids"]; !ok {
		log.Panic().Msg("Tried to access ForwardingCommittees relation of User which was not loaded.")
	}

	return m.forwardingCommittees
}

func (m *User) Gender() *Gender {
	if _, ok := m.loadedRelations["gender_id"]; !ok {
		log.Panic().Msg("Tried to access Gender relation of User which was not loaded.")
	}

	return m.gender
}

func (m *User) IsPresentInMeetings() []*Meeting {
	if _, ok := m.loadedRelations["is_present_in_meeting_ids"]; !ok {
		log.Panic().Msg("Tried to access IsPresentInMeetings relation of User which was not loaded.")
	}

	return m.isPresentInMeetings
}

func (m *User) MeetingUsers() []*MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingUsers relation of User which was not loaded.")
	}

	return m.meetingUsers
}

func (m *User) Options() []*Option {
	if _, ok := m.loadedRelations["option_ids"]; !ok {
		log.Panic().Msg("Tried to access Options relation of User which was not loaded.")
	}

	return m.options
}

func (m *User) Organization() Organization {
	if _, ok := m.loadedRelations["organization_id"]; !ok {
		log.Panic().Msg("Tried to access Organization relation of User which was not loaded.")
	}

	return *m.organization
}

func (m *User) PollCandidates() []*PollCandidate {
	if _, ok := m.loadedRelations["poll_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access PollCandidates relation of User which was not loaded.")
	}

	return m.pollCandidates
}

func (m *User) PollVoteds() []*Poll {
	if _, ok := m.loadedRelations["poll_voted_ids"]; !ok {
		log.Panic().Msg("Tried to access PollVoteds relation of User which was not loaded.")
	}

	return m.pollVoteds
}

func (m *User) Votes() []*Vote {
	if _, ok := m.loadedRelations["vote_ids"]; !ok {
		log.Panic().Msg("Tried to access Votes relation of User which was not loaded.")
	}

	return m.votes
}

func (m *User) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "committee_management_ids":
		for _, r := range m.committeeManagements {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "committee_ids":
		for _, r := range m.committees {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "delegated_vote_ids":
		for _, r := range m.delegatedVotes {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "forwarding_committee_ids":
		for _, r := range m.forwardingCommittees {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "gender_id":
		return m.gender.GetRelatedModelsAccessor()
	case "is_present_in_meeting_ids":
		for _, r := range m.isPresentInMeetings {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "meeting_user_ids":
		for _, r := range m.meetingUsers {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "option_ids":
		for _, r := range m.options {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "organization_id":
		return m.organization.GetRelatedModelsAccessor()
	case "poll_candidate_ids":
		for _, r := range m.pollCandidates {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "poll_voted_ids":
		for _, r := range m.pollVoteds {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	case "vote_ids":
		for _, r := range m.votes {
			if r.ID == id {
				return r.GetRelatedModelsAccessor()
			}
		}
	}

	return nil
}

func (m *User) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "committee_management_ids":
			m.committeeManagements = content.([]*Committee)
		case "committee_ids":
			m.committees = content.([]*Committee)
		case "delegated_vote_ids":
			m.delegatedVotes = content.([]*Vote)
		case "forwarding_committee_ids":
			m.forwardingCommittees = content.([]*Committee)
		case "gender_id":
			m.gender = content.(*Gender)
		case "is_present_in_meeting_ids":
			m.isPresentInMeetings = content.([]*Meeting)
		case "meeting_user_ids":
			m.meetingUsers = content.([]*MeetingUser)
		case "option_ids":
			m.options = content.([]*Option)
		case "organization_id":
			m.organization = content.(*Organization)
		case "poll_candidate_ids":
			m.pollCandidates = content.([]*PollCandidate)
		case "poll_voted_ids":
			m.pollVoteds = content.([]*Poll)
		case "vote_ids":
			m.votes = content.([]*Vote)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *User) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "committee_management_ids":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.committeeManagements = append(m.committeeManagements, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "committee_ids":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.committees = append(m.committees, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "delegated_vote_ids":
		var entry Vote
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.delegatedVotes = append(m.delegatedVotes, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "forwarding_committee_ids":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.forwardingCommittees = append(m.forwardingCommittees, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "gender_id":
		var entry Gender
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.gender = &entry

		result = entry.GetRelatedModelsAccessor()
	case "is_present_in_meeting_ids":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.isPresentInMeetings = append(m.isPresentInMeetings, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_user_ids":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUsers = append(m.meetingUsers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "option_ids":
		var entry Option
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.options = append(m.options, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.organization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "poll_candidate_ids":
		var entry PollCandidate
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCandidates = append(m.pollCandidates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_voted_ids":
		var entry Poll
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollVoteds = append(m.pollVoteds, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "vote_ids":
		var entry Vote
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.votes = append(m.votes, &entry)

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

func (m *User) Get(field string) interface{} {
	switch field {
	case "can_change_own_password":
		return m.CanChangeOwnPassword
	case "committee_ids":
		return m.CommitteeIDs
	case "committee_management_ids":
		return m.CommitteeManagementIDs
	case "default_password":
		return m.DefaultPassword
	case "default_vote_weight":
		return m.DefaultVoteWeight
	case "delegated_vote_ids":
		return m.DelegatedVoteIDs
	case "email":
		return m.Email
	case "first_name":
		return m.FirstName
	case "forwarding_committee_ids":
		return m.ForwardingCommitteeIDs
	case "gender_id":
		return m.GenderID
	case "id":
		return m.ID
	case "is_active":
		return m.IsActive
	case "is_demo_user":
		return m.IsDemoUser
	case "is_physical_person":
		return m.IsPhysicalPerson
	case "is_present_in_meeting_ids":
		return m.IsPresentInMeetingIDs
	case "last_email_sent":
		return m.LastEmailSent
	case "last_login":
		return m.LastLogin
	case "last_name":
		return m.LastName
	case "meeting_ids":
		return m.MeetingIDs
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "member_number":
		return m.MemberNumber
	case "option_ids":
		return m.OptionIDs
	case "organization_id":
		return m.OrganizationID
	case "organization_management_level":
		return m.OrganizationManagementLevel
	case "password":
		return m.Password
	case "poll_candidate_ids":
		return m.PollCandidateIDs
	case "poll_voted_ids":
		return m.PollVotedIDs
	case "pronoun":
		return m.Pronoun
	case "saml_id":
		return m.SamlID
	case "title":
		return m.Title
	case "username":
		return m.Username
	case "vote_ids":
		return m.VoteIDs
	}

	return nil
}

func (m *User) GetFqids(field string) []string {
	switch field {
	case "committee_management_ids":
		r := make([]string, len(m.CommitteeManagementIDs))
		for i, id := range m.CommitteeManagementIDs {
			r[i] = "committee/" + strconv.Itoa(id)
		}
		return r

	case "committee_ids":
		r := make([]string, len(m.CommitteeIDs))
		for i, id := range m.CommitteeIDs {
			r[i] = "committee/" + strconv.Itoa(id)
		}
		return r

	case "delegated_vote_ids":
		r := make([]string, len(m.DelegatedVoteIDs))
		for i, id := range m.DelegatedVoteIDs {
			r[i] = "vote/" + strconv.Itoa(id)
		}
		return r

	case "forwarding_committee_ids":
		r := make([]string, len(m.ForwardingCommitteeIDs))
		for i, id := range m.ForwardingCommitteeIDs {
			r[i] = "committee/" + strconv.Itoa(id)
		}
		return r

	case "gender_id":
		if m.GenderID != nil {
			return []string{"gender/" + strconv.Itoa(*m.GenderID)}
		}

	case "is_present_in_meeting_ids":
		r := make([]string, len(m.IsPresentInMeetingIDs))
		for i, id := range m.IsPresentInMeetingIDs {
			r[i] = "meeting/" + strconv.Itoa(id)
		}
		return r

	case "meeting_user_ids":
		r := make([]string, len(m.MeetingUserIDs))
		for i, id := range m.MeetingUserIDs {
			r[i] = "meeting_user/" + strconv.Itoa(id)
		}
		return r

	case "option_ids":
		r := make([]string, len(m.OptionIDs))
		for i, id := range m.OptionIDs {
			r[i] = "option/" + strconv.Itoa(id)
		}
		return r

	case "organization_id":
		return []string{"organization/" + strconv.Itoa(m.OrganizationID)}

	case "poll_candidate_ids":
		r := make([]string, len(m.PollCandidateIDs))
		for i, id := range m.PollCandidateIDs {
			r[i] = "poll_candidate/" + strconv.Itoa(id)
		}
		return r

	case "poll_voted_ids":
		r := make([]string, len(m.PollVotedIDs))
		for i, id := range m.PollVotedIDs {
			r[i] = "poll/" + strconv.Itoa(id)
		}
		return r

	case "vote_ids":
		r := make([]string, len(m.VoteIDs))
		for i, id := range m.VoteIDs {
			r[i] = "vote/" + strconv.Itoa(id)
		}
		return r
	}
	return []string{}
}

func (m *User) Update(data map[string]string) error {
	if val, ok := data["can_change_own_password"]; ok {
		err := json.Unmarshal([]byte(val), &m.CanChangeOwnPassword)
		if err != nil {
			return err
		}
	}

	if val, ok := data["committee_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommitteeIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["committee_ids"]; ok {
			m.committees = slices.DeleteFunc(m.committees, func(r *Committee) bool {
				return !slices.Contains(m.CommitteeIDs, r.ID)
			})
		}
	}

	if val, ok := data["committee_management_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommitteeManagementIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["committee_management_ids"]; ok {
			m.committeeManagements = slices.DeleteFunc(m.committeeManagements, func(r *Committee) bool {
				return !slices.Contains(m.CommitteeManagementIDs, r.ID)
			})
		}
	}

	if val, ok := data["default_password"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultPassword)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_vote_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultVoteWeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["delegated_vote_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DelegatedVoteIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["delegated_vote_ids"]; ok {
			m.delegatedVotes = slices.DeleteFunc(m.delegatedVotes, func(r *Vote) bool {
				return !slices.Contains(m.DelegatedVoteIDs, r.ID)
			})
		}
	}

	if val, ok := data["email"]; ok {
		err := json.Unmarshal([]byte(val), &m.Email)
		if err != nil {
			return err
		}
	}

	if val, ok := data["first_name"]; ok {
		err := json.Unmarshal([]byte(val), &m.FirstName)
		if err != nil {
			return err
		}
	}

	if val, ok := data["forwarding_committee_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ForwardingCommitteeIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["forwarding_committee_ids"]; ok {
			m.forwardingCommittees = slices.DeleteFunc(m.forwardingCommittees, func(r *Committee) bool {
				return !slices.Contains(m.ForwardingCommitteeIDs, r.ID)
			})
		}
	}

	if val, ok := data["gender_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.GenderID)
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

	if val, ok := data["is_active"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsActive)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_demo_user"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsDemoUser)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_physical_person"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsPhysicalPerson)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_present_in_meeting_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsPresentInMeetingIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["is_present_in_meeting_ids"]; ok {
			m.isPresentInMeetings = slices.DeleteFunc(m.isPresentInMeetings, func(r *Meeting) bool {
				return !slices.Contains(m.IsPresentInMeetingIDs, r.ID)
			})
		}
	}

	if val, ok := data["last_email_sent"]; ok {
		err := json.Unmarshal([]byte(val), &m.LastEmailSent)
		if err != nil {
			return err
		}
	}

	if val, ok := data["last_login"]; ok {
		err := json.Unmarshal([]byte(val), &m.LastLogin)
		if err != nil {
			return err
		}
	}

	if val, ok := data["last_name"]; ok {
		err := json.Unmarshal([]byte(val), &m.LastName)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingIDs)
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

	if val, ok := data["member_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.MemberNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["option_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["option_ids"]; ok {
			m.options = slices.DeleteFunc(m.options, func(r *Option) bool {
				return !slices.Contains(m.OptionIDs, r.ID)
			})
		}
	}

	if val, ok := data["organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["organization_management_level"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationManagementLevel)
		if err != nil {
			return err
		}
	}

	if val, ok := data["password"]; ok {
		err := json.Unmarshal([]byte(val), &m.Password)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["poll_candidate_ids"]; ok {
			m.pollCandidates = slices.DeleteFunc(m.pollCandidates, func(r *PollCandidate) bool {
				return !slices.Contains(m.PollCandidateIDs, r.ID)
			})
		}
	}

	if val, ok := data["poll_voted_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollVotedIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["poll_voted_ids"]; ok {
			m.pollVoteds = slices.DeleteFunc(m.pollVoteds, func(r *Poll) bool {
				return !slices.Contains(m.PollVotedIDs, r.ID)
			})
		}
	}

	if val, ok := data["pronoun"]; ok {
		err := json.Unmarshal([]byte(val), &m.Pronoun)
		if err != nil {
			return err
		}
	}

	if val, ok := data["saml_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.SamlID)
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

	if val, ok := data["username"]; ok {
		err := json.Unmarshal([]byte(val), &m.Username)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteIDs)
		if err != nil {
			return err
		}

		if _, ok := m.loadedRelations["vote_ids"]; ok {
			m.votes = slices.DeleteFunc(m.votes, func(r *Vote) bool {
				return !slices.Contains(m.VoteIDs, r.ID)
			})
		}
	}

	return nil
}

func (m *User) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

type Vote struct {
	DelegatedUserID *int    `json:"delegated_user_id"`
	ID              int     `json:"id"`
	MeetingID       int     `json:"meeting_id"`
	OptionID        int     `json:"option_id"`
	UserID          *int    `json:"user_id"`
	UserToken       string  `json:"user_token"`
	Value           *string `json:"value"`
	Weight          *string `json:"weight"`
	loadedRelations map[string]struct{}
	delegatedUser   *User
	meeting         *Meeting
	option          *Option
	user            *User
}

func (m *Vote) CollectionName() string {
	return "vote"
}

func (m *Vote) DelegatedUser() *User {
	if _, ok := m.loadedRelations["delegated_user_id"]; !ok {
		log.Panic().Msg("Tried to access DelegatedUser relation of Vote which was not loaded.")
	}

	return m.delegatedUser
}

func (m *Vote) Meeting() Meeting {
	if _, ok := m.loadedRelations["meeting_id"]; !ok {
		log.Panic().Msg("Tried to access Meeting relation of Vote which was not loaded.")
	}

	return *m.meeting
}

func (m *Vote) Option() Option {
	if _, ok := m.loadedRelations["option_id"]; !ok {
		log.Panic().Msg("Tried to access Option relation of Vote which was not loaded.")
	}

	return *m.option
}

func (m *Vote) User() *User {
	if _, ok := m.loadedRelations["user_id"]; !ok {
		log.Panic().Msg("Tried to access User relation of Vote which was not loaded.")
	}

	return m.user
}

func (m *Vote) GetRelated(field string, id int) *RelatedModelsAccessor {
	switch field {
	case "delegated_user_id":
		return m.delegatedUser.GetRelatedModelsAccessor()
	case "meeting_id":
		return m.meeting.GetRelatedModelsAccessor()
	case "option_id":
		return m.option.GetRelatedModelsAccessor()
	case "user_id":
		return m.user.GetRelatedModelsAccessor()
	}

	return nil
}

func (m *Vote) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "delegated_user_id":
			m.delegatedUser = content.(*User)
		case "meeting_id":
			m.meeting = content.(*Meeting)
		case "option_id":
			m.option = content.(*Option)
		case "user_id":
			m.user = content.(*User)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Vote) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "delegated_user_id":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.delegatedUser = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_id":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meeting = &entry

		result = entry.GetRelatedModelsAccessor()
	case "option_id":
		var entry Option
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.option = &entry

		result = entry.GetRelatedModelsAccessor()
	case "user_id":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.user = &entry

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

func (m *Vote) Get(field string) interface{} {
	switch field {
	case "delegated_user_id":
		return m.DelegatedUserID
	case "id":
		return m.ID
	case "meeting_id":
		return m.MeetingID
	case "option_id":
		return m.OptionID
	case "user_id":
		return m.UserID
	case "user_token":
		return m.UserToken
	case "value":
		return m.Value
	case "weight":
		return m.Weight
	}

	return nil
}

func (m *Vote) GetFqids(field string) []string {
	switch field {
	case "delegated_user_id":
		if m.DelegatedUserID != nil {
			return []string{"user/" + strconv.Itoa(*m.DelegatedUserID)}
		}

	case "meeting_id":
		return []string{"meeting/" + strconv.Itoa(m.MeetingID)}

	case "option_id":
		return []string{"option/" + strconv.Itoa(m.OptionID)}

	case "user_id":
		if m.UserID != nil {
			return []string{"user/" + strconv.Itoa(*m.UserID)}
		}
	}
	return []string{}
}

func (m *Vote) Update(data map[string]string) error {
	if val, ok := data["delegated_user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DelegatedUserID)
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

	if val, ok := data["option_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_token"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserToken)
		if err != nil {
			return err
		}
	}

	if val, ok := data["value"]; ok {
		err := json.Unmarshal([]byte(val), &m.Value)
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

func (m *Vote) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

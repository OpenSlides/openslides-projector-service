package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

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

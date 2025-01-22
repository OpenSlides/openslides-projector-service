package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

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

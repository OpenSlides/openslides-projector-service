package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

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

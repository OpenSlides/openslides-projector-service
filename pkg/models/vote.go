package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

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
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

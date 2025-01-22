package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

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
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

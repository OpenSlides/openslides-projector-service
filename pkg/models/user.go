package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

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

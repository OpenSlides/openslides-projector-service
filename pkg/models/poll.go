package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

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

func (m *Poll) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
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
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

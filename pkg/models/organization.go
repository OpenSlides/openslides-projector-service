package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

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
	mediafiles                 []*Mediafile
	organizationTags           []*OrganizationTag
	publishedMediafiles        []*Mediafile
	themes                     []*Theme
	theme                      *Theme
	archivedMeetings           []*Meeting
	committees                 []*Committee
	genders                    []*Gender
	templateMeetings           []*Meeting
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

func (m *Organization) Themes() []*Theme {
	if _, ok := m.loadedRelations["theme_ids"]; !ok {
		log.Panic().Msg("Tried to access Themes relation of Organization which was not loaded.")
	}

	return m.themes
}

func (m *Organization) Theme() Theme {
	if _, ok := m.loadedRelations["theme_id"]; !ok {
		log.Panic().Msg("Tried to access Theme relation of Organization which was not loaded.")
	}

	return *m.theme
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

func (m *Organization) TemplateMeetings() []*Meeting {
	if _, ok := m.loadedRelations["template_meeting_ids"]; !ok {
		log.Panic().Msg("Tried to access TemplateMeetings relation of Organization which was not loaded.")
	}

	return m.templateMeetings
}

func (m *Organization) Users() []*User {
	if _, ok := m.loadedRelations["user_ids"]; !ok {
		log.Panic().Msg("Tried to access Users relation of Organization which was not loaded.")
	}

	return m.users
}

func (m *Organization) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "active_meeting_ids":
			m.activeMeetings = content.([]*Meeting)
		case "mediafile_ids":
			m.mediafiles = content.([]*Mediafile)
		case "organization_tag_ids":
			m.organizationTags = content.([]*OrganizationTag)
		case "published_mediafile_ids":
			m.publishedMediafiles = content.([]*Mediafile)
		case "theme_ids":
			m.themes = content.([]*Theme)
		case "theme_id":
			m.theme = content.(*Theme)
		case "archived_meeting_ids":
			m.archivedMeetings = content.([]*Meeting)
		case "committee_ids":
			m.committees = content.([]*Committee)
		case "gender_ids":
			m.genders = content.([]*Gender)
		case "template_meeting_ids":
			m.templateMeetings = content.([]*Meeting)
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
	case "theme_ids":
		var entry Theme
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.themes = append(m.themes, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "theme_id":
		var entry Theme
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.theme = &entry

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
	case "template_meeting_ids":
		var entry Meeting
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.templateMeetings = append(m.templateMeetings, &entry)

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

	case "theme_ids":
		r := make([]string, len(m.ThemeIDs))
		for i, id := range m.ThemeIDs {
			r[i] = "theme/" + strconv.Itoa(id)
		}
		return r

	case "theme_id":
		return []string{"theme/" + strconv.Itoa(m.ThemeID)}

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

	case "template_meeting_ids":
		r := make([]string, len(m.TemplateMeetingIDs))
		for i, id := range m.TemplateMeetingIDs {
			r[i] = "meeting/" + strconv.Itoa(id)
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
	}

	if val, ok := data["archived_meeting_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ArchivedMeetingIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["committee_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommitteeIDs)
		if err != nil {
			return err
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
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

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
		m.SetRelated,
		m.SetRelatedJSON,
	}
}

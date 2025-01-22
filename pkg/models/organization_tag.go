package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

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

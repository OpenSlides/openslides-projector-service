package models

import (
	"encoding/json"
)

type ImportPreview struct {
	Created int             `json:"created"`
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Result  json.RawMessage `json:"result"`
	State   string          `json:"state"`
}

func (m *ImportPreview) CollectionName() string {
	return "import_preview"
}

func (m *ImportPreview) GetRelated(field string, id int) *RelatedModelsAccessor {
	return nil
}

func (m *ImportPreview) SetRelated(field string, content interface{}) {}

func (m *ImportPreview) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	return nil, nil
}

func (m *ImportPreview) Get(field string) interface{} {
	switch field {
	case "created":
		return m.Created
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "result":
		return m.Result
	case "state":
		return m.State
	}

	return nil
}

func (m *ImportPreview) GetFqids(field string) []string {
	return []string{}
}

func (m *ImportPreview) Update(data map[string]string) error {
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

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["result"]; ok {
		err := json.Unmarshal([]byte(val), &m.Result)
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

	return nil
}

func (m *ImportPreview) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

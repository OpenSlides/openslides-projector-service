package models

import (
	"encoding/json"
)

type ActionWorker struct {
	Created   int             `json:"created"`
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Result    json.RawMessage `json:"result"`
	State     string          `json:"state"`
	Timestamp int             `json:"timestamp"`
	UserID    int             `json:"user_id"`
}

func (m *ActionWorker) CollectionName() string {
	return "action_worker"
}

func (m *ActionWorker) GetRelated(field string, id int) *RelatedModelsAccessor {
	return nil
}

func (m *ActionWorker) SetRelated(field string, content interface{}) {}

func (m *ActionWorker) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	return nil, nil
}

func (m *ActionWorker) Get(field string) interface{} {
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
	case "timestamp":
		return m.Timestamp
	case "user_id":
		return m.UserID
	}

	return nil
}

func (m *ActionWorker) GetFqids(field string) []string {
	return []string{}
}

func (m *ActionWorker) Update(data map[string]string) error {
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

	if val, ok := data["timestamp"]; ok {
		err := json.Unmarshal([]byte(val), &m.Timestamp)
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

	return nil
}

func (m *ActionWorker) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.GetRelated,
		m.SetRelated,
		m.SetRelatedJSON,
		m.Update,
	}
}

package models

type IBaseModel interface {
	CollectionName() string
	Get(field string) interface{}
	GetFqids(field string) []string
	Update(data map[string]string) error
	SetRelated(string, interface{})
	SetRelatedJSON(string, []byte) (*RelatedModelsAccessor, error)
	GetRelatedModelsAccessor() *RelatedModelsAccessor
}

type RelatedModelsAccessor struct {
	GetFqids       func(field string) []string
	SetRelated     func(string, interface{})
	SetRelatedJSON func(string, []byte) (*RelatedModelsAccessor, error)
}

package api_key

type Store interface {
	Create(obj *ApiKey) (*ApiKey, error)
	Get(id string) (*ApiKey, error)
	GetByKey(key string) (*ApiKey, error)
	List() ([]ApiKey, error)
}

package api_key

type Service interface {
	Create(cmd *CreateApiKeyCommand) (*ApiKey, error)
	Get(cmd *GetApiKeyCommand) (*ApiKey, error)
	GetByKey(cmd *GetByApiKeyCommand) (*ApiKey, error)
	List(cmd *ListApiKeyCommand) ([]ApiKey, error)
}

type service struct {
	store Store
}

func NewService(s Store) Service {
	return &service{store: s}
}

func (s *service) Create(cmd *CreateApiKeyCommand) (*ApiKey, error) {
	return s.store.Create(&ApiKey{
		Name: cmd.Name,
	})
}

func (s *service) Get(cmd *GetApiKeyCommand) (*ApiKey, error) {
	return s.store.Get(cmd.Id)
}

func (s *service) GetByKey(cmd *GetByApiKeyCommand) (*ApiKey, error) {
	return s.store.GetByKey(cmd.Key)
}

func (s *service) List(cmd *ListApiKeyCommand) ([]ApiKey, error) {
	return s.store.List()
}

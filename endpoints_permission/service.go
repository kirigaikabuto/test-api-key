package endpoints_permission

type service struct {
	store Store
}

type Service interface {
	Create(cmd *CreateEndpointsPermissionCommand) (*EndpointsPermission, error)
	GetById(cmd *GetByIdEndpointsPermissionCommand) (*EndpointsPermission, error)
	GetByApiKeyId(cmd *GetByApiKeyIdEndpointsPermissionsCommand) (*EndpointsPermission, error)
	List(cmd *ListEndpointsPermissionsCommand) ([]EndpointsPermission, error)
	AddEndpointToEndpointsPermissions(cmd *AddEndpointToEndpointsPermissionsCommand) (*EndpointsPermission, error)
}

func NewService(s Store) Service {
	return &service{store: s}
}

func (s *service) Create(cmd *CreateEndpointsPermissionCommand) (*EndpointsPermission, error) {
	return s.store.Create(&EndpointsPermission{
		ApiKeyId:  cmd.ApiKeyId,
		Endpoints: cmd.Endpoints,
	})
}

func (s *service) GetById(cmd *GetByIdEndpointsPermissionCommand) (*EndpointsPermission, error) {
	return s.store.Get(cmd.Id)
}

func (s *service) GetByApiKeyId(cmd *GetByApiKeyIdEndpointsPermissionsCommand) (*EndpointsPermission, error) {
	return s.store.GetByApiKeyId(cmd.ApiKeyId)
}

func (s *service) List(cmd *ListEndpointsPermissionsCommand) ([]EndpointsPermission, error) {
	return s.store.List()
}

func (s *service) AddEndpointToEndpointsPermissions(cmd *AddEndpointToEndpointsPermissionsCommand) (*EndpointsPermission, error) {
	endpointsPerms, err := s.GetById(&GetByIdEndpointsPermissionCommand{Id: cmd.Id})
	if err != nil {
		return nil, err
	}
	endpointsPerms.Endpoints = append(endpointsPerms.Endpoints, cmd.Endpoint)
	return s.store.Update(&EndpointsPermissionUpdate{
		Id:        cmd.Id,
		Endpoints: &endpointsPerms.Endpoints,
	})
}

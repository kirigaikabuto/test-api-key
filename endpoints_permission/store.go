package endpoints_permission

type Store interface {
	Create(obj *EndpointsPermission) (*EndpointsPermission, error)
	Get(id string) (*EndpointsPermission, error)
	GetByApiKeyId(keyId string) (*EndpointsPermission, error)
	List() ([]EndpointsPermission, error)
	Update(obj *EndpointsPermissionUpdate) (*EndpointsPermission, error)
}

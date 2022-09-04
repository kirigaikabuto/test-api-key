package endpoints_permission

type CreateEndpointsPermissionCommand struct {
	ApiKeyId  string   `json:"api_key_id"`
	Endpoints []string `json:"endpoints"`
}

func (cmd *CreateEndpointsPermissionCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).Create(cmd)
}

type GetByIdEndpointsPermissionCommand struct {
	Id string `json:"id"`
}

func (cmd *GetByIdEndpointsPermissionCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).GetById(cmd)
}

type GetByApiKeyIdEndpointsPermissionsCommand struct {
	ApiKeyId string `json:"api_key_id"`
}

func (cmd *GetByApiKeyIdEndpointsPermissionsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).GetByApiKeyId(cmd)
}

type ListEndpointsPermissionsCommand struct {
}

func (cmd *ListEndpointsPermissionsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).List(cmd)
}

type AddEndpointToEndpointsPermissionsCommand struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
}

func (cmd *AddEndpointToEndpointsPermissionsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).AddEndpointToEndpointsPermissions(cmd)
}

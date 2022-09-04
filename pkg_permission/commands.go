package pkg_permission

type CreatePackagePermissionCommand struct {
	AccessZone string `json:"access_zone"`
	ApiKeyId   string `json:"api_key_id"`
	Action     string `json:"action"`
}

func (cmd *CreatePackagePermissionCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).CreatePackagePermission(cmd)
}

type GetPackagePermissionsByApiKeyIdCommand struct {
	ApiKeyId string `json:"api_key_id"`
}

func (cmd *GetPackagePermissionsByApiKeyIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).GetPackagePermissionByApiKeyId(cmd)
}

type ListPackagePermissionsCommand struct {
}

func (cmd *ListPackagePermissionsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).ListPackagePermissions(cmd)
}

type GetPackagePermissionByIdCommand struct {
	Id string `json:"id"`
}

func (cmd *GetPackagePermissionByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).GetPackagePermissionById(cmd)
}

type UpdatePackagePermissionCommand struct {
	*UpdatePackagePermission
}

func (cmd *UpdatePackagePermissionCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).UpdatePackagePermission(cmd)
}

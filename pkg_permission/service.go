package pkg_permission

type Service interface {
	CreatePackagePermission(cmd *CreatePackagePermissionCommand) (*PackagePermission, error)
	GetPackagePermissionByApiKeyId(cmd *GetPackagePermissionsByApiKeyIdCommand) ([]PackagePermission, error)
	ListPackagePermissions(cmd *ListPackagePermissionsCommand) ([]PackagePermission, error)
	GetPackagePermissionById(cmd *GetPackagePermissionByIdCommand) (*PackagePermission, error)
	UpdatePackagePermission(cmd *UpdatePackagePermissionCommand) (*PackagePermission, error)
}

type service struct {
	store Store
}

func (s *service) CreatePackagePermission(cmd *CreatePackagePermissionCommand) (*PackagePermission, error) {
	return s.store.Create(&PackagePermission{
		AccessZone: cmd.AccessZone,
		ApiKeyId:   cmd.ApiKeyId,
		Action:     cmd.Action,
	})
}

func (s *service) GetPackagePermissionByApiKeyId(cmd *GetPackagePermissionsByApiKeyIdCommand) ([]PackagePermission, error) {
	return s.store.GetByApiKeyId(cmd.ApiKeyId)
}

func (s *service) ListPackagePermissions(cmd *ListPackagePermissionsCommand) ([]PackagePermission, error) {
	return s.store.List()
}

func (s *service) GetPackagePermissionById(cmd *GetPackagePermissionByIdCommand) (*PackagePermission, error) {
	return s.store.GetById(cmd.Id)
}

func (s *service) UpdatePackagePermission(cmd *UpdatePackagePermission) (*PackagePermission, error) {
	return s.store.Update(&UpdatePackagePermission{
		Id:         cmd.Id,
		AccessZone: cmd.AccessZone,
		ApiKeyId:   cmd.ApiKeyId,
		Action:     cmd.Action,
	})
}

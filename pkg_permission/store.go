package pkg_permission

type Store interface {
	Create(obj *PackagePermission) (*PackagePermission, error)
	GetByApiKeyId(apiKeyId string) ([]PackagePermission, error)
	List() ([]PackagePermission, error)
	GetById(id string) (*PackagePermission, error)
	Update(obj *UpdatePackagePermission) (*PackagePermission, error)
}

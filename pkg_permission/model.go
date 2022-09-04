package pkg_permission

type PackagePermission struct {
	Id         string `json:"id"`
	AccessZone string `json:"access_zone"`
	ApiKeyId   string `json:"api_key_id"`
	Action     string `json:"action"`
}

type UpdatePackagePermission struct {
	Id         string  `json:"id"`
	AccessZone *string `json:"access_zone"`
	ApiKeyId   *string `json:"api_key_id"`
	Action     *string `json:"action"`
}

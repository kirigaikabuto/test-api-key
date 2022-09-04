package endpoints_permission

type EndpointsPermission struct {
	Id        string   `json:"id"`
	ApiKeyId  string   `json:"api_key_id"`
	Endpoints []string `json:"endpoints"`
}

type EndpointsPermissionUpdate struct {
	Id        string    `json:"id"`
	ApiKeyId  *string   `json:"api_key_id"`
	Endpoints *[]string `json:"endpoints"`
}

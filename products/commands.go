package products

type ListProductsCommand struct {
	ApiKeyId string `json:"-"`
}

func (cmd *ListProductsCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).ListProducts(cmd)
}

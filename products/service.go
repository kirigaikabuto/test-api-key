package products

import "github.com/kirigaikabuto/test-api-key/endpoints_permission"

type service struct {
	endpointsPermissionStore endpoints_permission.Store
}

type Service interface {
	ListProducts(cmd *ListProductsCommand) ([]Product, error)
}

func NewService(endPermStore endpoints_permission.Store) Service {
	return &service{endpointsPermissionStore: endPermStore}
}

func (s *service) ListProducts(cmd *ListProductsCommand) ([]Product, error) {
	blocked, err := s.EndpointIsBlocked("MakeListProducts", cmd.ApiKeyId)
	if err != nil {
		return nil, err
	}
	if blocked {
		return nil, ErrEndpointIsBlocked
	}
	var products []Product
	products = append(products, Product{
		Id:    "1",
		Name:  "product1",
		Price: 3.4,
	},
		Product{
			Id:    "2",
			Name:  "product2",
			Price: 3.5,
		},
	)
	return products, nil
}

func (s *service) EndpointIsBlocked(endpoint, apiKeyId string) (bool, error) {
	perms, err := s.endpointsPermissionStore.GetByApiKeyId(apiKeyId)
	if err != nil {
		return false, err
	}
	for _, v := range perms.Endpoints {
		if v == endpoint {
			return true, nil
		}
	}
	return false, nil
}

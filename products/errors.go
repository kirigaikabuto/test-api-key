package products

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrNotApiKeyId       = com.NewMiddleError(errors.New("no api key id"), 500, 100)
	ErrEndpointIsBlocked = com.NewMiddleError(errors.New("endpoint is blocked"), 500, 101)
)

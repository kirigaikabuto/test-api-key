package endpoints_permission

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrCreateEndpointsPermissionUnknown = com.NewMiddleError(errors.New("could not create endpoints permissions: unknown error"), 500, 3)
	ErrEndpointsPermissionNotFound      = com.NewMiddleError(errors.New("endpoints permission not found"), 500, 4)
	ErrNothingToUpdate                  = com.NewMiddleError(errors.New("nothing to update"), 400, 5)
)

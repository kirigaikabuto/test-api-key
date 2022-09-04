package pkg_permission

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrCreatePkgPermissionUnknown = com.NewMiddleError(errors.New("could not create package permissions: unknown error"), 500, 7)
	ErrPkgPermissionNotFound      = com.NewMiddleError(errors.New("package permission not found"), 500, 8)
	ErrNothingToUpdate            = com.NewMiddleError(errors.New("nothing to update"), 400, 9)
)

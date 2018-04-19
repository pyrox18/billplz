package billplz

import (
	"errors"
)

var (
	ErrCollectionNotFound         = errors.New("billplz: queried collection cannot be found")
	ErrCannotDeactivateCollection = errors.New("billplz: collection cannot be deactivated")
	ErrCannotActivateCollection   = errors.New("billplz: collection cannot be activated")
	ErrAdminPrivilegeRequired     = errors.New("billplz: admin privilege required")
	ErrUnauthorized               = errors.New("billplz: invalid API authorization key")
	ErrBillNotFound               = errors.New("billplz: bill not found")
)

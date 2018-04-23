package billplz

import (
	"errors"
)

var (
	// ErrCollectionNotFound is returned by Client.GetCollection and Client.GetOpenCollection
	// if the queried collection is not found.
	ErrCollectionNotFound = errors.New("billplz: queried collection cannot be found")

	// ErrCannotDeactivateCollection is returned by Client.DeactivateCollection if a collection
	// cannot be deactivated.
	ErrCannotDeactivateCollection = errors.New("billplz: collection cannot be deactivated")

	// ErrCannotActivateCollection is returned by Client.ActivateCollection if a collection
	// cannot be activated.
	ErrCannotActivateCollection = errors.New("billplz: collection cannot be activated")

	// ErrAdminPrivilegeRequired is returned if the request requires the 'ADMIN' setting to be
	// enabled in the user's Billplz account.
	ErrAdminPrivilegeRequired = errors.New("billplz: admin privilege required")

	// ErrUnauthorized is returned if an invalid API key is provided for authentication.
	ErrUnauthorized = errors.New("billplz: invalid API authorization key")

	// ErrBillNotFound is returned by Client.GetBill and Client.DeleteBill if a bill with the given
	// ID is not found.
	ErrBillNotFound = errors.New("billplz: bill not found")

	// ErrBankAccountNotFound is returned by Client.CheckRegistration if a bank
	// account with the given account number is not found.
	ErrBankAccountNotFound = errors.New("billplz: bank account not found")
)

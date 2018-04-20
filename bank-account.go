package billplz

import (
	"github.com/go-ozzo/ozzo-validation"
)

// BankAccount represents a bank account stored in the Billplz API.
type BankAccount struct {
	Name              string `json:"name,omitempty"`
	IDNumber          string `json:"id_no,omitempty"`
	AccountNumber     string `json:"acc_no,omitempty"`
	Code              string `json:"code,omitempty"`
	Organization      bool   `json:"organization,omitempty"`
	AuthorizationDate string `json:"authorization_date,omitempty"`
	Status            string `json:"status,omitempty"`
	ProcessedAt       string `json:"processed_at,omitempty"`
	RejectDescription string `json:"reject_desc,omitempty"`
}

// Validate validates a bank account's details for submission with Client.CreateBankAccount.
func (b *BankAccount) Validate() error {
	err := validation.Errors{
		"name":   validation.Validate(b.Name, validation.Required),
		"id_no":  validation.Validate(b.IDNumber, validation.Required),
		"acc_no": validation.Validate(b.AccountNumber, validation.Required),
		"code":   validation.Validate(b.Code, validation.Required),
	}.Filter()
	return err
}

// BankAccountCheckResponse represents the structure of the response body obtained with
// Client.CheckRegistration.
type BankAccountCheckResponse struct {
	Name string `json:"name,omitempty"`
}

// BankAccountList represents the structure of the response body obtained with Client.GetBankAccountIndex.
type BankAccountList struct {
	BankAccounts *[]BankAccount `json:"bank_verification_services,omitempty"`
}

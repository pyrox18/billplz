package billplz

import (
	"github.com/go-ozzo/ozzo-validation"
)

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

func (b *BankAccount) Validate() error {
	err := validation.Errors{
		"name":   validation.Validate(b.Name, validation.Required),
		"id_no":  validation.Validate(b.IDNumber, validation.Required),
		"acc_no": validation.Validate(b.AccountNumber, validation.Required),
		"code":   validation.Validate(b.Code, validation.Required),
	}.Filter()
	return err
}

type BankAccountCheckResponse struct {
	Name string `json:"name,omitempty"`
}

type BankAccountList struct {
	BankAccounts *[]BankAccount `json:"bank_verification_services,omitempty"`
}

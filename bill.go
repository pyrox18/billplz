package billplz

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Bill represents a bill contained within a collection.
type Bill struct {
	ID              string `json:"id,omitempty"`
	CollectionID    string `json:"collection_id,omitempty"`
	Paid            bool   `json:"paid,omitempty"`
	State           string `json:"state,omitempty"`
	Amount          uint   `json:"amount,omitempty"`
	PaidAmount      uint   `json:"paid_amount,omitempty"`
	DueAt           string `json:"due_at,omitempty"`
	Email           string `json:"email,omitempty"`
	Mobile          string `json:"mobile,omitempty"`
	Name            string `json:"name,omitempty"`
	URL             string `json:"url,omitempty"`
	Reference1Label string `json:"reference_1_label,omitempty"`
	Reference1      string `json:"reference_1,omitempty"`
	Reference2Label string `json:"reference_2_label,omitempty"`
	Reference2      string `json:"reference_2,omitempty"`
	Deliver         bool   `json:"deliver,omitempty"`
	RedirectURL     string `json:"redirect_url,omitempty"`
	CallbackURL     string `json:"callback_url,omitempty"`
	Description     string `json:"description,omitempty"`
}

// Validate validates a bill's details for submission with Client.CreateBill.
func (b *Bill) Validate() error {
	err := validation.Errors{
		"collection_id":     validation.Validate(b.CollectionID, validation.Required),
		"email":             validation.Validate(b.Email, is.Email),
		"mobile":            validation.Validate(b.Mobile, validation.Match(regexp.MustCompile(`\+?60\d{8,10}`))),
		"name":              validation.Validate(b.Name, validation.Required),
		"amount":            validation.Validate(b.Amount, validation.Required),
		"callback_url":      validation.Validate(b.CallbackURL, validation.Required, is.URL),
		"description":       validation.Validate(b.Description, validation.Required, validation.Length(1, 200)),
		"due_at":            validation.Validate(b.DueAt, validation.Date("2018-04-19")),
		"redirect_url":      validation.Validate(b.RedirectURL, is.URL),
		"reference_1_label": validation.Validate(b.Reference1Label, validation.Length(0, 20)),
		"reference_1":       validation.Validate(b.Reference1, validation.Length(0, 120)),
		"reference_2_label": validation.Validate(b.Reference2Label, validation.Length(0, 20)),
		"reference_2":       validation.Validate(b.Reference2, validation.Length(0, 120)),
	}.Filter()
	return err
}

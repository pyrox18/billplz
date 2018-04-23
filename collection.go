package billplz

import (
	"encoding/json"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Collection represents a set that contains many bills.
type Collection struct {
	ID           string        `json:"id,omitempty"`
	Title        string        `json:"title,omitempty"`
	Logo         *Logo         `json:"logo,omitempty"`
	SplitPayment *SplitPayment `json:"split_payment,omitempty"`
	Status       string        `json:"status,omitempty"`
}

func (c *Collection) validate() error {
	err := validation.Errors{
		"title":         validation.Validate(c.Title, validation.Required),
		"split_payment": c.SplitPayment.validate(),
	}.Filter()
	return err
}

// CollectionIndexResult represents the structure of the response body obtained with Client.GetCollectionIndex.
type CollectionIndexResult struct {
	Collections *[]Collection `json:"collections,omitempty"`
	Page        json.Number   `json:"page,omitempty"`
}

// OpenCollection represents a one-off payment form.
type OpenCollection struct {
	ID              string        `json:"id,omitempty"`
	Title           string        `json:"title,omitempty"`
	Description     string        `json:"description,omitempty"`
	Reference1Label string        `json:"reference_1_label,omitempty"`
	Reference2Label string        `json:"reference_2_label,omitempty"`
	EmailLink       string        `json:"email_link,omitempty"`
	Amount          uint          `json:"amount,omitempty"`
	FixedAmount     bool          `json:"fixed_amount,omitempty"`
	Tax             uint          `json:"tax,omitempty"`
	FixedQuantity   bool          `json:"fixed_quantity,omitempty"`
	PaymentButton   string        `json:"payment_button,omitempty"`
	Photo           *Photo        `json:"photo,omitempty"`
	SplitPayment    *SplitPayment `json:"split_payment,omitempty"`
	URL             string        `json:"url,omitempty"`
	Status          string        `json:"status,omitempty"`
}

func (o *OpenCollection) validate() error {
	err := validation.Errors{
		"title":             validation.Validate(o.Title, validation.Required, validation.Length(1, 50)),
		"description":       validation.Validate(o.Description, validation.Required, validation.Length(1, 200)),
		"reference_1_label": validation.Validate(o.Reference1Label, validation.Length(0, 20)),
		"reference_2_label": validation.Validate(o.Reference2Label, validation.Length(0, 20)),
		"email_link":        validation.Validate(o.EmailLink, is.Email),
		"payment_button":    validation.Validate(o.PaymentButton, validation.In("buy", "pay")),
		"split_payment":     o.SplitPayment.validate(),
	}.Filter()
	return err
}

// OpenCollectionIndexResult represents the structure of the response body obtained with
// Client.GetOpenCollectionIndex.
type OpenCollectionIndexResult struct {
	OpenCollections *[]OpenCollection `json:"open_collections,omitempty"`
	Page            json.Number       `json:"page,omitempty"`
}

// Logo represents a set of URLs to logo images for a collection.
type Logo struct {
	ThumbURL  string `json:"thumb_url,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

// Photo represents a set of URLs to images for an open collection.
type Photo struct {
	RetinaURL string `json:"retina_url,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

// SplitPayment represents data for a split payment made in collections or open collections.
type SplitPayment struct {
	Email       string `json:"email,omitempty"`
	FixedCut    uint   `json:"fixed_cut,omitempty"`
	VariableCut uint   `json:"variable_cut,omitempty"`
	SplitHeader bool   `json:"split_header,omitempty"`
}

func (s *SplitPayment) validate() error {
	err := validation.Errors{
		"email": validation.Validate(s.Email, is.Email),
	}.Filter()
	return err
}

package billplz

// PaymentMethod represents the data for a payment method related to a collection.
type PaymentMethod struct {
	Code   string `json:"code,omitempty"`
	Name   string `json:"name,omitempty"`
	Active bool   `json:"active,omitempty"`
}

// PaymentMethodList represents the structure of payment method data that is sent to and received
// from the Billplz API.
type PaymentMethodList struct {
	PaymentMethods *[]PaymentMethod `json:"payment_methods,omitempty"`
}

package billplz

type PaymentMethod struct {
	Code   string `json:"code,omitempty"`
	Name   string `json:"name,omitempty"`
	Active bool   `json:"active,omitempty"`
}

type PaymentMethodList struct {
	PaymentMethods *[]PaymentMethod `json:"payment_methods,omitempty"`
}

package billplz

type PaymentMethod struct {
	Code   string `json:"code,omitempty"`
	Name   string `json:"name,omitempty"`
	Active bool   `json:"active,omitempty"`
}

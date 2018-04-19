package billplz

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

type BankAccountCheckResponse struct {
	Name string `json:"name,omitempty"`
}

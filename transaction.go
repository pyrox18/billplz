package billplz

type Transaction struct {
	ID             string `json:"id,omitempty"`
	Status         string `json:"status,omitempty"`
	CompletedAt    string `json:"completed_at,omitempty"`
	PaymentChannel string `json:"payment_channel,omitempty"`
}

type BillTransactions struct {
	BillID       string         `json:"bill_id,omitempty"`
	Transactions *[]Transaction `json:"transactions,omitempty"`
	Page         uint           `json:"page,omitempty"`
}

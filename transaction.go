package billplz

import "encoding/json"

// Transaction represents a transaction made for a bill.
type Transaction struct {
	ID             string `json:"id,omitempty"`
	Status         string `json:"status,omitempty"`
	CompletedAt    string `json:"completed_at,omitempty"`
	PaymentChannel string `json:"payment_channel,omitempty"`
}

// BillTransactions represent the structure of a list of transactions received from the Billplz API.
type BillTransactions struct {
	BillID       string         `json:"bill_id,omitempty"`
	Transactions *[]Transaction `json:"transactions,omitempty"`
	Page         json.Number    `json:"page,omitempty"`
}

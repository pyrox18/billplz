package billplz

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

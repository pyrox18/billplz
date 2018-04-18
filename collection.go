package billplz

type Collection struct {
	ID           string        `json:"id,omitempty"`
	Title        string        `json:"title,omitempty"`
	Logo         *Logo         `json:"logo,omitempty"`
	SplitPayment *SplitPayment `json:"split_payment,omitempty"`
	Status       string        `json:"status,omitempty"`
}

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

type Logo struct {
	ThumbURL  string `json:"thumb_url,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type Photo struct {
	RetinaURL string `json:"retina_url,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type SplitPayment struct {
	Email       string `json:"email,omitempty"`
	FixedCut    uint   `json:"fixed_cut,omitempty"`
	VariableCut uint   `json:"variable_cut,omitempty"`
	SplitHeader bool   `json:"split_header,omitempty"`
}

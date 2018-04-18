package billplz

type Collection struct {
	ID           string        `json:"id,omitempty"`
	Title        string        `json:"title,omitempty"`
	Logo         *Logo         `json:"logo,omitempty"`
	SplitPayment *SplitPayment `json:"split_payment,omitempty"`
	Status       string        `json:"status,omitempty"`
}

type Logo struct {
	ThumbURL  string `json:"thumb_url,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type SplitPayment struct {
	Email       string `json:"email,omitempty"`
	FixedCut    uint   `json:"fixed_cut,omitempty"`
	VariableCut uint   `json:"variable_cut,omitempty"`
	SplitHeader bool   `json:"split_header,omitempty"`
}

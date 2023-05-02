package view

type Question struct {
	SpecialID string        `json:"special_id"`
	Statement string        `json:"statement"`
	Order     int           `json:"order"`
	Type      string        `json:"type"`
	Category  string        `json:"category"`
	Options   []OptionParam `json:"options"`
}

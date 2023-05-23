package view

type Question struct {
	ID        uint     `json:"id"`
	SpecialID string   `json:"special_id"`
	Statement string   `json:"statement"`
	Order     int      `json:"order"`
	Type      string   `json:"type"`
	Category  string   `json:"category"`
	Options   []Option `json:"options"`
	Answer    []Answer `json:"answer"`
}

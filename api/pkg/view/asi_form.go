package view

type ASIForm struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Edition     int    `json:"edition"`
	Version     string `json:"version"`
}

package view

type ASIForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Edition     int    `json:"edition"`
	Version     string `json:"version"`
}

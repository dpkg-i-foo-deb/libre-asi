package view

type Option struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	Value       int    `json:"value"`
}

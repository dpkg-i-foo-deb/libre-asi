package view

type Administrator struct {
	ID                 uint   `json:"id"`
	Email              string `json:"email"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	NeedsPasswordReset bool   `json:"needsPasswordReset"`
}

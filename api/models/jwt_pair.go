package models

type JWTPair struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh-token"`
}

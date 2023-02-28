package models

import "github.com/gofiber/fiber/v2"

type JWTPair struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh-token"`
}

type JwtCookies struct {
	AccessToken  *fiber.Cookie `json:"access_token"`
	RefreshToken *fiber.Cookie `json:"refresh_token"`
}

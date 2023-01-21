package models

import "github.com/gofiber/fiber/v2"

type JwtCookies struct {
	AccessToken  *fiber.Cookie `json:"access_token"`
	RefreshToken *fiber.Cookie `json:"refresh_token"`
}

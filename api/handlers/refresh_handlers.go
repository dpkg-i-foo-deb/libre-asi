package handlers

import (
	"libre-asi-api/auth"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func RefreshHandler(c *fiber.Ctx) error {

	var pair models.JWTPair
	var cookies models.JwtCookies
	var email string
	var role string
	var accessTk *fiber.Cookie
	var refreshTk *fiber.Cookie
	var err error

	email, err = auth.EmailFromToken(c.Cookies("refresh-token"))

	if err != nil {
		return util.HandleFiberError(c, errors.ErrInternalError)
	}

	role, err = auth.RoleFromToken(c.Cookies("refresh-token"))

	if err != nil {
		return util.HandleFiberError(c, errors.ErrInternalError)
	}

	pair, err = auth.GenerateJWTPair(email, role)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrInternalError)
	}

	accessTk = auth.GenerateAccessCookie(pair.Token)
	refreshTk = auth.GenerateRefreshCookie(pair.RefreshToken)

	cookies.AccessToken = accessTk
	cookies.RefreshToken = refreshTk

	return c.Status(200).JSON(cookies)

}

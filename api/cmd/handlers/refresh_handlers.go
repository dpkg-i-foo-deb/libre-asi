package handlers

import (
	"libre-asi-api/pkg/auth"
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/models"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func Refresh(c *fiber.Ctx) error {

	var pair models.JWTPair
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

	c.Cookie(accessTk)
	c.Cookie(refreshTk)

	return util.SendSuccess(c, 200, "Session refreshed")

}

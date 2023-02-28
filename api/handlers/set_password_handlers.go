package handlers

import (
	"libre-asi-api/auth"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func SetPasswordHandler(c *fiber.Ctx) error {

	var credentials models.PasswordChange

	token := c.Cookies("password-reset-token")

	if token == "" {
		return util.HandleFiberError(c, errors.ErrAccessDenied)
	}

	email, err := auth.EmailFromToken(token)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrAccessDenied)
	}

	err = c.BodyParser(&credentials)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	err = services.SetPasswordService(credentials, email)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "New password has been set")
}

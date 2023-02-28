package handlers

import (
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func SetPasswordHandler(c *fiber.Ctx) error {

	var credentials models.PasswordChange

	err := c.BodyParser(&credentials)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	return util.SendSuccess(c, 200, "New password has been set")
}

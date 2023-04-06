package handlers

import (
	"libre-asi-api/errors"
	"libre-asi-api/services"
	"libre-asi-api/util"
	"libre-asi-api/view"

	"github.com/gofiber/fiber/v2"
)

func CheckSetup(c *fiber.Ctx) error {

	err := services.CheckSetup()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Setup not needed")
}

func Setup(c *fiber.Ctx) error {

	var user view.Administrator

	if c.BodyParser(&user) != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	err := services.Setup(user)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Setup complete")
}

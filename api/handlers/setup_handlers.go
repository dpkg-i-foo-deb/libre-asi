package handlers

import (
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func CheckSetupHandler(c *fiber.Ctx) error {

	res := models.Response{
		Status:  string(models.OK),
		Message: "Setup is not needed",
	}

	err := services.CheckSetupService()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(res)
}

func SetupHandler(c *fiber.Ctx) error {

	res := models.Response{
		Status:  string(models.OK),
		Message: "Setup completed",
	}

	var user models.User

	if c.BodyParser(&user) != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	err := services.SetupService(user)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(201).JSON(res)
}

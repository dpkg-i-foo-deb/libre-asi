package handlers

import (
	"github.com/gofiber/fiber/v2"
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"
)

func GetAdministratorsHandler(c *fiber.Ctx) error {

	administrators, err := services.GetAdministratorsService()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(administrators)
}

func RegisterAdministratorHandler(c *fiber.Ctx) error {

	var newUser models.User

	err := c.BodyParser(&newUser)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	registeredUser, err := services.RegisterAdministratorService(newUser)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(201).JSON(registeredUser)
}

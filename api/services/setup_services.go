package services

import (
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func SetupService(c *fiber.Ctx) error {

	var res models.Response

	res.Status = string(models.OK)
	res.Message = "Setup Complete"

	if createAdmin(c) != nil {
		return c.Status(400).JSON(res)
	}

	return c.Status(201).JSON(res)
}

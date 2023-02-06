package services

import (
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func IndexService(c *fiber.Ctx) error {

	r := models.Response{
		Status:  string(models.OK),
		Message: "Welcome to LibreASI API!",
	}

	return c.Status(200).JSON(r)
}

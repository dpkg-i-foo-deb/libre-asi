package handlers

import (
	"libre-asi-api/services"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {

	response := services.IndexService()

	return c.Status(200).JSON(response)

}

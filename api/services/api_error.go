package services

import (
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func sendMessage(c *fiber.Ctx, status int, response *models.Response) {
	c.Status(status).JSON(response)
}

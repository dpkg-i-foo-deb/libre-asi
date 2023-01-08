package services

import (
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func sendAPIError(c *fiber.Ctx, message string, status int) {
	response.Status = string(models.STATUS_ERROR)
	response.Message = message
	c.Status(status).JSON(response)
}

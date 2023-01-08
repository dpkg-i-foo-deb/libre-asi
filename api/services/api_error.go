package services

import (
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func checkAPIError(err error, c *fiber.Ctx) {
	if err != nil {
		response.Status = string(models.STATUS_ERROR)
		response.Message = err.Error()

		c.Status(fiber.StatusInternalServerError).JSON(response)
	}

}

func sendAPIError(c *fiber.Ctx, message string, status int) {
	response.Status = string(models.STATUS_ERROR)
	response.Message = message
	c.Status(status).JSON(response)
}

package util

import (
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func SendSuccess(c *fiber.Ctx, status int, message string) error {

	res := models.Response{
		Status:  string(models.OK),
		Message: message,
	}

	return c.Status(200).JSON(res)

}

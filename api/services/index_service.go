package services

import (
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

var response models.Response = models.Response{
	Status:  "",
	Message: "",
}

func IndexService(c *fiber.Ctx) error {

	response.Status = "OK"
	response.Message = "Welcome to Libre-ASI API"

	return c.Status(200).JSON(response)
}

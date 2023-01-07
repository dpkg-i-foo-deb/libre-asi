package services

import (
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func LoginAdminService(c *fiber.Ctx) error {

	response.Status = string(models.STATUS_OK)
	response.Message = "Loggeed In"

	return c.Status(200).JSON(response)
}

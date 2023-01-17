package services

import (
	"libre-asi-api/database"
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetAdministratorsService(c *fiber.Ctx) error {

	var res models.Response
	res.Status = string(models.STATUS_ERROR)
	res.Message = "Something went wrong"

	var admins []models.User

	result := database.DB.Omit("people", "password").Preload("Administrators").Find(&admins)

	if result.Error != nil {
		return c.Status(500).JSON(res)
	}

	return c.Status(200).JSON(admins)
}

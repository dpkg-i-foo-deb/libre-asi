package services

import (
	"libre-asi-api/database"
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func WorldService(c *fiber.Ctx) error {

	var world []models.Country
	var res models.Response

	result := database.DB.Preload("States").Preload("Translations").Preload("States.Cities").Find(&world)

	if result.Error != nil {
		res.Status = string(models.ERROR)
		res.Message = "Something failed"
		return c.Status(500).JSON(res)
	}

	return c.Status(200).JSON(world)
}

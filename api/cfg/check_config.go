package cfg

import (
	"errors"
	"libre-asi-api/database"
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Config struct {
	Filter func(c *fiber.Ctx) bool
}

var ConfigDefault = Config{
	Filter: nil,
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var res models.Response

		res.Status = string(models.SETUP_REQUIRED)
		res.Message = "Libre-ASI requires set-up"

		//Check if the app requires configuration

		err := CheckExistingAdmin()

		if err != nil && c.Path() != "/set-up" {
			return c.Status(412).JSON(res)
		}

		if err == nil && c.Path() == "/set-up" && c.Route().Method == "POST" {
			res.Status = string(models.DENIED)
			res.Message = "Libre-ASI is already set-up"
			return c.Status(409).JSON(res)
		}

		return c.Next()
	}
}

func CheckExistingAdmin() error {

	var a models.Administrator

	if database.DB.First(&a).Error == gorm.ErrRecordNotFound {
		return errors.New("setup required")
	}

	return nil
}

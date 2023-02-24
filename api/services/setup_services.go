package services

import (
	"errors"
	"libre-asi-api/cfg"
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func SetupService(c *fiber.Ctx) error {

	var res models.Response

	res.Status = string(models.OK)
	res.Message = "Setup Complete"

	if createAdmin(c) != nil {
		return c.Status(400).JSON(res)
	}

	return c.Status(201).JSON(res)
}

func CheckSetupService() (models.Response, error) {

	var res models.Response

	res.Status = string(models.OK)
	res.Message = "Setup is OK"

	if cfg.CheckExistingAdmin() != nil {
		res.Status = string(models.SETUP_REQUIRED)
		res.Message = "Libre-ASI requires set-up"

		return res, errors.New("setup was required")
	}

	return res, nil
}

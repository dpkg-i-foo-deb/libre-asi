package handlers

import (
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func RegisterPatient(c *fiber.Ctx) error {

	var newPatient *models.Patient

	if err := c.BodyParser(&newPatient); err != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	newPatient, err := services.RegisterPatient(*newPatient)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(201).JSON(newPatient)

	return nil
}

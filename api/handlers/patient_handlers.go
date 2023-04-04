package handlers

import (
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"
	"strconv"

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
}

func GetPatients(c *fiber.Ctx) error {

	patients, err := services.GetPatients()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(patients)
}

func GetPatient(c *fiber.Ctx) error {

	id := c.Params("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	patient, err := services.GetPatient(uint(idInt))

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(patient)
}

func UpdatePatient(c *fiber.Ctx) error {

	var updatedPatient *models.Patient

	if c.BodyParser(&updatedPatient) != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	if err := services.UpdatePatient(*updatedPatient); err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Updated")
}

func DeletePatient(c *fiber.Ctx) error {

	id := c.Params("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	if err := services.DeletePatient(uint(idInt)); err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Deleted")
}

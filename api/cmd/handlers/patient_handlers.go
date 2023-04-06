package handlers

import (
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/view"
	"libre-asi-api/services"
	"libre-asi-api/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterPatient(c *fiber.Ctx) error {

	var newPatient *view.Patient

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

	var updatedPatient *view.Patient

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

func SetPatientPicture(c *fiber.Ctx) error {

	imageData := c.Body()

	id := c.Params("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	contentType := c.Get("Content-Type")

	if contentType != "image/jpeg" {
		return util.HandleFiberError(c, errors.ErrBadContentType)

	}

	if err := services.SetPatientPicture(idInt, imageData); err != nil {
		return util.HandleFiberError(c, err)

	}

	return util.SendSuccess(c, 200, "Picture set")

}

func GetPatientPicture(c *fiber.Ctx) error {

	id := c.Params("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	imageData, err := services.GetPatientPicture(idInt)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	c.Set("Content-Type", "image/jpeg")

	return c.Status(200).Send(imageData)
}

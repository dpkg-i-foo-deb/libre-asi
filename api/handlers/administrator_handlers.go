package handlers

import (
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAdministratorsHandler(c *fiber.Ctx) error {

	administrators, err := services.GetAdministratorsService()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(administrators)
}

func RegisterAdministratorHandler(c *fiber.Ctx) error {

	var newUser models.User

	err := c.BodyParser(&newUser)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	registeredUser, err := services.RegisterAdministratorService(newUser)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(201).JSON(registeredUser)
}

func UpdateAdministratorHandler(c *fiber.Ctx) error {

	var updatedAdmin models.User

	if c.BodyParser(&updatedAdmin) != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	if err := services.UpdateAdministratorService(updatedAdmin); err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Updated")
}

func DeleteAdministratorHandler(c *fiber.Ctx) error {

	id := c.Params("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	if err := services.DeleteAdministratorService(intId); err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Deleted")
}

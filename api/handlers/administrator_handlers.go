package handlers

import (
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAdministrators(c *fiber.Ctx) error {

	administrators, err := services.GetAdministrators()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(administrators)
}

func RegisterAdministrator(c *fiber.Ctx) error {

	var newAdmin models.Administrator

	err := c.BodyParser(&newAdmin)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	registeredAdmin, err := services.RegisterAdministrator(newAdmin)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(201).JSON(registeredAdmin)
}

func UpdateAdministrator(c *fiber.Ctx) error {

	var updatedAdmin models.Administrator

	if c.BodyParser(&updatedAdmin) != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	if err := services.UpdateAdministrator(updatedAdmin); err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Updated")
}

func DeleteAdministrator(c *fiber.Ctx) error {

	id := c.Params("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	if err := services.DeleteAdministrator(intId); err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Deleted")
}

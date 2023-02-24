package handlers

import (
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func CheckSetupHandler(c *fiber.Ctx) error {
	response, err := services.CheckSetupService()

	if err != nil {
		return util.HandleFiberError(c, err, response)
	}

	return c.Status(200).JSON(response)
}

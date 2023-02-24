package handlers

import (
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {

	err := services.IndexService()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Welcome to Libre-ASI API")
}

package handlers

import (
	"libre-asi-api/pkg/services"
	"libre-asi-api/pkg/util"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	err := services.Index()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 200, "Welcome to Libre-ASI API")
}

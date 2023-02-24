package handlers

import (
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func CheckSetupHandler(c *fiber.Ctx) error {

	res := models.Response{
		Status:  string(models.OK),
		Message: "Setup is not needed",
	}

	err := services.CheckSetupService()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(res)
}

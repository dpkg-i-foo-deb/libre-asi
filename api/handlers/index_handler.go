package handlers

import (
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {

	res := models.Response{
		Status:  string(models.OK),
		Message: "Welcome to Libre-ASI API",
	}

	err := services.IndexService()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(res)

}

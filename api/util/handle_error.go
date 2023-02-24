package util

import (
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleErorLog(err error) {
	if err != nil {
		log.Println(err)
	}
}

func HandleErrorStop(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HandleFiberError(c *fiber.Ctx, err error) error {

	status := 200
	var res models.Response

	switch err {
	case errors.ErrSetupRequired:
		status = 412
		res.Status = string(models.SETUP_REQUIRED)
		res.Message = "Libre-ASI Requires Set-Up"
	case errors.ErrCheckRequest:
		status = 400
		res.Status = string(models.CHECK_REQUEST)
		res.Message = "Check your request"
	}

	return c.Status(status).JSON(res)
}

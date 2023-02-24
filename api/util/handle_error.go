package util

import (
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

func HandleFiberError(c *fiber.Ctx, err error, res models.Response) error {

	status := 200

	switch res.Status {
	case string(models.SETUP_REQUIRED):
		status = 412
	}

	return c.Status(status).JSON(res)
}

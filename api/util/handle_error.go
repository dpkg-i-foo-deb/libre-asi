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
	return c.Status(500).JSON(res)
}

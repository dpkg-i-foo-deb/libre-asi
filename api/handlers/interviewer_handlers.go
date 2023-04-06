package handlers

import (
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func GetInterviewers(c *fiber.Ctx) error {

	interviewers, err := services.GetInterviewers()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(interviewers)
}

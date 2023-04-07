package handlers

import (
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/services"
	"libre-asi-api/pkg/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetInterviewer(c *fiber.Ctx) error {

	id := c.Params("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	interviewer, err := services.GetInterviewer(uint(idInt))

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(interviewer)
}
func GetInterviewers(c *fiber.Ctx) error {

	interviewers, err := services.GetInterviewers()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(interviewers)
}

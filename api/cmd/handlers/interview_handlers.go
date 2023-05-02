package handlers

import (
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/services"
	"libre-asi-api/pkg/util"
	"libre-asi-api/pkg/view"

	"github.com/gofiber/fiber/v2"
)

func StartInterview(c *fiber.Ctx) error {

	i := view.Interview{}

	err := c.BodyParser(&i)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	id, err := services.StartInterview(int(i.PatientID), int(i.InterviewerID))

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	i.ID = uint(id)

	return c.Status(201).JSON(i)

}

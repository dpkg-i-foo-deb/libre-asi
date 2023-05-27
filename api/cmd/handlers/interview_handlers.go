package handlers

import (
	"fmt"
	"libre-asi-api/pkg/auth"
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/services"
	"libre-asi-api/pkg/util"
	"libre-asi-api/pkg/view"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func StartInterview(c *fiber.Ctx) error {

	i := &view.Interview{}

	err := c.BodyParser(i)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	tk := c.Cookies("access-token")

	email, err := auth.EmailFromToken(tk)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrAccessDenied)
	}

	i, err = services.StartInterview(int(i.PatientID), email)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(201).JSON(i)

}

func GetInterviews(c *fiber.Ctx) error {
	interviews, err := services.GetInterviews()

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(interviews)
}

func GetInterview(c *fiber.Ctx) error {

	id := c.Params("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	interview, err := services.GetInterview(uint(idInt))

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(interview)
}

func NextQuestion(c *fiber.Ctx) error {

	i := &view.Interview{}

	err := c.BodyParser(i)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	i, err = services.NextQuestion(i)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(i)

}

func GetQuestion(c *fiber.Ctx) error {

	code := c.Params("code")

	question, err := services.GetQuestion(code)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(question)
}

func AnswerQuestion(c *fiber.Ctx) error {

	a := []view.Answer{}

	err := c.BodyParser(&a)

	if err != nil {

		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	fmt.Println(a[0].Comment)

	id := c.Params("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	err = services.AnswerQuestion(a, uint(idInt))

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.SendStatus(201)
}

func ComputeResults(c *fiber.Ctx) error {

	i := &view.Interview{}

	err := c.BodyParser(i)

	if err != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	r, err := services.ComputeResults(i)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return c.Status(200).JSON(r)

}

package handlers

import (
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {

	var res models.Response
	var requiresAdmin bool
	var role models.Role

	switch c.Params("role") {
	case "admin":
		role = models.ADMINISTRATOR
		requiresAdmin = true
		res.Message = "Administrator Created"
	case "interviewer":
		role = models.INTERVIEWER
		res.Message = "Interviewer Created"
		requiresAdmin = true
	case "attendant":
		return util.HandleFiberError(c, errors.ErrNotImplemmented)
	case "patient":
		return util.HandleFiberError(c, errors.ErrNotImplemmented)
	default:
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	if requiresAdmin && !isAdmin(c.Cookies("access-token")) {
		res.Status = string(models.DENIED)
		res.Message = "This session doesn't have enough privileges"
		return c.Status(401).JSON(&res)
	}

	return nil
}

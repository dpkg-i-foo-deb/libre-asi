package handlers

import (
	"libre-asi-api/auth"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {

	var requiresAdmin bool
	var role models.Role
	var user models.User

	switch c.Params("role") {
	case "admin":
		role = models.ADMINISTRATOR
		requiresAdmin = true
	case "interviewer":
		role = models.INTERVIEWER
		requiresAdmin = true
	case "attendant":
		return util.HandleFiberError(c, errors.ErrNotImplemmented)
	case "patient":
		return util.HandleFiberError(c, errors.ErrNotImplemmented)
	default:
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	if requiresAdmin && !auth.IsAdmin(c.Cookies("access-token")) {
		return util.HandleFiberError(c, errors.ErrAccessDenied)
	}

	if c.BodyParser(&user) != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	err := services.RegisterService(user, role)

	if err != nil {
		return util.HandleFiberError(c, err)
	}

	return util.SendSuccess(c, 201, "Account created successfully")
}

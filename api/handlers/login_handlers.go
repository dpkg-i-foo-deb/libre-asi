package handlers

import (
	"libre-asi-api/auth"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/services"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {

	var tk *models.JWTPair
	var pk *models.PasswordResetTk
	var err error

	switch c.Params("role") {
	case string(models.ADMINISTRATOR):
		_, tk, pk, err = loginAdmin(c)
	case string(models.INTERVIEWER):
		_, tk, pk, err = loginInterviewer(c)
	case "patient":
		return util.HandleFiberError(c, errors.ErrNotImplemmented)
	default:
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	if err != nil {

		if err == errors.ErrrNeedsPasswordReset {

			tkCookie := auth.GeneratePasswordResetCookie(pk.Token)

			c.Cookie(tkCookie)

			return c.SendStatus(428)

		}

		return util.HandleFiberError(c, err)
	}

	accessToken := auth.GenerateAccessCookie(tk.Token)
	refreshToken := auth.GenerateRefreshCookie(tk.RefreshToken)

	c.Cookie(accessToken)
	c.Cookie(refreshToken)

	return util.SendSuccess(c, 200, "Welcome back")

}

func loginAdmin(c *fiber.Ctx) (*models.Administrator, *models.JWTPair, *models.PasswordResetTk, error) {
	var a models.Administrator

	if c.BodyParser(&a) != nil {
		return nil, nil, nil, util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	return services.LoginAdminService(a)
}

func loginInterviewer(c *fiber.Ctx) (*models.Interviewer, *models.JWTPair, *models.PasswordResetTk, error) {
	var i models.Interviewer

	if c.BodyParser(&i) != nil {
		return nil, nil, nil, util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	return services.LoginInterviewerService(i)
}

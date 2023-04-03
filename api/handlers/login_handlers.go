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

	var u models.User
	var role models.Role

	if c.BodyParser(&u) != nil {
		return util.HandleFiberError(c, errors.ErrCheckRequest)
	}

	switch c.Params("role") {
	case string(models.ADMINISTRATOR):
		role = models.ADMINISTRATOR
	case string(models.INTERVIEWER):
		role = models.INTERVIEWER
	case "patient":
		return util.HandleFiberError(c, errors.ErrNotImplemmented)
	default:
		return util.HandleFiberError(c, errors.ErrBadRoute)
	}

	tk, err := services.LoginService(u, role)

	if err == errors.ErrrNeedsPasswordReset {

		tk, errTk := auth.GeneratePasswordResetToken(u.Email)

		if errTk != nil {
			return util.HandleFiberError(c, errors.ErrInternalError)
		}

		tkCookie := auth.GeneratePasswordResetCookie(tk.Token)

		c.Cookie(tkCookie)

		return c.SendStatus(428)

	}

	if err != nil {

		return util.HandleFiberError(c, err)
	}

	accessToken := auth.GenerateAccessCookie(tk.Token)
	refreshToken := auth.GenerateRefreshCookie(tk.RefreshToken)

	c.Cookie(accessToken)
	c.Cookie(refreshToken)

	return util.SendSuccess(c, 200, "Welcome back")

}

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
	var jwtResponse models.JwtCookies
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

		return c.Status(428).JSON(tkCookie)

	}

	if err != nil {

		return util.HandleFiberError(c, err)
	}

	accessToken := auth.GenerateAccessCookie(tk.Token)
	refreshToken := auth.GenerateRefreshCookie(tk.RefreshToken)

	jwtResponse.AccessToken = accessToken
	jwtResponse.RefreshToken = refreshToken

	//TODO use real cookies when Sveltekit allows easy cookie parsing
	//c.Cookie(auth)
	//c.Cookie(refresh)

	return c.Status(200).JSON(jwtResponse)

}

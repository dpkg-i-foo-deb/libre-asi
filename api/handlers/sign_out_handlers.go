package handlers

import (
	"libre-asi-api/auth"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
)

func SignOutHandler(c *fiber.Ctx) error {

	refresh := auth.GenerateFakeRefreshCookie()
	auth := auth.GenerateFakeAccessCookie()

	c.Cookie(auth)
	c.Cookie(refresh)

	return util.SendSuccess(c, 200, "Signed Out")
}

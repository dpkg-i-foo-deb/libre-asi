package services

import (
	"libre-asi-api/auth"
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func SignOutService(c *fiber.Ctx) error {
	refresh := auth.GenerateFakeRefreshCookie()
	auth := auth.GenerateFakeAccessCookie()

	c.Cookie(auth)
	c.Cookie(refresh)

	res := models.Response{
		Status:  string(models.STATUS_OK),
		Message: "Signed Out",
	}

	return c.Status(200).JSON(res)
}

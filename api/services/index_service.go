package services

import "github.com/gofiber/fiber/v2"

func IndexService(c *fiber.Ctx) error {
	return c.SendString("Welcome to Libre-ASI API")
}

package auth

import (
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
)

func ValidateAndContinue(c *fiber.Ctx) error {

	accessToken := c.Cookies("access-token")
	var response models.Response

	response.Status = string(models.STATUS_DENIED)
	response.Message = "The access token was not present"

	if accessToken == "" {
		c.Status(fiber.StatusUnauthorized).JSON(response)
		return nil
	}

	isValid, err := ValidateToken(accessToken)

	if isValid && err == nil {
		c.Next()
		return nil
	}

	response.Message = "The access token is not valid or has expired"
	response.Status = string(models.STATUS_DENIED)
	c.Status(fiber.StatusUnauthorized).JSON(response)
	return nil

}

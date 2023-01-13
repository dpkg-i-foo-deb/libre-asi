package auth

import (
	"libre-asi-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ValidateAccessToken(c *fiber.Ctx) error {

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

func ValidateRefreshToken(c *fiber.Ctx) error {
	tk := c.Cookies("refresh-token")

	var res models.Response
	res.Status = string(models.STATUS_DENIED)

	if tk == "" {
		res.Message = "The refresh token was not present"
		return c.Status(401).JSON(res)
	}

	isValid, err := ValidateToken(tk)

	if isValid && err == nil {
		c.Next()
		return nil
	}

	res.Message = "The refresh token has expired or is invalid"
	return c.Status(401).JSON(res)
}

func ValidateRefreshTokenDate(c *fiber.Ctx) error {

	//This should run AFTER validating the refresh token

	var res models.Response

	res.Status = string(models.STATUS_DENIED)
	res.Message = "The refresh token is still valid"

	tk := c.Cookies("refresh-token")

	claims, err := GetTokenClaims(tk)

	if err != nil {
		res.Status = string(models.STATUS_ERROR)
		res.Message = "Something failed"
		return c.Status(500).JSON(res)
	}

	if claims.ExpiresAt > time.Now().Unix()+int64(time.Minute)*5 {
		return c.Status(412).JSON(res)
	}

	c.Next()

	return nil
}

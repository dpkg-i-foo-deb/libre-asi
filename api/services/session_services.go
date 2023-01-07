package services

import (
	"database/sql"
	"libre-asi-api/auth"
	"libre-asi-api/database"
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func LoginAdminService(c *fiber.Ctx) error {

	var user models.User
	var queriedUser models.User
	var pair models.JWTPair

	err := c.BodyParser(&user)

	if err != nil {
		response.Status = string(models.STATUS_DENIED)
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = database.LoginStatement.QueryRow(user.Email).Scan(&queriedUser.Email, &queriedUser.Password)

	if err != nil && err != sql.ErrNoRows {
		response.Status = string(models.STATUS_ERROR)
		response.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if err == sql.ErrNoRows {
		response.Status = string(models.STATUS_DENIED)
		response.Message = "User or Password Incorrect"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	err = bcrypt.CompareHashAndPassword([]byte(queriedUser.Password), []byte(user.Password))

	if err != nil {
		response.Status = string(models.STATUS_DENIED)
		response.Message = "User or Password Incorrect"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	pair, err = auth.GenerateJWTPair(user.Email, string(models.ADMINISTRATOR))

	if err != nil {
		response.Status = string(models.STATUS_ERROR)
		response.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	accessCookie := auth.GenerateAccessCookie(pair.Token)
	refreshCookie := auth.GenerateRefreshCookie(pair.RefreshToken)

	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	response.Status = string(models.STATUS_OK)
	response.Message = "Logged In"

	return c.Status(200).JSON(response)
}

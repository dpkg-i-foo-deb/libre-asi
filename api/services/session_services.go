package services

import (
	"database/sql"
	"libre-asi-api/auth"
	"libre-asi-api/database"
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(c *fiber.Ctx) error {

	var user models.User
	var queriedUser models.User
	var pair models.JWTPair
	var stmt *sql.Stmt
	var role string

	err := c.BodyParser(&user)

	if err != nil {
		response.Status = string(models.STATUS_DENIED)
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	switch c.Params("role") {
	case "admin":
		stmt = database.LoginAdminStatement
		role = string(models.ADMINISTRATOR)

	case "interviewer":
		stmt = database.LoginInterviewerStatement
		role = string(models.INTERVIEWER)
	default:
		response.Status = string(models.STATUS_DENIED)
		response.Message = "Bad Route"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = stmt.QueryRow(user.Email).Scan(&queriedUser.Email, &queriedUser.Password)

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

	pair, err = auth.GenerateJWTPair(user.Email, role)

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

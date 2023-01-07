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

func RefreshTokenService(c *fiber.Ctx) error {

	refreshToken := c.Cookies("refresh-token")
	var newPair models.JWTPair
	var userEmail string
	var role string
	var newAccessCookie *fiber.Cookie
	var newRefreshCookie *fiber.Cookie

	response.Status = string(models.STATUS_DENIED)
	response.Message = "The refresh token was not present"

	if refreshToken == "" {

		c.Status(fiber.StatusUnauthorized).JSON(response)
		return nil
	}

	isValid, err := auth.ValidateToken(refreshToken)

	if isValid && err == nil {

		userEmail, err = auth.EmailFromToken(refreshToken)

		response.Status = string(models.STATUS_ERROR)
		response.Message = "Failed to generate the refresh token"

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		role, err = auth.RoleFromToken(refreshToken)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		newPair, err = auth.GenerateJWTPair(userEmail, role)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		newAccessCookie = auth.GenerateAccessCookie(newPair.Token)
		newRefreshCookie = auth.GenerateRefreshCookie(newPair.RefreshToken)

		c.Cookie(newAccessCookie)
		c.Cookie(newRefreshCookie)

		response.Status = string(models.STATUS_OK)
		response.Message = "Refreshed"

		return c.Status(200).JSON(response)

	} else {
		response.Status = string(models.STATUS_DENIED)
		response.Message = "The refresh token has expired or wasn't present"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

}

func SignOutService(c *fiber.Ctx) error {
	response.Status = string(models.STATUS_OK)
	response.Message = "Signed Out..."

	var newRefreshCookie *fiber.Cookie = auth.GenerateFakeRefreshCookie()
	var newAcessCookie *fiber.Cookie = auth.GenerateFakeAccessCookie()

	//Set the new cookies

	c.Cookie(newAcessCookie)
	c.Cookie(newRefreshCookie)

	return c.Status(fiber.StatusOK).JSON(response)
}

func RegisterService(c *fiber.Ctx) error {

	switch c.Params("role") {
	case "admin":
	case "interviewer":
	case "patient":

	default:
		response.Status = string(models.STATUS_DENIED)
		response.Message = "Bad Route"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Status = string(models.STATUS_OK)
	response.Message = "Registered correctly"
	return c.Status(fiber.StatusCreated).JSON(response)
}

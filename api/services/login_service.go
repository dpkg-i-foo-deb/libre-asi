package services

import (
	"errors"
	"libre-asi-api/auth"
	"libre-asi-api/database"
	"libre-asi-api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(c *fiber.Ctx) error {
	var res models.Response
	var u models.User
	var tk models.JWTPair

	err := c.BodyParser(&u)

	if err != nil {
		res.Status = string(models.STATUS_DENIED)
		res.Message = "Check JSON data"
		return c.Status(400).JSON(res)
	}

	switch c.Params("role") {
	case string(models.ADMINISTRATOR):
		err = loginAdmin(&u)
	case string(models.INTERVIEWER):
		err = loginInterviewer(&u)
	case "patient":
		err = errors.New("not implemmented")
	default:
		res = models.Response{
			Status:  string(models.STATUS_DENIED),
			Message: "Bad route",
		}
		return c.Status(400).JSON(&res)
	}

	if err != nil {
		res.Status = string(models.STATUS_DENIED)
		res.Message = "Invalid credentials or not enough privileges"

		return c.Status(401).JSON(&res)
	}

	tk, err = auth.GenerateJWTPair(u.Email, c.Params("role"))

	if err != nil {
		res.Status = string(models.STATUS_ERROR)
		res.Message = "Something went wrong"

		return c.Status(500).JSON(&res)
	}

	refresh := auth.GenerateRefreshCookie(tk.RefreshToken)
	auth := auth.GenerateAccessCookie(tk.Token)

	c.Cookie(auth)
	c.Cookie(refresh)

	res.Status = string(models.STATUS_OK)
	res.Message = "Logged in as " + c.Params("role")

	return c.Status(200).JSON(res)
}

func loginAdmin(u *models.User) error {

	var user models.User
	var admin models.Administrator

	res := database.DB.Where("email = ?", u.Email).First(&user)

	if res.Error != nil {
		return res.Error
	}

	res = database.DB.Where("user_id = ?", user.ID).First(&admin)

	if res.Error != nil {
		return res.Error
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
}

func loginInterviewer(u *models.User) error {
	var user models.User
	var person models.Person
	var inter models.Interviewer

	res := database.DB.Where("email = ?", u.Email).First(&user)

	if res.Error != nil {
		return res.Error
	}

	res = database.DB.Where("user_id = ?", user.ID).First(&person)

	if res.Error != nil {
		return res.Error
	}

	res = database.DB.Where("person_id = ?", person.ID).First(&inter)

	if res.Error != nil {
		return res.Error
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
}

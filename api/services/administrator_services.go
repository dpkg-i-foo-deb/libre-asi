package services

import (
	"libre-asi-api/database"
	"libre-asi-api/models"
	"libre-asi-api/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAdministratorsService(c *fiber.Ctx) error {

	var res models.Response
	res.Status = string(models.ERROR)
	res.Message = "Something went wrong"

	var admins []models.Administrator
	var ids []uint
	var users []models.User

	result := database.DB.Find(&admins)

	if result.Error != nil {
		return c.Status(500).JSON(res)
	}

	for i := range admins {
		ids = append(ids, admins[i].UserID)
	}

	result = database.DB.Omit("password", "administrators", "people").Find(&users, ids)

	if result.Error != nil {
		return c.Status(500).JSON(res)
	}

	return c.Status(200).JSON(&users)
}

func RegisterAdministratorService(c *fiber.Ctx) error {

	var res models.Response
	res.Status = string(models.ERROR)
	res.Message = "Something went wrong"

	var newUser models.User
	var queriedUser models.User

	err := c.BodyParser(&newUser)

	if err != nil {
		return c.Status(400).JSON(res)
	}

	result := database.DB.Where("email = ?", newUser.Email).First(&queriedUser)

	if result.Error == nil {
		res.Status = string(models.DENIED)
		res.Message = "Email already registered"
		return c.Status(409).JSON(res)
	}

	result = database.DB.Where("username = ?", newUser.Username).First(&queriedUser)

	if result.Error == nil {
		res.Status = string(models.DENIED)
		res.Message = "User name already registered"
		return c.Status(409).JSON(res)
	}

	err = database.DB.Transaction(func(tx *gorm.DB) error {

		password, err := util.MakeRandomPassword()

		if err != nil {
			return err
		}

		hashedPassword, err := util.HashPassword(password)

		if err != nil {
			return err
		}

		newUser.Password = hashedPassword

		result := database.DB.Omit("Administrators", "People").Create(&newUser)

		if result.Error != nil {
			return result.Error
		}

		id := newUser.ID

		var newAdmin = models.Administrator{
			UserID: id,
		}

		result = database.DB.Create(&newAdmin)

		if result.Error != nil {
			return result.Error
		}

		newUser.Password = password

		return nil
	})

	if err != nil {

		return c.Status(500).JSON(res)
	}

	return c.Status(201).JSON(newUser)
}

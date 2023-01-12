package services

import (
	"errors"
	"libre-asi-api/database"
	"libre-asi-api/models"
	"libre-asi-api/util"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterService(c *fiber.Ctx) error {

	var r models.Response
	var err error

	switch c.Params("role") {
	case "admin":
		err = createAdmin(c)
		r.Message = "Administrator Created"
	case "interviewer":
		err = createInterviewer(c)
		r.Message = "Interviewer Created"
	default:
		r = models.Response{
			Status:  string(models.STATUS_DENIED),
			Message: "Bad route",
		}
		return c.Status(400).JSON(&r)
	}

	if err != nil {
		r.Status = string(models.STATUS_ERROR)
		r.Message = "Something went wrong"
		log.Println(err)
		c.Status(500).JSON(r)
		return nil
	}

	r.Status = string(models.STATUS_OK)
	c.Status(201).JSON(r)

	return nil
}

func createAdmin(c *fiber.Ctx) error {
	var u models.User
	var p string

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		err := c.BodyParser(&u)

		if err != nil {
			return err
		}

		p, err = util.HashPassword(u.Password)

		if err != nil {
			return err
		}

		u.Password = p

		if len(u.Administrators) != 1 || u.Administrators == nil {
			return errors.New("tried to create more than one administrator")
		}

		res := database.DB.Omit("People").Create(&u)

		if res.Error != nil {
			return res.Error
		}

		return nil
	})

	return err
}

func createInterviewer(c *fiber.Ctx) error {
	var u models.User

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		err := c.BodyParser(&u)

		if err != nil {
			return err
		}

		if u.People == nil {
			return errors.New("nothing to register")
		}

		if len(u.People) != 1 {
			return errors.New("tried to create more than one person")
		} else {
			if len(u.People[0].Interviewers) > 1 || u.People[0].Interviewers == nil {
				return errors.New("tried to create more than one interviewer")
			}
		}

		r := database.DB.Omit("Administators",
			"Patients",
			"Attendants",
			"PersonID",
			"UserID",
			"Interpretations",
			"Reports",
			"ProfessionTranslations").Create(&u)

		if r.Error != nil {
			return r.Error
		}

		return nil
	})

	return err
}

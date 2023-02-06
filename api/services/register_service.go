package services

import (
	"errors"
	"libre-asi-api/auth"
	"libre-asi-api/database"
	"libre-asi-api/models"
	"libre-asi-api/util"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type registerFunc func(c *fiber.Ctx) error

func RegisterService(c *fiber.Ctx) error {

	var res models.Response
	var err error
	var fn registerFunc
	var requiresAdmin bool

	switch c.Params("role") {
	case "admin":
		fn = createAdmin
		requiresAdmin = true
		res.Message = "Administrator Created"
	case "interviewer":
		fn = createInterviewer
		res.Message = "Interviewer Created"
		requiresAdmin = true
	case "attendant":
		res.Message = "Not implemmented"
	case "patient":
		res.Message = "Not implemmented"
	default:
		res = models.Response{
			Status:  string(models.DENIED),
			Message: "Bad route",
		}
		return c.Status(400).JSON(&res)
	}

	if requiresAdmin && !isAdmin(c.Cookies("access-token")) {
		res.Status = string(models.DENIED)
		res.Message = "This session doesn't have enough privileges"
		return c.Status(401).JSON(&res)
	}

	err = fn(c)

	if err != nil {
		res.Status = string(models.ERROR)
		res.Message = "Check JSON data"
		log.Println(err)
		c.Status(400).JSON(res)
		return nil
	}

	res.Status = string(models.OK)
	c.Status(201).JSON(res)

	return nil
}

func createAdmin(c *fiber.Ctx) error {
	var u models.User
	var pass string

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		err := c.BodyParser(&u)

		if err != nil {
			return err
		}

		pass, err = util.HashPassword(u.Password)

		if err != nil {
			return err
		}

		u.Password = pass

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
	var pass string

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

		pass, err = util.HashPassword(u.Password)

		if err != nil {
			return err
		}

		u.Password = pass

		res := database.DB.Omit("Administators",
			"Patients",
			"Attendants",
			"PersonID",
			"UserID",
			"Interpretations",
			"Reports",
			"ProfessionTranslations").Create(&u)

		if res.Error != nil {
			return res.Error
		}

		return nil
	})

	return err
}

func isAdmin(tk string) bool {

	role, err := auth.RoleFromToken(tk)

	if err != nil {
		return false
	}

	if role != "admin" {
		return false
	}

	return true
}

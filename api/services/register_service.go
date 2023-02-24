package services

import (
	"libre-asi-api/auth"
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type registerFunc func(models.User) error

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
		//TODO enable interviewer login
		//fn = createInterviewer
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

	//TODO tidy up
	err = fn(models.User{})

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

func createAdmin(models.User) error {
	var u models.User

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		u.Administrators = []models.Administrator{
			{},
		}

		pass, err := util.HashPassword(u.Password)

		if err != nil {
			return errors.ErrInternalError
		}

		u.Password = pass

		if database.DB.Omit("People").Create(&u).Error != nil {
			return errors.ErrInternalError
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
			return errors.ErrCheckRequest
		}

		if u.People == nil {
			return errors.ErrNoData
		}

		if len(u.People) != 1 {
			return errors.ErrTooManyEntities
		} else {
			if len(u.People[0].Interviewers) > 1 || u.People[0].Interviewers == nil {
				return errors.ErrTooManyEntities
			}
		}

		pass, err = util.HashPassword(u.Password)

		if err != nil {
			return errors.ErrInternalError
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
			return errors.ErrInternalError
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

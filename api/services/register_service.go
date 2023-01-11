package services

import (
	"libre-asi-api/database"
	"libre-asi-api/models"
	"libre-asi-api/models/public"
	"libre-asi-api/util"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterService(c *fiber.Ctx) error {

	var r models.Response
	var u models.User
	//var a models.Administrator
	var pu public.PublicAdministrator
	var p string

	switch c.Params("role") {
	case "admin":
		err := database.DB.Transaction(func(tx *gorm.DB) error {
			err := c.BodyParser(&pu)

			if err != nil {
				return err
			}

			p, err = util.HashPassword(u.Password)

			if err != nil {
				return err
			}

			u = models.User{
				Email:    pu.Email,
				Username: pu.Username,
				Password: p,
				Administrators: []models.Administrator{
					{},
				},
			}

			res := database.DB.Create(&u)

			if res.Error != nil {
				return res.Error
			}

			r.Status = string(models.STATUS_OK)
			r.Message = "Administrator Created"

			c.Status(201).JSON(r)
			return nil
		})

		if err != nil {
			r.Status = string(models.STATUS_ERROR)
			r.Message = "Something went wrong"
			log.Println(err)
			c.Status(500).JSON(r)
		}

	default:
		r = models.Response{
			Status:  string(models.STATUS_DENIED),
			Message: "Bad route",
		}
		return c.Status(400).JSON(&r)
	}
	return nil
}

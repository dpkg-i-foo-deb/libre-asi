package services

import (
	"database/sql"
	"libre-asi-api/database"
	"libre-asi-api/models"
	"libre-asi-api/util"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterService(c *fiber.Ctx) error {

	var stmt *sql.Stmt
	var person models.Person
	var user models.User
	var hashedPassword string
	//var administrator models.Administrator
	var interviewer models.Interviewer
	var patient models.Patient
	var id int64

	switch c.Params("role") {
	case "admin":
		stmt = database.CreateAdminStatement
	case "interviewer":
		stmt = database.CreateInterviewerStatement
		err := c.BodyParser(&interviewer)
		if err != nil {
			sendAPIError(c, err.Error(), fiber.StatusBadRequest)
			log.Println(err.Error())
			return nil
		}
	case "patient":
		stmt = database.CreatePatientStatement
		err := c.BodyParser(&patient)
		if err != nil {
			log.Println(err.Error())
			sendAPIError(c, err.Error(), fiber.StatusBadRequest)
			return nil
		}

	default:
		sendAPIError(c, "Bad Route", fiber.StatusBadRequest)
		return nil
	}

	err := c.BodyParser(&user)

	if err != nil {
		log.Println(err.Error())
		sendAPIError(c, err.Error(), fiber.StatusBadRequest)
		return nil
	}

	hashedPassword, err = util.HashPassword(*user.Password)

	if err != nil {
		log.Println(err.Error())
		sendAPIError(c, "Something went wrong", fiber.StatusInternalServerError)
		return nil
	}

	transaction, err := database.DB.Begin()

	if err != nil {
		log.Println(err.Error())
		sendAPIError(c, "Something went wrong", fiber.StatusInternalServerError)
		return nil
	}

	err = transaction.Stmt(database.CreateUserStatement).QueryRow(
		user.Email,
		user.Username,
		hashedPassword,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		log.Println(err.Error())
		sendAPIError(c, "Something went wrong", fiber.StatusBadRequest)
		return nil
	}

	if c.Params("role") != "admin" {
		err = c.BodyParser(&person)

		if err != nil {
			log.Println(err.Error())
			transaction.Rollback()
			sendAPIError(c, err.Error(), fiber.StatusBadRequest)
			return nil
		}

		_, err = transaction.Stmt(database.CreatePersonStatement).Exec(
			id,
			person.FistName,
			person.SecondName,
			person.FirstSurname,
			person.SecondSurname,
			person.Age,
			person.Birthdate,
			person.PersonalID,
			time.Now(),
			time.Now(),
		)

		if err != nil {
			log.Println(err.Error())
			transaction.Rollback()
			sendAPIError(c, "Something went wrong", fiber.StatusInternalServerError)
			return nil
		}

		if c.Params("role") == "patient" {
			_, err = transaction.Stmt(stmt).Exec(
				id,
				time.Now(),
				time.Now(),
				patient.SocialSecurityNumber,
				patient.Race,
				patient.ReligiousPreference,
			)
		}

		if c.Params("role") == "interviewer" {
			_, err = transaction.Stmt(stmt).Exec(
				id,
				interviewer.RMA,
				time.Now(),
				time.Now(),
				interviewer.Profession,
			)
		}

		if err != nil {
			log.Println(err.Error())
			transaction.Rollback()
			sendAPIError(c, "Something went wrong", fiber.StatusInternalServerError)
			return nil
		}

		transaction.Commit()

	}

	response.Status = string(models.STATUS_OK)
	response.Message = "Registered correctly"
	return c.Status(fiber.StatusCreated).JSON(response)
}

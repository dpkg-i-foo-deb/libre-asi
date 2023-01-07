package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"libre-asi-api/models"
	"libre-asi-api/orm"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SignUpService(connection *fiber.Ctx) error {

	decoder := json.NewDecoder(bytes.NewReader(connection.Body()))
	var user models.User

	err := decoder.Decode(&user)

	if err != nil {

		connection.Status(fiber.StatusBadRequest).SendString("Malformed request received")
		return errors.New("bad signup request")

	}

	user.Password, err = hashPassword(user.Password)

	if err != nil {
		connection.Status(fiber.StatusInternalServerError).SendString("Failed to sign up, try again later")
		return errors.New("failed to hash password")
	}

	err = orm.DB.Create(&user).Error

	if err != nil {
		connection.Status(fiber.StatusInternalServerError).SendString("Failed to sign up, try again later")
		return errors.New("failed to create user on database")
	}

	connection.Status(fiber.StatusOK).SendString("Registered Successfully")

	return nil

}

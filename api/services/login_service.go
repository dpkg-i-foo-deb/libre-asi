package services

import (
	"golang.org/x/crypto/bcrypt"
	"libre-asi-api/auth"
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"
)

func LoginService(u models.User, role models.Role) (*models.JWTPair, error) {

	var err error

	switch role {
	case models.ADMINISTRATOR:
		err = loginAdmin(&u)
	case models.INTERVIEWER:
		err = loginInterviewer(&u)
	}

	if err != nil {
		return nil, err
	}

	tk, err := auth.GenerateJWTPair(u.Email, string(role))

	if err != nil {
		return nil, errors.ErrInternalError
	}

	return &tk, nil
}

func loginAdmin(u *models.User) error {

	var user models.User
	var admin models.Administrator

	if database.DB.Where("email = ?", u.Email).First(&user).Error != nil {
		return errors.ErrInvalidCredentials
	}

	if database.DB.Where("user_id = ?", user.ID).First(&admin).Error != nil {
		return errors.ErrInvalidCredentials
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)) != nil {
		return errors.ErrInvalidCredentials
	}

	if user.ResetPassword {
		return errors.ErrrNeedsPasswordReset
	}

	return nil
}

func loginInterviewer(u *models.User) error {
	var user models.User
	var person models.Person
	var inter models.Interviewer

	if database.DB.Where("email = ?", u.Email).First(&user).Error != nil {
		return errors.ErrInvalidCredentials
	}

	if database.DB.Where("user_id = ?", user.ID).First(&person).Error != nil {
		return errors.ErrInvalidCredentials
	}

	if database.DB.Where("person_id = ?", person.ID).First(&inter).Error != nil {
		return errors.ErrInvalidCredentials
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)) != nil {
		return errors.ErrInvalidCredentials
	}

	if user.ResetPassword {
		return errors.ErrrNeedsPasswordReset
	}

	return nil
}

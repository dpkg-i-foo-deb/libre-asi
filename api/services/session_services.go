package services

import (
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"

	"golang.org/x/crypto/bcrypt"
)

func SetPasswordService(credentials models.PasswordChange, email string) error {

	var user models.User

	if database.DB.Where("email = ?", email).First(&user).Error != nil {
		return errors.ErrAccessDenied
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.CurrentPassword)) != nil {
		return errors.ErrAccessDenied
	}

	hashedPassword, err := util.HashPassword(credentials.NewPassword)

	if err != nil {
		return errors.ErrInternalError
	}

	user.Password = hashedPassword
	user.ResetPassword = false

	if database.DB.Save(&user).Error != nil {
		return errors.ErrInternalError
	}

	return nil
}

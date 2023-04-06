package services

import (
	"libre-asi-api/pkg/auth"
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/util"
	"libre-asi-api/pkg/view"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginInterviewer(i view.Interviewer) (*models.JWTPair, *models.PasswordResetTk, error) {

	var queriedUser models.User
	var interviewer models.Interviewer

	if database.DB.Where("email = ?", i.Email).First(&queriedUser).Error != nil {
		return nil, nil, errors.ErrAccessDenied
	}

	if database.DB.Where("user_id = ?", queriedUser.ID).First(&interviewer).Error != nil {
		return nil, nil, errors.ErrAccessDenied
	}

	if err := bcrypt.CompareHashAndPassword([]byte(queriedUser.Password), []byte(i.Password)); err != nil {
		return nil, nil, errors.ErrAccessDenied
	}

	if queriedUser.NeedsPasswordReset {
		token, err := auth.GeneratePasswordResetToken(i.Email)

		if err != nil {
			return nil, nil, errors.ErrInternalError
		}

		return nil, &token, errors.ErrrNeedsPasswordReset
	}

	token, err := auth.GenerateJWTPair(i.Email, string(models.INTERVIEWER))

	if err != nil {
		return nil, nil, errors.ErrInternalError
	}

	return &token, nil, nil

}

func SetInterviewerPassword(email string, credentials models.PasswordChange) error {

	var interviewer models.Interviewer
	var user models.User
	var person models.Person

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrEntityNotFound
		}
		return errors.ErrInternalError
	}

	if err := database.DB.Where("user_id", user.ID).First(&person).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrEntityNotFound
		}
		return errors.ErrInternalError
	}

	if err := database.DB.Where("person_id", person.ID).First(&interviewer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrEntityNotFound
		}
		return errors.ErrInternalError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.CurrentPassword)); err != nil {
		return errors.ErrAccessDenied
	}

	p, err := util.HashPassword(credentials.NewPassword)

	if err != nil {
		return errors.ErrInternalError
	}

	user.Password = p

	if err := database.DB.Save(&user).Error; err != nil {
		return errors.ErrInternalError
	}

	return nil
}

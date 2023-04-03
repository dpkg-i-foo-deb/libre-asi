package services

import (
	"libre-asi-api/auth"
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginInterviewer(i models.Interviewer) (*models.Interviewer, *models.JWTPair, *models.PasswordResetTk, error) {
	var interviewer models.Interviewer

	if database.DB.Where("email = ?", i.Email).First(&interviewer).Error != nil {
		return nil, nil, nil, errors.ErrNoData
	}

	if err := bcrypt.CompareHashAndPassword([]byte(interviewer.Password), []byte(i.Password)); err != nil {
		return nil, nil, nil, errors.ErrAccessDenied
	}

	if interviewer.ResetPassword {
		token, err := auth.GeneratePasswordResetToken(i.Email)

		if err != nil {
			return nil, nil, nil, errors.ErrInternalError
		}

		return &interviewer, nil, &token, nil
	}

	token, err := auth.GenerateJWTPair(i.Email, string(models.INTERVIEWER))

	if err != nil {
		return nil, nil, nil, errors.ErrInternalError
	}

	return &interviewer, &token, nil, nil
}

func SetInterviewerPassword(email string, credentials models.PasswordChange) error {

	var found models.Interviewer

	if err := database.DB.Where("email = ?", email).First(&found).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrNoData
		}
		return errors.ErrInternalError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(credentials.CurrentPassword)); err != nil {
		return errors.ErrAccessDenied
	}

	p, err := util.HashPassword(credentials.NewPassword)

	if err != nil {
		return errors.ErrInternalError
	}

	found.Password = p

	if err := database.DB.Save(&found).Error; err != nil {
		return errors.ErrInternalError
	}

	return nil
}

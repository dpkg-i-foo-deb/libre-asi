package services

import (
	"libre-asi-api/auth"
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"

	"golang.org/x/crypto/bcrypt"
)

func LoginInterviewerService(i models.Interviewer) (*models.Interviewer, *models.JWTPair, *models.PasswordResetTk, error) {
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

package services

import (
	"libre-asi-api/auth"
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"

	"golang.org/x/crypto/bcrypt"
)

func LoginInterviewerService(i models.Interviewer) (error, *models.Interviewer, *models.JWTPair, *models.PasswordResetTk) {
	var interviewer models.Interviewer

	if database.DB.Where("email = ?", i.Email).First(&interviewer).Error != nil {
		return errors.ErrNoData, nil, nil, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(interviewer.Password), []byte(i.Password)); err != nil {
		return errors.ErrAccessDenied, nil, nil, nil
	}

	if interviewer.ResetPassword {
		token, err := auth.GeneratePasswordResetToken(i.Email)

		if err != nil {
			return errors.ErrInternalError, nil, nil, nil
		}

		return nil, &interviewer, nil, &token
	}

	token, err := auth.GenerateJWTPair(i.Email, string(models.INTERVIEWER))

	if err != nil {
		return errors.ErrInternalError, nil, nil, nil
	}

	return nil, &interviewer, &token, nil
}

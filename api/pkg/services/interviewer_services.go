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

func GetInterviewers() ([]view.Interviewer, error) {

	var interviewers []view.Interviewer
	if err := database.DB.
		Joins("LEFT JOIN people ON interviewers.person_id = people.id").
		Joins("LEFT JOIN users ON people.user_id = users.id").
		Select("interviewers.id, users.email,users.username,'', users.needs_password_reset,people.first_name, people.last_name, people.first_surname, people.last_surname, people.birthdate, people.age, people.personal_id,interviewers.rma").
		Find(&interviewers).
		Error; err != nil {
		return nil, errors.ErrInternalError
	}
	return interviewers, nil
}
func GetInterviewer(id uint) (*view.Interviewer, error) {
	var interviewer view.Interviewer
	var dbInterviewer models.Interviewer

	if database.DB.Where("ID = ?", id).First(&dbInterviewer).Error != nil {
		return nil, errors.ErrEntityNotFound
	}

	if err := database.DB.
		Joins("LEFT JOIN people ON interviewer.person_id = people.id").
		Joins("LEFT JOIN users ON people.user_id = users.id").
		Select("interviewers.id, people.first_name, people.last_name, people.first_surname, people.last_surname, people.birthdate, people.age, people.personal_id, users.email, users.username, users.needs_password_reset").
		Where("interviewers.id = ?", id).
		First(&interviewer).
		Error; err != nil {
		return nil, errors.ErrInternalError
	}

	return &interviewer, nil
}
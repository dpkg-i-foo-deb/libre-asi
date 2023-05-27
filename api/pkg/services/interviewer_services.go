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
	var queriedPerson models.Person
	var interviewer models.Interviewer

	if database.DB.Where("email = ?", i.Email).First(&queriedUser).Error != nil {
		return nil, nil, errors.ErrAccessDenied
	}

	if database.DB.Where("user_id = ?", queriedUser.ID).First(&queriedPerson).Error != nil {
		return nil, nil, errors.ErrAccessDenied
	}

	if database.DB.Where("person_id = ?", queriedPerson.ID).First(&interviewer).Error != nil {
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
	user.NeedsPasswordReset = false

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
		Where("interviewers.deleted_at IS NULL").
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
		Joins("LEFT JOIN people ON interviewers.person_id = people.id").
		Joins("LEFT JOIN users ON people.user_id = users.id").
		Select("interviewers.id, people.first_name, people.last_name, people.first_surname, people.last_surname, people.birthdate, people.age, people.personal_id, users.email, users.username, users.needs_password_reset").
		Where("interviewers.id = ? AND interviewers.deleted_at IS NULL", id).
		First(&interviewer).
		Error; err != nil {
		return nil, errors.ErrInternalError
	}

	return &interviewer, nil
}

func RegisterInterviewer(i view.Interviewer) (*view.Interviewer, error) {

	var user models.User
	var person models.Person
	var interviewer models.Interviewer

	if err := database.DB.Where("email = ?", i.Email).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.ErrConflict
		}
	}

	if err := database.DB.Where("personal_id = ?", i.PersonalID).First(&person).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.ErrConflict
		}
	}

	user.Email = i.Email
	p, err := util.MakeRandomPassword()

	if err != nil {
		return nil, errors.ErrInternalError
	}

	hash, err := util.HashPassword(p)

	if err != nil {
		return nil, errors.ErrInternalError
	}

	user.Password = hash
	user.NeedsPasswordReset = true
	user.Username = i.Username
	user.Email = i.Email

	person.FirstName = i.FirstName
	person.PersonalID = i.PersonalID

	interviewer = models.Interviewer{}

	if err := database.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&user).Error; err != nil {
			return err

		}

		person.UserID = user.ID

		if err := tx.Create(&person).Error; err != nil {
			return err

		}

		interviewer.PersonID = person.ID

		if err := tx.Create(&interviewer).Error; err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, errors.ErrInternalError
	}

	i.Password = p

	return &i, nil

}

func DeleteInterviewer(id int) error {

	var i models.Interviewer

	if database.DB.Where("id = ?", id).First(&i).Error != nil {
		return errors.ErrEntityNotFound
	}

	if database.DB.Delete(&i).Error != nil {
		return errors.ErrInternalError
	}

	return nil
}

package services

import (
	"gorm.io/gorm"
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"
)

func RegisterService(user models.User, role models.Role) error {

	switch role {
	case models.ADMINISTRATOR:
		return createAdmin(user, true)
	case models.INTERVIEWER:
		return createInterviewer(user, true)
	default:
		return errors.ErrCheckRequest
	}

}

func createAdmin(u models.User, passwordReset bool) error {

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		u.Administrators = []models.Administrator{
			{},
		}

		pass, err := util.HashPassword(u.Password)

		if err != nil {
			return errors.ErrInternalError
		}

		u.Password = pass

		u.ResetPassword = passwordReset

		if database.DB.Omit("People").Create(&u).Error != nil {
			return errors.ErrInternalError
		}

		return nil
	})

	return err
}

func createInterviewer(u models.User, passwordReset bool) error {

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		if u.People == nil {
			return errors.ErrNoData
		}

		if len(u.People) != 1 {
			return errors.ErrTooManyEntities
		} else {
			if len(u.People[0].Interviewers) > 1 || u.People[0].Interviewers == nil {
				return errors.ErrTooManyEntities
			}
		}

		pass, err := util.HashPassword(u.Password)

		if err != nil {
			return errors.ErrInternalError
		}

		u.Password = pass

		u.ResetPassword = passwordReset

		res := database.DB.Omit("Administators",
			"Patients",
			"Attendants",
			"PersonID",
			"UserID",
			"Interpretations",
			"Reports",
			"ProfessionTranslations").Create(&u)

		if res.Error != nil {
			return errors.ErrInternalError
		}

		return nil
	})

	return err
}

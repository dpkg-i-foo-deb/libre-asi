package services

import (
	"gorm.io/gorm"
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"
)

func GetAdministratorsService() ([]models.User, error) {

	var admins []models.Administrator
	var ids []uint
	var users []models.User

	if database.DB.Find(&admins).Error != nil {
		return nil, errors.ErrInternalError
	}

	for i := range admins {
		ids = append(ids, admins[i].UserID)
	}

	if database.DB.Omit("password", "administrators", "people").Find(&users, ids).Error != nil {
		return nil, errors.ErrInternalError
	}

	return users, nil
}

func RegisterAdministratorService(newUser models.User) (*models.User, error) {

	var queriedUser models.User

	if database.DB.Where("email = ?", newUser.Email).First(&queriedUser).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	if database.DB.Where("username = ?", newUser.Username).First(&queriedUser).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		password, err := util.MakeRandomPassword()

		if err != nil {
			return errors.ErrInternalError
		}

		hashedPassword, err := util.HashPassword(password)

		if err != nil {
			return errors.ErrInternalError
		}

		newUser.Password = hashedPassword

		if database.DB.Omit("Administrators", "People").Create(&newUser).Error != nil {
			return errors.ErrInternalError
		}

		id := newUser.ID

		var newAdmin = models.Administrator{
			UserID: id,
		}

		if database.DB.Create(&newAdmin).Error != nil {
			return errors.ErrInternalError
		}

		newUser.Password = password

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

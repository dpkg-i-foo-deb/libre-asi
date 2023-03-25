package services

import (
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"

	"gorm.io/gorm"
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
		newUser.ResetPassword = true

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

func UpdateAdministratorService(updatedAdmin models.User) error {

	var found models.User

	if err := database.DB.Where("ID = ?", updatedAdmin.ID).First(&found).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrEntityNotFound
		}
		return errors.ErrInternalError
	}

	if err := database.DB.Where("user_id = ?", updatedAdmin.ID).First(&models.Administrator{}).Error; err != nil {
		return errors.ErrBadRoute
	}

	if err := database.DB.Model(&updatedAdmin).Select("email", "username").Updates(&updatedAdmin).Error; err != nil {
		return errors.ErrInternalError
	}

	return nil
}

func DeleteAdministratorService(id int) error {

	var found models.User

	if err := database.DB.Where("ID = ?", id).First(&found).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrEntityNotFound
		}
		return errors.ErrInternalError
	}

	if err := database.DB.Delete(&found).Error; err != nil {
		return errors.ErrInternalError
	}

	return nil
}

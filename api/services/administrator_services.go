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

func LoginAdminService(a models.Administrator) (*models.Administrator, *models.JWTPair, *models.PasswordResetTk, error) {

	var admin models.Administrator

	if database.DB.Where("email = ?", a.Email).First(&admin).Error != nil {
		return nil, nil, nil, errors.ErrNoData
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(a.Password)); err != nil {
		return nil, nil, nil, errors.ErrAccessDenied
	}

	if admin.ResetPassword {
		token, err := auth.GeneratePasswordResetToken(a.Email)

		if err != nil {
			return nil, nil, nil, errors.ErrInternalError
		}

		return &admin, nil, &token, err
	}

	token, err := auth.GenerateJWTPair(a.Email, string(models.ADMINISTRATOR))

	if err != nil {
		return nil, nil, nil, errors.ErrInternalError
	}

	return &admin, &token, nil, nil
}

func GetAdministratorsService() ([]models.Administrator, error) {

	var admins []models.Administrator

	if database.DB.Find(&admins).Error != nil {
		return nil, errors.ErrInternalError
	}

	return admins, nil
}

func RegisterAdministratorService(newAdmin models.Administrator) (*models.Administrator, error) {

	var queriedAdmin models.Administrator

	if database.DB.Where("email = ?", newAdmin.Email).First(&queriedAdmin).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	if database.DB.Where("username = ?", newAdmin.Username).First(&queriedAdmin).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	p, err := util.HashPassword(newAdmin.Password)

	if err != nil {
		return nil, errors.ErrInternalError
	}

	newAdmin.Password = p

	if err := database.DB.Create(&newAdmin).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	return &newAdmin, nil
}

func UpdateAdministratorService(updatedAdmin models.Administrator) error {

	var found models.Administrator

	if err := database.DB.Where("ID = ?", updatedAdmin.ID).First(&found).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrEntityNotFound
		}
		return errors.ErrInternalError
	}

	if err := database.DB.Model(&updatedAdmin).Updates(&updatedAdmin).Error; err != nil {
		return errors.ErrInternalError
	}

	return nil
}

func DeleteAdministratorService(id int) error {

	var found models.Administrator

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

func SetAdministratorPassword(email string, credentials models.PasswordChange) error {

	var found models.Administrator

	if err := database.DB.Where("email = ?", email).First(&found).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrEntityNotFound
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

	if err := database.DB.Model(&found).Updates(&found).Error; err != nil {
		return errors.ErrInternalError
	}

	return nil
}

package services

import (
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/util"

	"gorm.io/gorm"
)

func RegisterPatient(newPatient models.Patient) (*models.Patient, error) {

	var p string
	var found models.Patient
	var err error

	if err = database.DB.Where("email = ?", newPatient.Email).First(&found).Error; err != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	if database.DB.Where("username = ?", newPatient.Username).First(&found).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	p, err = util.MakeRandomPassword()

	if err != nil {
		return nil, errors.ErrInternalError
	}

	newPatient.Password = p

	newPatient.NeedsPasswordReset = true

	newPatient.Password, err = util.HashPassword(newPatient.Password)

	if err != nil {
		return nil, errors.ErrInternalError
	}

	if err = database.DB.Create(&newPatient).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	newPatient.Password = p

	return &newPatient, gorm.ErrNotImplemented
}

func GetPatients() ([]models.Patient, error) {

	var patients []models.Patient

	if database.DB.Find(&patients).Error != nil {
		return nil, errors.ErrInternalError
	}

	return patients, nil
}

package services

import (
	"libre-asi-api/database"
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"libre-asi-api/view"

	"gorm.io/gorm"
)

func RegisterPatient(newPatient view.Patient) (*view.Patient, error) {

	var user models.User
	var person models.Person
	var patient models.Patient

	if database.DB.Where("email = ?", newPatient.Email).First(&user).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	if database.DB.Where("personal_id = ?", newPatient.PersonalID).First(&person).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	if database.DB.Where("social_security_number = ?", newPatient.SocialSecurityNumber).First(&patient).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	user.Email = newPatient.Email
	user.Password = "unhashed"
	user.NeedsPasswordReset = true
	user.Username = newPatient.Username

	person.FirstName = newPatient.FirstName
	person.LastName = newPatient.LastName
	person.FirstSurname = newPatient.FirstSurname
	person.LastSurname = newPatient.LastSurname
	person.Birthdate = newPatient.Birthdate
	person.Age = newPatient.Age
	person.PersonalID = newPatient.PersonalID

	//TODO race and religious preference
	patient.SocialSecurityNumber = newPatient.SocialSecurityNumber
	patient.Interviews = []models.Interview{}

	if err := database.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		person.UserID = user.ID

		if err := tx.Create(&person).Error; err != nil {
			return err
		}

		patient.PersonID = person.ID

		if err := tx.Create(&patient).Error; err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, errors.ErrInternalError
	}

	return &newPatient, nil

}

func GetPatients() ([]view.Patient, error) {

	var patients []view.Patient

	if err := database.DB.
		Joins("LEFT JOIN people ON patients.person_id = people.id").
		Joins("LEFT JOIN users ON people.user_id = users.id").
		Select("patients.id, people.first_name, people.last_name, people.first_surname, people.last_surname, people.birthdate, people.age, people.personal_id, users.email, users.username, users.needs_password_reset, patients.social_security_number").
		Find(&patients).
		Error; err != nil {
		return nil, errors.ErrInternalError
	}

	return patients, nil

}

func GetPatient(id uint) (*view.Patient, error) {

	var patient view.Patient
	var dbPatient models.Patient

	if database.DB.Where("ID = ?", id).First(&dbPatient).Error != nil {
		return nil, errors.ErrEntityNotFound
	}

	if err := database.DB.
		Joins("LEFT JOIN people ON patients.person_id = people.id").
		Joins("LEFT JOIN users ON people.user_id = users.id").
		Select("patients.id, people.first_name, people.last_name, people.first_surname, people.last_surname, people.birthdate, people.age, people.personal_id, users.email, users.username, users.needs_password_reset, patients.social_security_number").
		Where("patients.id = ?", id).
		First(&patient).
		Error; err != nil {
		return nil, errors.ErrInternalError
	}

	return &patient, nil
}

func UpdatePatient(updatedPatient models.Patient) error {

	var found models.Patient

	if err := database.DB.Where("ID = ?", updatedPatient.ID).First(&found).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrEntityNotFound
		}

		return errors.ErrInternalError
	}

	if err := database.DB.Omit("password").Save(updatedPatient).Error; err != nil {
		return errors.ErrInternalError
	}

	return nil
}

func DeletePatient(id uint) error {

	var found models.Patient

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

func SetPatientPicture(id int, imageData []byte) error {

	var patient models.Patient

	var person models.Person

	if err := database.DB.Where("ID = ?", id).First(&patient).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrEntityNotFound
		}

		return errors.ErrInternalError
	}

	if err := database.DB.Where("ID = ?", patient.PersonID).First(&person).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			return errors.ErrEntityNotFound

		}

		return errors.ErrInternalError

	}

	person.Picture = imageData

	if err := database.DB.Save(&person).Error; err != nil {

		return errors.ErrInternalError

	}

	return nil
}

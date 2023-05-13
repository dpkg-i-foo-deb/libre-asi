package test

import (
	"libre-asi-api/pkg/models"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreatePatient(t *testing.T) {
	//Create User
	user := &models.User{
		Email:    "patient@example.com",
		Username: "Patient1",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatalf("Error creating user: %d", err)
	}
	//Create Person
	person := &models.Person{
		FirstName:    "Mauricio",
		LastName:     "",
		FirstSurname: "Marulanda",
		LastSurname:  "Carvajal",
		Birthdate:    time.Date(2003, time.May, 9, 0, 0, 0, 0, time.Local),
		UserID:       user.ID,
	}
	if err := DB.Create(&person).Error; err != nil {
		t.Fatalf("Error creating person: %d", err)
	}

	//create patient
	patient := &models.Patient{
		SocialSecurityNumber: "0123456899",
		PersonID:             person.ID,
	}
	if err := DB.Create(&patient).Error; err != nil {
		t.Fatalf("Error creating patient: %d", err)
	}
}

func TestUpdatePatient(t *testing.T) {
	//Create User
	user := &models.User{
		Email:    "patient2@example.com",
		Username: "Patient2",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatalf("Error creating user: %d", err)
	}
	//Create Person
	person := &models.Person{
		FirstName:    "Fernando",
		LastName:     "",
		FirstSurname: "Galvis",
		LastSurname:  "Bedoya",
		Birthdate:    time.Date(1988, time.December, 7, 0, 0, 0, 0, time.Local),
		UserID:       user.ID,
	}
	if err := DB.Create(&person).Error; err != nil {
		t.Fatalf("Error creating person: %d", err)
	}

	//create patient
	patient := &models.Patient{
		SocialSecurityNumber: "7896543210",
		PersonID:             person.ID,
	}
	if err := DB.Create(&patient).Error; err != nil {
		t.Fatalf("Error creating patient: %d", err)
	}

	// update the Patient
	newPersonID := uint(444)
	patient.PersonID = newPersonID
	if err := DB.Save(patient).Error; err != nil {
		t.Fatal(err)
	}
	// check if the patient was updated successfully
	var result models.Patient
	if err := DB.First(&result, patient.ID).Error; err != nil {
		t.Fatal(err)
	}
	if result.PersonID != newPersonID {
		t.Errorf("expected PersonID=%d, got %d", newPersonID, result.PersonID)
	}

}

func TestDeletePatient(t *testing.T) {

	//create user for Patient
	user := &models.User{
		Email:    "patient3@example.com",
		Username: "patient3",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatal(err)
	}
	//Create Person
	person := &models.Person{
		FirstName:    "Pedro",
		LastName:     "Daniel",
		FirstSurname: "Valencia",
		LastSurname:  "Arias",
		Birthdate:    time.Date(1985, time.February, 2, 0, 0, 0, 0, time.Local),
		UserID:       user.ID,
	}
	if err := DB.Create(person).Error; err != nil {
		t.Fatal(err)
	}
	//Create Patient
	patient := &models.Patient{
		SocialSecurityNumber: "7532146985",
	}

	if err := DB.Create(patient).Error; err != nil {
		t.Fatal(err)
	}

	// delete the patient
	if err := DB.Delete(patient).Error; err != nil {
		t.Fatal(err)
	}

	// check if the patient was deleted successfully
	var result models.Patient
	if err := DB.First(&result, patient.ID).Error; err != gorm.ErrRecordNotFound {
		t.Errorf("expected record not found error, got %v", err)
	}

	// check if the specific patient was deleted
	var count int64
	if err := DB.Model(&models.Patient{}).Where("id = ?", patient.ID).Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("expected 0 records, got %d", count)
	}
}

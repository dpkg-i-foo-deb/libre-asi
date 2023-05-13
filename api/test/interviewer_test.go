package test

import (
	"libre-asi-api/pkg/models"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreateInterviewer(t *testing.T) {

	//Create User
	user := &models.User{
		Email:    "interviewer@example.com",
		Username: "Interviewer1",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatalf("Error creating user: %d", err)
	}
	//Create Person
	person := &models.Person{
		FirstName:    "Laura",
		LastName:     "Camila",
		FirstSurname: "López",
		LastSurname:  "Gutierrez",
		Birthdate:    time.Date(2003, time.May, 9, 0, 0, 0, 0, time.Local),
		UserID:       user.ID,
	}
	if err := DB.Create(&person).Error; err != nil {
		t.Fatalf("Error creating person: %d", err)
	}
	//Create Interviewer
	interviewer := &models.Interviewer{
		RMA:      "1111111",
		PersonID: person.ID,
	}

	if err := DB.Create(&interviewer).Error; err != nil {
		t.Fatalf("Error creating interviewer: %d", err)
	}
}

func TestUpdateInterviewer(t *testing.T) {
	//create user for Interviewer
	user := &models.User{
		Email:    "interviewer2@example.com",
		Username: "Interviewer2",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatal(err)
	}
	//Create Person
	person := &models.Person{
		FirstName:    "Miguel",
		LastName:     "Jose",
		FirstSurname: "Posada",
		LastSurname:  "Rodriguez",
		Birthdate:    time.Date(2001, time.March, 7, 0, 0, 0, 0, time.Local),
		UserID:       user.ID,
	}
	if err := DB.Create(person).Error; err != nil {
		t.Fatal(err)
	}
	//Create Interviewer
	interviewer := &models.Interviewer{
		RMA:      "1234567",
		PersonID: person.ID,
	}

	if err := DB.Create(interviewer).Error; err != nil {
		t.Fatal(err)
	}
	// update the Interviewer
	newFirstName := "Maria"
	person.FirstName = newFirstName
	if err := DB.Save(person).Error; err != nil {
		t.Fatal(err)
	}
	newPersonID := uint(222)
	interviewer.PersonID = newPersonID
	if err := DB.Save(interviewer).Error; err != nil {
		t.Fatal(err)
	}

	// check if the Interviewer was updated successfully
	var result models.Interviewer
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Fatal(err)
	}
	if result.PersonID != newPersonID {
		t.Errorf("expected PersonID=%d, got %d", newPersonID, result.PersonID)
	}

}

func TestDeleteInterviewer(t *testing.T) {

	//create user for Interviewer
	user := &models.User{
		Email:    "interviewer3@example.com",
		Username: "Interviewer3",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatal(err)
	}
	//Create Person
	person := &models.Person{
		FirstName:    "Juan",
		LastName:     "Pablo",
		FirstSurname: "Sepulveda",
		LastSurname:  "Hoyos",
		Birthdate:    time.Date(2007, time.July, 5, 0, 0, 0, 0, time.Local),
		UserID:       user.ID,
	}
	if err := DB.Create(person).Error; err != nil {
		t.Fatal(err)
	}
	//Create Interviewer
	interviewer := &models.Interviewer{
		RMA:      "1234567",
		PersonID: person.ID,
	}

	if err := DB.Create(interviewer).Error; err != nil {
		t.Fatal(err)
	}

	// delete the interviewer
	if err := DB.Delete(interviewer).Error; err != nil {
		t.Fatal(err)
	}

	// check if the interviewer was deleted successfully
	var result models.Interviewer
	if err := DB.First(&result, interviewer.ID).Error; err != gorm.ErrRecordNotFound {
		t.Errorf("expected record not found error, got %v", err)
	}

	// check if the specific interviewer was deleted
	var count int64
	if err := DB.Model(&models.Interviewer{}).Where("id = ?", interviewer.ID).Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("expected 0 records, got %d", count)
	}
}

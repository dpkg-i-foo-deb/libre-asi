package test

import (
	"libre-asi-api/pkg/models"
	"testing"
	"time"
)

func TestCreateInterview(t *testing.T) {
	//Create Interview
	interview := &models.Interview{
		StartDate: time.Date(2023, time.May, 13, 0, 0, 0, 0, time.Local),
		EndDate:   time.Date(2003, time.May, 13, 0, 0, 0, 0, time.Local),
	}
	if err := DB.Create(&interview).Error; err != nil {
		t.Fatalf("Error creating interview: %d", err)
	}
}

func TestDoInterview(t *testing.T) {

	//Create User for Interviewer
	user := &models.User{
		Email:    "interviewer4@example.com",
		Username: "Interviewer4",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatalf("Error creating user: %d", err)
	}
	//Create Person
	person := &models.Person{
		FirstName:    "German",
		LastName:     "",
		FirstSurname: "Vargas",
		LastSurname:  "Bedoya",
		Birthdate:    time.Date(2001, time.November, 9, 0, 0, 0, 0, time.Local),
		UserID:       user.ID,
	}
	if err := DB.Create(&person).Error; err != nil {
		t.Fatalf("Error creating person: %d", err)
	}
	//Create Interviewer
	interviewer := &models.Interviewer{
		RMA:      "7412589630",
		PersonID: person.ID,
	}

	if err := DB.Create(&interviewer).Error; err != nil {
		t.Fatalf("Error creating interviewer: %d", err)
	}
	interviewerList := []models.Interviewer{*interviewer}

	//Create User 2 for patient
	user2 := &models.User{
		Email:    "patient4@example.com",
		Username: "Patient4",
		Password: "password",
	}
	if err := DB.Create(user2).Error; err != nil {
		t.Fatalf("Error creating user: %d", err)
	}
	//Create Person for patient
	person2 := &models.Person{
		FirstName:    "Cristian",
		LastName:     "Stiven",
		FirstSurname: "Hernandez",
		LastSurname:  "Meneses",
		Birthdate:    time.Date(1988, time.April, 28, 0, 0, 0, 0, time.Local),
		UserID:       user2.ID,
	}
	if err := DB.Create(&person2).Error; err != nil {
		t.Fatalf("Error creating person: %d", err)
	}

	//create patient
	patient := &models.Patient{
		SocialSecurityNumber: "9513574613",
		PersonID:             person2.ID,
	}
	if err := DB.Create(&patient).Error; err != nil {
		t.Fatalf("Error creating patient: %d", err)
	}
	//Create Question for interview
	question := &models.Question{
		Order:       12,
		SpecialCode: "I12",
	}
	if err := DB.Create(&question).Error; err != nil {
		t.Fatalf("Error creating question: %d", err)
	}
	//Create Options for Interview
	option := &models.Option{
		Description: "Casado",
		Value:       1,
		QuestionID:  question.ID,
	}
	if err := DB.Create(&option).Error; err != nil {
		t.Fatalf("Error creating option: %d", err)
	}
	option2 := &models.Option{
		Description: "Soltero",
		Value:       2,
		QuestionID:  question.ID,
	}
	optionList := []models.Option{*option, *option2}

	//Create Interview
	interview := &models.Interview{
		StartDate:       time.Date(2023, time.May, 14, 0, 0, 0, 0, time.Local),
		EndDate:         time.Date(2003, time.May, 14, 0, 0, 0, 0, time.Local),
		PatientID:       patient.ID,
		Interviewers:    interviewerList,
		CurrentQuestion: question.SpecialCode,
		Answers:         optionList,
	}
	if err := DB.Create(&interview).Error; err != nil {
		t.Fatalf("Error creating interview: %d", err)
	}

	//Create answers for Interview
	answer := &models.InterviewAnswers{
		Answer:      "Cuál es su estado civil actual?",
		InterviewID: interview.ID,
	}
	if err := DB.Create(&answer).Error; err != nil {
		t.Fatalf("Error creating answer: %d", err)
	}

}

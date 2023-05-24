package test

import (
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/services"
	"libre-asi-api/pkg/view"
	"testing"
)

func TestRegisterInterviewer(t *testing.T) {

	testCases := []struct {
		name  string
		input view.Interviewer
	}{
		{
			name: "Registration case 1",
			input: view.Interviewer{
				Email:      "Interviewer11@example.com",
				PersonalID: "1234567890",
				FirstName:  "John",
				Username:   "john123",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create User
			var user models.User
			user.Email = tc.input.Email
			user.Username = tc.input.Username
			database.DB.Create(&user)

			// Create Person
			var person models.Person
			person.FirstName = tc.input.FirstName
			person.PersonalID = tc.input.PersonalID
			person.UserID = user.ID
			database.DB.Create(&person)

			// Create Interviewer
			var interviewer models.Interviewer
			interviewer.PersonID = person.ID
			database.DB.Create(&interviewer)

			// Run the function
			_, err := services.RegisterInterviewer(tc.input)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestGetInterviewers(t *testing.T) {

	interviewerData := []view.Interviewer{
		{
			Email:      "Interviewer12@example.com",
			PersonalID: "1234567894",
			FirstName:  "Javier",
			Username:   "javi555",
		},
		{
			Email:      "Interviewer13@example.com",
			PersonalID: "0987654321",
			FirstName:  "Jane",
			Username:   "jane123",
		},
	}

	for _, interviewer := range interviewerData {

		_, err := services.RegisterInterviewer(interviewer)
		if err != nil {
			t.Fatalf("Error setting up test data: %v", err)
		}
	}

	interviewers, err := services.GetInterviewers()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(interviewers) != 3 {
		t.Errorf("Expected %d interviewers, got %d", len(interviewerData), len(interviewers))
	}
}

func TestGetInterviewer(t *testing.T) {

	testID := uint(1)

	interviewer, err := services.GetInterviewer(testID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if interviewer == nil {
		t.Errorf("Expected a non-nil interviewer, got nil")
	}
}

func TestDeleteInterviewer(t *testing.T) {

	interviewerID := 1

	err := services.DeleteInterviewer(interviewerID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, err = services.GetInterviewer(uint(interviewerID))
	if err == nil {
		t.Errorf("Expected an error when retrieving a deleted interviewer, but got none")
	}
}

// func TestLoginInterviewer(t *testing.T) {

// 	interviewerData := view.Interviewer{
// 		Email:    "Interviewer12@example.com",
// 		Username: "javi555",
// 	}

// 	_, _, err := services.LoginInterviewer(interviewerData)
// 	if err != nil {
// 		t.Fatalf("Unexpected error: %v", err)
// 	}

// }

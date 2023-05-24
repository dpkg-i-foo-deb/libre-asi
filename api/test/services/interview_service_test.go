package test

// func TestStartInterview(t *testing.T) {
// 	asiForm := models.AsiForm{}
// 	if err := database.DB.First(&asiForm).Error; err != nil {
// 		t.Fatalf("Failed to get ASI form: %v", err)
// 	}

// 	interviewer := models.Interviewer{}
// 	if err := database.DB.First(&interviewer).Error; err != nil {
// 		t.Fatalf("Failed to get interviewer: %v", err)
// 	}

// 	patient := models.Patient{}
// 	if err := database.DB.Create(&patient).Error; err != nil {
// 		t.Fatalf("Failed to create patient: %v", err)
// 	}

// 	interview, err := services.StartInterview(patient.ID, interviewer.ID)

// 	if err != nil {
// 		t.Fatalf("Failed to start interview: %v", err)
// 	}

// 	if interview == nil {
// 		t.Fatal("Expected interview to be created, got nil")
// 	}

// 	if interview.StartDate.IsZero() {
// 		t.Fatal("Expected interview StartDate to be set")
// 	}

// 	if interview.PatientID != patient.ID {
// 		t.Fatalf("Expected interview PatientID to be %d, got %d", patient.ID, interview.PatientID)
// 	}

// 	if interview.InterviewerID != interviewer.ID {
// 		t.Fatalf("Expected interview InterviewerID to be %d, got %d", interviewer.ID, interview.InterviewerID)
// 	}

// 	if interview.AsiFormID != asiForm.ID {
// 		t.Fatalf("Expected interview AsiFormID to be %d, got %d", asiForm.ID, interview.AsiFormID)
// 	}

// 	if err := database.DB.Delete(&patient).Error; err != nil {
// 		t.Fatalf("Failed to delete patient: %v", err)
// 	}
// }

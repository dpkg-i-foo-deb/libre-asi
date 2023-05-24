package test

import (
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/services"
	"libre-asi-api/pkg/view"
	"testing"
	"time"
)

func TestRegisterPatient(t *testing.T) {
	// Define the patient data
	newPatientData := view.Patient{
		Email:                "Patient12@example.com",
		FirstName:            "Miguel",
		LastName:             "Mauricio",
		FirstSurname:         "Arbelaez",
		LastSurname:          "Toro",
		Birthdate:            time.Now(),
		Age:                  25,
		PersonalID:           "1234567899",
		SocialSecurityNumber: "1122334454",
	}

	// Now run RegisterPatient
	_, err := services.RegisterPatient(newPatientData)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestGetPatients(t *testing.T) {

	var count int64
	database.DB.Model(&models.Patient{}).Where("deleted_at is NULL").Count(&count)

	patients, err := services.GetPatients()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Compara la longitud del slice de pacientes devuelto con el conteo de la base de datos
	if int64(len(patients)) != count {
		t.Errorf("Expected %d patients, got %d", count, len(patients))
	}
}

func TestGetPatient(t *testing.T) {

	existingPatientID := uint(1)

	patient, err := services.GetPatient(existingPatientID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if patient == nil {
		t.Errorf("Expected a non-nil patient, got nil")
	} else {
		t.Logf("Patient: %+v", patient)
	}
}

func TestUpdatePatient(t *testing.T) {

	// Obtén un paciente existente de la base de datos
	existingPatientID := uint(1)
	existingPatient, err := services.GetPatient(existingPatientID)
	if err != nil {
		t.Fatalf("Error retrieving existing patient: %v", err)
	}

	existingPatient.FirstName = "UpdatedName"

	err = services.UpdatePatient(*existingPatient)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	updatedPatient, err := services.GetPatient(existingPatientID)
	if err != nil {
		t.Fatalf("Error retrieving updated patient: %v", err)
	}

	if updatedPatient.FirstName != "UpdatedName" {
		t.Errorf("Expected FirstName to be 'UpdatedName', got '%s'", updatedPatient.FirstName)
	}

}

func TestDeletePatient(t *testing.T) {
	// ID del paciente existente en la base de datos
	existingPatientID := uint(1)

	// Ejecuta DeletePatient con el ID del paciente existente
	err := services.DeletePatient(existingPatientID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Opcionalmente, puedes intentar obtener el paciente eliminado de la base de datos
	deletedPatient, err := services.GetPatient(existingPatientID)
	if err == nil && deletedPatient != nil {
		t.Errorf("Expected patient to be deleted, but found patient with ID %d", deletedPatient.ID)
	}
}

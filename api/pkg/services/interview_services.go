package services

import (
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/models"
	"time"
)

func StartInterview(patientID int, interviewerID int) (int, error) {

	i := models.Interview{}

	asiForm := models.AsiForm{}

	if err := database.DB.First(&asiForm).Error; err != nil {
		return -1, errors.ErrInternalError
	}

	i.AsiFormID = asiForm.ID

	interviewer := models.Interviewer{}

	if err := database.DB.Where("id = ?", interviewerID).First(&interviewer).Error; err != nil {
		return -1, errors.ErrEntityNotFound
	}

	patient := models.Patient{}

	if err := database.DB.Where("id = ?", patientID).First(&patient).Error; err != nil {
		return -1, errors.ErrEntityNotFound
	}

	i.PatientID = patient.ID

	i.Interviewers = []models.Interviewer{interviewer}

	i.StartDate = time.Now()

	i.Language = "ES"

	return int(i.ID), nil
}

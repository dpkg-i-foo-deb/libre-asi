package view

import "time"

type Interview struct {
	ID              uint      `json:"id"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate"`
	PauseAt         time.Time `json:"pausedAt"`
	ResumedAt       time.Time `json:"resumedAt"`
	PatientID       uint      `json:"patient"`
	InterviewerID   uint      `json:"interviewer"`
	AsiFormID       uint      `json:"AsiForm"`
	CurrentQuestion string    `json:"currentQuestion"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Interview struct {
	gorm.Model
	StartDate       time.Time     `json:"startDate"`
	EndDate         time.Time     `json:"endDate"`
	PausedAt        time.Time     `json:"pausedAt"`
	ResumedAt       time.Time     `json:"resumedAt"`
	PatientID       uint          `json:"patient"`
	Interviewers    []Interviewer `json:"interviewer" gorm:"many2many:interviewer_interviews"`
	Attendants      []Attendant   `json:"attendants" gorm:"many2many:attendants_interviews"`
	Answers         []Option      `json:"answers" gorm:"many2many:interview_answers"`
	AsiFormID       uint          `json:"AsiForm"`
	Language        string        `json:"language"`
	CurrentSection  string        `json:"currentSection"`
	CurrentQuestion string        `json:"currentQuestion"`
}

type InterviewAnswers struct {
	InterviewID uint `json:"interview" gorm:"primaryKey"`
	QuestionID  uint `json:"questionID" gorm:"primaryKey"`
	OptionID    uint `json:"option" gorm:"primaryKey"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Answer      int    `json:"answer"`
	Commentary  string `json:"commentary"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Interviewer struct {
	gorm.Model
	//TODO add person fields
	User
	Professions     []Profession `gorm:"many2many:interviewer_professions;" json:"professions"`
	Interpretations []Interview  `json:"interpretations" gorm:"many2many:interview_interpretations"`
	Reports         []Interview  `json:"reports" gorm:"many2many:interview_reports"`
	RMA             string       `json:"rma"`
}

type InterviewInterpretations struct {
	InterviewID    uint `json:"interview" gorm:"primaryKey"`
	InterviewerID  uint `json:"interviewer" gorm:"primaryKey"`
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Interpretation string `json:"interpretation"`
}

type InterviewReports struct {
	InterviewID   uint      `json:"interview" gorm:"primaryKey"`
	InterviewerID uint      `json:"interviewer" gorm:"primaryKey"`
	CreatedAt     time.Time `gorm:"primaryKey"`
	DeletedAt     gorm.DeletedAt
	Text          string `json:"text"`
}

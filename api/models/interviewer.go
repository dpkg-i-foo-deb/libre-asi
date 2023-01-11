package models

import (
	"time"

	"gorm.io/gorm"
)

type Interviewer struct {
	gorm.Model
	PersonID        uint         `json:"person"`
	Professions     []Profession `gorm:"many2many:interviewer_professions;" json:"professions"`
	Interpretations []Interview  `json:"interpretations" gorm:"many2many:interview_interpretations"`
	RMA             string       `json:"rma"`
}

type InterviewInterpretations struct {
	InterviewID    uint `json:"interview" gorm:"primaryKey"`
	InterviewerID  uint `json:"interviewer" gorm:"primaryKey"`
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Interpretation string `json:"interpretation"`
}

package models

import "gorm.io/gorm"

type Interviewer struct {
	gorm.Model
	PersonID    uint         `json:"person"`
	Professions []Profession `gorm:"many2many:interviewer_professions;" json:"professions"`
	RMA         string       `json:"rma"`
}

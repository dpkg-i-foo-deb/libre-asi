package models

import "gorm.io/gorm"

type QuestionType struct {
	gorm.Model
	Type      string     `json:"type"`
	Questions []Question `json:"questions"`
}

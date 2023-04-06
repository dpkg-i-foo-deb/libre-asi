package models

import "gorm.io/gorm"

type QuestionTranslations struct {
	gorm.Model
	Language            string `json:"language"`
	QuestionID          uint   `json:"question"`
	IsDefault           bool   `json:"isDefault"`
	Statement           string `json:"statement"`
	SimplifiedStatement string `json:"simplifiedStatement"`
}

package models

import "gorm.io/gorm"

type QuestionCategoryTranslations struct {
	gorm.Model
	LanguageID         uint   `json:"language"`
	QuestionCategoryID uint   `json:"questionCategory"`
	IsDefault          bool   `json:"isDefault"`
	Category           string `json:"category"`
	Description        string `json:"description"`
}

package models

import "gorm.io/gorm"

type AsiFormTranslations struct {
	gorm.Model
	Language    string `json:"language"`
	AsiFormID   uint   `json:"asiForm"`
	IsDefault   bool   `json:"isDefault"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

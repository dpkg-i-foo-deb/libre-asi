package models

import "gorm.io/gorm"

type ProfessionTranslations struct {
	gorm.Model
	LanguageID   uint   `json:"language"`
	ProfessionID uint   `json:"professionId"`
	IsDefault    bool   `json:"isDefault"`
	Profession   string `json:"profession"`
}

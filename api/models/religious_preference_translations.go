package models

import "gorm.io/gorm"

type ReligiousPreferenceTranslations struct {
	gorm.Model
	LanguageID            uint   `json:"language"`
	ReligiousPreferenceID uint   `json:"religiousPreference"`
	IsDefault             bool   `json:"isDefault"`
	Preference            string `json:"preference"`
}

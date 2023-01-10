package models

import "gorm.io/gorm"

type CountryTranslations struct {
	gorm.Model
	LanguageID  uint   `json:"language"`
	CountryID   uint   `json:"country"`
	IsDefault   bool   `json:"isDefault"`
	CountryName string `json:"countryName"`
}

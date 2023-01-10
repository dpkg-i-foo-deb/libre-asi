package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Code                            string `json:"code" gorm:"unique"`
	Description                     string `json:"description"`
	RegionTranslations              []RegionTranslations
	CountryTranslations             []CountryTranslations
	StateTranslations               []StateTranslations
	ReligiousPreferenceTranslations []ReligiousPreferenceTranslations
	RaceTranslations                []RaceTranslations
	ProfessionTranslations          []ProfessionTranslations
	AsiFormTranslations             []AsiFormTranslations
}

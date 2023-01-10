package models

import "gorm.io/gorm"

type ReligiousPreference struct {
	gorm.Model
	Preference             string                            `gorm:"unique not null"`
	PreferenceTranslations []ReligiousPreferenceTranslations `json:"preferenceTranslations"`
}

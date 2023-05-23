package models

import "gorm.io/gorm"

type ReligiousPreference struct {
	gorm.Model
	Preference string `gorm:"unique not null"`
	//Patients               []Patient
	PreferenceTranslations []ReligiousPreferenceTranslations `json:"preferenceTranslations"`
}

package models

import "gorm.io/gorm"

type Race struct {
	gorm.Model
	Race             string `json:"race"`
	Patients         []Patient
	RaceTranslations []RaceTranslations `json:"raceTranslations"`
}

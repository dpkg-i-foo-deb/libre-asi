package models

import "gorm.io/gorm"

type Race struct {
	gorm.Model
	Race             string             `json:"race"`
	RaceTranslations []RaceTranslations `json:"raceTranslations"`
}

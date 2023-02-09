package models

import "gorm.io/gorm"

type RaceTranslations struct {
	gorm.Model
	Language  string `json:"language"`
	RaceID    uint   `json:"raceId"`
	IsDefault bool   `json:"isDefault"`
	Race      string `json:"race"`
}

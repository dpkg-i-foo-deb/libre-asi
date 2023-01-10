package models

import "gorm.io/gorm"

type RaceTranslations struct {
	gorm.Model
	LanguageID uint   `json:"language"`
	RaceID     uint   `json:"raceId"`
	IsDefault  bool   `json:"isDefault"`
	Race       string `json:"race"`
}

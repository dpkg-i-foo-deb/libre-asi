package models

import "gorm.io/gorm"

type AsiForm struct {
	gorm.Model
	Edition             int                   `json:"edition"`
	Version             string                `json:"version"`
	AsiFormTranslations []AsiFormTranslations `json:"translations"`
}

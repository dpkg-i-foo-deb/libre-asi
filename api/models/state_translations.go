package models

import "gorm.io/gorm"

type StateTranslations struct {
	gorm.Model
	LanguageID uint   `json:"language"`
	StateID    uint   `json:"state"`
	StateName  string `json:"stateName"`
}

package models

import "gorm.io/gorm"

type State struct {
	gorm.Model
	Code              string              `json:"code"`
	CountryID         uint                `json:"country"`
	StateTranslations []StateTranslations `json:"stateTranslations"`
	Cities            []City              `json:"cities"`
}

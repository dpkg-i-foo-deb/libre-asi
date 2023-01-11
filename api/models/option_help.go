package models

import "gorm.io/gorm"

type OptionHelp struct {
	gorm.Model
	OptionID   uint   `json:"option"`
	LanguageID uint   `json:"language"`
	Order      int    `json:"order"`
	Example    string `json:"example"`
	Help       string `json:"help"`
}

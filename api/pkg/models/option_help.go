package models

import "gorm.io/gorm"

type OptionHelp struct {
	gorm.Model
	OptionID uint   `json:"option"`
	Language string `json:"language"`
	Order    int    `json:"order"`
	Example  string `json:"example"`
	Help     string `json:"help"`
}

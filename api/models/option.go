package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	QuestionID            uint         `json:"question"`
	LanguageID            uint         `json:"language"`
	Order                 int          `json:"order"`
	Description           string       `json:"description"`
	SimplifiedDescription string       `json:"simplifiedDescription"`
	Help                  []OptionHelp `json:"help"`
}

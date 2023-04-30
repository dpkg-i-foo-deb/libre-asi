package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	QuestionID            uint         `json:"question"`
	Language              string       `json:"language"`
	Order                 int          `json:"order"`
	Description           string       `json:"description"`
	SimplifiedDescription string       `json:"simplifiedDescription"`
	Help                  []OptionHelp `json:"help"`
}

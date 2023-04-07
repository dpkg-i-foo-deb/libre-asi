package models

import "gorm.io/gorm"

type Profession struct {
	gorm.Model
	Profession             string                   `json:"profession"`
	ProfessionTranslations []ProfessionTranslations `json:"professionTranslations"`
}

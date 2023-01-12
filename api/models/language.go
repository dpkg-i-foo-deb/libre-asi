package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Code                            string `json:"code" gorm:"unique"`
	Description                     string `json:"description"`
	ReligiousPreferenceTranslations []ReligiousPreferenceTranslations
	RaceTranslations                []RaceTranslations
	ProfessionTranslations          []ProfessionTranslations
	AsiFormTranslations             []AsiFormTranslations
	QuestionCategoryTranslations    []QuestionCategoryTranslations
	QuestionTranslations            []QuestionTranslations
	Options                         []Option
	OptionHelp                      []OptionHelp
	Interviews                      []Interview
}

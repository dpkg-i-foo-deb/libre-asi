package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	QuestionCategoryID   uint                   `json:"questionCategory"`
	Order                int                    `json:"order"`
	SpecialCode          string                 `json:"specialCode"`
	QuestionTranslations []QuestionTranslations `json:"translations"`
	Options              []Option               `json:"options"`
	QuestionTypeID       uint                   `json:"type"`
	AsiFormID            uint                   `json:"asiForm"`
	DependencyID         *uint
	Dependencies         []Question `json:"dependencies" gorm:"foreignKey:DependencyID"`
}

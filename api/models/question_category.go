package models

import "gorm.io/gorm"

type QuestionCategory struct {
	gorm.Model
	Category                     string                         `json:"category"`
	QuestionCategoryTranslations []QuestionCategoryTranslations `json:"translations"`
}

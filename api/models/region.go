package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	Code               string               `json:"code" gorm:"unique"`
	RegionTranslations []RegionTranslations `json:"regionTranslations"`
	Countries          []Country            `json:"countries"`
}

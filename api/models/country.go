package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Code                string                `json:"code"`
	CountryTranslations []CountryTranslations `json:"countryTranslations"`
	RegionID            uint                  `json:"region"`
	States              []State               `json:"states"`
}

package models

import "gorm.io/gorm"

type RegionTranslations struct {
	gorm.Model
	LanguageID uint   `json:"language"`
	RegionID   uint   `json:"region"`
	IsDefault  bool   `json:"isDefault"`
	RegionName string `json:"regionName"`
}

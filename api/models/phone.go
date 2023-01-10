package models

import "gorm.io/gorm"

type Phone struct {
	gorm.Model
	Number   uint `json:"number"`
	PersonID uint `json:"person"`
}

package models

import "gorm.io/gorm"

type Attendant struct {
	gorm.Model
	PersonID uint `json:"person"`
}

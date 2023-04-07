package models

import "gorm.io/gorm"

type Administrator struct {
	gorm.Model
	UserID uint `json:"userID" gorm:"not null;default:null"`
}

package models

import "gorm.io/gorm"

type Administrator struct {
	gorm.Model
	User
}

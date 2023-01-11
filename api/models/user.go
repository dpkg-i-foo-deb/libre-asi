package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `json:"email" gorm:"unique;not null; default:null"`
	Username       string `json:"username" gorm:"not null;default:null"`
	Password       string `json:"password" gorm:"not null;default:null"`
	Administrators []Administrator
	People         []Person
}

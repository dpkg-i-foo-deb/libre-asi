package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `json:"email" gorm:"unique"`
	Username       string `json:"username"`
	Password       string `json:"password" gorm:"not null;default:null"`
	Administrators []Administrator
	People         []Person
}

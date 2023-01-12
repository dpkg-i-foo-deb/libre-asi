package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string          `json:"email" gorm:"unique;not null; default:null"`
	Username       string          `json:"username" gorm:"not null;default:null;unique"`
	Password       string          `json:"password" gorm:"not null;default:null"`
	Administrators []Administrator `json:"administrators"`
	People         []Person        `json:"people"`
}

type Role string

const (
	ADMINISTRATOR Role = "admin"
	INTERVIEWER   Role = "interviewer"
)

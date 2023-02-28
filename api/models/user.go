package models

import "gorm.io/gorm"

type Role string

const (
	ADMINISTRATOR Role = "admin"
	INTERVIEWER   Role = "interviewer"
	NONE          Role = "none"
)

type User struct {
	gorm.Model
	Email          string          `json:"email" gorm:"unique;not null; default:null"`
	Username       string          `json:"username" gorm:"not null;default:null;unique"`
	Password       string          `json:"password" gorm:"not null;default:null"`
	ResetPassword  bool            `json:"reset_password" gorm:"not null;default:false"`
	Administrators []Administrator `json:"administrators"`
	People         []Person        `json:"people"`
}

type PasswordChange struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

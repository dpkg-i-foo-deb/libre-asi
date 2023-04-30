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
	Email              string          `json:"email"`
	Username           string          `json:"username" gorm:"not null;default:null;unique"`
	Password           string          `json:"password" gorm:"not null;default:null"`
	NeedsPasswordReset bool            `json:"resetPassword" gorm:"not null;default:false"`
	Administrators     []Administrator `json:"administrators" gorm:"foreignKey:UserID"`
	People             []Person        `json:"people" gorm:"foreignKey:UserID"`
}

type PasswordChange struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

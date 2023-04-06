package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName    string        `json:"firstName" gorm:"not null;default:null"`
	LastName     string        `json:"lastName"`
	FirstSurname string        `json:"firstSurname"`
	LastSurname  string        `json:"lastSurname"`
	Birthdate    time.Time     `json:"birthdate"`
	Age          int           `json:"age"`
	PersonalID   string        `json:"personalID"`
	Addresses    []Address     `json:"addresses"`
	Phones       []Phone       `json:"phones" gorm:"foreignkey:PersonID;association_foreignkey:ID"`
	Interviewers []Interviewer `json:"interviewers" gorm:"foreignKey:PersonID"`
	Patients     []Patient     `json:"patients" gorm:"foreignKey:PersonID"`
	UserID       uint          `json:"userID" gorm:"not null;default:null"`
}

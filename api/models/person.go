package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName    string    `json:"firstName" gorm:"not null;default:null"`
	LastName     string    `json:"lastName"`
	FirstSurname string    `json:"firstSurname"`
	LastSurname  string    `json:"lastSurname"`
	Birthdate    time.Time `json:"birthdate"`
	Age          int       `json:"age"`
	UserID       uint      `json:"user"`
	Addresses    []Address `json:"addresses"`
	Phones       []Phone   `json:"phones"`
}

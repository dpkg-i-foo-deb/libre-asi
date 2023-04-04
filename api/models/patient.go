package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	SocialSecurityNumber  string `json:"socialSecurityNumber"`
	ReligiousPreferenceID uint   `json:"religiousPreference"`
	RaceID                uint   `json:"race"`
	Person
	User
	Interviews []Interview `json:"interviews"`
}

package view

import "time"

type Patient struct {
	//TODO add the rest of the fields
	ID                      uint      `json:"id"`
	FirstName               string    `json:"firstName"`
	LastName                string    `json:"lastName"`
	FirstSurname            string    `json:"firstSurname"`
	LastSurname             string    `json:"lastSurname"`
	Birthdate               time.Time `json:"birthdate"`
	Age                     int       `json:"age"`
	PersonalID              string    `json:"personalID"`
	Username                string    `json:"username"`
	Email                   string    `json:"email"`
	NeedsPasswordReset      bool      `json:"needsPasswordReset"`
	ReligiousPreferenceName string    `json:"religiousPreferenceName"`
	RaceName                string    `json:"raceName"`
	SocialSecurityNumber    string    `json:"socialSecurityNumber"`
}

package view

type Interviewer struct {
	//TODO Interpretations, interviews and professions.
	ID                 uint      `json:"id"`
	Email              string    `json:"email"`
	Username           string    `json:"username"`
	Password           string    `json:"password"`
	NeedsPasswordReset bool      `json:"needsPasswordReset"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	FirstSurname       string    `json:"firstSurname"`
	LastSurname        string    `json:"lastSurname"`
	Birthdate          string    `json:"birthdate"`
	Age                int       `json:"age"`
	PersonalID         string    `json:"personalID"`
	Addresses          []Address `json:"addresses"`
	Phones             []Phone   `json:"phones"`
}

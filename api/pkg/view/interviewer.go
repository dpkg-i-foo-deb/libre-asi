package view

type Interviewer struct {
	ID                 uint   `json:"ID"`
	Email              string `json:"email"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	NeedsPasswordReset bool   `json:"needsPasswordReset"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	FirstSurname       string `json:"firstSurname"`
	LastSurname        string `json:"lastSurname"`
	Birthdate          string `json:"birthdate"`
	Age                int    `json:"age"`
	PersonalID         string `json:"personalID"`
	RMA                string `json:"rma"`
}

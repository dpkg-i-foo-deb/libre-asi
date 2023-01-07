package models

type Role string

const (
	ADMINISTRATOR   Role = "ADMIN"
	INTERVIEWER     Role = "INTERVIEWER"
	PATIENT_MANAGER Role = "PATIENT_MGR"
)

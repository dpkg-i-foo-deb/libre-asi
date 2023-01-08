package models

import (
	"database/sql"
	"time"
)

type Person struct {
	ID            int64          `json:"id"`
	FistName      *string        `json:"first_name"`
	SecondName    *string        `json:"second_name"`
	FirstSurname  *string        `json:"first_surname"`
	SecondSurname *string        `json:"second_surname"`
	Age           int            `json:"age"`
	Birthdate     time.Time      `json:"birhdate"`
	PersonalID    sql.NullString `json:"personal_id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

package models

import "time"

type Interviewer struct {
	ID         int64     `json:"id"`
	RMA        *string   `json:"rma"`
	Profession *string   `json:"profession"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

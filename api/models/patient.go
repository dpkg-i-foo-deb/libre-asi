package models

import (
	"time"
)

type Patient struct {
	ID                   int64     `json:"id"`
	SocialSecurityNumber *string   `json:"social_security_number"`
	Race                 *string   `json:"race"`
	ReligiousPreference  *string   `json:"religious_preference"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

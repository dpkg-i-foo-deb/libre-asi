// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameReligiousPreference = "religious_preference"

// ReligiousPreference mapped from table <religious_preference>
type ReligiousPreference struct {
	Preference string    `gorm:"column:preference;primaryKey" json:"preference"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ReligiousPreference's table name
func (*ReligiousPreference) TableName() string {
	return TableNameReligiousPreference
}

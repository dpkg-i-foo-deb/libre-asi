// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameRace = "race"

// Race mapped from table <race>
type Race struct {
	Race      string    `gorm:"column:race;primaryKey" json:"race"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName Race's table name
func (*Race) TableName() string {
	return TableNameRace
}

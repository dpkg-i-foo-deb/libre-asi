package models

import (
	"time"
)

type Administrator struct {
	ID        int32     `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

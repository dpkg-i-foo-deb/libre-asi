package models

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email     *string   `gorm:"column:email;not null" json:"email"`
	Username  *string   `gorm:"column:username;not null" json:"username"`
	Password  *string   `gorm:"column:password;not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

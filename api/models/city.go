package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Name      string `json:"name"`
	StateID   uint   `json:"state"`
	Addresses []Address
}

package models

import "gorm.io/gorm"

type AddressType struct {
	gorm.Model
	Type      string `json:"type" gorm:"unique"`
	Addresses []Address
}

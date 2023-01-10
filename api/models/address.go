package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Address       string `json:"address"`
	PostalCode    string `json:"postalCode"`
	AddressTypeID uint   `json:"type"`
	CityID        uint   `json:"city"`
}

package models

type City struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	StateID   uint   `json:"state"`
	Addresses []Address
}

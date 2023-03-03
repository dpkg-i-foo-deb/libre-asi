package models

// Adapted to use https://github.com/dr5hn/countries-states-cities-database
type City struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	StateID   uint   `json:"state"`
	Addresses []Address
}

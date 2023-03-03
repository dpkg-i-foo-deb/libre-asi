package models

// Adapted to use https://github.com/dr5hn/countries-states-cities-database
type State struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `json:"name"`
	StateCode string  `json:"state_code"`
	Type      *string `json:"type"`
	CountryID uint    `json:"country_id"`
	Cities    []City  `json:"cities"`
}

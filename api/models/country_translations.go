package models

import "gorm.io/gorm"

// Adapted to use https://github.com/dr5hn/countries-states-cities-database
type CountryTranslations struct {
	gorm.Model
	CountryID uint   `json:"country"`
	Kr        string `json:"kr"`
	Pt_BR     string `json:"pt-BR"`
	Pt        string `json:"pt"`
	Nl        string `json:"nl"`
	Hr        string `json:"hr"`
	Fa        string `json:"fa"`
	De        string `json:"de"`
	Es        string `json:"es"`
	Fr        string `json:"fr"`
	Ja        string `json:"ja"`
	It        string `json:"it"`
	Cn        string `json:"cn"`
	Tr        string `json:"tr"`
}

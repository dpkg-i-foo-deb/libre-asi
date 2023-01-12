package models

//Adapted to use data from https://github.com/dr5hn/countries-states-cities-database
type Country struct {
	ID           uint                `gorm:"primaryKey" json:"id"`
	Name         string              `json:"name"`
	Iso3         string              `json:"iso3"`
	Iso2         string              `json:"iso2"`
	NumericCode  string              `json:"numeric_code"`
	PhoneCode    string              `json:"phone_code"`
	Emoji        string              `json:"emoji"`
	Region       string              `json:"region"`
	States       []State             `json:"states"`
	Translations CountryTranslations `json:"translations"`
}

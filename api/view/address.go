package view

type Address struct {
	ID          uint   `json:"id"`
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	AddressType uint   `json:"type"`
	City        uint   `json:"city"`
}

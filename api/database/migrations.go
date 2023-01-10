package database

import "libre-asi-api/models"

func migrateModels() {
	DB.AutoMigrate(&models.User{},
		&models.Language{},
		&models.Region{},
		&models.RegionTranslations{},
		&models.Country{},
		&models.CountryTranslations{},
		&models.State{},
		&models.StateTranslations{},
		&models.City{},
		&models.AddressType{},
		&models.Address{},
		&models.Administrator{},
		&models.Person{},
		&models.Phone{})
}

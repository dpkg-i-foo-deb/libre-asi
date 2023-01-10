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
		&models.Phone{},
		&models.ReligiousPreference{},
		&models.ReligiousPreferenceTranslations{},
		&models.Race{},
		&models.RaceTranslations{},
		&models.Patient{},
		&models.Profession{},
		&models.ProfessionTranslations{},
		&models.Interviewer{},
		&models.Attendant{},
		&models.AsiForm{},
		&models.AsiFormTranslations{},
		&models.QuestionCategory{},
		&models.QuestionCategoryTranslations{})
}

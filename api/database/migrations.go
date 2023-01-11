package database

import (
	"libre-asi-api/models"
	"libre-asi-api/util"
)

func migrateModels() {
	err = DB.SetupJoinTable(&models.Interview{}, "Answers", &models.InterviewAnswers{})

	util.HandleErrorStop(err)

	err = DB.SetupJoinTable(&models.Interviewer{}, "Interpretations", &models.InterviewInterpretations{})

	util.HandleErrorStop(err)

	err = DB.SetupJoinTable(&models.Interviewer{}, "Reports", &models.InterviewReports{})

	err := DB.AutoMigrate(&models.User{},
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
		&models.QuestionCategoryTranslations{},
		&models.QuestionType{},
		&models.Question{},
		&models.QuestionTranslations{},
		&models.Option{},
		&models.OptionHelp{},
		&models.Interview{},
		&models.InterviewAnswers{},
		&models.InterviewInterpretations{},
		&models.InterviewReports{},
	)

	util.HandleErrorStop(err)
}

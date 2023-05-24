package test

import (
	"libre-asi-api/pkg/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func setup() {

	var err error

	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = DB.SetupJoinTable(&models.Interview{}, "Answers", &models.InterviewAnswers{})

	if err != nil {
		log.Fatal("Failed to setup InterviewAnswers join table")
	}

	err = DB.SetupJoinTable(&models.Interviewer{}, "Interpretations", &models.InterviewInterpretations{})

	if err != nil {
		log.Fatal("Failed to setup InterviewInterpretations join table")
	}

	err = DB.SetupJoinTable(&models.Interviewer{}, "Reports", &models.InterviewReports{})

	if err != nil {
		log.Fatal("Failed to setup InterviewReports join table")
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Country{},
		&models.CountryTranslations{},
		&models.State{},
		&models.City{},
		&models.AddressType{},
		&models.Address{},
		&models.Administrator{},
		&models.Phone{},
		&models.ReligiousPreference{},
		&models.ReligiousPreferenceTranslations{},
		&models.Race{},
		&models.RaceTranslations{},
		&models.Person{},
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

	if err != nil {
		log.Fatal("Failed to migrate models")
	}

}

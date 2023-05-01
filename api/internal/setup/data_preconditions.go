package setup

import (
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/models"
	"log"

	"gorm.io/gorm"
)

func CheckPreconditions() {
	checkAsiForms()
	checkQuestionCategories()
	checkQuestionTypes()
	checkQuestions()
}

func checkAsiForms() {
	form := models.AsiForm{}

	if err := database.DB.First(&form).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("AsiForms not found... Trying to create")
			registerAsiForms()
			log.Println("AsiForms created")
		} else {
			log.Fatal("Error while checking AsiForms", err)
		}
	}

}

func checkQuestionCategories() {
	category := models.QuestionCategory{}

	if err := database.DB.First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Question categories not found... Trying to create")

			registerQuestionCategories()

			log.Println("Question categories created")
		} else {
			log.Fatal("Error while checking question categories", err)
		}

	}
}

func checkQuestionTypes() {
	qtype := models.QuestionType{}

	if err := database.DB.First(&qtype).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Question types not found... Trying to create")

			registerQuestionTypes()

			log.Println("Question types created")
		} else {
			log.Fatal("Error while checking question types", err)
		}
	}
}

func checkQuestions() {
	question := models.Question{}

	if err := database.DB.First(&question).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			log.Println("Questions not found... Trying to create")

			registerQuestions()

			log.Println("Questions created")
		} else {
			log.Fatal("Error while checking questions", err)
		}
	}

}

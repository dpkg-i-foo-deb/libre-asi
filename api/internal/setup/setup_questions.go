package setup

import (
	"encoding/json"
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/util"
	"libre-asi-api/pkg/view"
	"os"
)

func registerQuestionCategories() {
	questionCategories := []view.QuestionCategoryParam{}

	fileName := "data/question_categories.json"

	file, err := os.Open(fileName)

	util.HandleErrorStop(err)

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&questionCategories)

	util.HandleErrorStop(err)

	for _, questionCategory := range questionCategories {

		questionCategoryDb := models.QuestionCategory{}

		questionCategoryDb.Category = questionCategory.Name

		err := database.DB.Create(&questionCategoryDb).Error

		util.HandleErrorStop(err)
	}
}

func registerQuestionTypes() {

	questionTypes := []view.QuestionTypeParam{}

	fileName := "data/question_types.json"

	file, err := os.Open(fileName)

	util.HandleErrorStop(err)

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&questionTypes)

	util.HandleErrorStop(err)

	for _, questionType := range questionTypes {

		questionTypeDb := models.QuestionType{}

		questionTypeDb.Type = questionType.Name

		err := database.DB.Create(&questionTypeDb).Error

		util.HandleErrorStop(err)
	}

}

func registerQuestions() {

	questionTypes := []models.QuestionType{}

	questionCategories := []models.QuestionCategory{}

	asiForm := models.AsiForm{}

	if err := database.DB.Find(&questionTypes).Error; err != nil {
		util.HandleErrorStop(err)
	}

	if err := database.DB.Find(&questionCategories).Error; err != nil {
		util.HandleErrorStop(err)
	}

	if err := database.DB.First(&asiForm).Error; err != nil {
		util.HandleErrorStop(err)
	}

	fileName := "data/questions.json"

	file, err := os.Open(fileName)

	util.HandleErrorStop(err)

	decoder := json.NewDecoder(file)

	questions := []view.QuestionParam{}

	err = decoder.Decode(&questions)

	util.HandleErrorStop(err)

	for _, question := range questions {

		questionDb := models.Question{}

		questionDb.SpecialCode = question.SpecialID
		questionDb.Order = question.Order
		questionDb.AsiFormID = asiForm.ID

		for _, questionType := range questionTypes {
			if questionType.Type == question.Type {
				questionDb.QuestionTypeID = questionType.ID
			}
		}

		for _, questionCategory := range questionCategories {
			if questionCategory.Category == question.Category {
				questionDb.QuestionCategoryID = questionCategory.ID
			}
		}

		questionDb.QuestionTranslations = []models.QuestionTranslations{
			{Language: "ES", Statement: question.Statement, IsDefault: false},
		}

		questionDb.Options = []models.Option{}

		for _, option := range question.Options {
			questionDb.Options = append(questionDb.Options, models.Option{
				Language:    "ES",
				Order:       option.Order,
				Value:       option.Value,
				Description: option.Description})
		}

		if err := database.DB.Create(&questionDb).Error; err != nil {
			util.HandleErrorStop(err)
		}

	}

}

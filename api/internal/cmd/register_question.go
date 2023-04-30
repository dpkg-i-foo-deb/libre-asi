package cmd

import (
	"encoding/json"
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/util"
	"libre-asi-api/pkg/view"
	"os"

	"github.com/spf13/cobra"
)

var registerQuestionCmd = &cobra.Command{
	Use:   "register-question",
	Short: "Register a question",
	Run:   registerQuestion,
}

func init() {
	rootCmd.AddCommand(registerQuestionCmd)
}

func registerQuestion(cmd *cobra.Command, args []string) {

	questionTypes := []models.QuestionType{}

	questionCategories := []models.QuestionCategory{}

	if err := database.DB.Find(&questionTypes).Error; err != nil {
		util.HandleErrorStop(err)
	}

	if err := database.DB.Find(&questionCategories).Error; err != nil {
		util.HandleErrorStop(err)
	}

	fileName := args[0]

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

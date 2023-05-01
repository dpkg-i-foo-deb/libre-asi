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

var registerQuestionCategoryCmd = &cobra.Command{
	Use:   "register-question-category",
	Short: "Register Question Category",
	Run:   registerQuestionCategory,
}

func init() {
	rootCmd.AddCommand(registerQuestionCategoryCmd)
}

func registerQuestionCategory(cmd *cobra.Command, args []string) {
	questionCategories := []view.QuestionCategoryParam{}

	fileName := args[0]

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

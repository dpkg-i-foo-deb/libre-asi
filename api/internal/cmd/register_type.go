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

var registerQuestionTypeCmd = &cobra.Command{
	Use:   "register-question-type",
	Short: "Register a question type",
	Run:   registerQuestionType,
}

func init() {
	rootCmd.AddCommand(registerQuestionTypeCmd)
}

func registerQuestionType(cmd *cobra.Command, args []string) {

	questionTypes := []view.QuestionTypeParam{}

	fileName := args[0]

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

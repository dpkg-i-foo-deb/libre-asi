package global

import (
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/util"
)

var QuestionCategories = []models.QuestionCategory{}

func initQuestionCategories() {
	err := database.DB.Find(&QuestionCategories).Error

	util.HandleErrorStop(err)

}

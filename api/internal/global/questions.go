package global

import (
	"fmt"
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/util"
)

var INFQuestions = []models.Question{}

func initINFQuestions() {

	err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
		Where("qc.category = 'INF'").
		Find(&INFQuestions).Error

	util.HandleErrorStop(err)

	fmt.Println(len(INFQuestions))
}

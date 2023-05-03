package database

import (
	"libre-asi-api/pkg/models"
	"log"
)

func GetQuestions(category string) []string {

	questions := []models.Question{}

	if err := DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
		Where("qc.name = ?", category).
		Order("questions.order ASC").
		Find(&questions).Error; err != nil {
		log.Fatal(err)
	}

	var codes []string

	for _, question := range questions {
		codes = append(codes, question.SpecialCode)
	}

	return codes
}

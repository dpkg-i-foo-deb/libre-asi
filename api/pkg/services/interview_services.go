package services

import (
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/view"
	"time"

	"gorm.io/gorm"
)

var INFQuestions = []models.Question{}

func GetInterviews() (*[]view.Interview, error) {

	interviews := []view.Interview{}

	interviewsDB := []models.Interview{}

	if err := database.DB.Find(&interviewsDB).Error; err != nil {

		if err != gorm.ErrRecordNotFound {
			return nil, errors.ErrInternalError
		}

	}

	for _, interview := range interviewsDB {
		interviews = append(interviews, view.Interview{
			ID:              interview.ID,
			StartDate:       interview.StartDate,
			EndDate:         interview.EndDate,
			PauseAt:         interview.PauseAt,
			ResumedAt:       interview.ResumedAt,
			PatientID:       interview.PatientID,
			InterviewerID:   interview.Interviewers[0].ID,
			AsiFormID:       interview.AsiFormID,
			CurrentQuestion: interview.CurrentQuestion,
		})
	}

	return &interviews, nil
}

func StartInterview(patientID int, interviewerID int) (int, error) {

	i := models.Interview{}

	asiForm := models.AsiForm{}

	if err := database.DB.First(&asiForm).Error; err != nil {
		return -1, errors.ErrInternalError
	}

	i.AsiFormID = asiForm.ID

	interviewer := models.Interviewer{}

	if err := database.DB.Where("id = ?", interviewerID).First(&interviewer).Error; err != nil {
		return -1, errors.ErrEntityNotFound
	}

	patient := models.Patient{}

	if err := database.DB.Where("id = ?", patientID).First(&patient).Error; err != nil {
		return -1, errors.ErrEntityNotFound
	}

	i.PatientID = patient.ID

	i.Interviewers = []models.Interviewer{interviewer}

	i.StartDate = time.Now()

	i.Language = "ES"

	if err := database.DB.Create(&i).Error; err != nil {
		return -1, errors.ErrInternalError
	}

	return int(i.ID), nil
}

func GetQuestion(code string) (*view.Question, error) {

	q := models.Question{}

	if err := database.DB.Preload("Options").Preload("QuestionTranslations").Where("special_code = ?", code).Find(&q).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrEntityNotFound
		}
		return nil, errors.ErrInternalError
	}

	questionType := models.QuestionType{}

	if err := database.DB.Where("id = ?", q.QuestionTypeID).First(&questionType).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	questionCategory := models.QuestionCategory{}

	if err := database.DB.Where("id = ?", q.QuestionCategoryID).First(&questionCategory).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	question := view.Question{}

	question.SpecialID = q.SpecialCode

	question.Statement = q.QuestionTranslations[0].Statement

	question.Order = q.Order

	question.Type = questionType.Type

	question.Category = questionCategory.Category

	for _, option := range q.Options {

		question.Options = append(question.Options, view.OptionParam{
			Order:       option.Order,
			Value:       option.Value,
			Description: option.Description,
		})

	}

	return &question, nil
}

func NextQuestion(interview *view.Interview) (*view.Interview, error) {

	i := models.Interview{}

	if err := database.DB.Where("id = ?", interview.ID).First(&i).Error; err != nil {

		return nil, errors.ErrEntityNotFound

	}

	handleNextQuestion(i)

	interview.CurrentQuestion = i.CurrentQuestion

	return interview, nil

}

func handleNextQuestion(i models.Interview) error {

	switch i.CurrentSection {
	case "INF":
		return handleINF(i)

	}

	return nil
}

func handleINF(i models.Interview) error {

	if len(INFQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'INF'").
			Find(&INFQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "I14" {
		i.CurrentSection = "AL"
		i.CurrentQuestion = "A1"
		return nil
	}

	for index, question := range INFQuestions {

		if question.SpecialCode == i.CurrentQuestion {
			i.CurrentQuestion = INFQuestions[index+1].SpecialCode
			break
		}

	}

	return nil

}

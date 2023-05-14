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
var ALQuestions = []models.Question{}

func GetInterviews() (*[]view.Interview, error) {

	interviews := []view.Interview{}

	interviewsDB := []models.Interview{}

	if err := database.DB.Preload("Interviewers").Find(&interviewsDB).Error; err != nil {

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

func StartInterview(patientID int, interviewerID int) (*view.Interview, error) {

	i := models.Interview{}

	iv := view.Interview{}

	asiForm := models.AsiForm{}

	if err := database.DB.First(&asiForm).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	i.AsiFormID = asiForm.ID

	interviewer := models.Interviewer{}

	if err := database.DB.First(&interviewer).Error; err != nil {
		return nil, errors.ErrEntityNotFound
	}

	patient := models.Patient{}

	if err := database.DB.Where("id = ?", patientID).First(&patient).Error; err != nil {
		return nil, errors.ErrEntityNotFound
	}

	i.PatientID = patient.ID

	i.Interviewers = []models.Interviewer{interviewer}

	i.StartDate = time.Now()

	i.Language = "ES"

	i.CurrentQuestion = "I12"
	i.CurrentSection = "INF"

	if err := database.DB.Create(&i).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	iv.ID = i.ID
	iv.StartDate = i.StartDate
	iv.EndDate = i.EndDate
	iv.PauseAt = i.PauseAt
	iv.ResumedAt = i.ResumedAt
	iv.PatientID = i.PatientID
	iv.InterviewerID = i.Interviewers[0].ID
	iv.AsiFormID = i.AsiFormID
	iv.CurrentQuestion = i.CurrentQuestion

	return &iv, nil
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

	if err := handleNextQuestion(&i); err != nil {
		return nil, err
	}

	interview.CurrentQuestion = i.CurrentQuestion

	if err := database.DB.Save(&i).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	return interview, nil

}

func handleNextQuestion(i *models.Interview) error {

	switch i.CurrentSection {
	case "INF":
		return handleINF(i)
	case "AL":
		return handleAL(i)

	}

	return nil
}

func handleINF(i *models.Interview) error {

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

func handleAL(i *models.Interview) error {

	if len(ALQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'AL'").
			Find(&ALQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "A12" {
		i.CurrentSection = "MED"
		i.CurrentQuestion = "SF1"
		return nil
	}

	for index, question := range ALQuestions {

		if question.SpecialCode == i.CurrentQuestion {
			i.CurrentQuestion = ALQuestions[index+1].SpecialCode
			break
		}

	}

	return nil
}

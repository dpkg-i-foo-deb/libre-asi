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
var MEDQuestions = []models.Question{}
var EMPQuestions = []models.Question{}
var DRUQuestions = []models.Question{}
var LAWQuestions = []models.Question{}
var FAMQuestions = []models.Question{}
var PSYQuestions = []models.Question{}
var VALQuestions = []models.Question{}

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
			PausedAt:        interview.PausedAt,
			ResumedAt:       interview.ResumedAt,
			PatientID:       interview.PatientID,
			InterviewerID:   interview.Interviewers[0].ID,
			AsiFormID:       interview.AsiFormID,
			CurrentQuestion: interview.CurrentQuestion,
		})
	}

	return &interviews, nil
}

func GetInterview(id uint) (*view.Interview, error) {

	interview := view.Interview{}

	if err := database.DB.Table("interviews").
		Select("interviews.id", "interviews.start_date", "interviews.end_date", "interviews.paused_at", "interviews.resumed_at",
			"interviews.patient_id", "interviews.asi_form_id", "interviews.current_question").
		Where("interviews.id = ?", id).
		Scan(&interview).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrEntityNotFound
		} else {
			return nil, errors.ErrInternalError
		}
	}

	return &interview, nil
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
	iv.PausedAt = i.PausedAt
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
	case "MED":
		return handleMED(i)
	case "EMP":
		return handleEMP(i)
	case "DRU":
		return handleDRU(i)
	case "LAW":
		return handleLAW(i)
	case "FAM":
		return handleFAM(i)
	case "PSY":
		return handlePSY(i)
	case "VAL":
		return handleVAL(i)

	}

	return nil
}

func handleINF(i *models.Interview) error {

	if len(INFQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'INF'").
			Order(`questions."order" ASC`).
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
			Order(`questions."order" ASC`).
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

func handleMED(i *models.Interview) error {

	if len(MEDQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'MED'").
			Order(`questions."order" ASC`).
			Find(&MEDQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "SF28" {
		i.CurrentSection = "EMP"
		i.CurrentQuestion = "E1"
		return nil
	}

	for index, question := range MEDQuestions {

		if question.SpecialCode == i.CurrentQuestion {
			i.CurrentQuestion = MEDQuestions[index+1].SpecialCode
			break
		}

	}

	return nil
}

func handleEMP(i *models.Interview) error {

	if len(EMPQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'EMP'").
			Order(`questions."order" ASC`).
			Find(&EMPQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "E36" {
		i.CurrentSection = "DRU"
		i.CurrentQuestion = "D1"
		return nil
	}

	for index, question := range EMPQuestions {

		if question.SpecialCode == i.CurrentQuestion {
			i.CurrentQuestion = EMPQuestions[index+1].SpecialCode
			break
		}

	}

	return nil
}

func handleDRU(i *models.Interview) error {

	if len(DRUQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'DRU'").
			Order(`questions."order" ASC`).
			Find(&DRUQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "D60" {
		i.CurrentQuestion = "L1"
		i.CurrentSection = "LAW"
		return nil
	}

	for index, question := range DRUQuestions {

		if question.SpecialCode == i.CurrentQuestion {
			i.CurrentQuestion = DRUQuestions[index+1].SpecialCode
			break
		}

	}

	return nil

}

func handleLAW(i *models.Interview) error {

	if len(LAWQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'LAW'").
			Order(`questions."order" ASC`).
			Find(&LAWQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "L32" {
		i.CurrentQuestion = "F1"
		i.CurrentSection = "FAM"
		return nil
	}

	for index, question := range LAWQuestions {

		if question.SpecialCode == i.CurrentQuestion {
			i.CurrentQuestion = LAWQuestions[index+1].SpecialCode
			break
		}

	}

	return nil

}

func handleFAM(i *models.Interview) error {

	if len(FAMQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'FAM'").
			Order(`questions."order" ASC`).
			Find(&FAMQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "F54" {
		i.CurrentQuestion = "P1"
		i.CurrentSection = "PSY"
		return nil
	}

	for index, question := range FAMQuestions {

		if question.SpecialCode == i.CurrentQuestion {
			i.CurrentQuestion = FAMQuestions[index+1].SpecialCode
			break
		}

	}

	return nil
}

func handlePSY(i *models.Interview) error {

	if len(PSYQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'PSY'").
			Order(`questions."order" ASC`).
			Find(&PSYQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "P21" {
		i.CurrentQuestion = "V1"
		i.CurrentSection = "VAL"
	}

	for index, question := range PSYQuestions {

		if question.SpecialCode == i.CurrentQuestion {
			i.CurrentQuestion = PSYQuestions[index+1].SpecialCode
			break
		}

	}

	return nil
}

func handleVAL(i *models.Interview) error {

	if len(VALQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'VAL'").
			Order(`questions."order" ASC`).
			Find(&VALQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "V1" {
		i.CurrentQuestion = "#"
		i.CurrentSection = "END"
		return nil
	}

	for index, question := range VALQuestions {

		if question.SpecialCode == i.CurrentQuestion {
			i.CurrentQuestion = VALQuestions[index+1].SpecialCode
			break
		}

	}

	return nil

}

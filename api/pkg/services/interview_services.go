package services

import (
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/view"
	"math"
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

func StartInterview(patientID int, interviewerEmail string) (*view.Interview, error) {

	i := models.Interview{}

	iv := view.Interview{}

	asiForm := models.AsiForm{}

	if err := database.DB.First(&asiForm).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	i.AsiFormID = asiForm.ID

	id := 0

	interviewer := models.Interviewer{}

	if err := database.DB.Table("interviewers").
		Select("interviewers.id").
		Joins("join people p ON interviewers.person_id = p.id").
		Joins("join users u ON p.user_id = u.id").
		Where("u.email = ?", interviewerEmail).
		Scan(&id).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrEntityNotFound
		}
		return nil, errors.ErrInternalError
	}

	if err := database.DB.Where("id = ?", id).First(&interviewer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrEntityNotFound
		}
		return nil, errors.ErrInternalError
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
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrEntityNotFound
		}
		return nil, errors.ErrInternalError
	}

	questionCategory := models.QuestionCategory{}

	if err := database.DB.Where("id = ?", q.QuestionCategoryID).First(&questionCategory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrEntityNotFound
		}
		return nil, errors.ErrInternalError
	}

	question := view.Question{}

	question.ID = q.ID

	question.SpecialID = q.SpecialCode

	question.Statement = q.QuestionTranslations[0].Statement

	question.Order = q.Order

	question.Type = questionType.Type

	question.Category = questionCategory.Category

	for _, option := range q.Options {

		question.Options = append(question.Options, view.Option{
			ID:          option.ID,
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

func PreviousQuestion(interview *view.Interview) (*view.Interview, error) {

	i := models.Interview{}

	if err := database.DB.Where("id = ?", interview.ID).First(&i).Error; err != nil {
		return nil, errors.ErrEntityNotFound
	}

	//Handle previous question

	interview.CurrentQuestion = i.CurrentQuestion

	if err := database.DB.Save(&i).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	return interview, nil
}

func AnswerQuestion(answers []view.Answer, interviewID uint) error {

	if err := database.DB.Transaction(func(tx *gorm.DB) error {

		for _, answer := range answers {
			i := models.Interview{}
			q := models.Question{}
			o := models.Option{}

			if err := database.DB.Where("id = ?", interviewID).First(&i).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return errors.ErrEntityNotFound
				}
				return errors.ErrInternalError
			}

			if err := database.DB.Where("id = ?", answer.OptionID).First(&o).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return errors.ErrEntityNotFound
				}
				return errors.ErrInternalError
			}

			if err := database.DB.Where("id = ?", answer.QuestionID).First(&q).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return errors.ErrEntityNotFound
				}
				return errors.ErrInternalError
			}

			ans := models.InterviewAnswers{}

			if err := database.DB.
				Where("interview_id = ? AND question_id = ? ", interviewID, answer.QuestionID).
				First(&ans).
				Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					return errors.ErrInternalError
				}
			} else {
				if err := database.DB.Unscoped().Delete(&ans).Error; err != nil {
					return errors.ErrInternalError
				}
			}

			ans.InterviewID = interviewID
			ans.OptionID = uint(answer.OptionID)
			ans.QuestionID = uint(answer.QuestionID)
			ans.Answer = answer.Value
			ans.Commentary = answer.Comment

			if err := database.DB.Save(&ans).Error; err != nil {
				return errors.ErrInternalError
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func handleNextQuestion(i *models.Interview) error {

	switch i.CurrentSection {
	case "INF":
		return handleNextINF(i)
	case "AL":
		return handleNextAL(i)
	case "MED":
		return handleNextMED(i)
	case "EMP":
		return handleNextEMP(i)
	case "DRU":
		return handleNextDRU(i)
	case "LAW":
		return handleNextLAW(i)
	case "FAM":
		return handleNextFAM(i)
	case "PSY":
		return handleNextPSY(i)
	case "VAL":
		return handleNextVAL(i)

	}

	return nil
}

func handleNextINF(i *models.Interview) error {

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
		i.CurrentQuestion = "A1A"
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

func handleNextAL(i *models.Interview) error {

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

func handleNextMED(i *models.Interview) error {

	if len(MEDQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'MED'").
			Order(`questions."order" ASC`).
			Find(&MEDQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "SF28B" {
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

func handleNextEMP(i *models.Interview) error {

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

func handleNextDRU(i *models.Interview) error {

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

func handleNextLAW(i *models.Interview) error {

	if len(LAWQuestions) == 0 {
		if err := database.DB.Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("qc.category = 'LAW'").
			Order(`questions."order" ASC`).
			Find(&LAWQuestions).
			Error; err != nil {
			return errors.ErrInternalError
		}
	}

	if i.CurrentQuestion == "L32B" {
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

func handleNextFAM(i *models.Interview) error {

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

func handleNextPSY(i *models.Interview) error {

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

func handleNextVAL(i *models.Interview) error {

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

func ComputeResults(i *view.Interview) (*view.Results, error) {

	answers := []models.InterviewAnswers{}

	if err := database.DB.Where("interview_id = ?", i.ID).Find(&answers).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrEntityNotFound
		}
		return nil, errors.ErrInternalError
	}

	druTMRAW, err := computeDruTMRAW(answers)

	if err != nil {
		return nil, err

	}

	famChildTMRAW, err := computeFamilyChildTMRAW(answers)

	if err != nil {
		return nil, err
	}

	alcoholTMRAW, err := computeAlcoholTMRAW(answers)

	if err != nil {
		return nil, err
	}

	psyTMRAW, err := computePsyTMRAW(answers)

	if err != nil {
		return nil, err
	}

	medTMRAW, err := computeMedTMRAW(answers)

	if err != nil {
		return nil, err
	}

	lawTMRAW, err := computeLawTMRAW(answers)

	if err != nil {
		return nil, err
	}

	empTMRAW, err := computeEmpTMRAW(answers)

	if err != nil {
		return nil, err
	}

	famSupportTMRAW, err := computeFamilySupportTMRAW(answers)

	if err != nil {
		return nil, err
	}

	famProblemTMRAW, err := computeFamilyProblemTMRAW(answers)

	if err != nil {
		return nil, err
	}

	results := view.Results{
		DrugScale:                float32(druTMRAW),
		FamilyChildScale:         float32(famChildTMRAW),
		AlcoholScale:             float32(alcoholTMRAW),
		PsychScale:               float32(psyTMRAW),
		MedicalScale:             float32(medTMRAW),
		LegalScale:               float32(lawTMRAW),
		EmploymentScale:          float32(empTMRAW),
		FamilySocialSupportScale: float32(famSupportTMRAW),
		FamilySocialProblemScale: float32(famProblemTMRAW),
	}

	return &results, nil

}

func computeDruTMRAW(answers []models.InterviewAnswers) (float64, error) {

	druTMRAW := 0.0

	d25_33dn := 0

	d25_33en := 0

	for _, answer := range answers {

		q := models.Question{}

		if err := database.DB.
			Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("questions.id = ? AND qc.category = 'DRU'", answer.QuestionID).
			First(&q).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return -1, errors.ErrInternalError
		}

		if q.SpecialCode == "D25D" || q.SpecialCode == "D25E" || q.SpecialCode == "D26E" || q.SpecialCode == "D27D" || q.SpecialCode == "D28D" || q.SpecialCode == "D28E" || q.SpecialCode == "D29D" || q.SpecialCode == "D30D" || q.SpecialCode == "D31D" || q.SpecialCode == "D31E" || q.SpecialCode == "D32D" || q.SpecialCode == "D32E" || q.SpecialCode == "D33D" || q.SpecialCode == "D33E" || q.SpecialCode == "D47" || q.SpecialCode == "D48" {
			d25_33dn += answer.Answer
		}

		if q.SpecialCode == "D25E" {
			d25_33en += answer.Answer
		}

		if q.SpecialCode == "D26E" {
			d25_33en += answer.Answer
		}

		if q.SpecialCode == "D27D" {
			if answer.Answer > 1 {
				d25_33en += 1
			}
		}

		if q.SpecialCode == "D28E" {
			d25_33en += answer.Answer
		}

		if q.SpecialCode == "D29D" {
			if answer.Answer > 1 {
				d25_33en += 1
			}
		}

		if q.SpecialCode == "D30D" {
			if answer.Answer > 1 {
				d25_33en += 1
			}
		}

		if q.SpecialCode == "D31E" {
			d25_33en += answer.Answer
		}

		if q.SpecialCode == "D32E" {
			d25_33en += answer.Answer
		}

		if q.SpecialCode == "D33E" {
			d25_33en += answer.Answer
		}

		if q.SpecialCode == "D39" {
			if answer.Answer >= 1 && answer.Answer <= 5 {
				druTMRAW += 1
			}

			if answer.Answer >= 6 && answer.Answer <= 15 {
				druTMRAW += 2
			}

			if answer.Answer >= 16 && answer.Answer <= 25 {
				druTMRAW += 3
			}

			if answer.Answer >= 26 {
				druTMRAW += 4
			}
		}

		if q.SpecialCode == "D40" {
			if answer.Answer >= 15 && answer.Answer <= 30 {
				druTMRAW += 1
			}

			if answer.Answer >= 8 && answer.Answer <= 14 {
				druTMRAW += 2
			}

			if answer.Answer >= 4 && answer.Answer <= 7 {
				druTMRAW += 3
			}

			if answer.Answer >= 0 && answer.Answer <= 3 {
				druTMRAW += 4
			}
		}

		if q.SpecialCode == "D41" {
			if answer.Answer >= 1 && answer.Answer <= 299999 {
				druTMRAW += 1
			}

			if answer.Answer >= 300000 && answer.Answer <= 600000 {
				druTMRAW += 2
			}

			if answer.Answer >= 600001 && answer.Answer <= 1000000 {
				druTMRAW += 3
			}

			if answer.Answer >= 1000001 {
				druTMRAW += 4
			}
		}

		if q.SpecialCode == "D42" || q.SpecialCode == "D43" || q.SpecialCode == "D44" || q.SpecialCode == "D45" {
			druTMRAW += math.Sqrt(8) * float64(answer.Answer)
		}

		if q.SpecialCode == "D46" {
			if answer.Answer >= 1 && answer.Answer <= 5 {
				druTMRAW += 1
			}

			if answer.Answer >= 6 && answer.Answer <= 15 {
				druTMRAW += 2
			}

			if answer.Answer >= 16 && answer.Answer <= 25 {
				druTMRAW += 3
			}

			if answer.Answer >= 26 {
				druTMRAW += 4
			}
		}

	}

	if d25_33dn >= 1 && d25_33dn <= 15 {
		druTMRAW += 1
	}

	if d25_33dn >= 16 && d25_33dn <= 30 {
		druTMRAW += 2
	}

	if d25_33dn >= 31 && d25_33dn <= 60 {
		druTMRAW += 3
	}

	if d25_33dn >= 61 {
		druTMRAW += 4
	}

	druTMRAW += float64(d25_33en)

	return druTMRAW, nil
}

func computeFamilyChildTMRAW(answers []models.InterviewAnswers) (float64, error) {

	familyChildTMRAW := 0.0

	for _, answer := range answers {

		q := models.Question{}

		if err := database.DB.
			Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("questions.id = ? AND qc.category = 'FAM'", answer.QuestionID).
			First(&q).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return -1, errors.ErrInternalError
		}

		if q.SpecialCode == "F46" || q.SpecialCode == "F49" {

			if answer.Answer >= 1 {
				familyChildTMRAW += math.Sqrt(8)
			}

		}

		if q.SpecialCode == "F47" || q.SpecialCode == "F48" || q.SpecialCode == "F49" {
			familyChildTMRAW += float64(answer.Answer)
		}

	}

	return familyChildTMRAW, nil
}

func computeAlcoholTMRAW(answers []models.InterviewAnswers) (float64, error) {

	alcoholTMRAW := 0.0

	for _, answer := range answers {

		q := models.Question{}

		if err := database.DB.
			Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("questions.id = ? AND qc.category = 'DRU'", answer.QuestionID).
			First(&q).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return -1, errors.ErrInternalError
		}

		if q.SpecialCode == "D13" || q.SpecialCode == "D15" || q.SpecialCode == "D21" {
			if answer.Answer >= 1 && answer.Answer <= 5 {
				alcoholTMRAW += 1
			}

			if answer.Answer >= 6 && answer.Answer <= 15 {
				alcoholTMRAW += 2
			}

			if answer.Answer >= 16 && answer.Answer <= 25 {
				alcoholTMRAW += 3
			}

			if answer.Answer >= 26 {
				alcoholTMRAW += 4
			}
		}

		if q.SpecialCode == "D14" {
			if answer.Answer >= 15 && answer.Answer <= 30 {
				alcoholTMRAW += 1
			}

			if answer.Answer >= 8 && answer.Answer <= 14 {
				alcoholTMRAW += 2
			}

			if answer.Answer >= 4 && answer.Answer <= 7 {
				alcoholTMRAW += 3
			}

			if answer.Answer >= 0 && answer.Answer <= 3 {
				alcoholTMRAW += 4
			}
		}

		if q.SpecialCode == "D16" {

			if answer.Answer >= 1 && answer.Answer <= 60000 {
				alcoholTMRAW += 1
			}

			if answer.Answer >= 60001 && answer.Answer <= 150000 {
				alcoholTMRAW += 2
			}

			if answer.Answer >= 150001 && answer.Answer <= 400000 {
				alcoholTMRAW += 3
			}

			if answer.Answer >= 400001 {
				alcoholTMRAW += 4
			}
		}

		if q.SpecialCode == "D17" || q.SpecialCode == "D18" || q.SpecialCode == "D19" || q.SpecialCode == "D20" {
			if answer.Answer == 1 {
				alcoholTMRAW += math.Sqrt(8)
			}
		}

		if q.SpecialCode == "D22" || q.SpecialCode == "D23" {
			alcoholTMRAW += float64(answer.Answer)
		}

	}

	return alcoholTMRAW, nil
}

func computePsyTMRAW(answers []models.InterviewAnswers) (float64, error) {

	psyTMRAW := 0.0

	p16dn := make([]int, 0)

	for _, answer := range answers {

		q := models.Question{}

		if err := database.DB.
			Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("questions.id = ? AND (qc.category = 'DRU' OR qc.category = 'PSY' OR qc.category = 'FAM' OR qc.category = 'AL')", answer.QuestionID).
			First(&q).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return -1, errors.ErrInternalError
		}

		if q.SpecialCode == "A4B" {
			if answer.Answer > 1 {
				psyTMRAW += math.Sqrt(8)
			}
		}

		if q.SpecialCode == "F8B" || q.SpecialCode == "F39" || q.SpecialCode == "P20" || q.SpecialCode == "P21" {
			psyTMRAW += float64(answer.Answer)
		}

		if q.SpecialCode == "P8B" || q.SpecialCode == "P9B" || q.SpecialCode == "P10B" || q.SpecialCode == "P11B" || q.SpecialCode == "P12B" || q.SpecialCode == "P13B" || q.SpecialCode == "P14B" || q.SpecialCode == "P15B" || q.SpecialCode == "P16B" {
			psyTMRAW += math.Sqrt(3) * float64(answer.Answer)
		}

		if q.SpecialCode == "P11C" || q.SpecialCode == "P13C" || q.SpecialCode == "P14C" || q.SpecialCode == "P15C" || q.SpecialCode == "P16C" {
			p16dn = append(p16dn, answer.Answer)
		}

		if q.SpecialCode == "P18" || q.SpecialCode == "P19" {
			if answer.Answer >= 1 && answer.Answer <= 5 {
				psyTMRAW += 1
			}

			if answer.Answer >= 6 && answer.Answer <= 15 {
				psyTMRAW += 2
			}

			if answer.Answer >= 16 && answer.Answer <= 25 {
				psyTMRAW += 3
			}

			if answer.Answer >= 26 {
				psyTMRAW += 4
			}
		}

	}

	psyTMRAW += float64(min(p16dn))

	return psyTMRAW, nil
}

func computeMedTMRAW(answers []models.InterviewAnswers) (float64, error) {

	medTMRAW := 0.0

	for _, answer := range answers {

		q := models.Question{}

		if err := database.DB.
			Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("questions.id = ? AND (qc.category = 'PSY' OR qc.category = 'FAM' OR qc.category = 'AL')", answer.QuestionID).
			First(&q).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return -1, errors.ErrInternalError
		}

		if q.SpecialCode == "A3B" {
			if answer.Answer > 1 {
				medTMRAW += 1
			}
		}

		if q.SpecialCode == "F49" || q.SpecialCode == "SF22" || q.SpecialCode == "SF23" || q.SpecialCode == "SF24" {
			medTMRAW += float64(answer.Answer)
		}

		if q.SpecialCode == "SF20" || q.SpecialCode == "SF21" {
			if answer.Answer >= 1 && answer.Answer <= 5 {
				medTMRAW += 1
			}

			if answer.Answer >= 6 && answer.Answer <= 15 {
				medTMRAW += 2
			}

			if answer.Answer >= 16 && answer.Answer <= 25 {
				medTMRAW += 3
			}

			if answer.Answer >= 26 {
				medTMRAW += 4
			}
		}

	}

	return medTMRAW, nil
}

func computeLawTMRAW(answers []models.InterviewAnswers) (float64, error) {

	lawTMRAW := 0.0

	l27_29bn := 0

	for _, answer := range answers {

		q := models.Question{}

		if err := database.DB.
			Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("questions.id = ? AND (qc.category='LAW' OR qc.category='EMP')", answer.QuestionID).
			First(&q).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return -1, errors.ErrInternalError
		}

		if q.SpecialCode == "E29" {
			if answer.Answer >= 1 && answer.Answer <= 180000 {
				lawTMRAW += 1
			}

			if answer.Answer >= 181001 && answer.Answer <= 370000 {
				lawTMRAW += 2
			}

			if answer.Answer >= 370001 && answer.Answer <= 1000000 {
				lawTMRAW += 3
			}

			if answer.Answer >= 1000001 {
				lawTMRAW += 4
			}
		}

		if q.SpecialCode == "L26B" || q.SpecialCode == "L28B" || q.SpecialCode == "L30B" {
			if answer.Answer >= 1 {
				lawTMRAW += math.Sqrt(8)
			}
		}

		if q.SpecialCode == "L27B" || q.SpecialCode == "L29B" {
			if answer.Answer > 1 {
				l27_29bn = 1
			}
		}

		if q.SpecialCode == "L31" {
			if answer.Answer >= 1 && answer.Answer <= 5 {
				lawTMRAW += 1
			}

			if answer.Answer >= 6 && answer.Answer <= 15 {
				lawTMRAW += 2
			}

			if answer.Answer >= 16 && answer.Answer <= 25 {
				lawTMRAW += 3
			}

			if answer.Answer >= 26 {
				lawTMRAW += 4
			}
		}

	}

	lawTMRAW += math.Sqrt(8) * float64(l27_29bn)

	return lawTMRAW, nil
}

func computeEmpTMRAW(answers []models.InterviewAnswers) (float64, error) {
	empTMRAW := 0.0

	for _, answer := range answers {

		q := models.Question{}

		if err := database.DB.
			Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("questions.id = ? AND qc.category='EMP'", answer.QuestionID).
			First(&q).
			Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return -1, errors.ErrInternalError
		}

		if q.SpecialCode == "E10" {
			if answer.Answer >= 3 && answer.Answer <= 4 {
				empTMRAW += 1
			}
		}

		if q.SpecialCode == "E19" {
			if answer.Answer >= 16 && answer.Answer <= 22 {
				empTMRAW += 1
			}

			if answer.Answer >= 6 && answer.Answer <= 15 {
				empTMRAW += 2
			}

			if answer.Answer >= 1 && answer.Answer <= 5 {
				empTMRAW += 3
			}
		}

		if q.SpecialCode == "E20" {
			if answer.Answer >= 370001 && answer.Answer <= 1000000 {
				empTMRAW += 1
			}

			if answer.Answer >= 181001 && answer.Answer <= 370000 {
				empTMRAW += 2
			}

			if answer.Answer >= 1 && answer.Answer <= 180000 {
				empTMRAW += 3
			}

			if answer.Answer == 0 {
				empTMRAW += 4
			}
		}

		if q.SpecialCode == "E21" {
			if answer.Answer >= 1 && answer.Answer <= 5 {
				empTMRAW += math.Sqrt(1.6)
			}

			if answer.Answer >= 6 {
				empTMRAW += math.Sqrt(1.6) * 2
			}
		}

	}

	return empTMRAW, nil
}

func computeFamilySupportTMRAW(answers []models.InterviewAnswers) (float64, error) {

	famSupportTMRAW := 0.0

	ff3 := 0.0
	ff4 := 0.0
	ff5 := 0.0
	ff9 := 0.0

	for _, answer := range answers {

		q := models.Question{}

		if err := database.DB.
			Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("questions.id = ? AND qc.category='FAM'", answer.QuestionID).
			First(&q).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return -1, errors.ErrInternalError
		}

		if q.SpecialCode == "F3A" || q.SpecialCode == "F3B" || q.SpecialCode == "F3C" {
			ff3 += float64(answer.Answer)
		}

		if q.SpecialCode == "F4A" || q.SpecialCode == "F4B" || q.SpecialCode == "F4C" {
			ff4 += float64(answer.Answer)
		}

		if q.SpecialCode == "F5A" || q.SpecialCode == "F5B" || q.SpecialCode == "F5C" {
			ff5 += float64(answer.Answer)
		}

		if q.SpecialCode == "F9A" || q.SpecialCode == "F9B" || q.SpecialCode == "F9C" {
			ff9 += float64(answer.Answer)
		}

	}

	famSupportTMRAW += ff3 * math.Sqrt(1.6)
	famSupportTMRAW += ff4 * math.Sqrt(1.6)
	famSupportTMRAW += ff5 * math.Sqrt(1.6)
	famSupportTMRAW += ff9 * math.Sqrt(1.6)

	return famSupportTMRAW, nil
}

func computeFamilyProblemTMRAW(answers []models.InterviewAnswers) (float64, error) {

	famProblemTMRAW := 0.0

	ff6 := 0.0
	ff7 := 0.0

	for _, answer := range answers {

		q := models.Question{}

		if err := database.DB.
			Joins("JOIN question_categories qc ON qc.id = questions.question_category_id").
			Where("questions.id = ? AND qc.category='FAM'", answer.QuestionID).
			First(&q).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return -1, errors.ErrInternalError
		}

		if q.SpecialCode == "F11" {
			famProblemTMRAW += float64(answer.Answer) * math.Sqrt(8)
		}

		if q.SpecialCode == "F14" || q.SpecialCode == "F15" {
			famProblemTMRAW += float64(answer.Answer)
		}

		if q.SpecialCode == "F6A" || q.SpecialCode == "F6B" || q.SpecialCode == "F6C" {
			ff6 += float64(answer.Answer)
		}

		if q.SpecialCode == "F7A" || q.SpecialCode == "F7B" || q.SpecialCode == "F7C" {
			ff7 += float64(answer.Answer)
		}

	}

	famProblemTMRAW += ff6 * math.Sqrt(1.6)
	famProblemTMRAW += ff7 * math.Sqrt(1.6)

	return famProblemTMRAW, nil
}

func min(a []int) int {
	minVal := math.MaxInt

	for _, v := range a {
		if v < minVal {
			minVal = v
		}
	}

	return minVal
}

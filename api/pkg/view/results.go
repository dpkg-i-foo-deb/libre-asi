package view

type Results struct {
	DrugScale                float32 `json:"drugScale"`
	FamilyChildScale         float32 `json:"familyChildScale"`
	AlcoholScale             float32 `json:"alcoholScale"`
	PsychScale               float32 `json:"psychScale"`
	MedicalScale             float32 `json:"medicalScale"`
	LegalScale               float32 `json:"legalScale"`
	EmploymentScale          float32 `json:"employmentScale"`
	FamilySocialSupportScale float32 `json:"familySocialSupportScale"`
	FamilySocialProblemScale float32 `json:"familySocialProblemScale"`
}

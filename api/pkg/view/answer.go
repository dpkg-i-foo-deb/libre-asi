package view

type Answer struct {
	Value      int    `json:"value"`
	QuestionID int    `json:"questionID"`
	OptionID   int    `json:"optionID"`
	Comment    string `json:"comment"`
}

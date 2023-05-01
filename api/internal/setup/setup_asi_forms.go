package setup

import (
	"encoding/json"
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/util"
	"libre-asi-api/pkg/view"
	"os"
)

func registerAsiForms() {
	asiForms := []view.ASIForm{}

	fileName := "data/asi_forms.json"

	file, err := os.Open(fileName)

	util.HandleErrorStop(err)

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&asiForms)

	util.HandleErrorStop(err)

	for _, asiForm := range asiForms {

		asiFormDb := models.AsiForm{}

		asiFormDb.Edition = asiForm.Edition
		asiFormDb.Version = asiForm.Version

		asiFormDb.AsiFormTranslations = []models.AsiFormTranslations{
			{Language: "ES",
				IsDefault:   false,
				Name:        asiForm.Name,
				Description: asiForm.Description},
		}

		if err := database.DB.Create(&asiFormDb).Error; err != nil {
			util.HandleErrorStop(err)
		}
	}
}

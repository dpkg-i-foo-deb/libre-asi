package services

import (
	"libre-asi-api/models"
)

func IndexService() models.Response {

	r := models.Response{
		Status:  string(models.OK),
		Message: "Welcome to LibreASI API!",
	}

	return r
}

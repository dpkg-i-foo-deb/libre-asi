package services

import (
	"libre-asi-api/cfg"
	"libre-asi-api/errors"
	"libre-asi-api/models"
)

func SetupService(user models.User) (models.Response, error) {

	var res models.Response

	res.Status = string(models.OK)
	res.Message = "Setup Complete"

	if createAdmin(user) != nil {

		res.Status = string(models.CHECK_REQUEST)
		res.Message = "Check Request"

		return res
	}

	return c.Status(201).JSON(res)
}

func CheckSetupService() error {

	if cfg.CheckExistingAdmin() != nil {
		return errors.ErrSetupRequired
	}

	return nil
}

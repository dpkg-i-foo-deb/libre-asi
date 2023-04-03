package services

import (
	"libre-asi-api/cfg"
	"libre-asi-api/errors"
	"libre-asi-api/models"
)

func SetupService(admin models.Administrator) error {
	_, err := RegisterAdministratorService(admin)

	return err
}

func CheckSetupService() error {

	if cfg.CheckExistingAdmin() != nil {
		return errors.ErrSetupRequired
	}

	return nil
}

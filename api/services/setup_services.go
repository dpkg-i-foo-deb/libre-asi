package services

import (
	"libre-asi-api/cfg"
	"libre-asi-api/errors"
	"libre-asi-api/models"
)

func Setup(admin models.Administrator) error {
	_, err := RegisterAdministrator(admin)

	return err
}

func CheckSetup() error {

	if cfg.CheckExistingAdmin() != nil {
		return errors.ErrSetupRequired
	}

	return nil
}

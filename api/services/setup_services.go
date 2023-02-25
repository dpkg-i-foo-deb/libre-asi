package services

import (
	"libre-asi-api/cfg"
	"libre-asi-api/errors"
	"libre-asi-api/models"
)

func SetupService(user models.User) error {
	return createAdmin(user, false)
}

func CheckSetupService() error {

	if cfg.CheckExistingAdmin() != nil {
		return errors.ErrSetupRequired
	}

	return nil
}

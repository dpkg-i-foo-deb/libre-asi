package services

import (
	"libre-asi-api/cfg"
	"libre-asi-api/errors"
	"libre-asi-api/view"
)

func Setup(admin view.Administrator) error {
	_, err := RegisterAdministrator(admin, true)

	return err
}

func CheckSetup() error {

	if cfg.CheckExistingAdmin() != nil {
		return errors.ErrSetupRequired
	}

	return nil
}

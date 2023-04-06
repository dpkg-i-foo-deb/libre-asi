package services

import (
	"libre-asi-api/pkg/auth"
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/errors"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/util"
	"libre-asi-api/pkg/view"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginAdmin(a view.Administrator) (*models.JWTPair, *models.PasswordResetTk, error) {

	var user models.User
	var administrator models.Administrator

	if database.DB.Where("email = ?", a.Email).First(&user).Error != nil {
		return nil, nil, errors.ErrAccessDenied
	}

	if database.DB.Where("user_id = ?", user.ID).First(&administrator).Error != nil {
		return nil, nil, errors.ErrAccessDenied
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(a.Password)); err != nil {
		return nil, nil, errors.ErrAccessDenied
	}

	if user.NeedsPasswordReset {
		token, err := auth.GeneratePasswordResetToken(a.Email)

		if err != nil {
			return nil, nil, errors.ErrInternalError
		}

		return nil, &token, errors.ErrrNeedsPasswordReset
	}

	token, err := auth.GenerateJWTPair(a.Email, string(models.ADMINISTRATOR))

	if err != nil {
		return nil, nil, errors.ErrInternalError
	}

	return &token, nil, nil
}

// TODO test this
func GetAdministrators() ([]view.Administrator, error) {

	var admins []view.Administrator

	if err := database.DB.Table("users").
		Joins("JOIN administrators ON administrators.user_id = users.id").
		Select("administrators.id, users.email, users.username, '', users.needs_password_reset").
		Scan(&admins).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	return admins, nil

}

func RegisterAdministrator(newAdmin view.Administrator, isFirst bool) (*view.Administrator, error) {

	var p string
	var queriedUser models.User
	var err error

	if database.DB.Where("email = ?", newAdmin.Email).First(&queriedUser).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	if database.DB.Where("username = ?", newAdmin.Username).First(&queriedUser).Error != gorm.ErrRecordNotFound {
		return nil, errors.ErrConflict
	}

	if !isFirst {

		p, err = util.MakeRandomPassword()

		if err != nil {
			return nil, errors.ErrInternalError
		}

		newAdmin.Password = p

		newAdmin.NeedsPasswordReset = true
	}

	newAdmin.Password, err = util.HashPassword(newAdmin.Password)

	if err != nil {
		return nil, errors.ErrInternalError
	}

	if err = database.DB.Transaction(func(tx *gorm.DB) error {

		var user models.User

		user.Email = newAdmin.Email
		user.Username = newAdmin.Username
		user.Password = newAdmin.Password
		user.NeedsPasswordReset = newAdmin.NeedsPasswordReset

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		var admin models.Administrator

		admin.UserID = user.ID

		if err := tx.Create(&admin).Error; err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, errors.ErrInternalError
	}

	return &newAdmin, nil
}

func UpdateAdministrator(a view.Administrator) (*view.Administrator, error) {
	var admin models.Administrator
	var user models.User

	if database.DB.Where("ID = ?", a.ID).First(&admin).Error != nil {
		return nil, errors.ErrEntityNotFound
	}

	if database.DB.Where("ID = ?", admin.UserID).First(&user).Error != nil {
		return nil, errors.ErrEntityNotFound
	}

	user.Username = a.Username

	if err := database.DB.Omit("password").Save(&user).Error; err != nil {
		return nil, errors.ErrInternalError
	}

	return &view.Administrator{
		ID:                 admin.ID,
		Email:              user.Email,
		Username:           user.Username,
		Password:           "",
		NeedsPasswordReset: false,
	}, nil
}

func DeleteAdministrator(id int) error {

	var queriedAdmin models.Administrator

	if database.DB.Where("ID = ?", id).First(&queriedAdmin).Error != nil {
		return errors.ErrEntityNotFound
	}

	if err := database.DB.Delete(&queriedAdmin).Error; err != nil {
		return errors.ErrInternalError
	}

	return nil
}

func SetAdministratorPassword(email string, credentials models.PasswordChange) error {

	var queriedUser models.User

	if database.DB.Preload("AdministratorsCre").Where("email = ?", email).First(&queriedUser).Error != nil {
		return errors.ErrEntityNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(queriedUser.Password), []byte(credentials.CurrentPassword)); err != nil {
		return errors.ErrAccessDenied
	}

	isAdmin := false

	for _, admin := range queriedUser.Administrators {
		if admin.UserID == queriedUser.ID {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		return errors.ErrAccessDenied
	}

	p, err := util.HashPassword(credentials.NewPassword)

	if err != nil {
		return errors.ErrInternalError
	}

	queriedUser.Password = p

	if err := database.DB.Save(&queriedUser).Error; err != nil {
		return errors.ErrInternalError
	}

	return nil
}

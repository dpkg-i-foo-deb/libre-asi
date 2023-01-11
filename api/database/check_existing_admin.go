package database

import (
	"libre-asi-api/models"
	"libre-asi-api/util"
	"log"
	"os"

	"gorm.io/gorm"
)

func checkExistingAdmin() {
	var a models.Administrator
	r := DB.Take(&a)

	if r.Error == gorm.ErrRecordNotFound {
		log.Println("No administrator account exists, creating a new one")
		err = DB.Transaction(createAdmin)
		util.HandleErrorStop(err)
	} else {
		util.HandleErrorStop(r.Error)
	}
}

func createAdmin(db *gorm.DB) error {

	u := models.User{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Username: os.Getenv("ADMIN_USERNAME"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Administrators: []models.Administrator{
			{},
		},
	}

	r := DB.Create(&u)

	if r.Error != nil {
		return r.Error
	}

	log.Println("Administrator user Created")
	log.Println("Admin email: " + u.Email)
	log.Println("Admin username: " + u.Username)
	log.Println("Admin password: " + u.Password)

	return nil
}

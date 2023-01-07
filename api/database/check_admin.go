package database

import (
	"database/sql"
	"libre-asi-api/models"
	"libre-asi-api/util"
	"log"
	"os"
	"time"
)

func CheckAdmin() {
	var newUser models.User
	var id int64
	var newAdmin models.Administrator

	err := CheckAdminStatement.QueryRow().Scan(&newUser)

	if err != nil && err != sql.ErrNoRows {
		util.HandleErrorStop(err)
	}

	if err == sql.ErrNoRows {

		hashedPassword, err := util.HashPassword(os.Getenv("ADMIN_PASSWORD"))

		util.HandleErrorStop(err)

		newUser.Email = os.Getenv("ADMIN_EMAIL")
		newUser.Username = "ADMIN"
		newUser.Password = hashedPassword
		newUser.CreatedAt = time.Now()
		newUser.UpdatedAt = time.Now()

		transaction, err := DB.Begin()

		util.HandleErrorStop(err)

		err = transaction.Stmt(CreateUserAdminStatement).QueryRow(
			newUser.Email,
			newUser.Username,
			newUser.Password,
			newUser.CreatedAt,
			newUser.UpdatedAt).Scan(&id)

		if err != nil {
			transaction.Rollback()
			util.HandleErrorStop(err)
		}

		newAdmin.ID = int32(id)
		newAdmin.CreatedAt = time.Now()
		newAdmin.UpdatedAt = time.Now()

		_, err = transaction.Stmt(CreateAdminStatement).Query(
			newAdmin.ID,
			newAdmin.CreatedAt,
			newAdmin.UpdatedAt)

		if err != nil {
			transaction.Rollback()
			util.HandleErrorStop(err)
		}

		transaction.Commit()

		log.Println("The Admin login password is: " + os.Getenv("ADMIN_PASSWORD"))
	}
}

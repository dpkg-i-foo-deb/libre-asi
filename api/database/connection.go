package database

import (
	"libre-asi-api/util"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	var dsn = os.Getenv("DB_CONNECTION")
	DB, err = gorm.Open(postgres.Open(dsn))

	util.HandleErrorStop(err)

	migrateModels()

	checkExistingAdmin()

	checkWorld()
}

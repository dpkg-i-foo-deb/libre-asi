package database

import (
	"libre-asi-api/util"
	"os"

	//FOR PRODUCTION
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//FOR DEVELOPMENT
	sqlite "github.com/glebarez/sqlite"
)

var DB *gorm.DB
var err error

func init() {

	if os.Getenv("PROD") == "true" {
		var dsn = os.Getenv("DB_CONNECTION")
		DB, err = gorm.Open(postgres.Open(dsn))
	} else {
		DB, err = gorm.Open(sqlite.Open("libre-asi.db"), &gorm.Config{})
	}

	util.HandleErrorStop(err)

	migrateModels()

	//checkExistingAdmin()

	checkWorld()
}

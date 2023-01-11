package database

import (
	"libre-asi-api/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	const dsn = "host=localhost user=libre_asi password=libre_asi dbname=libre_asi port=5432 sslmode=disable TimeZone=America/Bogota"
	DB, err = gorm.Open(postgres.Open(dsn))

	util.HandleErrorStop(err)

	migrateModels()

	checkExistingAdmin()
}

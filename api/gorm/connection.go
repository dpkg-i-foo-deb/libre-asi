package gorm

import (
	"libre-asi-api/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Start() {
	dsn := "host=libre-asi-database user=libre_asi password=libre_asi dbname=libre_asi port=5432 sslmode=disable TimeZone=America/Bogota"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	util.HandleErrorStop(err)
}

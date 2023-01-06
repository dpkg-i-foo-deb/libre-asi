package gorm

import (
	"libre-asi-api/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	dsn := "host=libre-asi-database user=libre_asi password=libre_asi dbname=libre_asi port=5432 sslmode=disable TimeZone=America/Bogota"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	util.HandleErrorStop(err)
}

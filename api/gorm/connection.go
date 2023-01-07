package gorm

import (
	"libre-asi-api/util"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Start() {
	dsn := os.Getenv("DB_CONNECTION")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	util.HandleErrorStop(err)
}

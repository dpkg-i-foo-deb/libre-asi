package orm

import (
	"libre-asi-api/util"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Start() {
	dsn := os.Getenv("DB_CONNECTION")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	util.HandleErrorStop(err)

}

package test

import (
	"log"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func postRun() {

	err := os.Remove("libre-asi.db")

	if err != nil {
		log.Fatal(err)
	}

}

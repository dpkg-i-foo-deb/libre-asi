package main

import (
	"libre-asi-api/app"
	"libre-asi-api/gorm"
	_ "libre-asi-api/routes"
)

func main() {
	gorm.Start()
	app.Start()
}

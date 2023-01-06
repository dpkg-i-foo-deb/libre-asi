package main

import (
	"libre-asi-api/app"
	"libre-asi-api/gorm"
)

func main() {
	app.Start()
	gorm.Start()
}

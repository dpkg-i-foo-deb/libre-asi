package main

import (
	"libre-asi-api/app"
	"libre-asi-api/database"
	_ "libre-asi-api/routes"
)

func main() {
	database.CheckAdmin()
	app.Start()
}

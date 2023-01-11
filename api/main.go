package main

import (
	"libre-asi-api/app"
	_ "libre-asi-api/database"
	_ "libre-asi-api/routes"
)

func main() {
	app.Start()
}

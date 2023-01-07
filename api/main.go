package main

import (
	"libre-asi-api/app"
	"libre-asi-api/orm"
	_ "libre-asi-api/routes"
)

func main() {
	orm.Start()
	app.Start()
}

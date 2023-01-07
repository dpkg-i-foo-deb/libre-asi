package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/services"
)

func indexRoute() {
	app.AddGet("/", services.IndexService)
}

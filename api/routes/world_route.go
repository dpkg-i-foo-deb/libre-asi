package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/services"
)

func worldRoute() {
	app.AddGet("/world", services.WorldService)
}

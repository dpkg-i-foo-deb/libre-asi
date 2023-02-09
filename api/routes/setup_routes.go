package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/services"
)

func setupRoutes() {
	app.AddPost("/set-up", services.SetupService)
	app.AddGet("/set-up", services.CheckSetupService)
}

package routes

import (
	"libre-asi-api/handlers"
)

func setupRoutes() {

	server.Get("/set-up", handlers.CheckSetupHandler)

	//app.AddPost("/set-up", services.SetupService)
}

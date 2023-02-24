package routes

import (
	"libre-asi-api/handlers"
)

func setupRoutes() {
	server.Get("/set-up", handlers.CheckSetupHandler)
	server.Post("/set-up", handlers.SetupHandler)
}

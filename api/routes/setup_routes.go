package routes

import (
	"libre-asi-api/handlers"
)

func setupRoutes() {
	server.Get("/set-up", handlers.CheckSetup)
	server.Post("/set-up", handlers.Setup)
}

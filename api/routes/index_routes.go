package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func indexRoute() {
	server.Get("/", auth.ValidateAccessToken, handlers.IndexHandler)
}

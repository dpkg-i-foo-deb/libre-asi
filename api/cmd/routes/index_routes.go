package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/cmd/handlers"
)

func indexRoute() {
	server.Get("/", auth.ValidateAccessToken, handlers.Index)
}

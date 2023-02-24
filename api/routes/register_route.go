package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func registerRoute() {
	server.Post("/register/:role", auth.ValidateAccessToken, handlers.RegisterHandler)
}

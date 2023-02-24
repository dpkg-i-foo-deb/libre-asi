package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func signOutRoute() {
	server.Post("/sign-out", auth.ValidateAccessToken, handlers.SignOutHandler)
}

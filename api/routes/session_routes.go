package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func sessionRoutes() {
	server.Post("/login/:role", handlers.LoginHandler)
	server.Post("/password-reset", auth.ValidatePasswordResetToken, handlers.SetPasswordHandler)
	server.Post("/sign-out", auth.ValidateAccessToken, handlers.SignOutHandler)
}

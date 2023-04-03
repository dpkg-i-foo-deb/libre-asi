package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func sessionRoutes() {
	server.Post("/login/:role", handlers.Login)
	server.Post("/password-reset/:role", auth.ValidatePasswordResetToken, handlers.SetPassword)
	server.Post("/sign-out", auth.ValidateAccessToken, handlers.SignOut)
}

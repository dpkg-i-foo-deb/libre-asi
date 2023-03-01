package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func passwordResetRoute() {
	server.Post("/password-reset", auth.ValidatePasswordResetToken, handlers.SetPasswordHandler)
}

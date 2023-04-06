package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/cmd/handlers"
)

func refreshRoute() {
	server.Post("/refresh", auth.ValidateRefreshToken, auth.ValidateRefreshTokenDate, handlers.Refresh)
}

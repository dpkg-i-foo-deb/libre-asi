package routes

import (
	"libre-asi-api/cmd/handlers"
	"libre-asi-api/pkg/auth"
)

func refreshRoute() {
	server.Post("/refresh", auth.ValidateRefreshToken, auth.ValidateRefreshTokenDate, handlers.Refresh)
}

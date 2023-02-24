package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func refreshRoute() {
	server.Post("/refresh", auth.ValidateRefreshToken, auth.ValidateRefreshTokenDate, handlers.RefreshHandler)
}

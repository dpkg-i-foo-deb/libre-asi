package routes

import (
	"libre-asi-api/cmd/handlers"
	"libre-asi-api/pkg/auth"
)

func indexRoute() {
	server.Get("/", auth.ValidateAccessToken, handlers.Index)
}

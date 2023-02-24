package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func adminRoutes() {
	server.Get("/administrators", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.GetAdministratorsHandler)
	server.Post("/administrators", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.RegisterAdministratorHandler)
}

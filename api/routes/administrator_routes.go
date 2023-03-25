package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/handlers"
)

func adminRoutes() {
	server.Get("/administrators", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.GetAdministratorsHandler)
	server.Post("/administrators", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.RegisterAdministratorHandler)
	server.Patch("/administrators/", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.UpdateAdministratorHandler)
	server.Delete("/administrators/:id", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.DeleteAdministratorHandler)
}

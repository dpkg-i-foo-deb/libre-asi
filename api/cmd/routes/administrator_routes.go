package routes

import (
	"libre-asi-api/auth"
	"libre-asi-api/cmd/handlers"
)

func adminRoutes() {
	server.Get("/administrators", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.GetAdministrators)
	server.Post("/administrators", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.RegisterAdministrator)
	server.Patch("/administrators/", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.UpdateAdministrator)
	server.Delete("/administrators/:id", auth.ValidateAccessToken, auth.ValidateAdministratorRole, handlers.DeleteAdministrator)
}

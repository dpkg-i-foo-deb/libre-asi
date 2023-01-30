package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/auth"
	"libre-asi-api/services"
)

func adminRoutes() {
	app.AddGet("/administrators", auth.ValidateAccessToken, auth.ValidateAdministratorRole, services.GetAdministratorsService)
	app.AddPost("/administrators", auth.ValidateAccessToken, auth.ValidateAdministratorRole, services.RegisterAdministratorService)
}

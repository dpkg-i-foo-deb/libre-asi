package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/services"
)

func loginAdminRoute() {
	app.AddPost("/login/admin", services.LoginAdminService)
}

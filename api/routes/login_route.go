package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/services"
)

func loginRoute() {
	app.AddPost("/login/:role", services.LoginService)
}
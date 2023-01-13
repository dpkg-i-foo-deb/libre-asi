package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/auth"
	"libre-asi-api/services"
)

func registerRoute() {
	app.AddPost("/register/:role", auth.ValidateAccessToken, services.RegisterService)
}

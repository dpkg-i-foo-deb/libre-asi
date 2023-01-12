package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/auth"
	"libre-asi-api/services"
)

func signOutRoute() {
	app.AddPost("/sign-out", auth.ValidateAndContinue, services.SignOutService)
}

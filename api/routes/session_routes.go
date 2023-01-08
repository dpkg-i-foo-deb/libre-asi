package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/services"
)

func loginRoute() {
	app.AddPost("/login/:role", services.LoginService)
}

func refreshRoute() {
	app.AddGet("/refresh", services.RefreshTokenService)
}

func signOutRoute() {
	app.AddPost("/sign-out", services.SignOutService)
}

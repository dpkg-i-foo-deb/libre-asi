package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/services"
)

func signUpRoute() {
	app.AddPost("/sign-up", services.SignUpService)
}

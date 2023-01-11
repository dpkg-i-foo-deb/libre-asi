package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/services"
)

func registerRoute() {
	app.AddPost("/register/:role", services.RegisterService)
}

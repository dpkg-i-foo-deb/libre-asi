package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/auth"
	"libre-asi-api/services"
)

func indexRoute() {
	app.AddGet("/", auth.ValidateAccessToken, services.IndexService)
}

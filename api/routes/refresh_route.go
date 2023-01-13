package routes

import (
	"libre-asi-api/app"
	"libre-asi-api/auth"
	"libre-asi-api/services"
)

func refreshRoute() {
	app.AddPost("/refresh", auth.CheckRefreshAndContinue, services.RefreshService)
}

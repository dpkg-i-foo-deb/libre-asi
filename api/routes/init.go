package routes

import "github.com/gofiber/fiber/v2"

var server *fiber.App

func SetRoutes(app *fiber.App) {
	server = app
	indexRoute()
}

func init() {
	//	registerRoute()
	//
	// loginRoute()
	// signOutRoute()
	// worldRoute()
	// refreshRoute()
	// adminRoutes()
	// setupRoutes()
}

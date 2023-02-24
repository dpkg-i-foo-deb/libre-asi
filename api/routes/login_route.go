package routes

import "libre-asi-api/handlers"

func loginRoute() {
	server.Post("/login/:role", handlers.LoginHandler)
}

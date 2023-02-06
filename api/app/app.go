package app

import (
	"libre-asi-api/cfg"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var app *fiber.App

func init() {
	app = fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     os.Getenv("CORS_ORIGINS"),
		AllowMethods:     "GET,POST,OPTIONS,PUT,DELETE,PATCH",
		Next:             nil,
		AllowHeaders:     "",
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	//TODO enable limiter when frontend is stable
	//app.Use(limiter.New())

	app.Use(cfg.New(cfg.ConfigDefault))

}

func Start() {
	log.Fatal(app.Listen(os.Getenv("API_PORT")))
}

func AddGet(route string, handlers ...fiber.Handler) {

	app.Get(route, handlers...)

}

func AddPost(route string, handlers ...fiber.Handler) {

	app.Post(route, handlers...)

}

func AddPut(route string, handlers ...fiber.Handler) {

	app.Put(route, handlers...)

}

func AddPatch(route string, handlers ...fiber.Handler) {

	app.Patch(route, handlers...)

}

func AddDelete(route string, handlers ...fiber.Handler) {

	app.Delete(route, handlers...)

}

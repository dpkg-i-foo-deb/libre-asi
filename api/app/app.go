package app

import (
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
}

func Start() {
	log.Fatal(app.Listen(os.Getenv("API_PORT")))
}

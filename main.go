package main

import (
	"statistics-service/src/database"
	"statistics-service/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowCredentials: true,
		},
	))

	routes.SetUp(app)

	app.Listen(":3000")
}

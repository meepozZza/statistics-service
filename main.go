package main

import (
	"github.com/meepozZza/statistics-service/src/database"
	"github.com/meepozZza/statistics-service/src/routes"

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

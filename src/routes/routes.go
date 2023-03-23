package routes

import (
	"github/meepozZza/statistics-server/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	api := app.Group("api")

	api.Get("requests", controllers.RequestController{}.Index)
	api.Post("requests", controllers.RequestController{}.Store)

	api.Get("requests/report", controllers.RequestController{}.Report)
}

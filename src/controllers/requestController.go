package controllers

import (
	"github.com/meepozZza/statistics-service/src/database"
	"github.com/meepozZza/statistics-service/src/models"
	"github.com/meepozZza/statistics-service/src/services"

	"github.com/gofiber/fiber/v2"
)

type RequestController struct {
	requests []models.Request
	request  models.Request
}

func (controller RequestController) Index(c *fiber.Ctx) error {
	database.DB.Find(&controller.requests)

	return c.JSON(controller.requests)
}

func (controller RequestController) Store(c *fiber.Ctx) error {
	if err := c.BodyParser(&controller.request); err != nil {
		return err
	}

	database.DB.Create(&controller.request)

	return c.JSON(controller.request)
}

func (controller RequestController) Report(c *fiber.Ctx) error {
	return c.JSON(services.Report{
		DAU:    services.CalculateDAU(),
		MAU:    services.CalculateMAU(),
		VD:     services.CalculateVD(),
		Views:  services.CalculateViews(),
		Visits: services.CalculateVisits(),
		AC:     services.CalculateAC(),
		RT:     services.CalculateRT(),
	})
}

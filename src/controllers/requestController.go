package controllers

import (
	"statistics-service/src/database"
	"statistics-service/src/models"

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

	// rows, err := database.SqlDB.Query("SELECT t.DAU as DAU FROM (select count(distinct user_id) as DAU, toDate(created_at) as day from requests group by day ) as t LIMIT 1")

	// if err != nil {
	// 	return err
	// }

	// var data models.Report
	// // var age int8

	// err = scan.Row(&data, rows)

	// if err != nil {
	// 	return err
	// }
	// rows, err = database.SqlDB.Query("SELECT t.DAU as MAU FROM (select count(distinct user_id) as DAU, toDate(created_at) as day from requests group by day ) as t LIMIT 1")

	// if err != nil {
	// 	return err
	// }

	// err = scan.Row(&data, rows)

	// if err != nil {
	// 	return err
	// }

	return c.JSON(models.Report{
		DAU: services.getDAU(),
	})
}

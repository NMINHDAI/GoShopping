package controllers

import (
	"GoShopping/database"
	"GoShopping/middleware"
	"GoShopping/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// add new order to the database
func AddOrder(c *fiber.Ctx) error {
	var id = middleware.CheckCookie(c)

	if id == 0 {
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	order := models.Order{
		UserId:   id,
		CreateAt: time.Now(),
		Status:   "Initial",
	}

	database.DB.Create(&order)

	return c.JSON(order)
}

// View all order
func GetAllOrder(c *fiber.Ctx) error {
	var order []models.Order

	result := database.DB.Find(&order)

	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}
	return c.JSON(order)
}

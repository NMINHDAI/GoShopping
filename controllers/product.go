package controllers

import (
	"GoShopping/database"
	"GoShopping/models"

	"github.com/gofiber/fiber/v2"
)

// get all product
func Product(c *fiber.Ctx) error {
	var product []models.Product

	result := database.DB.Find(&product)

	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}
	return c.JSON(product)
}

// get product with id
func GetProduct(c *fiber.Ctx) error {
	var id = c.Params("id")

	var product models.Product

	result := database.DB.First(&product, id)

	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}
	return c.JSON(product)
}

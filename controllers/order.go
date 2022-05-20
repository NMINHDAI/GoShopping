package controllers

import (
	"GoShopping/database"
	"GoShopping/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// add new order to the database
func AddOrder(c *fiber.Ctx) error {
	// return current user
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	var id = user.Id

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

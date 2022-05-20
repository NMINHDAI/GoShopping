package main

import (
	"GoShopping/database"
	"GoShopping/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()

	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	app.Use(
		logger.New(), // add Logger middleware
	)

	routes.Setup(app)

	app.Listen("localhost:8000")
}

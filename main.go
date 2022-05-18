package main

import (
	"GoShopping/database"
	"GoShopping/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	routes.Setup(app)

	app.Listen("localhost:8000")
}

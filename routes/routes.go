package routes

import (
	"GoShopping/controllers"
	"GoShopping/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// User route
	app.Post("/user/register", middleware.CheckCookie(), controllers.Register)
	app.Post("/user/login", controllers.Login)
	app.Get("/user/me", controllers.User)
	// app.Put("/user/update", controllers.Update)
	// app.Delete("/user/update", controllers.Delete)
	app.Post("/user/logout", controllers.Logout)

	// Product route
	app.Get("/products", controllers.Product)
	app.Get("/products/:id", controllers.GetProduct)

	// Order route
	app.Post("/addorder", controllers.AddOrder)
	app.Get("/order", controllers.GetAllOrder)

	// shop route (paid product database)
	app.Post("/addshop", controllers.AddShop)
}

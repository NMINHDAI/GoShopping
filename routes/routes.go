package routes

import (
	"GoShopping/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// User route
	app.Post("/user/register", controllers.Register)
	app.Post("/user/login", controllers.Login)
	app.Get("/user/me", controllers.User)
	// app.Put("/user/update", controllers.Update)
	// app.Delete("/user/update", controllers.Delete)
	app.Post("/user/logout", controllers.Logout)

	// Product route

}

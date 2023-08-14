package main

import (
	"go-mvc/controllers"

	"github.com/gofiber/fiber/v2"
)

func WebRoutes(app *fiber.App) {
	app.Get("/login", controllers.ViewLogin)
	app.Post("/login", controllers.SubmitLogin)
	app.Get("/register", controllers.ViewRegister)
	app.Post("/register", controllers.SubmitRegister)
}

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
}

package main

import (
	"go-mvc/controllers"

	"github.com/gofiber/fiber/v2"
)

func WebRoutes(app *fiber.App) {
	app.Get("/", controllers.PostsIndex)
}

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api") // /api
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
}

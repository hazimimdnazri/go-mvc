package main

import (
	"go-mvc/initializers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	// Define templating engine
	engine := html.New("./views", ".html")

	// Setting up the fiber framework
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure public folder
	app.Static("/", "./public")

	// Routing goes here
	WebRoutes(app)
	ApiRoutes(app)

	// Return error if route not found
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})

	// Define port
	app.Listen(":" + os.Getenv("PORT"))
}

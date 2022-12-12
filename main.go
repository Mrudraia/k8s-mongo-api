package main

import (
	"k8s-mongo-api/configs"
	"k8s-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongodb "})
	// })

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":6000")
}

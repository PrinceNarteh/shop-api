package order

import "github.com/gofiber/fiber/v2"

func OrderRoutes(app *fiber.App) {
	orderRoutes := app.Group("/api/orders")
	orderRoutes.Post("/", CreateOrder)
}

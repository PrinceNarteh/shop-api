package product

import "github.com/gofiber/fiber/v2"

func ProductRoutes(app *fiber.App) {
	productRoutes := app.Group("/api/products")
	productRoutes.Post("/", CreateProduct)
	productRoutes.Get("/", GetProducts)
	productRoutes.Get("/:productId", GetProduct)
}

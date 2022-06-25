package main

import (
	"log"

	"shop_api/pkg/config"
	"shop_api/pkg/order"
	"shop_api/pkg/product"
	"shop_api/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func status(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"status": "Ok"})
}

func main() {
	config.ConnectDb(&order.Order{}, &product.Product{}, &user.User{})
	app := fiber.New()

	app.Get("/api/status", status)
	user.UserRoutes(app)
	product.ProductRoutes(app)

	log.Fatal(app.Listen(":4000"))
}

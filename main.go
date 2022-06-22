package main

import (
	"log"

	"shop_api/config"

	"github.com/gofiber/fiber/v2"
)

func status(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{"status": "Ok"})
}

func main() {
	config.ConnectDb()
	app := fiber.New()

	app.Get("/api/status", status)

	log.Fatal(app.Listen(":4000"))
}

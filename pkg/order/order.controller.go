package order

import (
	"shop_api/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(ctx *fiber.Ctx) error {
	order := new(Order)

	if err := ctx.BodyParser(order); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Create(order)
	return ctx.Status(fiber.StatusCreated).JSON(order)
}

package order

import (
	"github.com/gofiber/fiber/v2"
)

func CreateOrder(ctx *fiber.Ctx) error {
	order := new(Order)

	if err := ctx.BodyParser(order); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Create(order)
	return ctx.Status(fiber.StatusCreated).JSON(order)
}

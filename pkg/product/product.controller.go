package product

import "github.com/gofiber/fiber/v2"

func CreateProduct(ctx *fiber.Ctx) error {
	product := new(Product)

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Create(product)
	return ctx.Status(fiber.StatusCreated).JSON(product)
}

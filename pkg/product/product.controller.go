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

func GetProducts(ctx *fiber.Ctx) error {
	products := make([]Product, 0)
	FindProducts(&products)
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func GetProduct(ctx *fiber.Ctx) error {
	product := new(Product)

	productId, err := ctx.ParamsInt("productId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please make sure to provide product ID as integer"})
	}

	if err = FindProduct(product, productId); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(product)
}

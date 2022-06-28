package product

import (
	"shop_api/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(ctx *fiber.Ctx) error {
	product := new(Product)

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Create(product)
	return ctx.Status(fiber.StatusCreated).JSON(product)
}

func GetProducts(ctx *fiber.Ctx) error {
	products := make([]Product, 0)
	config.DB.Find(&products)
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func GetProduct(ctx *fiber.Ctx) error {
	product := new(Product)

	productId, err := ctx.ParamsInt("productId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please make sure to provide product ID as integer"})
	}

	config.DB.Find(product, productId)
	if product.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	product := new(Product)
	type UpdateProduct struct {
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		Price        float32 `json:"price"`
		SerialNumber string  `json:"serialNumber"`
	}

	productId, err := ctx.ParamsInt("productId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	body := new(UpdateProduct)
	if err = ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Find(product, productId)
	if product.ID == 0 {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"error": err.Error()})
	}

	if body.Price != 0 {
		product.Price = body.Price
	}
	if body.SerialNumber != "" {
		product.SerialNumber = body.SerialNumber
	}
	if body.Description != "" {
		product.Description = body.Description
	}
	if body.Name != "" {
		product.Name = body.Name
	}

	config.DB.Save(product)
	return ctx.Status(fiber.StatusOK).JSON(product)
}

func DeleteProduct(ctx *fiber.Ctx) error {
	product := new(Product)

	productId, err := ctx.ParamsInt("productId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Find(product, productId)
	if product.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Delete(product)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Successfully deleted product."})
}

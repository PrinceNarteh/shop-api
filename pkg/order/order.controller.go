package order

import (
	"fmt"
	"shop_api/pkg/config"
	"shop_api/pkg/product"
	"shop_api/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(ctx *fiber.Ctx) error {
	order := new(Order)

	if err := ctx.BodyParser(order); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := new(user.User)
	config.DB.Find(&user, order.UserId)
	if user.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("user with id %d not found", order.UserId)})
	}

	product := new(product.Product)
	config.DB.Find(&product, order.ProductId)
	if product.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("product with id %d not found", order.ProductId)})
	}

	if order.Quantity == 0 {
		order.Quantity = 1
	}

	config.DB.Create(order)
	orderResponse := Order{
		ID:      order.ID,
		User:    *user,
		Product: *product,
	}
	return ctx.Status(fiber.StatusCreated).JSON(orderResponse)
}

package user

import (
	"shop_api/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func CreateUserHandler(ctx *fiber.Ctx) error {
	var user User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	config.Database.Db.Create(&user)
	return ctx.Status(201).JSON(user)
}

package user

import (
	"log"
	"shop_api/pkg/config"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(ctx *fiber.Ctx) error {
	var user User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hashedPassword)

	config.Database.Db.Create(&user)
	return ctx.Status(201).JSON(user)
}

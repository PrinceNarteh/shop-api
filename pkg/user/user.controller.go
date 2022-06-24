package user

import (
	"fmt"
	"log"
	"shop_api/pkg/config"
	"shop_api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserHandler(ctx *fiber.Ctx) error {
	var user User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	result := config.Database.Db.Where("email = ?", user.Email).Find(&user)
	if result.RowsAffected > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already in use."})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hashedPassword)
	config.Database.Db.Create(&user)

	return ctx.Status(201).JSON(user)
}

func LoginUserHandler(ctx *fiber.Ctx) error {
	var user User
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	result := config.Database.Db.Where("email = ?", body.Email).Find(&user)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials 1"})
	}

	isMatch := utils.VerifyPassword(body.Password, user.Password)
	fmt.Println(isMatch)

	if !isMatch {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials 2"})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

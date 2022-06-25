package user

import (
	"log"
	"shop_api/pkg/config"
	"shop_api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserHandler(ctx *fiber.Ctx) error {
	user := new(User)

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	errors := utils.ValidateStruct(*user)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
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
	user := new(User)
	var data struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	errors := utils.ValidateStruct(data)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	result := config.Database.Db.Where("email = ?", data.Email).Find(&user)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials 1"})
	}

	isMatch := utils.VerifyPassword(data.Password, user.Password)
	if !isMatch {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials 2"})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func GetUsers(ctx *fiber.Ctx) error {
	users := make([]User, 0)
	config.Database.Db.Find(&users)
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func GetUser(ctx *fiber.Ctx) error {
	user := new(User)

	userId, err := ctx.ParamsInt("userId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please make sure to provide user ID as integer"})
	}

	result := config.Database.Db.Find(&user, userId)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User Not Found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

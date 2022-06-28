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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errors := utils.ValidateStruct(*user)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	config.DB.Find(user, "email = ?", user.Email)
	if user.ID > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already in used"})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hashedPassword)
	config.DB.Create(&user)

	userResponse := CreateUserResponse(user)

	return ctx.Status(fiber.StatusCreated).JSON(userResponse)
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

	config.DB.Find(user, "email = ?", data.Email)
	if user.ID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	isMatch := utils.VerifyPassword(data.Password, user.Password)
	if !isMatch {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	userResponse := CreateUserResponse(user)

	return ctx.Status(fiber.StatusOK).JSON(userResponse)
}

func GetUsers(ctx *fiber.Ctx) error {
	users := make([]User, 0)
	config.DB.Find(&users)

	userResponses := make([]UserResponse, 0)
	for _, user := range users {
		userResponse := CreateUserResponse(&user)
		userResponses = append(userResponses, userResponse)
	}

	return ctx.Status(fiber.StatusOK).JSON(userResponses)
}

func GetUser(ctx *fiber.Ctx) error {
	user := new(User)

	userId, err := ctx.ParamsInt("userId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please make sure to provide user ID as integer"})
	}

	config.DB.Find(user, userId)
	if user.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	userResponse := CreateUserResponse(user)

	return ctx.Status(fiber.StatusOK).JSON(userResponse)
}

func UpdateUser(ctx *fiber.Ctx) error {
	user := new(User)
	type UpdateUser struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}
	data := new(UpdateUser)

	userId, err := ctx.ParamsInt("userId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please make sure to provide user ID as integer"})
	}

	config.DB.Find(user, userId)
	if user.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if data.Email != "" {
		user.Email = data.Email
	}
	if data.FirstName != "" {
		user.FirstName = data.FirstName
	}
	if data.LastName != "" {
		user.LastName = data.LastName
	}

	config.DB.Find(user)
	userResponse := CreateUserResponse(user)

	return ctx.Status(fiber.StatusOK).JSON(userResponse)
}

func DeleteUser(ctx *fiber.Ctx) error {
	user := new(User)

	userId, err := ctx.ParamsInt("userId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please make sure to provide user ID as integer"})
	}

	config.DB.Find(user, userId)
	if user.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Find(user)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Successfully deleted user."})
}

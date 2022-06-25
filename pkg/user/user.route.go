package user

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {
	userRoutes := app.Group("/api/users")
	userRoutes.Post("/register", RegisterUserHandler)
	userRoutes.Post("/login", LoginUserHandler)
	userRoutes.Get("/", GetUsers)
	userRoutes.Get("/:userId", GetUser)
	userRoutes.Patch("/:userId", UpdateUser)
	userRoutes.Delete("/:userId", DeleteUser)
}

package user

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {
	userRouters := app.Group("/api/users")
	userRouters.Post("/register", RegisterUserHandler)
	userRouters.Post("/login", LoginUserHandler)
	userRouters.Get("/", GetUsers)
	userRouters.Get("/:userId", GetUser)
	userRouters.Patch("/:userId", UpdateUser)
}

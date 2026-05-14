package routes

import (
	"database/sql"

	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handlers"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/persistence"
	"github.com/gofiber/fiber/v3"
)

func SetupRouter(db *sql.DB) *fiber.App {
	app := fiber.New()

	userRepo := persistence.NewUserRepository(db)
	userService := application.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	api := app.Group("/api")
	api.Get("/health", handlers.GetHealth)

	SetupUserRoutes(api, userHandler)

	return app
}

func SetupUserRoutes(fiber fiber.Router, handler *handlers.UserHandler) {
	//Basic CRUD
	fiber.Post("/users", handler.Create)
	fiber.Get("/users/:id", handler.GetByID)
	fiber.Put("/users/:id", handler.Update)
	fiber.Delete("/users/:id", handler.Delete)

	//Buscas
	fiber.Post("/users/by-email", handler.GetByEmail)

	//Exists
	fiber.Post("/users/exists-by-email", handler.ExistsByEmail)
	fiber.Post("/users/exists-by-cpf", handler.ExistsByCPF)
}

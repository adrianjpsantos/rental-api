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
	api.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API funcionando",
		})
	})

	SetupUserRoutes(api, userHandler)

	return app
}

func SetupUserRoutes(fiber fiber.Router, handler *handlers.UserHandler) {

	fiber.Post("/users/exists-by-email", handler.ExistsByEmail) // ou Get, dependendo do caso
}

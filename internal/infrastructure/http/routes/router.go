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

func SetupUserRoutes(router fiber.Router, handler *handlers.UserHandler) {
	//Basic CRUD
	router.Post("/users", handler.Create)
	router.Get("/users/:id", handler.GetByID)
	router.Put("/users/:id", handler.Update)
	router.Delete("/users/:id", handler.Delete)

	//Buscas
	router.Post("/users/by-email", handler.GetByEmail)

	//Exists
	router.Post("/users/exists-by-email", handler.ExistsByEmail)
	router.Post("/users/exists-by-cpf", handler.ExistsByCPF)
}

func SetupReviewRoutes(router fiber.Router, handler *handlers.ReviewHandler) {
	router.Post("/reviews", handler.Create)
	router.Get("/reviews/:id", handler.GetByID)

	//Other Reads
	router.Get("/rentals/:rental_id/reviews", handler.GetByRentalID)

	router.Get("/users/:user_id/reviews/received", handler.GetReceivedReviews)
	router.Get("/users/:user_id/reviews/given", handler.GetGivenReviews)

	router.Get("/users/:user_id/reviews", handler.GetUserReviews)
}

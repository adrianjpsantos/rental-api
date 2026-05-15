package routes

import (
	"database/sql"

	_ "github.com/adrianjpsantos/rental-api/docs"
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handlers"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/persistence"
	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

func SetupRouter(db *sql.DB) *fiber.App {
	app := fiber.New()

	userRepo := persistence.NewUserRepository(db)
	userService := application.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	api := app.Group("/api")
	api.Get("/health", handlers.GetHealth)

	SetupSwagger(app)
	SetupUserRoutes(api, userHandler)

	return app
}

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swaggo.HandlerDefault)

	// Customize the UI by passing a Config
	app.Get("/docs/*", swaggo.New(swaggo.Config{
		URL:               "http://example.com/doc.json",
		DeepLinking:       false,
		DocExpansion:      "none",
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
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

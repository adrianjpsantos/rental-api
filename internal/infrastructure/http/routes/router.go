package routes

import (
	"database/sql"

	_ "github.com/adrianjpsantos/rental-api/docs"
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handlers"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/persistence"
	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

func SetupRouter(db *sql.DB) *fiber.App {
	app := fiber.New()

	//Repositories, Services e Handlers
	repositories := persistence.NewAllRepositories(db)
	services := application.NewAllServices(repositories)
	apiHandlers := handlers.NewAllHandlers(services)

	api := app.Group("/api")
	api.Get("/health", handlers.GetHealth)

	SetupSwagger(app)
	SetupUserRoutes(api, apiHandlers.UserHandler)
	SetupReviewRoutes(api, apiHandlers.ReviewHandler)
	SetupRentalRoutes(api, apiHandlers.RentalHandler)
	SetupItemRoutes(api, apiHandlers.ItemHandler)
	//SetupAvailabilityRoutes(api, apiHandlers.AvailabilityHandler)
	//SetupCategoryRoutes(api, apiHandlers.CategoryHandler)

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
	router.Get("/users/:user_id", handler.GetByID)
	router.Put("/users/:user_id", handler.Update)
	router.Delete("/users/:user_id", handler.Delete)

	//Buscas
	router.Post("/users/by-email", handler.GetByEmail)

	//Exists
	router.Post("/users/exists-by-email", handler.ExistsByEmail)
	router.Post("/users/exists-by-cpf", handler.ExistsByCPF)
}

func SetupReviewRoutes(router fiber.Router, handler *handlers.ReviewHandler) {
	router.Post("/reviews", handler.Create)
	router.Get("/reviews/:review_id", handler.GetByID)

	//Other Reads
	router.Get("/rentals/:rental_id/reviews", handler.GetByRentalID)

	router.Get("/users/:user_id/reviews/received", handler.GetReceivedReviews)
	router.Get("/users/:user_id/reviews/given", handler.GetGivenReviews)

	router.Get("/users/:user_id/reviews", handler.GetUserReviews)
}

func SetupRentalRoutes(router fiber.Router, handler *handlers.RentalHandler) {
	router.Post("/rentals", handler.Create)
	router.Get("/rentals/:rental_id", handler.GetByID)
	router.Put("/rentals/:rental_id/status", handler.UpdateStatus)
	router.Put("/rentals/:rental_id/cancel", handler.Cancel)

	//Other Reads
	router.Get("/users/:user_id/rentals", handler.GetAllUserRentals)
	router.Get("/users/:user_id/rentals/lessee", handler.GetUserRentalsAsLessee)
	router.Get("/users/:user_id/rentals/lessor", handler.GetUserRentalsAsLessor)
}

func SetupItemRoutes(router fiber.Router, handler item.InterfaceItemHandler) {
	router.Post("/items", handler.CreateItem)
	router.Get("/items/:item_id", handler.GetItemByID)
	router.Put("/items/:item_id", handler.UpdateItem)
	router.Delete("/items/:item_id", handler.DeleteItem)

	//Other Reads
	router.Post("/items/search", handler.ListItems)
}

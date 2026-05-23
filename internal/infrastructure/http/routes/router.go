package routes

import (
	"database/sql"

	_ "github.com/adrianjpsantos/rental-api/docs"
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handlers"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/repositories"
	"github.com/adrianjpsantos/rental-api/internal/pkg/middleware/authentication"

	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

func SetupRouter(db *sql.DB) *fiber.App {
	app := fiber.New()

	repositories := repositories.NewAllRepositories(db)
	services := application.NewAllServices(repositories)
	apiHandlers := handlers.NewAllHandlers(services)

	SetupSwagger(app)

	api := app.Group("/api")

	api.Get("/health", handlers.GetHealth)

	public := api.Group("/")

	private := api.Group(
		"/",
		authentication.AuthMiddleware(
			services.TokenService,
		),
	)

	authGroup := public.Group("/auth")

	userGroup := private.Group("/users")
	reviewGroup := private.Group("/reviews")
	rentalGroup := private.Group("/rentals")
	itemGroup := private.Group("/items")
	availabilityGroup := private.Group("/availability")
	categoryGroup := private.Group("/categories")

	SetupAuthRoutes(authGroup, apiHandlers.AuthHandler)

	SetupUserRoutes(userGroup, apiHandlers.UserHandler)

	SetupReviewRoutes(
		reviewGroup,
		userGroup,
		rentalGroup,
		apiHandlers.ReviewHandler,
	)

	SetupRentalRoutes(
		rentalGroup,
		userGroup,
		apiHandlers.RentalHandler,
	)

	SetupItemRoutes(
		itemGroup,
		apiHandlers.ItemHandler,
	)

	SetupAvailabilityRoutes(
		availabilityGroup,
		itemGroup,
		apiHandlers.AvailabilityHandler,
	)

	SetupCategoryRoutes(
		categoryGroup,
		apiHandlers.CategoryHandler,
	)

	return app
}

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swaggo.HandlerDefault)

	app.Get("/docs/*", swaggo.New(swaggo.Config{
		URL:               "http://example.com/doc.json",
		DeepLinking:       false,
		DocExpansion:      "none",
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
}

func SetupAuthRoutes(
	router fiber.Router,
	handler *handlers.AuthHandler,
) {
	router.Post("/login", handler.Authenticate)
	router.Post("/refresh", handler.Refresh)
	router.Post("/logout", handler.Logout)
}

func SetupUserRoutes(
	router fiber.Router,
	handler *handlers.UserHandler,
) {
	router.Post("/", handler.Create)
	router.Get("/:user_id", handler.GetByID)
	router.Put("/:user_id", handler.Update)
	router.Delete("/:user_id", handler.Delete)

	router.Get("/by-email", handler.GetByEmail)

	router.Get("/exists-by-email", handler.ExistsByEmail)
	router.Get("/exists-by-cpf", handler.ExistsByCPF)
}

func SetupReviewRoutes(
	reviewRouter fiber.Router,
	userRouter fiber.Router,
	rentalRouter fiber.Router,
	handler *handlers.ReviewHandler,
) {
	reviewRouter.Post("/", handler.Create)
	reviewRouter.Get("/:review_id", handler.GetByID)

	rentalRouter.Get("/:rental_id/reviews", handler.GetByRentalID)

	userRouter.Get("/:user_id/reviews", handler.GetUserReviews)
	userRouter.Get("/:user_id/reviews/received", handler.GetReceivedReviews)
	userRouter.Get("/:user_id/reviews/given", handler.GetGivenReviews)
}

func SetupRentalRoutes(
	rentalRouter fiber.Router,
	userRouter fiber.Router,
	handler *handlers.RentalHandler,
) {
	rentalRouter.Post("/", handler.Create)
	rentalRouter.Get("/:rental_id", handler.GetByID)
	rentalRouter.Put("/:rental_id/status", handler.UpdateStatus)
	rentalRouter.Put("/:rental_id/cancel", handler.Cancel)

	userRouter.Get("/:user_id/rentals", handler.GetAllUserRentals)
	userRouter.Get("/:user_id/rentals/lessee", handler.GetUserRentalsAsLessee)
	userRouter.Get("/:user_id/rentals/lessor", handler.GetUserRentalsAsLessor)
}

func SetupItemRoutes(
	router fiber.Router,
	handler *handlers.ItemHandler,
) {
	router.Post("/", handler.CreateItem)
	router.Get("/:item_id", handler.GetItemByID)
	router.Put("/:item_id", handler.UpdateItem)
	router.Delete("/:item_id", handler.DeleteItem)

	router.Get("/", handler.ListItems)
}

func SetupAvailabilityRoutes(
	availabilityRouter fiber.Router,
	itemRouter fiber.Router,
	handler *handlers.AvailabilityHandler,
) {
	availabilityRouter.Post("/", handler.Create)
	availabilityRouter.Get("/:availability_id", handler.GetByID)
	availabilityRouter.Delete("/:availability_id", handler.Delete)

	itemRouter.Get("/:item_id/availability", handler.GetByItemID)
	itemRouter.Post("/:item_id/check-availability", handler.CheckAvailability)
}

func SetupCategoryRoutes(
	router fiber.Router,
	handler *handlers.CategoryHandler,
) {
	router.Post("/", handler.CreateCategory)
	router.Get("/", handler.ListCategories)
	router.Get("/:category_id", handler.GetCategoryByID)
	router.Put("/:category_id", handler.UpdateCategory)
	router.Delete("/:category_id", handler.DeleteCategory)
}

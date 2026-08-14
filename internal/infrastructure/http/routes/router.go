package routes

import (
	"database/sql"

	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handlers"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/repositories"
	"github.com/adrianjpsantos/rental-api/internal/pkg/middleware/authentication"
	"github.com/adrianjpsantos/rental-api/internal/security/validator"
	"github.com/adrianjpsantos/rental-api/internal/uow"
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

func SetupRouter(db *sql.DB) *fiber.App {
	app := fiber.New()

	validator.Init()
	uow := uow.NewUow(db)
	repositories := repositories.NewAllRepositories(db)
	services := application.NewAllServices(uow, repositories)
	apiHandlers := handlers.NewAllHandlers(services)

	SetupSwagger(app)

	api := app.Group("/api")

	api.Get("/health", handlers.GetHealth)

	v1 := api.Group("/v1")

	authGroup := v1.Group("/auth")
	SetupAuthRoutes(authGroup, apiHandlers.AuthHandler)

	authMiddleware := authentication.AuthMiddleware(
		services.SessionService,
	)

	userGroup := v1.Group("/user", authMiddleware)
	reviewGroup := v1.Group("/review", authMiddleware)
	rentalGroup := v1.Group("/rental", authMiddleware)
	itemGroup := v1.Group("/item")
	availabilityGroup := v1.Group("/availability", authMiddleware)
	categoryGroup := v1.Group("/category", authMiddleware)

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
		authMiddleware,
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
	router.Post("/register", handler.Register)
	router.Post("/login", handler.LoginLocal)
	router.Post("/refresh", handler.Refresh)
	router.Post("/logout", handler.Logout)

}

func SetupUserRoutes(
	router fiber.Router,
	handler *handlers.UserHandler,
) {

	router.Get("/me", handler.GetCurrentUser)
	router.Put("/me", handler.UpdateCurrentUser)
	router.Delete("/me", handler.DeleteAccount)

	router.Get("/:user_id", handler.GetByID)
	//router.Put("/:user_id", handler.Update)
	//router.Delete("/:user_id", handler.Delete)
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
	authMiddleware fiber.Handler,
) {
	router.Post("/", handler.CreateItem, authMiddleware)
	router.Get("/:item_id", handler.GetItemByID, authMiddleware)
	router.Put("/:item_id", handler.UpdateItem, authMiddleware)
	router.Delete("/:item_id", handler.DeleteItem, authMiddleware)

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

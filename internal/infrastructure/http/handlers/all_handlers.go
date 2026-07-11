package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	"github.com/gofiber/fiber/v3"
)

type AllHandlers struct {
	UserHandler         *UserHandler
	ReviewHandler       *ReviewHandler
	RentalHandler       *RentalHandler
	ItemHandler         *ItemHandler
	AvailabilityHandler *AvailabilityHandler
	CategoryHandler     *CategoryHandler
	AuthHandler         *AuthHandler
}

func NewAllHandlers(services *application.AllServices) *AllHandlers {
	userHandler := NewUserHandler(
		services.UserService,
	)

	reviewHandler := NewReviewHandler(
		services.ReviewService,
	)

	rentalHandler := NewRentalHandler(
		services.RentalService,
	)

	itemHandler := NewItemHandler(
		services.ItemService,
	)

	availabilityHandler := NewAvailabilityHandler(
		services.AvailabilityService,
	)

	categoryHandler := NewCategoryHandler(
		services.CategoryService,
	)

	authHandler := NewAuthHandler(services.AuthService, services.SessionService)

	return &AllHandlers{
		UserHandler:         userHandler,
		ReviewHandler:       reviewHandler,
		RentalHandler:       rentalHandler,
		ItemHandler:         itemHandler,
		AvailabilityHandler: availabilityHandler,
		CategoryHandler:     categoryHandler,
		AuthHandler:         authHandler,
	}
}

func ResponseSuccess(c fiber.Ctx, data any) error {
	return c.JSON(Response{
		Success: true,
		Data:    data,
	})
}

func ResponseAuthError(c fiber.Ctx, status int, err string) error {
	secureCookie := config.LoadConfig().IsProduction()
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   secureCookie,
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
	})

	return c.Status(status).JSON(Response{
		Success: false,
		Error:   err,
	})
}

func ResponseAuthSuccess(c fiber.Ctx, output *authenticate.AuthenticateOutput) error {
	secureCookie := config.LoadConfig().IsProduction()
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    output.RefreshToken,
		HTTPOnly: true,
		Secure:   secureCookie,
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
	})

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"access_token": output.AccessToken,
		},
	})
}

func ResponseError(c fiber.Ctx, status int, err string) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Error:   err,
	})
}

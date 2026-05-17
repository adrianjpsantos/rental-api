package application

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/token"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/repositories"
)

type AllServices struct {
	UserService         *UserService
	ReviewService       *ReviewService
	RentalService       *RentalService
	ItemService         *ItemService
	AvailabilityService *AvailabilityService
	CategoryService     *CategoryService
	AuthService         authenticate.InterfaceAuthenticateService
	TokenService        token.InterfaceTokenService
}

func NewAllServices(repos *repositories.AllRepositories) *AllServices {
	userService := NewUserService(repos.UserRepo)

	reviewService := NewReviewService(
		repos.ReviewRepo,
		repos.UserRepo,
	)

	rentalService := NewRentalService(
		repos.RentalRepo,
	)

	itemService := NewItemService(
		repos.ItemRepo,
	)

	availabilityService := NewAvailabilityService(
		repos.AvailabilityRepo,
	)

	categoryService := NewCategoryService(
		repos.CategoryRepo,
	)

	tokenService := NewTokenService()

	authService := NewAuthService(repos.UserRepo, tokenService)

	return &AllServices{
		UserService:         userService,
		ReviewService:       reviewService,
		RentalService:       rentalService,
		ItemService:         itemService,
		AvailabilityService: availabilityService,
		CategoryService:     categoryService,
		TokenService:        tokenService,
		AuthService:         authService,
	}
}

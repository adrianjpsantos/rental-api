package application

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/repositories"
)

type AllServices struct {
	UserService         user.Service
	ReviewService       review.Service
	RentalService       rental.Service
	ItemService         item.Service
	AvailabilityService availability.Service
	CategoryService     category.Service
	AuthService         authenticate.Service
	SessionService      session.Service
}

func NewAllServices(repos *repositories.AllRepositories) *AllServices {

	googleService := NewGoogleService(config.LoadConfig().GoogleClientID)
	userService := NewUserService(repos.UserRepo)

	reviewService := NewReviewService(
		repos.ReviewRepo,
		userService,
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

	sessionService := NewSessionService(repos.SessionRepo)

	authService := NewAuthService(
		userService,
		sessionService,
		googleService,
	)

	return &AllServices{
		UserService:         userService,
		ReviewService:       reviewService,
		RentalService:       rentalService,
		ItemService:         itemService,
		AvailabilityService: availabilityService,
		CategoryService:     categoryService,
		SessionService:      sessionService,
		AuthService:         authService,
	}
}

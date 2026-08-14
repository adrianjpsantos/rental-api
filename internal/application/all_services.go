package application

import (
	authaccount "github.com/adrianjpsantos/rental-api/internal/domain/auth_account"
	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/domain/profile"
	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/repositories"
	"github.com/adrianjpsantos/rental-api/internal/uow"
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
	AuthAccount         authaccount.Service
	Profile             profile.Service
}

func NewAllServices(uow uow.UnitOfWork, repos *repositories.AllRepositories) *AllServices {

	userService := NewUserService(repos.User)

	authAccountService := NewAuthAccountService(repos.AuthAccount)

	reviewService := NewReviewService(
		repos.Review,
		userService,
	)

	rentalService := NewRentalService(
		repos.Rental,
	)

	itemService := NewItemService(
		repos.Item,
	)

	availabilityService := NewAvailabilityService(
		repos.Availability,
	)

	categoryService := NewCategoryService(
		repos.Category,
	)

	sessionService := NewSessionService(repos.Session)

	authService := NewAuthService(
		uow,
		sessionService,
	)

	profileService := NewProfileService(
		repos.Profile,
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
		AuthAccount:         authAccountService,
		Profile:             profileService,
	}
}

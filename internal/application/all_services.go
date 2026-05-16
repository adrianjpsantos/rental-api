package application

import "github.com/adrianjpsantos/rental-api/internal/infrastructure/persistence"

type AllServices struct {
	UserService         *UserService
	ReviewService       *ReviewService
	RentalService       *RentalService
	ItemService         *ItemService
	AvailabilityService *AvailabilityService
	CategoryService     *CategoryService
}

func NewAllServices(repos *persistence.AllRepositories) *AllServices {
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

	return &AllServices{
		UserService:         userService,
		ReviewService:       reviewService,
		RentalService:       rentalService,
		ItemService:         itemService,
		AvailabilityService: availabilityService,
		CategoryService:     categoryService,
	}
}

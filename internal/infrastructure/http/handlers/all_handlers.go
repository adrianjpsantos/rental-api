package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
)

type AllHandlers struct {
	UserHandler   *UserHandler
	ReviewHandler *ReviewHandler
	RentalHandler *RentalHandler
	ItemHandler   item.InterfaceItemHandler
	//AvailabilityHandler *AvailabilityHandler
	//CategoryHandler     *CategoryHandler
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

	//availabilityHandler := NewAvailabilityHandler(
	//	services.AvailabilityService,
	//)

	//categoryHandler := NewCategoryHandler(
	//	services.CategoryService,
	//)

	return &AllHandlers{
		UserHandler:   userHandler,
		ReviewHandler: reviewHandler,
		RentalHandler: rentalHandler,
		ItemHandler:   itemHandler,
		//	AvailabilityHandler: availabilityHandler,
		//	CategoryHandler:     categoryHandler,
	}
}

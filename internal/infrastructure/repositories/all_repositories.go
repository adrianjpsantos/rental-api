package repositories

import (
	"database/sql"

	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
)

type AllRepositories struct {
	UserRepo         user.Repository
	ReviewRepo       review.Repository
	RentalRepo       rental.Repository
	ItemRepo         item.Repository
	AvailabilityRepo availability.Repository
	CategoryRepo     category.Repository
	SessionRepo      session.Repository
}

func NewAllRepositories(db *sql.DB) *AllRepositories {
	userRepo := NewUserRepository(db)
	reviewRepo := NewReviewRepository(db)
	rentalRepo := NewRentalRepository(db)
	itemRepo := NewItemRepository(db)
	availabilityRepo := NewAvailabilityRepository(db)
	categoryRepo := NewCategoryRepository(db)
	sessionRepo := NewSessionRepository(db)

	return &AllRepositories{
		UserRepo:         userRepo,
		ReviewRepo:       reviewRepo,
		RentalRepo:       rentalRepo,
		ItemRepo:         itemRepo,
		AvailabilityRepo: availabilityRepo,
		CategoryRepo:     categoryRepo,
		SessionRepo:      sessionRepo,
	}
}

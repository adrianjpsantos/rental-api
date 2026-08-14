package repositories

import (
	authaccount "github.com/adrianjpsantos/rental-api/internal/domain/auth_account"
	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/domain/profile"
	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
)

type AllRepositories struct {
	User         user.Repository
	Review       review.Repository
	Rental       rental.Repository
	Item         item.Repository
	Availability availability.Repository
	Category     category.Repository
	Session      session.Repository
	Profile      profile.Repository
	AuthAccount  authaccount.Repository
}

func NewAllRepositories(db DBTX) *AllRepositories {
	userRepo := NewUserRepository(db)
	reviewRepo := NewReviewRepository(db)
	rentalRepo := NewRentalRepository(db)
	itemRepo := NewItemRepository(db)
	availabilityRepo := NewAvailabilityRepository(db)
	categoryRepo := NewCategoryRepository(db)
	sessionRepo := NewSessionRepository(db)
	authAccountRepo := NewAuthAccountRepository(db)
	profileRepo := NewProfileRepository(db)

	return &AllRepositories{
		User:         userRepo,
		Review:       reviewRepo,
		Rental:       rentalRepo,
		Item:         itemRepo,
		Availability: availabilityRepo,
		Category:     categoryRepo,
		Session:      sessionRepo,
		AuthAccount:  authAccountRepo,
		Profile:      profileRepo,
	}
}

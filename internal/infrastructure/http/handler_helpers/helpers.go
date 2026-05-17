package handler_helpers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func ParseRentalStatusQuery(c fiber.Ctx) (*rental.Status, error) {
	q := c.Query("status")

	if q == "" {
		return nil, nil
	}

	return rental.ParseRentalStatus(q)
}

func ParseReviewTypeQuery(c fiber.Ctx) (*review.ReviewType, error) {
	q := c.Query("type")

	if q == "" {
		return nil, nil
	}

	return review.ParseReviewType(q)
}

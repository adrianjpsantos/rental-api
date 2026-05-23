package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/parses"
	"github.com/gofiber/fiber/v3"
)

type ReviewHandler struct {
	service review.Service
}

func NewReviewHandler(service review.Service) *ReviewHandler {
	return &ReviewHandler{service: service}
}

// CRUD - CREATE
func (h *ReviewHandler) Create(c fiber.Ctx) error {

	input, err := parses.ParseBody[review.ReviewCreateInput](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Create(c.Context(), input)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Created Review",
	},
	)
}

// CRUD - READ
func (h *ReviewHandler) GetByID(c fiber.Ctx) error {
	reviewId, err := parses.ParseUUIDParam(c, "review_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingReview, err := h.service.GetByID(c.Context(), *reviewId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"review": existingReview,
	})
}

// CRUD - READ - get by rental id
func (h *ReviewHandler) GetByRentalID(c fiber.Ctx) error {
	rentalId, err := parses.ParseUUIDParam(c, "rental_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingReview, err := h.service.GetByRentalID(c.Context(), *rentalId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"review": existingReview,
	},
	)
}

// CRUD - READ - get by reviewed id ( user )
func (h *ReviewHandler) GetReceivedReviews(c fiber.Ctx) error {
	userId, err := parses.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	rtQuery := parses.ParseStringQuery(c, "type")
	reviewType, err := review.ParseReviewType(*rtQuery)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	existingReviews, err := h.service.GetReceivedReviews(c.Context(), *userId, reviewType)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"reviews": existingReviews, // listagem de avaliações
	},
	)
}

func (h *ReviewHandler) GetGivenReviews(c fiber.Ctx) error {
	userId, err := parses.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	rtQuery := parses.ParseStringQuery(c, "type")
	reviewType, err := review.ParseReviewType(*rtQuery)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	existingReviews, err := h.service.GetGivenReviews(c.Context(), *userId, reviewType)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"reviews": existingReviews, // listagem de avaliações
	},
	)
}

func (h *ReviewHandler) GetUserReviews(c fiber.Ctx) error {
	userId, err := parses.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	rtQuery := parses.ParseStringQuery(c, "type")
	reviewType, err := review.ParseReviewType(*rtQuery)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	existingReviews, err := h.service.GetUserReviews(c.Context(), *userId, reviewType)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"reviews": existingReviews, // listagem de avaliações
	},
	)
}

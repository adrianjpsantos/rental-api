package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handler_helpers"
	"github.com/gofiber/fiber/v3"
)

type ReviewHandler struct {
	service *application.ReviewService
}

func NewReviewHandler(service *application.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

// CRUD - CREATE
func (h *ReviewHandler) Create(c fiber.Ctx) error {
	var req review.ReqCreate

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "JSON Inválido",
		})
	}

	newReview := req.NewReview

	createdReview, err := h.service.Register(c.Context(), newReview.RentalID, newReview.ReviewerID, newReview.ReviewedID, newReview.ItemID, newReview.Rating, newReview.Comment, newReview.ReviewType)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"created-review-id": createdReview.Id,
		},
	})
}

// CRUD - READ
func (h *ReviewHandler) GetByID(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "review_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	existingReview, err := h.service.GetByID(c.Context(), reqId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"review": existingReview,
		},
	})
}

// CRUD - READ - get by rental id
func (h *ReviewHandler) GetByRentalID(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "rental_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	existingReview, err := h.service.GetByRentalID(c.Context(), reqId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"review": existingReview,
		},
	})
}

// CRUD - READ - get by reviewed id ( user )
func (h *ReviewHandler) GetReceivedReviews(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	reviewType, err := handler_helpers.ParseReviewTypeQuery(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	existingReviews, err := h.service.GetReceivedReviews(c.Context(), reqId, reviewType)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"reviews": existingReviews, // listagem de avaliações
		},
	})
}

func (h *ReviewHandler) GetGivenReviews(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	reviewType, err := handler_helpers.ParseReviewTypeQuery(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	existingReviews, err := h.service.GetGivenReviews(c.Context(), reqId, reviewType)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"reviews": existingReviews, // listagem de avaliações
		},
	})
}

func (h *ReviewHandler) GetUserReviews(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	reviewType, err := handler_helpers.ParseReviewTypeQuery(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	existingReviews, err := h.service.GetUserReviews(c.Context(), reqId, reviewType)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"reviews": existingReviews, // listagem de avaliações
		},
	})
}

package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
	reqId := c.Params("id")
	if _, err := uuid.Parse(reqId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	existingReview, err := h.service.GetByID(c.Context(), uuid.MustParse(reqId))

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
	reqId := c.Params("rental_id")
	if _, err := uuid.Parse(reqId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	existingReview, err := h.service.GetByRentalID(c.Context(), uuid.MustParse(reqId))

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
	reqId := c.Params("user_id")
	if _, err := uuid.Parse(reqId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	var reviewType *review.ReviewType
	if q := c.Query("type"); q != "" {
		rt, err := review.ParseReviewType(q)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Success: false,
				Error:   err.Error(),
			})
		}

		reviewType = rt
	}

	existingReviews, err := h.service.GetReceivedReviews(c.Context(), uuid.MustParse(reqId), reviewType)

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
	reqId := c.Params("user_id")
	if _, err := uuid.Parse(reqId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	var reviewType *review.ReviewType
	if q := c.Query("type"); q != "" {
		rt, err := review.ParseReviewType(q)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Success: false,
				Error:   err.Error(),
			})
		}

		reviewType = rt
	}

	existingReviews, err := h.service.GetGivenReviews(c.Context(), uuid.MustParse(reqId), reviewType)

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
	reqId := c.Params("user_id")
	if _, err := uuid.Parse(reqId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	var reviewType *review.ReviewType
	if q := c.Query("type"); q != "" {
		rt, err := review.ParseReviewType(q)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(Response{
				Success: false,
				Error:   err.Error(),
			})
		}

		reviewType = rt
	}

	existingReviews, err := h.service.GetUserReviews(c.Context(), uuid.MustParse(reqId), reviewType)

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

package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/parses"
	"github.com/gofiber/fiber/v3"
)

type RentalHandler struct {
	service rental.Service
}

func NewRentalHandler(service rental.Service) *RentalHandler {
	return &RentalHandler{service: service}
}

func (h *RentalHandler) Create(c fiber.Ctx) error {

	newRental, err := parses.ParseBody[rental.RentalCreateInput](c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Create(c.Context(), newRental)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Rental Created",
	},
	)
}

// CRUD - READ
func (h *RentalHandler) GetByID(c fiber.Ctx) error {
	rentalId, err := parses.ParseUUIDParam(c, "rental_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingRental, err := h.service.GetByID(c.Context(), *rentalId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"rental": existingRental,
	},
	)
}

func (h *RentalHandler) UpdateStatus(c fiber.Ctx) error {
	rentalId, err := parses.ParseUUIDParam(c, "rental_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	newStatus, err := parses.ParseBody[rental.RentalStatus](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.UpdateStatus(c.Context(), *rentalId, &newStatus)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Updated Status",
	},
	)
}

func (h *RentalHandler) Cancel(c fiber.Ctx) error {
	rentalId, err := parses.ParseUUIDParam(c, "rental_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	cancellationReason, err := parses.ParseBody[string](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Cancel(c.Context(), *rentalId, cancellationReason)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Canceled Rental",
	},
	)
}

func (h *RentalHandler) GetAllUserRentals(c fiber.Ctx) error {
	userId, err := parses.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	statusQuery := parses.ParseStringQuery(c, "status")
	rentalStatus, err := rental.ParseRentalStatus(*statusQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	existingRentals, err := h.service.GetAllUserRentals(c.Context(), *userId, rentalStatus)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"rentals": existingRentals, // listagem de rentals do usuário como locador
	},
	)
}

func (h *RentalHandler) GetUserRentalsAsLessor(c fiber.Ctx) error {
	userId, err := parses.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	statusQuery := parses.ParseStringQuery(c, "status")
	rentalStatus, err := rental.ParseRentalStatus(*statusQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	existingRentals, err := h.service.ListByLessor(c.Context(), *userId, rentalStatus)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"rentals": existingRentals, // listagem de rentals do usuário como locador
	},
	)
}

func (h *RentalHandler) GetUserRentalsAsLessee(c fiber.Ctx) error {
	userId, err := parses.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	statusQuery := parses.ParseStringQuery(c, "status")
	rentalStatus, err := rental.ParseRentalStatus(*statusQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	existingRentals, err := h.service.ListByLessee(c.Context(), *userId, rentalStatus)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"rentals": existingRentals, // listagem de rentals do usuário como locador
	},
	)
}

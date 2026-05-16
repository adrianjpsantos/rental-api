package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handler_helpers"
	"github.com/gofiber/fiber/v3"
)

type RentalHandler struct {
	service *application.RentalService
}

func NewRentalHandler(service *application.RentalService) *RentalHandler {
	return &RentalHandler{service: service}
}

func (h *RentalHandler) Create(c fiber.Ctx) error {

	newRental, err := handler_helpers.ParseCreateRental(c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	createdRental, err := h.service.Register(c.Context(), newRental.ItemID, newRental.LesseeID, newRental.LessorID, newRental.StartDate, newRental.EndDate, newRental.TotalAmount, newRental.DeliveryMethod, newRental.Notes)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"rental_id": createdRental.Id,
	},
	)
}

// CRUD - READ
func (h *RentalHandler) GetByID(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "rental_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingRental, err := h.service.GetByID(c.Context(), reqId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"rental": existingRental,
	},
	)
}

func (h *RentalHandler) UpdateStatus(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "rental_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	var req rental.ReqUpdateStatus
	if err := c.Bind().Body(&req); err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	updatedRental, err := h.service.UpdateStatus(c.Context(), reqId, &req.NewStatus)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"updated_rental_status": updatedRental.Status,
	},
	)
}

func (h *RentalHandler) Cancel(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "rental_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	var req rental.ReqCancel
	if err := c.Bind().Body(&req); err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	canceledRental, err := h.service.Cancel(c.Context(), reqId, req.CancellationReason)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"canceled_id": canceledRental.Id, // listagem de rentals do usuário como locador
	},
	)
}

func (h *RentalHandler) GetAllUserRentals(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	rentalStatus, err := handler_helpers.ParseRentalStatusQuery(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	existingRentals, err := h.service.GetAllUserRentals(c.Context(), reqId, rentalStatus)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"rentals": existingRentals, // listagem de rentals do usuário como locador
	},
	)
}

func (h *RentalHandler) GetUserRentalsAsLessor(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	rentalStatus, err := handler_helpers.ParseRentalStatusQuery(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	existingRentals, err := h.service.ListByLessor(c.Context(), reqId, rentalStatus)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"rentals": existingRentals, // listagem de rentals do usuário como locador
	},
	)
}

func (h *RentalHandler) GetUserRentalsAsLessee(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	rentalStatus, err := handler_helpers.ParseRentalStatusQuery(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	existingRentals, err := h.service.ListByLessee(c.Context(), reqId, rentalStatus)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"rentals": existingRentals, // listagem de rentals do usuário como locador
	},
	)
}

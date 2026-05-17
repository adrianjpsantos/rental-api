package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/request"
	"github.com/gofiber/fiber/v3"
)

type AvailabilityHandler struct {
	service *application.AvailabilityService
}

func (h *AvailabilityHandler) CheckAvailability(c fiber.Ctx) error {

	toCheck, err := request.ParseDatesAndItemID(c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	exists, err := h.service.ExistBlockingSlot(c.Context(), toCheck.ItemID, toCheck.StartDate, toCheck.EndDate)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"exists": exists,
	})

}

func (h *AvailabilityHandler) Create(c fiber.Ctx) error {

	newSlot, err := request.ParseAvailabilitySlotCreateInput(c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	createdSlot, err := h.service.Register(c.Context(), newSlot)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"availability_slot_id": createdSlot.Id,
	})
}

func (h *AvailabilityHandler) Delete(c fiber.Ctx) error {
	slotId, err := request.ParseUUIDParam(c, "slot_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	err = h.service.Delete(c.Context(), slotId)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"deleted_id": slotId,
	})
}

func (h *AvailabilityHandler) GetByID(c fiber.Ctx) error {
	slotId, err := request.ParseUUIDParam(c, "slot_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingSlot, err := h.service.GetByID(c.Context(), slotId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"availability_slot": existingSlot,
	})

}

func (h *AvailabilityHandler) GetByItemID(c fiber.Ctx) error {
	itemId, err := request.ParseUUIDParam(c, "item_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingSlots, err := h.service.ListByItemID(c.Context(), itemId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"availability_slots": existingSlots,
	})
}

func NewAvailabilityHandler(service *application.AvailabilityService) availability.InterfaceAvailabilityHandler {
	return &AvailabilityHandler{
		service: service,
	}
}

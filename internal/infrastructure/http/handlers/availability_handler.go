package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/parses"
	"github.com/gofiber/fiber/v3"
)

type AvailabilityHandler struct {
	service availability.Service
}

func (h *AvailabilityHandler) CheckAvailability(c fiber.Ctx) error {

	toCheck, err := parses.ParseBody[availability.AvailabilityFilter](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	exists, err := h.service.ExistsBlockingSlot(c.Context(), toCheck.ItemID, toCheck.StartDate, toCheck.EndDate)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"exists": exists,
	})

}

func (h *AvailabilityHandler) Create(c fiber.Ctx) error {

	newSlot, err := parses.ParseBody[availability.AvailabilitySlotCreateInput](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Create(c.Context(), newSlot)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Slot Created",
	})
}

func (h *AvailabilityHandler) Delete(c fiber.Ctx) error {
	slotId, err := parses.ParseUUIDParam(c, "slot_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	err = h.service.Delete(c.Context(), *slotId)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"deleted_id": slotId,
	})
}

func (h *AvailabilityHandler) GetByID(c fiber.Ctx) error {
	slotId, err := parses.ParseUUIDParam(c, "slot_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingSlot, err := h.service.GetByID(c.Context(), *slotId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"availability_slot": existingSlot,
	})

}

func (h *AvailabilityHandler) GetByItemID(c fiber.Ctx) error {
	itemId, err := parses.ParseUUIDParam(c, "item_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingSlots, err := h.service.ListByItemID(c.Context(), *itemId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"availability_slots": existingSlots,
	})
}

func NewAvailabilityHandler(service availability.Service) *AvailabilityHandler {
	return &AvailabilityHandler{
		service: service,
	}
}

package handler_helpers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/gofiber/fiber/v3"
)

func ParseDatesAndItemID(c fiber.Ctx) (availability.DatesAndItemID, error) {
	var req availability.RequestCheckAvailability
	if err := c.Bind().Body(&req); err != nil {
		return availability.DatesAndItemID{}, err
	}

	return req.Tocheck, nil
}

func ParseCreateAvailabilitySlot(c fiber.Ctx) (availability.AvailabilitySlotCreateInput, error) {
	var req availability.RequestCreate
	if err := c.Bind().Body(&req); err != nil {
		return availability.AvailabilitySlotCreateInput{}, err
	}

	return req.NewSlot, nil
}

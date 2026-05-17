package request

import (
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type DatesAndItemID struct {
	ItemID    uuid.UUID `json:"item_id"`
	StartDate time.Time
	EndDate   time.Time
}

type RequestCheckAvailabilitySlotInput struct {
	Tocheck DatesAndItemID `json:"to_check"`
}

type RequestCreateAvalabilitySlotInput struct {
	NewSlot availability.AvailabilitySlotCreateInput `json:"new_slot"`
}

type RequestGetByItemIDInput struct {
	ItemID uuid.UUID `json:"item_id"`
}

func ParseDatesAndItemID(c fiber.Ctx) (DatesAndItemID, error) {
	var req RequestCheckAvailabilitySlotInput
	if err := c.Bind().Body(&req); err != nil {
		return DatesAndItemID{}, err
	}

	return req.Tocheck, nil
}

func ParseAvailabilitySlotCreateInput(c fiber.Ctx) (availability.AvailabilitySlotCreateInput, error) {
	var req RequestCreateAvalabilitySlotInput
	if err := c.Bind().Body(&req); err != nil {
		return availability.AvailabilitySlotCreateInput{}, err
	}

	return req.NewSlot, nil
}

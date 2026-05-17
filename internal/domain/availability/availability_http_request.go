package availability

import "github.com/google/uuid"

type DatesAndItemID struct {
	ItemID uuid.UUID `json:"item_id"`
	Dates  SlotDates `json:"dates"`
}

type RequestCheckAvailability struct {
	Tocheck DatesAndItemID `json:"to_check"`
}

type RequestCreate struct {
	NewSlot AvailabilitySlotCreateInput `json:"new_slot"`
}

type RequestByItemID struct {
	ItemID uuid.UUID `json:"item_id"`
}

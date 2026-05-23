package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/parses"
	"github.com/gofiber/fiber/v3"
)

type ItemHandler struct {
	service item.Service
}

func (h *ItemHandler) CreateItem(c fiber.Ctx) error {

	newItem, err := parses.ParseBody[item.ItemCreateInput](c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Create(c.Context(), newItem)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Item Created",
	},
	)
}

func (h *ItemHandler) DeleteItem(c fiber.Ctx) error {
	panic("unimplemented")
}

func (h *ItemHandler) GetItemByID(c fiber.Ctx) error {
	itemId, err := parses.ParseUUIDParam(c, "item_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingItem, err := h.service.GetByID(c.Context(), *itemId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"item": existingItem,
	},
	)
}

func (h *ItemHandler) ListItems(c fiber.Ctx) error {
	filters, err := parses.ParseItemFilter(c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	existingItems, err := h.service.ListByFilters(c.Context(), *filters)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"items": existingItems,
	})
}

func (h *ItemHandler) UpdateItem(c fiber.Ctx) error {
	itemId, err := parses.ParseUUIDParam(c, "item_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	toUpdate, err := parses.ParseBody[item.ItemUpdateInput](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Update(c.Context(), *itemId, toUpdate)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Item updated",
	},
	)
}

func NewItemHandler(service item.Service) *ItemHandler {
	return &ItemHandler{service: service}
}

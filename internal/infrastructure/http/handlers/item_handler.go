package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/request"
	"github.com/gofiber/fiber/v3"
)

type ItemHandler struct {
	service *application.ItemService
}

func (h *ItemHandler) CreateItem(c fiber.Ctx) error {

	newItem, err := request.ParseCreateItemInput(c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	createItem, err := h.service.Register(c.Context(), newItem)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"item_id": createItem.Id,
	},
	)
}

func (h *ItemHandler) DeleteItem(c fiber.Ctx) error {
	panic("unimplemented")
}

func (h *ItemHandler) GetItemByID(c fiber.Ctx) error {
	reqId, err := request.ParseUUIDParam(c, "item_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingItem, err := h.service.GetByID(c.Context(), reqId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"item": existingItem,
	},
	)
}

func (h *ItemHandler) ListItems(c fiber.Ctx) error {
	filters, err := request.ParseItemFilterInput(c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	existingItems, err := h.service.ListByFilters(c.Context(), filters)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"items": existingItems,
	})
}

func (h *ItemHandler) UpdateItem(c fiber.Ctx) error {
	reqId, err := request.ParseUUIDParam(c, "item_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	toUpdate, err := request.ParseUpdateItemInput(c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	updatedItem, err := h.service.Update(c.Context(), reqId, toUpdate)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"updated_item_id": updatedItem.Id,
	},
	)
}

func NewItemHandler(service *application.ItemService) item.InterfaceItemHandler {
	return &ItemHandler{service: service}
}

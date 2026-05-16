package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handler_helpers"
	"github.com/gofiber/fiber/v3"
)

type ItemHandler struct {
	service *application.ItemService
}

func (h *ItemHandler) CreateItem(c fiber.Ctx) error {

	var req item.ReqCreate

	if err := c.Bind().Body(&req); err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	newItem := req.NewItem

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
	reqId, err := handler_helpers.ParseUUIDParam(c, "item_id")
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
	filters, err := handler_helpers.ParseItemFilters(c)
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
	reqId, err := handler_helpers.ParseUUIDParam(c, "item_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	var req item.ReqUpdate
	if err := c.Bind().Body(&req); err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	updatedItem, err := h.service.Update(c.Context(), reqId, &req.ToUpdate)

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

package request

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/gofiber/fiber/v3"
)

type RequestCreateItemInput struct {
	NewItem item.ItemCreateInput `json:"new_item"`
}

type RequestUpdateItemInput struct {
	ToUpdate item.ItemUpdate `json:"to_update"`
}

type RequestListByFilterInput struct {
	Filter item.ItemFilter `json:"filter"`
}

func ParseCreateItemInput(c fiber.Ctx) (item.ItemCreateInput, error) {
	var req RequestCreateItemInput
	if err := c.Bind().Body(&req); err != nil {
		return item.ItemCreateInput{}, err
	}

	return req.NewItem, nil
}

func ParseUpdateItemInput(c fiber.Ctx) (item.ItemUpdate, error) {
	var req RequestUpdateItemInput
	if err := c.Bind().Body(&req); err != nil {
		return item.ItemUpdate{}, err
	}

	return req.ToUpdate, nil
}

func ParseItemFilterInput(c fiber.Ctx) (item.ItemFilter, error) {
	var req RequestListByFilterInput
	if err := c.Bind().Body(&req); err != nil {
		return item.ItemFilter{}, err
	}

	return req.Filter, nil
}

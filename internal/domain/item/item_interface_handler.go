package item

import "github.com/gofiber/fiber/v3"

type InterfaceItemHandler interface {
	CreateItem(ctx fiber.Ctx) error
	GetItemByID(ctx fiber.Ctx) error
	ListItems(ctx fiber.Ctx) error // With pagination and filters
	UpdateItem(ctx fiber.Ctx) error
	DeleteItem(ctx fiber.Ctx) error
}

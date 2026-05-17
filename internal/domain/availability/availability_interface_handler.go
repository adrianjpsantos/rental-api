package availability

import "github.com/gofiber/fiber/v3"

type InterfaceAvailabilityHandler interface {
	Create(c fiber.Ctx) error
	GetByID(c fiber.Ctx) error
	GetByItemID(c fiber.Ctx) error
	CheckAvailability(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
}

package category

import "github.com/gofiber/fiber/v3"

type InterfaceCategoryHandler interface {
	CreateCategory(c fiber.Ctx) error

	GetCategoryByID(c fiber.Ctx) error

	ListCategories(c fiber.Ctx) error

	UpdateCategory(c fiber.Ctx) error

	DeleteCategory(c fiber.Ctx) error

	CategoryExists(c fiber.Ctx) error
}

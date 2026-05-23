package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/parses"

	"github.com/gofiber/fiber/v3"
)

type CategoryHandler struct {
	Service category.Service
}

func (h *CategoryHandler) CategoryExists(c fiber.Ctx) error {
	categoryName, err := parses.ParseBody[string](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	exists, err := h.Service.Exists(c.Context(), categoryName)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"exists": exists,
	})
}

func (h *CategoryHandler) CreateCategory(c fiber.Ctx) error {
	newCat, err := parses.ParseBody[category.CategoryCreateInput](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}
	err = h.Service.Create(c.Context(), newCat)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Category Created",
	})
}

func (h *CategoryHandler) DeleteCategory(c fiber.Ctx) error {
	categoryId, err := parses.ParseUUIDParam(c, "category_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	err = h.Service.Delete(c.Context(), *categoryId)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"category_id": categoryId,
	})
}

func (h *CategoryHandler) GetCategoryByID(c fiber.Ctx) error {
	categoryId, err := parses.ParseUUIDParam(c, "category_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingCategory, err := h.Service.GetByID(c.Context(), *categoryId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"category": existingCategory,
	})
}

// ListCategories implements [category.InterfaceCategoryHandler].
func (h *CategoryHandler) ListCategories(c fiber.Ctx) error {
	categories, err := h.Service.List(c.Context())
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"categories": categories,
	})
}

func (h *CategoryHandler) UpdateCategory(c fiber.Ctx) error {
	categoryId, err := parses.ParseUUIDParam(c, "category_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	toUpdateCat, err := parses.ParseBody[category.CategoryUpdate](c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.Service.Update(c.Context(), *categoryId, toUpdateCat)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Updated Category",
	})
}

func NewCategoryHandler(service category.Service) *CategoryHandler {
	return &CategoryHandler{
		Service: service,
	}
}

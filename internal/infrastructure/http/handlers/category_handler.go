package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handler_helpers"
	"github.com/gofiber/fiber/v3"
)

type CategoryHandler struct {
	Service *application.CategoryService
}

func (h *CategoryHandler) CategoryExists(c fiber.Ctx) error {
	categoryName, err := handler_helpers.ParseExistsCategory(c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	exists, err := h.Service.CategoryExists(c.Context(), categoryName)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"exists": exists,
	})
}

func (h *CategoryHandler) CreateCategory(c fiber.Ctx) error {
	newCat, err := handler_helpers.ParseCreateCategory(c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}
	createdCategory, err := h.Service.Register(c.Context(), newCat)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"category_id": createdCategory.ID,
	})
}

func (h *CategoryHandler) DeleteCategory(c fiber.Ctx) error {
	categoryId, err := handler_helpers.ParseUUIDParam(c, "category_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	err = h.Service.Delete(c.Context(), categoryId)
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"category_id": categoryId,
	})
}

func (h *CategoryHandler) GetCategoryByID(c fiber.Ctx) error {
	categoryId, err := handler_helpers.ParseUUIDParam(c, "category_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingCategory, err := h.Service.GetByID(c.Context(), categoryId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"category": existingCategory,
	})
}

// ListCategories implements [category.InterfaceCategoryHandler].
func (h *CategoryHandler) ListCategories(c fiber.Ctx) error {
	categories, err := h.Service.ListCategories(c.Context())
	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"categories": categories,
	})
}

func (h *CategoryHandler) UpdateCategory(c fiber.Ctx) error {
	categoryId, err := handler_helpers.ParseUUIDParam(c, "category_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	toUpdateCat, err := handler_helpers.ParseUpdateCategory(c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	updatedCategory, err := h.Service.Update(c.Context(), categoryId, toUpdateCat)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"updated_category": updatedCategory,
	})
}

func NewCategoryHandler(service *application.CategoryService) category.InterfaceCategoryHandler {
	return &CategoryHandler{
		Service: service,
	}
}

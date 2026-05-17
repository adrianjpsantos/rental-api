package request

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/gofiber/fiber/v3"
)

type RequestCreateCategoryInput struct {
	NewCategory category.CategoryCreateInput `json:"new_category"`
}

type RequestExistsCategoryInput struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type RequestUpdateCategoryInput struct {
	ToUpdate category.CategoryUpdate `json:"to_update"`
}

func ParseCategoryCreateInput(c fiber.Ctx) (category.CategoryCreateInput, error) {
	var req RequestCreateCategoryInput
	if err := c.Bind().Body(&req); err != nil {
		return category.CategoryCreateInput{}, err
	}

	return req.NewCategory, nil
}

func ParseExistsCategory(c fiber.Ctx) (string, error) {
	var req RequestExistsCategoryInput
	if err := c.Bind().Body(&req); err != nil {
		return "", err
	}

	return req.CategoryName, nil
}

func ParseCategoryUpdate(c fiber.Ctx) (category.CategoryUpdate, error) {
	var req RequestUpdateCategoryInput
	if err := c.Bind().Body(&req); err != nil {
		return category.CategoryUpdate{}, err
	}

	return req.ToUpdate, nil
}

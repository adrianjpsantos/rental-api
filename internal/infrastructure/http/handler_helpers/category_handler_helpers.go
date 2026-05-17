package handler_helpers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/gofiber/fiber/v3"
)

func ParseCreateCategory(c fiber.Ctx) (category.CategoryCreateInput, error) {
	var req category.RequestCreateCategory
	if err := c.Bind().Body(&req); err != nil {
		return category.CategoryCreateInput{}, err
	}

	return req.NewCategory, nil
}

func ParseExistsCategory(c fiber.Ctx) (string, error) {
	var req category.RequestExistsCategory
	if err := c.Bind().Body(&req); err != nil {
		return "", err
	}

	return req.CategoryName, nil
}

func ParseUpdateCategory(c fiber.Ctx) (category.CategoryUpdate, error) {
	var req category.RequestUpdateCategory
	if err := c.Bind().Body(&req); err != nil {
		return category.CategoryUpdate{}, err
	}

	return req.ToUpdate, nil
}

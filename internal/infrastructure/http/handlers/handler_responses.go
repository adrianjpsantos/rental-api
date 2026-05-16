package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func ResponseSuccess(c fiber.Ctx, data any) error {
	return c.JSON(Response{
		Success: true,
		Data:    data,
	})
}

func ResponseError(c fiber.Ctx, status int, err string) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Error:   err,
	})
}

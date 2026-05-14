package handlers

import "github.com/gofiber/fiber/v3"

func GetHealth(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "API funcionando",
	})
}

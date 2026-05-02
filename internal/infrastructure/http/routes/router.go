package routes

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API funcionando",
		})
	})

	return app
}

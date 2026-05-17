package request

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func ParseUUIDParam(c fiber.Ctx, key string) (uuid.UUID, error) {
	id := c.Params(key)
	return uuid.Parse(id)
}

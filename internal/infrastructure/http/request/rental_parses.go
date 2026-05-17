package request

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/gofiber/fiber/v3"
)

func ParseCreateRental(c fiber.Ctx) (rental.Rental, error) {
	var req rental.ReqCreate
	if err := c.Bind().Body(&req); err != nil {
		return rental.Rental{}, err
	}

	return req.NewRental, nil
}

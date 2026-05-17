package request

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/gofiber/fiber/v3"
)

type RequestAuthenticate struct {
	AuthenticateInput authenticate.AuthenticateInput `json:"authenticate_input"`
}

func ParseAuthenticateInput(c fiber.Ctx) (authenticate.AuthenticateInput, error) {
	var req RequestAuthenticate
	if err := c.Bind().Body(&req); err != nil {
		return authenticate.AuthenticateInput{}, err
	}

	return req.AuthenticateInput, nil
}

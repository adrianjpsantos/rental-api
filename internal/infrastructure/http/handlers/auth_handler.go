package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/parses"
	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	authService authenticate.Service
}

func NewAuthHandler(
	authService authenticate.Service,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) SignUp(c fiber.Ctx) error {
	input, err := parses.ParseBody[user.UserCreateInput](c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	output, err := h.authService.SignUp(c.Context(), input)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseAuthSuccess(c, output)
}

func (h *AuthHandler) Authenticate(c fiber.Ctx) error {
	input, err := parses.ParseBody[authenticate.AuthenticateInput](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	output, err := h.authService.Authenticate(c.Context(), input)
	if err != nil {

		if err == authenticate.ErrInvalidCredentials {
			return ResponseError(c, fiber.StatusUnauthorized, err.Error())
		}

		return ResponseError(c, fiber.StatusInternalServerError, "Erro no servidor")
	}

	return ResponseAuthSuccess(c, output)
}

func (h *AuthHandler) Logout(c fiber.Ctx) error {
	c.ClearCookie("refresh_token")

	return ResponseSuccess(c, fiber.Map{
		"message": "Logout bem-sucedido",
	})
}

func (h *AuthHandler) Refresh(c fiber.Ctx) error {

	return nil
}

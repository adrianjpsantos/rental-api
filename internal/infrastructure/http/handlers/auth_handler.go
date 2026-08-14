package handlers

import (
	"errors"
	"fmt"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/parses"
	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	AuthService    authenticate.Service
	SessionService session.Service
}

func NewAuthHandler(
	authService authenticate.Service, sessionService session.Service,
) *AuthHandler {
	return &AuthHandler{
		AuthService:    authService,
		SessionService: sessionService,
	}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	input, err := parses.ParseBody[authenticate.RegisterInput](c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.AuthService.Register(c.Context(), input)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Registro bem-sucedido",
	})
}

func (h *AuthHandler) LoginLocal(c fiber.Ctx) error {
	input, err := parses.ParseBody[authenticate.AuthenticateInput](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	output, err := h.AuthService.LoginLocal(c.Context(), input)
	if err != nil {

		if err == authenticate.ErrInvalidCredentials {
			return ResponseError(c, fiber.StatusUnauthorized, err.Error())
		}

		return ResponseError(c, fiber.StatusInternalServerError, "Erro no servidor")
	}

	return ResponseAuthSuccess(c, output)
}

func (h *AuthHandler) Logout(c fiber.Ctx) error {
	secureCookie := config.LoadConfig().IsProduction()
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		Secure:   secureCookie,
		SameSite: "Strict",
		MaxAge:   -1,
	})

	return ResponseSuccess(c, fiber.Map{
		"message": "Logout bem-sucedido",
	})
}

func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")

	if refreshToken == "" {
		return ResponseError(c, fiber.StatusUnauthorized, "Refresh token não fornecido")
	}

	fmt.Println("Refresh token: ", refreshToken)

	accessToken, err := h.AuthService.RefreshAccessToken(c.Context(), refreshToken)

	if err != nil {
		if errors.Is(err, session.ErrSessionNotFound) || errors.Is(err, session.ErrSessionExpired) || errors.Is(err, session.ErrInvalidRefreshToken) {
			secureCookie := config.LoadConfig().IsProduction()
			c.Cookie(&fiber.Cookie{
				Name:     "refresh_token",
				Value:    "",
				Path:     "/",
				HTTPOnly: true,
				Secure:   secureCookie,
				SameSite: "Strict",
				MaxAge:   -1,
			})

			return ResponseError(c, fiber.StatusUnauthorized, err.Error())
		}

		return ResponseError(c, fiber.StatusInternalServerError, "Erro no servidor")
	}
	return ResponseSuccess(c, fiber.Map{
		"access_token": accessToken,
	})
}

func (h *AuthHandler) HandleGoogleLogin(c fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

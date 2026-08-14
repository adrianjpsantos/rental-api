package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/parses"
	"github.com/adrianjpsantos/rental-api/internal/pkg/middleware/authentication"
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{service: service}
}

// CRUD Básico
func (h *UserHandler) Create(c fiber.Ctx) error {
	newUser, err := parses.ParseBody[user.CreateInput](c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	id, err := h.service.Create(c.Context(), newUser.Role)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Created User",
		"user_id": id,
	})
}

func (h *UserHandler) GetByID(c fiber.Ctx) error {
	userId, err := parses.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingUser, err := h.service.GetByID(c.Context(), *userId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"user": existingUser,
	},
	)
}
func (h *UserHandler) Update(c fiber.Ctx) error {

	userId, err := parses.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	input, err := parses.ParseBody[user.User](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Update(c.Context(), *userId, input)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "updated user",
	},
	)

}
func (h *UserHandler) Delete(c fiber.Ctx) error {

	userId, err := parses.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	err = h.service.Delete(c.Context(), *userId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"deleted_id": userId,
	},
	)
}

// Current User
func (h *UserHandler) GetCurrentUser(c fiber.Ctx) error {
	claims, err := authentication.GetAuthenticatedUser(c)
	if err != nil {
		return ResponseError(c, fiber.StatusUnauthorized, "Favor fazer login novamente")
	}

	existingUser, err := h.service.GetByID(c.Context(), claims.UserId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"user": existingUser,
	},
	)
}
func (h *UserHandler) UpdateCurrentUser(c fiber.Ctx) error {

	claims, err := authentication.GetAuthenticatedUser(c)
	if err != nil {
		return ResponseError(c, fiber.StatusUnauthorized, "Favor fazer login novamente")
	}

	input, err := parses.ParseBody[user.User](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Update(c.Context(), claims.UserId, input)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "updated user",
	},
	)

}

// SOFT DELETE
func (h *UserHandler) DeleteAccount(c fiber.Ctx) error {

	claims, err := authentication.GetAuthenticatedUser(c)
	if err != nil {
		return ResponseError(c, fiber.StatusUnauthorized, "Favor fazer login novamente")
	}

	err = h.service.Delete(c.Context(), claims.UserId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"deleted_id": claims,
	},
	)
}

package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/parses"
	"github.com/adrianjpsantos/rental-api/internal/pkg/middleware/authentication"
	"github.com/adrianjpsantos/rental-api/internal/security/validator"
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
	newUser, err := parses.ParseBody[user.UserCreateInput](c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Create(c.Context(), newUser)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"message": "Created User",
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

	input, err := parses.ParseBody[user.UserUpdateInput](c)
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

// Buscas
func (h *UserHandler) GetByEmail(c fiber.Ctx) error {
	email := parses.ParseStringQuery(c, "email")
	err := validator.Get().Var(email, "required,email")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, user.ErrInvalidEmail.Error())
	}

	existingUser, err := h.service.GetByEmail(c.Context(), *email)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"user": existingUser,
	},
	)
}

// Exists
func (h *UserHandler) ExistsByEmail(c fiber.Ctx) error {
	email := parses.ParseStringQuery(c, "email")

	err := validator.Get().Var(email, "required,email")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, user.ErrInvalidEmail.Error())
	}

	exists, err := h.service.ExistsByEmail(c.Context(), *email)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"exists": exists,
	},
	)
}

func (h *UserHandler) ExistsByCPF(c fiber.Ctx) error {
	cpf := parses.ParseStringQuery(c, "cpf")

	err := validator.Get().Var(cpf, "required,cpf")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, user.ErrInvalidCPF.Error())
	}

	exists, err := h.service.ExistsByCPF(c.Context(), *cpf)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"exists": exists,
	},
	)
}

// Current User
func (h *UserHandler) GetCurrentUser(c fiber.Ctx) error {
	claims, err := authentication.GetAuthenticatedUser(c)
	if err != nil {
		return ResponseError(c, fiber.StatusUnauthorized, "Favor fazer login novamente")
	}

	existingUser, err := h.service.GetByID(c.Context(), claims.UserID)

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

	input, err := parses.ParseBody[user.UserUpdateInput](c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	err = h.service.Update(c.Context(), claims.UserID, input)

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

	err = h.service.Delete(c.Context(), claims.Payload().UserID)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"deleted_id": claims,
	},
	)
}

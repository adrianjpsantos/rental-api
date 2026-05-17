package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/http/handler_helpers"
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type UserHandler struct {
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CRUD Básico
func (h *UserHandler) Create(c fiber.Ctx) error {
	newUser, err := handler_helpers.ParseCreateUser(c)

	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	createdUser, err := h.service.Register(c.Context(), newUser)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"created_user_id": createdUser.Id,
	})
}

func (h *UserHandler) GetByID(c fiber.Ctx) error {
	userId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	existingUser, err := h.service.GetByID(c.Context(), userId)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"user": existingUser,
	},
	)
}
func (h *UserHandler) Update(c fiber.Ctx) error {

	userId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	toUpdate, err := handler_helpers.ParseUpdateUser(c)
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	updatedUser, err := h.service.UpdateProfile(c.Context(), userId, toUpdate)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"updated_user": updatedUser,
	},
	)

}
func (h *UserHandler) Delete(c fiber.Ctx) error {

	userId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "ID Inválido")
	}

	err = h.service.Delete(c.Context(), userId)

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
	var req user.ReqByEmail

	if err := c.Bind().Body(&req); err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	existingUser, err := h.service.GetByEmail(c.Context(), req.Email)

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
	var req user.ReqByEmail

	if err := c.Bind().Body(&req); err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	exists, err := h.service.ExistsByEmail(c.Context(), req.Email)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"exists": exists,
	},
	)
}

func (h *UserHandler) ExistsByCPF(c fiber.Ctx) error {
	var req user.ReqByCPF

	if err := c.Bind().Body(&req); err != nil {
		return ResponseError(c, fiber.StatusBadRequest, "JSON Inválido")
	}

	exists, err := h.service.ExistsByCPF(c.Context(), req.Cpf)

	if err != nil {
		return ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return ResponseSuccess(c, fiber.Map{
		"exists": exists,
	},
	)
}

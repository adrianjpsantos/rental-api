package handlers

import (
	"github.com/adrianjpsantos/rental-api/internal/application"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
func (h *UserHandler) Create(c fiber.Ctx) error { return nil }
func (h *UserHandler) GetByID(c fiber.Ctx) error {
	var req user.ReqById

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "JSON Inválido",
		})
	}

	existingUser, err := h.service.GetByID(c.Context(), uuid.MustParse(req.Id))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"user": existingUser,
		},
	})
}
func (h *UserHandler) Update(c fiber.Ctx) error { return nil }
func (h *UserHandler) Delete(c fiber.Ctx) error { return nil }

// Buscas
func (h *UserHandler) GetByEmail(c fiber.Ctx) error {
	var req user.ReqByEmail

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "JSON Inválido",
		})
	}

	existingUser, err := h.service.GetByEmail(c.Context(), req.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"user": existingUser,
		},
	})
}

// Exists
func (h *UserHandler) ExistsByEmail(c fiber.Ctx) error {
	var req user.ReqByEmail

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "JSON Inválido",
		})
	}

	exists, err := h.service.ExistsByEmail(c.Context(), req.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"exists": exists,
		},
	})
}

func (h *UserHandler) ExistsByCPF(c fiber.Ctx) error {
	var req user.ReqByCPF

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "JSON Inválido",
		})
	}

	exists, err := h.service.ExistsByCPF(c.Context(), req.Cpf)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"exists": exists,
		},
	})
}

// Cache
func (h *UserHandler) UpdateReputationCache(c fiber.Ctx) error       { return nil }
func (h *UserHandler) UpdateTotalRentalCache(c fiber.Ctx) error      { return nil }
func (h *UserHandler) UpdateTotalItemsRentedCache(c fiber.Ctx) error { return nil }

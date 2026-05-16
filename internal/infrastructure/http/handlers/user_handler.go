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
	var req user.ReqCreate

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "JSON Inválido",
		})
	}

	newUser := req.NewUser
	newUser.GenerateHashedPassword()

	createdUser, err := h.service.Register(c.Context(), newUser.Name, newUser.Email, newUser.PasswordHash, newUser.CPF, newUser.Phone, newUser.BirthDate, newUser.Role)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"created_user_id": createdUser.Id,
		},
	})

}
func (h *UserHandler) GetByID(c fiber.Ctx) error {
	reqId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	existingUser, err := h.service.GetByID(c.Context(), reqId)

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
func (h *UserHandler) Update(c fiber.Ctx) error {

	reqId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	var req user.ReqUpdate

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "JSON Inválido",
		})
	}

	updatedUser, err := h.service.UpdateProfile(c.Context(), reqId, req.ToUpdate)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"updated_user": updatedUser,
		},
	})

}
func (h *UserHandler) Delete(c fiber.Ctx) error {

	reqId, err := handler_helpers.ParseUUIDParam(c, "user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Success: false,
			Error:   "ID Inválido",
		})
	}

	err = h.service.Delete(c.Context(), reqId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data: fiber.Map{
			"deleted_id": reqId,
		},
	})
}

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

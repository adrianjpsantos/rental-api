package handler_helpers

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/gofiber/fiber/v3"
)

func ParseCreateUser(c fiber.Ctx) (user.UserCreateInput, error) {
	var req user.ReqCreate
	if err := c.Bind().Body(&req); err != nil {
		return user.UserCreateInput{}, err
	}

	return req.NewUser, nil
}

func ParseUpdateUser(c fiber.Ctx) (user.UserUpdate, error) {
	var req user.ReqUpdate
	if err := c.Bind().Body(&req); err != nil {
		return user.UserUpdate{}, err
	}

	return req.ToUpdate, nil
}

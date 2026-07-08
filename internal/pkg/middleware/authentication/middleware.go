package authentication

import (
	"fmt"
	"strings"

	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/gofiber/fiber/v3"
)

const AuthenticatedUserContextKey = "authenticated_user"

func AuthMiddleware(
	sessionService session.Service,
) fiber.Handler {

	return func(c fiber.Ctx) error {

		authHeader := strings.TrimSpace(
			c.Get("Authorization"),
		)

		if authHeader == "" {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				ErrMissingAuthorizationHeader.Error(),
			)
		}

		const bearerPrefix = "Bearer "

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				ErrInvalidAuthorizationHeader.Error(),
			)
		}

		sessionToken := strings.TrimSpace(
			strings.TrimPrefix(authHeader, bearerPrefix),
		)

		if sessionToken == "" {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				ErrInvalidSessionToken.Error(),
			)
		}

		claims, err := sessionService.ValidateAccessToken(sessionToken)
		if err != nil {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				err.Error(),
			)
		}

		c.Locals(
			AuthenticatedUserContextKey,
			claims,
		)

		return c.Next()
	}
}

func GetAuthenticatedUser(
	c fiber.Ctx,
) (*session.Claims, error) {

	user, ok := c.Locals(AuthenticatedUserContextKey).(*session.Claims)

	if !ok {
		fmt.Println("Erro GETAUTHUSER")
		return nil, ErrAuthenticatedUserNotFound
	}

	return user, nil
}

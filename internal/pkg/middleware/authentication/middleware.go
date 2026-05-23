package authentication

import (
	"strings"

	"github.com/adrianjpsantos/rental-api/internal/domain/token"
	"github.com/gofiber/fiber/v3"
)

const AuthenticatedUserContextKey = "authenticated_user"

func AuthMiddleware(
	tokenService token.Service,
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

		claims, err := tokenService.ValidateAccessToken(sessionToken)
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
) (*token.Claims, error) {

	user, ok := c.Locals(AuthenticatedUserContextKey).(*token.Claims)

	if !ok {
		return nil, ErrAuthenticatedUserNotFound
	}

	return user, nil
}

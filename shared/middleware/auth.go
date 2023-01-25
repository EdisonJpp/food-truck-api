package middleware

import (
	"food-truck-api/shared/presenter"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SIGN")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(presenter.ErrorResponsePresenter("Missing or malformed JWT"))
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(presenter.ErrorResponsePresenter("Invalid or expired JWT"))
}

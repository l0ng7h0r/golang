package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/l0ng7h0r/golang/internal/usecase"
	"github.com/l0ng7h0r/golang/pkg/security"
)

type AuthMiddleware struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthMiddleware(authUsecase *usecase.AuthUsecase) *AuthMiddleware {
	return &AuthMiddleware{authUsecase: authUsecase}
}

func (m *AuthMiddleware) Auth(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	token = strings.TrimPrefix(token, "Bearer ")
	claims, err := security.ParseToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	c.Locals("user_id", (*claims)["user_id"])
	c.Locals("roles", (*claims)["roles"])

	return c.Next()
}
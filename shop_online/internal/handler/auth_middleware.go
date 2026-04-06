package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/l0ng7h0r/golang/internal/usecase"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
		Roles []string `json:"roles"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.authUsecase.Register(req.Email, req.Password, req.Roles)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	accessToken, refreshToken, err := h.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"access_token": accessToken, "refresh_token": refreshToken})
}

func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	accessToken, refreshToken, err := h.authUsecase.Refresh(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"access_token": accessToken, "refresh_token": refreshToken})
}
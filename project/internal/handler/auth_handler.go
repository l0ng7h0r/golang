package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/l0ng7h0r/internal/usecase"
)


type AuthHandler struct {
	AuthUsecase *usecase.AuthUsecase
}

func NewAthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{AuthUsecase: authUsecase}
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

	err := h.AuthUsecase.Register(req.Email, req.Password, req.Roles)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}


func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := h.AuthUsecase.Login(req.Email, req.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
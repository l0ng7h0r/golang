package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/l0ng7h0r/internal/domain"
	"github.com/l0ng7h0r/internal/usecase"
)

type SellerHandler struct {
	SellerUsecase *usecase.SellerUsecase
}

func NewSellerHandler(sellerUsecase *usecase.SellerUsecase) *SellerHandler {
	return &SellerHandler{SellerUsecase: sellerUsecase}
}

// สมัครเป็น seller (ต้อง login แล้ว)
func (h *SellerHandler) CreateSeller(c fiber.Ctx) error {
	var req struct {
		ShopName    string `json:"shop_name"`
		Description string `json:"description"`
		Phone       string `json:"phone"`
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// ดึง user_id จาก JWT token
	userID, ok := c.Locals("user_id").(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	seller := &domain.Seller{
		UserID:      int64(userID),
		ShopName:    req.ShopName,
		Description: req.Description,
		Phone:       req.Phone,
	}

	err := h.SellerUsecase.CreateSeller(seller)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "seller created successfully"})
}

// ดูข้อมูล seller ตาม id
func (h *SellerHandler) GetSeller(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	seller, err := h.SellerUsecase.GetSellerByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "seller not found"})
	}

	return c.Status(fiber.StatusOK).JSON(seller)
}

// ดูข้อมูล seller ของตัวเอง (จาก JWT token)
func (h *SellerHandler) GetMySellerProfile(c fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	seller, err := h.SellerUsecase.GetSellerByUserID(int64(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "seller profile not found"})
	}

	return c.Status(fiber.StatusOK).JSON(seller)
}

// แก้ไขข้อมูล seller
func (h *SellerHandler) UpdateSeller(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	var req struct {
		ShopName    string `json:"shop_name"`
		Description string `json:"description"`
		Phone       string `json:"phone"`
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	seller := &domain.Seller{
		ID:          id,
		ShopName:    req.ShopName,
		Description: req.Description,
		Phone:       req.Phone,
	}

	err = h.SellerUsecase.UpdateSeller(seller)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "seller updated successfully"})
}

// ลบ seller
func (h *SellerHandler) DeleteSeller(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	err = h.SellerUsecase.DeleteSeller(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "seller deleted successfully"})
}

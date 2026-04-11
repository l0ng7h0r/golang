package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/l0ng7h0r/internal/domain"
	"github.com/l0ng7h0r/internal/usecase"
)

type ProductHandler struct {
	ProductUsecase *usecase.ProductUsecase
	SellerUsecase  *usecase.SellerUsecase
}

func NewProductHandler(productUsecase *usecase.ProductUsecase, sellerUsecase *usecase.SellerUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: productUsecase,
		SellerUsecase:  sellerUsecase,
	}
}

func (h *ProductHandler) CreateProduct(c fiber.Ctx) error {
	var req struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// ดึง user_id จาก JWT แล้วหา seller_id
	userID, ok := c.Locals("user_id").(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	seller, err := h.SellerUsecase.GetSellerByUserID(int64(userID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "seller profile not found"})
	}

	product := &domain.Product{
		SellerID: seller.ID,
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.ProductUsecase.CreateProduct(product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func (h *ProductHandler) GetProduct(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	product, err := h.ProductUsecase.GetProduct(int64(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *ProductHandler) GetProducts(c fiber.Ctx) error {
	products, err := h.ProductUsecase.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(products)
}

func (h *ProductHandler) UpdateProduct(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var req struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	product := &domain.Product{
		ID:    int64(id),
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}
	err = h.ProductUsecase.UpdateProduct(product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})	
}

func (h *ProductHandler) DeleteProduct(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.ProductUsecase.DeleteProduct(int64(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
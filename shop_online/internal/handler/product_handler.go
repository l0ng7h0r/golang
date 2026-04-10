package handler

import (
	"strconv"
	
	"github.com/gofiber/fiber/v3"
	"github.com/l0ng7h0r/golang/internal/domain"
	"github.com/l0ng7h0r/golang/internal/usecase"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
	sellerUsecase *usecase.SellerUsecase
}

func NewProductHandler(productUsecase *usecase.ProductUsecase, sellerUsecase *usecase.SellerUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase, sellerUsecase: sellerUsecase}
}

func (h *ProductHandler) CreateProduct(c fiber.Ctx) error {
	var req struct {
		SellerID string `json:"seller_id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Price float64 `json:"price"`
		Stock int `json:"stock"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.productUsecase.CreateProduct(req.SellerID, req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Product created successfully"})
}

func (h *ProductHandler) GetProductByID(c fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.productUsecase.GetProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(product)
}

func (h *ProductHandler) GetAllProducts(c fiber.Ctx) error {
	products, err := h.productUsecase.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (h *ProductHandler) GetProductsBySeller(c fiber.Ctx) error {
	id := c.Params("id")
	products, err := h.productUsecase.GetProductsBySeller(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (h *ProductHandler) DeleteProduct(c fiber.Ctx) error {
	id := c.Params("id")
	err := h.productUsecase.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}

func (h *ProductHandler) UpdateProduct(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var req struct {
		SellerID string `json:"seller_id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Price float64 `json:"price"`
		Stock int `json:"stock"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.productUsecase.UpdateProduct(string(id), &domain.Product{
		SellerID: req.SellerID,
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		Stock: req.Stock,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Product updated successfully"})
}


package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/l0ng7h0r/internal/domain"
	"github.com/l0ng7h0r/internal/usecase"
)

type OrderHandler struct {
	OrderUsecase *usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		OrderUsecase: orderUsecase,
	}
}

func (h *OrderHandler)CreateOrders(c fiber.Ctx) error {
    var req struct {
        UserID     int64               `json:"user_id"`
        OrderItems []*domain.OrderItem `json:"order_items"`
    }

    if err := c.Bind().Body(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    order := &domain.Order{
        UserID: req.UserID,
    }

    // แก้ CreateOrder → CreateOrders
    err := h.OrderUsecase.CreateOrders(order, req.OrderItems)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": order})
}

func (h *OrderHandler)GetOrdersByID(c fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	order, order_items, err := h.OrderUsecase.GetOrdersByID(int64(id))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"order": order,
		 "order_items": order_items,
	})
}


func (h *OrderHandler)GetOrdersByUserId(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	orders, err := h.OrderUsecase.GetOrdersByUserId(int64(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": orders,
	})
}

func (h *OrderHandler)UpdateOrders(c fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
    }

	var req struct {
		Status string `json:"status"`
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.OrderUsecase.UpdateOrders(&domain.Order{ID: id, Status: req.Status})
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func (h *OrderHandler)DeleteOrders(c fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	err = h.OrderUsecase.DeleteOrders(int64(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})

}
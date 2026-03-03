package adapters

import(
	"long/entities"
	"long/usecases"
	"github.com/gofiber/fiber/v3"
)


type HttpOrderHandler struct {
	orderUsecase usecases.OrderUsecase
}

func NewHttpOrderHandler(orderUsecase usecases.OrderUsecase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUsecase: orderUsecase}
}

func (h *HttpOrderHandler) CreateOrder(c fiber.Ctx) error {
	var order entities.Order
	if err := c.Bind().Body(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := h.orderUsecase.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
package usecase

import (
	"errors"
	"github.com/l0ng7h0r/internal/domain"
	"github.com/l0ng7h0r/internal/repository"
)


type OrderUsecase struct {
	OrderRepo * repository.OrderRepository
}

func NewOrderUsecase(repo *repository.OrderRepository) *OrderUsecase{
	return &OrderUsecase{
		OrderRepo: repo,
	}
}

func (u *OrderUsecase)CreateOrders(order *domain.Order, orderItems []*domain.OrderItem) error {
	if len(orderItems) == 0 {
        return errors.New("order must have at least 1 item")
    }

    // เช็ค quantity แต่ละ item
    for _, item := range orderItems {
        if item.Quantity <= 0 {
            return errors.New("quantity must be greater than 0")
        }
	}
	return u.OrderRepo.CreateOrders(order, orderItems)
}

func (u *OrderUsecase)GetOrdersByID(id int64) (*domain.Order, []*domain.OrderItem, error) {
	return u.OrderRepo.GetOrdersByID(id)
}

func (u *OrderUsecase)GetOrdersByUserId(id int64) ([]*domain.Order, error) {
	return u.OrderRepo.GetOrdersByUserId(id)
}

func (u *OrderUsecase)UpdateOrders(order *domain.Order) error {
	return u.OrderRepo.UpdateOrders(order)
}

func (u *OrderUsecase)DeleteOrders(id int64) error {
	return u.OrderRepo.DeleteOrders(id)
}
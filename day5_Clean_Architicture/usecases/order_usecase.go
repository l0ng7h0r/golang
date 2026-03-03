package usecases

import (
	"errors"
	"long/entities"
)


type OrderUsecase interface {
	CreateOrder(order entities.Order) error
}


type OrderService struct  {
	repo OrderRepository
}


func NewOrderService(repo OrderRepository) OrderUsecase {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order entities.Order) error {

	if order.Total <= 0 {
		return errors.New("Total must be greater than 0")
	}

	return s.repo.Save(order)
}
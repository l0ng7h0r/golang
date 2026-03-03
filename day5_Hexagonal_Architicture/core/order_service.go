package core

import "errors"

type OrderService interface {
	CreateOrder(order Order) error
}

type orderServiceImpl struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo}
}

func (s *orderServiceImpl) CreateOrder(order Order) error {
	//Bussiness logic function
	if order.Total <= 0 {
		return errors.New("Total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}
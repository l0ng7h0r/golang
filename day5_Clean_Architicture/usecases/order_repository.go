package usecases

import "long/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
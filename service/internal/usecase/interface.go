package usecase

import (
	"orderAPI/service/internal/domain/order"
)

type Order interface{
	GetByUID(string) (*order.Order, error)
	Save(*order.Order) error
}
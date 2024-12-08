package adapter

import "orderAPI/service/internal/domain/order"

type Cache interface {
	Get(uid string) (*order.Order, bool)
	Set(order *order.Order)
}
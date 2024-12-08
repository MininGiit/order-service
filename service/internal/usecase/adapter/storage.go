package adapter

import(
	"orderAPI/service/internal/domain/order"
)

type Storage interface{
	GetByUID(uid string) (*order.Order, error)
	Save(*order.Order) error
}
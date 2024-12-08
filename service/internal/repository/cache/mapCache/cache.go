package mapCache

import (
	"orderAPI/service/internal/domain/order"
)

type Cache struct {
	data 	map[string] order.Order
	limit	uint
}

func New(limit uint) *Cache {
	data := make(map[string] order.Order, limit)
	return &Cache{
		data:	data,
		limit:	limit,
	}
}

func (c *Cache) Get(uid string) (*order.Order, bool) {
	order, ok := c.data[uid]
	return &order, ok
}

func (c *Cache) Set(order *order.Order){
	c.data[order.OrderUID] = *order
}
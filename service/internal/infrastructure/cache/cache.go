package cache

import (
	"orderAPI/service/internal/domain/order"
	"container/list"
	"sync"
)

type Item struct {
	Key   string
	Value *order.Order
}

type Cache struct {
	size		int
	maxSize		int
	timeQueue	list.List	//содержит только ключи data
	mutex		sync.Mutex
	data		map[string] *order.Order
}

func New(maxSize int) *Cache {
	data := make(map[string] *order.Order, maxSize)
	timeQueue := list.New()
	return &Cache{
		size: 		0,
		maxSize: 	maxSize,
		timeQueue: 	*timeQueue,
		data: 		data,
	}
}

func (c *Cache) Get(uid string) (*order.Order, bool) {
	c.mutex.Lock()
	order, ok := c.data[uid]
	c.mutex.Unlock()
	return order, ok
}

func (c *Cache) Set(order *order.Order){
	if c.size >= c.maxSize {
		c.mutex.Lock()
		front := c.timeQueue.Front()
		key := front.Value.(string)
		delete(c.data, key)
		c.size--
		c.mutex.Unlock()
	}
	c.mutex.Lock()
	c.data[order.OrderUID] = order
	c.timeQueue.PushBack(order.OrderUID)
	c.size++
	c.mutex.Unlock()
}

func (c *Cache) SetOrders(orders []*order.Order) {
	for _, order := range orders {
		c.Set(order)
	}
}


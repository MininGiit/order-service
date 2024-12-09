package repository

import "orderAPI/service/internal/domain/order"

type Cache interface {
	Get(uid string) (*order.Order, bool)
	Set(order *order.Order)
	SetOrders([]order.Order)
}

type Storage interface{
	GetByUID(uid string) (*order.Order, error)
	Save(*order.Order) error
	GetAll() ([]order.Order, error)
}

type Repo struct {
	storage	Storage
	cache 	Cache
}

func New(storage Storage, cache Cache) *Repo{
	orders, err := storage.GetAll()
	if err != nil {
		return &Repo{storage: storage, cache: cache}
	}
	cache.SetOrders(orders)
	return &Repo{storage: storage, cache: cache}
}

func (r *Repo) GetByUID(uid string) (*order.Order, error) {
	order, ok := r.cache.Get(uid)
	if ok {
		return order, nil
	}
	order, err := r.storage.GetByUID(uid)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repo) Save(order *order.Order) error {
	err := r.storage.Save(order)
	if err != nil {
		return err
	}
	r.cache.Set(order)
	return nil
}
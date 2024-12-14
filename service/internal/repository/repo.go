package repository

import (
	"orderAPI/service/internal/domain/order"
	"log"
)

type Cache interface {
	Get(uid string) (*order.Order, bool)
	Set(order *order.Order)
	SetOrders([]*order.Order)
}

type Storage interface{
	GetByUID(uid string) (*order.Order, error)
	Save(*order.Order) error
	GetAll() ([]*order.Order, error)
}

type Repo struct {
	storage	Storage
	cache 	Cache
}

func New(storage Storage, cache Cache) *Repo{
	orders, err := storage.GetAll()
	if err != nil {
		log.Println("failed to load data from cache:", err)
		return &Repo{storage: storage, cache: cache}
	}
	cache.SetOrders(orders)
	log.Println("data loaded from cache")
	return &Repo{storage: storage, cache: cache}
}

func (r *Repo) GetByUID(uid string) (*order.Order, error) {
	order, ok := r.cache.Get(uid)
	if ok {
		log.Println("order retrieved from cache")
		return order, nil
	}
	order, err := r.storage.GetByUID(uid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("order retrieved from DB")
	return order, nil
}

func (r *Repo) Save(order *order.Order) error {
	err := r.storage.Save(order)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("order saved to cache")
	r.cache.Set(order)
	return nil
}
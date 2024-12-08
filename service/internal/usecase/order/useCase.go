package order

import (
	"orderAPI/service/internal/domain/order"
	"orderAPI/service/internal/usecase/adapter"
)

type OrderUseCase struct {
	storage adapter.Storage
	cache	adapter.Cache
}

func New(storage adapter.Storage, cache adapter.Cache) *OrderUseCase{
	return &OrderUseCase{
		storage: storage,
		cache: cache,
	}
}

func (uc *OrderUseCase) GetByUID(uid string) (*order.Order, error) {
	//реализовать поход в кеш
	order, ok := uc.cache.Get(uid)
	if ok {
		return order, nil
	}
	order, err := uc.storage.GetByUID(uid)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (uc *OrderUseCase) Save(order *order.Order) error {
	uc.cache.Set(order)
	err := uc.storage.Save(order)
	return err
}

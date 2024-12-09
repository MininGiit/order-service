package order

import (
	"orderAPI/service/internal/domain/order"
	"orderAPI/service/internal/repository"
)

type OrderUseCase struct {
	repo repository.Repo	
}

func New(r repository.Repo) *OrderUseCase{
	return &OrderUseCase{repo: r}
}

func (uc *OrderUseCase) GetByUID(uid string) (*order.Order, error) {
	return uc.repo.GetByUID(uid)
}

func (uc *OrderUseCase) Save(order *order.Order) error {
	return uc.repo.Save(order)	
}

package tests

import (
	//"orderAPI/service/internal/domain/order"
	"orderAPI/service/internal/infrastructure/cache"
	"testing"
)

func TestSet(t *testing.T) {
	cacheMaxSize := 10
	cache := cache.New(cacheMaxSize)
	order1 := generateOrder()
	order2 := generateOrder()
	cache.Set(order1)
	cache.Set(order2)
	if cache.GetSize() != 2 {
        t.Errorf("size of cache = %d; want 2", cache.GetSize())
    }
	
}

func TestGet(t *testing.T) {
	cacheMaxSize := 10
	cache := cache.New(cacheMaxSize)
	order1 := generateOrder()
	cache.Set(order1)
	order1FromCache, ok := cache.Get(order1.OrderUID)
	if !ok {
		t.Errorf("data with UID:%s not found", order1.OrderUID)
	}
	if !CompareOrders(order1FromCache, order1){
		t.Errorf("UID didnn't match")
	}
}

func TestLRU(t *testing.T) {
	cacheMaxSize := 2
	cache := cache.New(cacheMaxSize)
	order1 := generateOrder()
	order2 := generateOrder()
	order3 := generateOrder()
	cache.Set(order1)
	cache.Set(order2)
	cache.Set(order3)

	_, ok := cache.Get(order1.OrderUID)
	if ok {
		t.Errorf("data with UID:%s not found; expexted that it was not found", order1.OrderUID)
	}
}
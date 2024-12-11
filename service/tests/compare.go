package tests

import (
	"orderAPI/service/internal/domain/order"
	"reflect"
)

// "fmt"
// "reflect"

// Функция для сравнения двух структур Order
func CompareOrders(a, b *order.Order) bool {
    if a.OrderUID != b.OrderUID ||
        a.TrackNumber != b.TrackNumber ||
        a.Entry != b.Entry ||
        a.Locale != b.Locale ||
        a.InternalSig != b.InternalSig ||
        a.CustomerID != b.CustomerID ||
        a.DeliveryService != b.DeliveryService ||
        a.ShardKey != b.ShardKey ||
        a.SmId != b.SmId ||
        a.OofShard != b.OofShard ||
        !reflect.DeepEqual(a.DateCreated, b.DateCreated) {
        return false
    }

    if !ComparePayments(a.Payment, b.Payment) || !CompareDeliveries(a.Delivery, b.Delivery) {
        return false
    }

    if len(a.Items) != len(b.Items) {
        return false
    }

    if !CompareItems(a.Items, b.Items) {
        return false
    }

    return true
}

func ComparePayments(a, b *order.Payment) bool {
    if a.Transaction != b.Transaction ||
        a.RequestID != b.RequestID ||
        a.Currency != b.Currency ||
        a.Provider != b.Provider ||
        a.Amount != b.Amount ||
        a.PaymentID != b.PaymentID ||
        a.Bank != b.Bank ||
        a.DeliveryCost != b.DeliveryCost ||
        a.GoodsTotal != b.GoodsTotal ||
        a.CustomFee != b.CustomFee {
        return false
    }
    return true
}

func CompareDeliveries(a, b *order.Delivery) bool {
    if a.Name != b.Name ||
        a.Phone != b.Phone ||
        a.Zip != b.Zip ||
        a.City != b.City ||
        a.Address != b.Address ||
        a.Region != b.Region ||
        a.Email != b.Email {
        return false
    }
    return true
}

func CompareItems(a, b []order.Item) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i].ChrtID != b[i].ChrtID ||
        a[i].TrackNumber != b[i].TrackNumber ||
        a[i].Price != b[i].Price ||
        a[i].Rid != b[i].Rid ||
        a[i].Name != b[i].Name ||
        a[i].Sale != b[i].Sale ||
        a[i].Size != b[i].Size ||
        a[i].TotalPrice != b[i].TotalPrice ||
        a[i].NmID != b[i].NmID ||
        a[i].Brand != b[i].Brand ||
        a[i].Status != b[i].Status {
        return false
    }
	}
    
    return true
}
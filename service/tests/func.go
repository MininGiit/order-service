package tests

import (
	"fmt"
	"math/rand"
	"orderAPI/service/internal/domain/order"
	"time"
)

func randomString(n int) string {
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    s := make([]rune, n)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}

func generateOrder() *order.Order {
    rand.Seed(time.Now().UnixNano())

    newOrder := &order.Order{}

	newOrder.OrderUID = fmt.Sprintf("%stest", randomString(8))
	newOrder.TrackNumber = fmt.Sprintf("WBILTEST%s", randomString(12))
	newOrder.Entry = "WBIL"

    delivery := &order.Delivery{
        Name:    randomString(5),
        Phone:   fmt.Sprintf("+%d%d", rand.Int(), rand.Int()),
        Zip:     fmt.Sprintf("%d", rand.Int()),
        City:	 randomString(5),
        Address: randomString(5),
        Region:  randomString(5),
        Email:   randomString(5),
    }
    newOrder.Delivery = delivery

    payment := &order.Payment{
        Transaction:  fmt.Sprintf("%stest", randomString(16)),
        RequestID:    fmt.Sprintf("%stest", randomString(16)),
        Currency:     "USD",
        Provider:     "wbpay",
        Amount:       uint(rand.Intn(10000)),
        PaymentID:    uint64(time.Now().Unix()),
        Bank:         "alpha",
        DeliveryCost: uint(rand.Intn(2000)),
        GoodsTotal:   uint(rand.Intn(500)),
        CustomFee:    uint(rand.Intn(50)),
    }
    newOrder.Payment = payment

    item := order.Item{
        ChrtID:      uint64(rand.Int63n(999999999)),
        TrackNumber: randomString(12),
        Price:       uint(rand.Intn(1000)),
        Rid:         randomString(16),
        Name:        randomString(5),
        Sale:        uint(rand.Intn(90)),
        Size:        fmt.Sprintf("%stest", randomString(10)),
        TotalPrice:  uint(rand.Intn(300)),
        NmID:        uint(rand.Int31n(999999)),
        Brand:       randomString(5),
        Status:      uint(rand.Intn(400)),
    }
    newOrder.Items = append(newOrder.Items, item)

    newOrder.Locale = "en"
    newOrder.InternalSig = ""
    newOrder.CustomerID = "test"
    newOrder.DeliveryService = "meest"
    newOrder.ShardKey = "9"
    newOrder.SmId = uint64(rand.Int63n(999999999))
    newOrder.DateCreated = time.Now().UTC()
    newOrder.OofShard = "1"

    return newOrder
}

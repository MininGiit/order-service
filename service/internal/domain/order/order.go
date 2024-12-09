package order

import (
	"time"
)

type Order struct {
	OrderUID 		string		`json:"order_uid"`
	TrackNumber		string		`json:"track_number"`
	Entry			string		`json:"entry"`
	Delivery		*Delivery	`json:"delivery"`
	Payment			*Payment	`json:"payment"`
	Items			[]Item		`json:"items"`
	Locale			string		`json:"locale"`
	InternalSig		string		`json:"internal_signature"`
	CustomerID		string		`json:"customer_id"`
	DeliveryService	string		`json:"delivery_service"`
	ShardKey		string		`json:"shardkey"`
	SmId			uint64		`json:"sm_id"`
	DateCreated		time.Time 	`json:"date_created"`
	OofShard		string		`json:"oof_shard"`
}

// func NewOrder(
// 	orderUID	string, 
// 	trackNumber	string, 
// 	entry 		string, 
// 	delivery 	*Delivery, 
// 	payment 	*Payment,
// 	) *Order {
// 	return &Order{
// 		OrderUID: 		orderUID,
// 		TrackNumber: 	trackNumber,
// 		Entry: 			entry,
// 		Delivery: 		delivery,
// 		Payment: 		payment,	
// 	}

// }
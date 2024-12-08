package order

type Payment struct {
	Transaction		string	`json:"transaction"`
	RequestID		string	`json:"request_id"`
	Currency		string	`json:"currency"`
	Provider		string	`json:"provider"`
	Amount			uint	`json:"amount"`
	PaymentID		uint64	`json:"payment_dt"`
	Bank			string	`json:"bank"`
	DeliveryCost	uint	`json:"delivery_cost"`
	GoodsTotal		uint	`json:"goods_total"`
	CustomFee		uint	`json:"custom_fee"`
}

// func NewPayment(
// 	transaction		string,
// 	requestID		string,
// 	currency		string,
// 	provider		string,
// 	amount			uint,
// 	paymentID		uint64,
// 	bank			string,
// 	deliveryCost	uint,
// 	goodsTotal		uint,
// 	customFee		uint,
// ) *Payment{
// 	return &Payment{
// 		transaction: 	transaction,
// 		requestID: 		requestID,
// 		currency:		currency,
// 		provider:		provider,
// 		amount:			amount,
// 		paymentID:		paymentID,
// 		bank:			bank,
// 		deliveryCost:	deliveryCost,
// 		goodsTotal:		goodsTotal,
// 		customFee:		customFee,
// 	}
// }

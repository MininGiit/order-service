package order

type Payment struct {
	Transaction		string	`json:"transaction" validate:"required"`
	RequestID		string	`json:"request_id"`
	Currency		string	`json:"currency" validate:"required"`
	Provider		string	`json:"provider" validate:"required"`
	Amount			uint	`json:"amount" validate:"required"`
	PaymentID		uint64	`json:"payment_dt" validate:"required"`
	Bank			string	`json:"bank" validate:"required"`
	DeliveryCost	uint	`json:"delivery_cost" validate:"required"`
	GoodsTotal		uint	`json:"goods_total" validate:"required"`
	CustomFee		uint	`json:"custom_fee"`
}
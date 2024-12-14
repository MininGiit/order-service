package order

type Item struct {
	ChrtID		uint64	`json:"chrt_id" validate:"required"`
	TrackNumber	string	`json:"track_number" validate:"required"`
	Price 		uint	`json:"price" validate:"required"`
	Rid			string	`json:"rid" validate:"required"`	
	Name		string	`json:"name" validate:"required"`
	Sale		uint	`json:"sale" validate:"required"`
	Size		string	`json:"size" validate:"required"`
	TotalPrice	uint	`json:"total_price" validate:"required"`
	NmID		uint	`json:"nm_id" validate:"required"`
	Brand		string	`json:"brand" validate:"required"`
	Status		uint	`json:"status" validate:"required"`
}

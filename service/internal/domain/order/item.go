package order

type Item struct {
	ChrtID		uint64	`json:"chrt_id"`
	TrackNumber	string	`json:"track_number"`
	Price 		uint	`json:"price"`
	Rid			string	`json:"rid"`	
	Name		string	`json:"name"`
	Sale		string	`json:"sale"`
	Size		uint	`json:"size"`
	TotalPrice	uint	`json:"total_price"`
	NmID		uint	`json:"nm_id"`
	Brand		string	`json:"brand"`
	Status		uint	`json:"status"`
}

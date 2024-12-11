package order


type Delivery struct {
	Name	string	`json:"name"`
	Phone	string	`json:"phone"`
	Zip		string	`json:"zip"`
	City	string	`json:"city"`
	Address	string	`json:"address"`
	Region	string	`json:"region"`
	Email	string	`json:"email"`
}

func NewDelivery(name, 
	phone, 
	zip,
	city, 
	address, 
	region, 
	email string) *Delivery {
		return &Delivery{
			Name: 	name,
			Phone: 	phone,
			Zip:	zip,
			City:	city,
			Address: address,
			Region: region,
			Email: email,
		}
}
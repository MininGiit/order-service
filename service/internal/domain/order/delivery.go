package order


type Delivery struct {
	Name	string	`json:"name"`
	Phone	string	`json:"phone"`
	Zip		string	`json:"zip"`
	City	string	`json:"city"`
	Addres	string	`json:"addres"`
	Region	string	`json:"region"`
	Email	string	`json:"email"`
}

func NewDelivery(name, 
	phone, 
	zip,
	city, 
	addres, 
	region, 
	email string) *Delivery {
		return &Delivery{
			Name: 	name,
			Phone: 	phone,
			Zip:	zip,
			City:	city,
			Addres: addres,
			Region: region,
			Email: email,
		}
}
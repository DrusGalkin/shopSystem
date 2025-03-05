package shop_system

type category string

type Product struct {
	name     string
	price    uint
	count    uint
	Category category
}

func NewProduct(Name string, Price uint, Count uint, Cat category) *Product {
	return &Product{
		name:     Name,
		count:    Count,
		price:    Price,
		Category: Cat,
	}
}

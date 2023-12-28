package product

import "time"

type product struct {
	id          int
	name        string
	description string
	price       float32
	SKU         string
	createdOn   string
	UpdatedOn   string
	DeletedOn   string
}

func getProducts() []*product {

	return product_list
}

var product_list = []*product{
	{
		id:          1,
		name:        "expresso",
		description: "black coffee",
		price:       2.45,
		SKU:         "abc123",
		createdOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   ""},

	{id: 2,
		name:        "latte",
		description: "Forthy shit coffee",
		price:       3.5,
		SKU:         "abcd12",
		createdOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   ""},
}

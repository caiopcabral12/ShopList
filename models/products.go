package models

import (
	db "WishList/database"
)

type Products struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func SearchProducts() []Products {
	db := db.ConnectDB()

	selectAllProducts, err := db.Query("Select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Products{}
	products := []Products{}

	for selectAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func CreateProduct(name, description string, price float64, amount int) {
	db := db.ConnectDB()

	createProducts, err := db.Prepare("INSERT INTO products (name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	createProducts.Exec(
		name,
		description,
		price,
		amount)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProducts, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProducts.Exec(id)
	defer db.Close()
}

/*func EditProduct(id, name, description string, price float64, amount int) {
	db := db.ConnectDB()

	editProducts, err := db.Prepare("UPDATE products  SET (name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}


	defer db.Close()

}*/

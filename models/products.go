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

	selectAllProducts, err := db.Query("Select * from products ORDER BY id ASC")
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

func EditProduct(id string) Products {
	db := db.ConnectDB()

	showProducts, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	ed := Products{}

	for showProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = showProducts.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}
		ed.Id = id
		ed.Name = name
		ed.Description = description
		ed.Price = price
		ed.Amount = amount
	}
	defer db.Close()
	return ed

}
func UpdateProduct(id int, name, description string, price float64, amount int) {
	db := db.ConnectDB()

	updateProducts, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, amount=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProducts.Exec(
		name,
		description,
		price,
		amount,
		id)

	defer db.Close()
}

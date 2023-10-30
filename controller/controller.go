package controller

import (
	pd "WishList/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := pd.SearchProducts()
	temp.ExecuteTemplate(w, "index", allProducts)

}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "newProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			if err != nil {
				log.Println("Error while converting price!", err)
			}
		}

		amountConv, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error while converting Amount!", err)
		}

		pd.CreateProduct(name, description, priceConv, amountConv)

	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	getID := r.URL.Query().Get("id")
	pd.DeleteProduct(getID)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "edit", nil)
}

/*func Edit(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	amount := r.FormValue("amount")

	priceConv, err := strconv.ParseFloat(price, 64)
	if err != nil {
		if err != nil {
			log.Println("Error while converting price!", err)
		}
	}

	amountConv, err := strconv.Atoi(amount)
	if err != nil {
		log.Println("Error while converting Amount!", err)
	}

	getID := r.URL.Query().Get("id")
	pd.EditProduct(getID, name, description, priceConv, amountConv)
	http.Redirect(w, r, "/", 301)
}
*/

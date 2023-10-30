package routes

import (
	ct "WishList/controller"
	"net/http"
)

func Route() {
	http.HandleFunc("/", ct.Index)
	http.HandleFunc("/newProduct", ct.NewProduct)
	http.HandleFunc("/insert", ct.Insert)

}

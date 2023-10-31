package routes

import (
	ct "WishList/controller"
	"net/http"
)

func Route() {
	http.HandleFunc("/", ct.Index)
	http.HandleFunc("/newProduct", ct.NewProduct)
	http.HandleFunc("/insert", ct.Insert)
	http.HandleFunc("/delete", ct.Delete)
	http.HandleFunc("/edit", ct.Edit)
	http.HandleFunc("/update", ct.Update)
}

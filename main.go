package main

import (
	rt "WishList/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	rt.Route()
	http.ListenAndServe(":3333", nil)

}

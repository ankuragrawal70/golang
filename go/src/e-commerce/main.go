package main

import "net/http"
import "e-commerce/store"

func main() {
	//port := 8000
	http.HandleFunc("/", store.AllProducts)
	http.HandleFunc("/product/", store.ProductById)
	http.ListenAndServe(":7000", nil)

}







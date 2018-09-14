package store

import (
	"net/http"
	"encoding/json"
)

func AllProducts(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(GetAllProducts())
}


func ProductById(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(FilterProductById(2))
}



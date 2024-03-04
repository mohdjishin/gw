package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// getProducts handles HTTP requests and returns a list of products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("getProducts called...")
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 1200.99},
		{ID: 2, Name: "Smartphone", Price: 999.99},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

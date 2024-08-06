package main

import (
	"api_products/internal/db"
	"api_products/pkg/discount"
	"api_products/pkg/pagination"
	"api_products/pkg/product"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var DB = db.New("internal/data")

func main() {
	http.HandleFunc("/products", getProducts)
	fmt.Println("Server is running on port :3040")
	log.Fatal(http.ListenAndServe(":3040", nil))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	category := getStringParam("category", r, "")
	priceLessThan := getUintParam("priceLessThan", r, 0)
	pageNum := getUintParam("page", r, 1)
	pageSize := getUintParam("page_size", r, 5)

	filter := db.FilterParams{
		Category:      category,
		PriceLessThen: uint(priceLessThan),

		PageNum:  uint(pageNum),
		PageSize: uint(pageSize),
	}
	products := DB.FindProducts(filter)
	log.Printf("found products: %d\n", len(products))
	discounts := DB.GetDiscounts()
	productsOutput := make([]product.ProductResponse, len(products))
	for i, p := range products {
		log.Printf("Found product %+v", p)
		price := discount.GetPrices(discounts, &p)
		productOut := product.ProductResponse{
			SKU:      p.SKU,
			Name:     p.Name,
			Category: p.Category,
			Price:    price,
		}
		productsOutput[i] = productOut

	}
	log.Printf("Ouput products: %d; %+v\n", len(productsOutput), productsOutput)
	output := product.ProductListResponse{
		Products:   productsOutput,
		Pagination: pagination.Pagination{PageSize: uint(pageSize), PageNum: uint(pageNum)},
	}
	log.Printf("Output: %+v", output)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func getUintParam(name string, r *http.Request, default_value int) int {

	param := r.URL.Query().Get(name)
	if param == "" {
		return default_value
	}
	if value, err := strconv.Atoi(param); err == nil && value >= 0 {
		return value
	}
	return default_value
}
func getStringParam(name string, r *http.Request, default_value string) string {

	param := r.URL.Query().Get(name)
	if param != "" {
		return param
	}
	return default_value
}

package product

import (
	"api_products/pkg/pagination"
	"api_products/pkg/price"
)

type Product struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    uint   `json:"price"`
}
type ProductResponse struct {
	SKU      string      `json:"sku"`
	Name     string      `json:"name"`
	Category string      `json:"category"`
	Price    price.Price `json:"price"`
}
type ProductListResponse struct {
	Products   []ProductResponse     `json:"products"`
	Pagination pagination.Pagination `json:"pagination"`
}

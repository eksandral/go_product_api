package db

import (
	"api_products/pkg/discount"
	"api_products/pkg/product"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DB struct {
	Products  []product.Product
	Discounts []discount.Discount
}
type FilterParams struct {
	Category      string
	PriceLessThen uint
	PageNum       uint
	PageSize      uint
}

func New(dirname string) *DB {
	products_file, discounts_file := dirname+"/products.json", dirname+"/discounts.json"
	products, err := loadProducts(products_file)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}
	var discounts []discount.Discount
	discounts, err = loadDiscounts(discounts_file)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}
	return &DB{Products: products, Discounts: discounts}
}

func loadProducts(filename string) ([]product.Product, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var out []product.Product
	err = json.Unmarshal(data, &out)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	return out, nil
}
func loadDiscounts(filename string) ([]discount.Discount, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var out []discount.Discount
	err = json.Unmarshal(data, &out)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	return out, nil
}
func (self *DB) FindProducts(filter FilterParams) []product.Product {
	log.Printf("%+v\n", filter)
	var out []product.Product
	for _, p := range self.Products {
		if filter.Category == "" || p.Category != filter.Category {
			continue
		}
		if filter.PriceLessThen > 0 && p.Price > filter.PriceLessThen {
			continue
		}

		out = append(out, p)
	}
	if len(out) == 0 {
		return out
	}
	pages := uint(len(out)/int(filter.PageSize) + 1)
	currentPage := filter.PageNum
	if filter.PageNum > pages {
		currentPage = uint(pages)
	}
	start, end := (currentPage-1)*filter.PageSize, currentPage*filter.PageSize
	log.Printf("%v, %v; %d\n", start, end, len(out))
	return out[start:end]

}
func (self *DB) GetDiscounts() *[]discount.Discount {
	return &self.Discounts
}

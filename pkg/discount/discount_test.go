package discount

import (
	"api_products/pkg/price"
	"api_products/pkg/product"
	"testing"
)

func TestAppliedDiscount(t *testing.T) {
	p := product.Product{
		Name:     "1",
		SKU:      "1",
		Price:    200,
		Category: "1",
	}
	d := Discount{
		Kind:      "category",
		Predicate: "1",
		Value:     25,
	}
	ds := []Discount{d}
	calculatedPrice := GetPrices(&ds, &p)
	if calculatedPrice.Original != p.Price {
		t.Errorf("Expected %d price, got %d", p.Price, calculatedPrice.Original)
	}
	finalPrice := p.Price - p.Price*d.Value/100
	if calculatedPrice.Final != finalPrice {
		t.Errorf("Expected final price %d , got %d", finalPrice, calculatedPrice.Final)
	}
	if calculatedPrice.DiscountPercentage != price.DiscountPercentage(d.Value) {
		t.Errorf("Expected DiscountPercentage %d , got %d", d.Value, calculatedPrice.DiscountPercentage)
	}
	if calculatedPrice.Currency != "EUR" {
		t.Errorf("Expected Currency %s , got %s", "EUR", calculatedPrice.Currency)
	}
}
func TestNoDiscount(t *testing.T) {
	p := product.Product{
		Name:     "1",
		SKU:      "1",
		Price:    200,
		Category: "1",
	}
	d := Discount{
		Kind:      "category",
		Predicate: "2",
		Value:     25,
	}
	ds := []Discount{d}
	calculatedPrice := GetPrices(&ds, &p)
	if calculatedPrice.Original != p.Price {
		t.Errorf("Expected %d price, got %d", p.Price, calculatedPrice.Original)
	}
	if calculatedPrice.Final != p.Price {
		t.Errorf("Expected final price %d , got %d", p.Price, calculatedPrice.Final)
	}
	if calculatedPrice.DiscountPercentage != 0 {
		t.Errorf("Expected DiscountPercentage %d , got %d", d.Value, calculatedPrice.DiscountPercentage)
	}
	if calculatedPrice.Currency != "EUR" {
		t.Errorf("Expected Currency %s , got %s", "EUR", calculatedPrice.Currency)
	}
}

func TestMaxDiscount(t *testing.T) {
	p := product.Product{
		Name:     "1",
		SKU:      "1",
		Price:    200,
		Category: "1",
	}
	d1 := Discount{
		Kind:      "category",
		Predicate: "1",
		Value:     25,
	}
	d2 := Discount{
		Kind:      "sku",
		Predicate: "1",
		Value:     30,
	}
	ds := []Discount{d1, d2}
	calculatedPrice := GetPrices(&ds, &p)
	if calculatedPrice.Original != p.Price {
		t.Errorf("Expected %d price, got %d", p.Price, calculatedPrice.Original)
	}
	finalPrice := p.Price - p.Price*d2.Value/100
	if calculatedPrice.Final != finalPrice {
		t.Errorf("Expected final price %d , got %d", finalPrice, calculatedPrice.Final)
	}
	if calculatedPrice.DiscountPercentage != price.DiscountPercentage(d2.Value) {
		t.Errorf("Expected DiscountPercentage %d , got %d", d2.Value, calculatedPrice.DiscountPercentage)
	}
	if calculatedPrice.Currency != "EUR" {
		t.Errorf("Expected Currency %s , got %s", "EUR", calculatedPrice.Currency)
	}
}

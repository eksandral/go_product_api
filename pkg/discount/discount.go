package discount

import (
	"api_products/pkg/price"
	"api_products/pkg/product"
)

type Discount struct {
	Kind      string `json:"type"`
	Predicate string `json:"predicate"`
	Value     uint   `json:"value"`
}

func GetPrices(discounts *[]Discount, product *product.Product) price.Price {
	discount := uint(0)
	for _, d := range *discounts {
		switch d.Kind {
		case "category":
			if d.Predicate == product.Category && discount < d.Value {
				discount = d.Value
			}

			continue
		case "sku":
			if d.Predicate == product.SKU && discount < d.Value {
				discount = d.Value
			}

			continue
		}
	}
	finalPrice := product.Price - discount*product.Price/100

	return price.Price{
		Original:           product.Price,
		Final:              finalPrice,
		DiscountPercentage: price.DiscountPercentage(discount),
		Currency:           "EUR",
	}

}

# Product REST API

This is simple implamentation of a product information API
- get list of products with `{sku,name,category,price:{origianl,final,discount_pct,currency}}` data
- pagination:
    - limit list by `page_size` query param or `5` by default
    - show page usinng `page` query param or `1` by default

- filter by: 
    - `category` query param, *required*, it will show empty list if no category specified
    - `priceLessThan` query param, _optional_, it will filter by original price, but *not* discounted


## how to use
Run server by `go run cmd/server/server.go` command.
run a request `curl "http://127.0.0.1:3040/products?category=boots"` or in you favorite browser or in postman

## Configuration
The list of available products is in `interanl/data/products.json` file
The list of available discounts is in `internal/data/discounts.json` file



package dto

type Product struct {
	Id          string  `json:"id"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	LeftInStock int     `json:"left_in_stock"`
}

type BasicResponse struct{}

type GetProductRequest struct {
	Id       string `path:"product_id"`
	Currency string `query:"currency"`
}

type GetProductsRequest struct {
	Limit    int64  `query:"limit"`
	Offset   int64  `query:"offset"`
	Currency string `query:"currency"`
}

type CreateProductRequest struct {
	Description string  `json:"description"`
	PriceInUsd  float64 `json:"price_in_usd"`
	LeftInStock int     `json:"left_in_stock"`
}

type UpdateProductRequest struct {
	Id          string  `path:"product_id"`
	Description string  `json:"description,omitempty"`
	PriceInUsd  float64 `json:"price_in_usd"`
	LeftInStock int     `json:"left_in_stock"`
}

type DeleteProductRequest struct {
	Id string `path:"product_id"`
}

type CreateProductResponse struct {
	Product
}

type UpdateProductResponse struct {
	Product
}

type DeleteProductResponse struct {
	BasicResponse
}

type GetProductResponse struct {
	Product
}

type GetProductsResponse struct {
	Products []Product `json:"products"`
	Length   int64     `json:"length"`
}

package orderProducts

import (
	"time"
)

type OrderProducts struct {
	OrderId   string    `db:"order_id"`
	ProductId string    `db:"product_id"`
	CreatedAt time.Time `db:"created_at"`
}

type BasicResponse struct{}

type AddOrderProductsRequest struct {
	OrderId   string
	ProductId string
}
type AddOrderProductsResponse struct {
	BasicResponse
}
type DeleteOrderProductsRequest struct {
	OrderId   string
	ProductId string
}
type DeleteOrderProductsResponse struct {
	BasicResponse
}
type GetOrderProductsRequest struct {
	OrderId string
	Limit   int64
	Offset  int64
}
type GetOrderProductsResponse struct {
	OrderProducts *[]OrderProducts
	Length        int64
}

package dto

import (
	"github.com/rinatkh/test_2022/internal/products/models/dto"
	"time"
)

type Order struct {
	Id        string        `json:"uuid"`
	Products  []dto.Product `json:"products,omitempty"`
	Length    int64         `json:"length"`
	Price     float64       `json:"price"`
	Currency  string        `json:"currency"`
	CreatedAt time.Time     `json:"created_at"`
}

type BasicResponse struct{}

type GetOrderRequest struct {
	Id             string `path:"order_id"`
	UserId         string `query:"user_id"`
	Currency       string `query:"currency"`
	LimitProducts  int64  `query:"limit_products"`
	OffsetProducts int64  `query:"offset_products"`
}

type GetOrdersRequest struct {
	UserId         string `query:"user_id"`
	Limit          int64  `query:"limit"`
	Offset         int64  `query:"offset"`
	Currency       string `query:"currency"`
	LimitProducts  int64  `query:"limit_products"`
	OffsetProducts int64  `query:"offset_products"`
}

type CreateOrderRequest struct {
	UserId     string   `json:"user_id"`
	ProductIDs []string `json:"products_ids,omitempty"`
	Currency   string   `json:"currency"`
}

type UpdateOrderRequest struct {
	Id        string `path:"order_id"`
	UserId    string `json:"user_id"`
	ProductID string `json:"products_id"`
	Currency  string `json:"currency"`
}

type DeleteOrderRequest struct {
	Id     string `path:"order_id"`
	UserId string `query:"user_id"`
}

type CreateOrderResponse struct {
	Order
}

type UpdateOrderResponse struct {
	Order
}

type DeleteOrderResponse struct {
	BasicResponse
}

type GetOrderResponse struct {
	Order
}

type GetOrdersResponse struct {
	Orders []Order `json:"Orders"`
	Length int64   `json:"length"`
}

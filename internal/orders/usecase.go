package orders

import "github.com/rinatkh/test_2022/internal/orders/models/dto"

type UseCase interface {
	CreateOrder(params *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error)
	AddOrderProducts(params *dto.UpdateOrderRequest) (*dto.UpdateOrderResponse, error)
	DeleteOrderProducts(params *dto.UpdateOrderRequest) (*dto.UpdateOrderResponse, error)
	DeleteOrder(params *dto.DeleteOrderRequest) (*dto.DeleteOrderResponse, error)
	GetOrder(params *dto.GetOrderRequest) (*dto.GetOrderResponse, error)
	GetOrders(params *dto.GetOrdersRequest) (*dto.GetOrdersResponse, error)
}

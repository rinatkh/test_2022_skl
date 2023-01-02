package convert

import (
	"github.com/rinatkh/test_2022/internal/orders/models/core"
	ordersDTO "github.com/rinatkh/test_2022/internal/orders/models/dto"
	productsDTO "github.com/rinatkh/test_2022/internal/products/models/dto"
)

func Order2DTO(order *core.Order, products *[]productsDTO.Product, length int64, price float64, currency string) ordersDTO.Order {
	result := ordersDTO.Order{
		Id:        order.Id,
		Products:  *products,
		Length:    length,
		Price:     price,
		Currency:  currency,
		CreatedAt: order.CreatedAt,
	}
	return result
}

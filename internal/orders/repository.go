package orders

import "github.com/rinatkh/test_2022/internal/orders/models/core"

type OrderRepository interface {
	GetOrder(id, userId string) (*core.Order, error)
	GetOrders(userId string, limit, offset int64) (*[]core.Order, int64, error)
	DeleteOrder(order *core.Order) error
	CreateOrder(order *core.Order) (*core.Order, error)
}

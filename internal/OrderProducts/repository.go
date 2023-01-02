package orderProducts

type OrderProductsRepository interface {
	GetOrderProducts(orderId string, limit, offset int64) (*[]OrderProducts, int64, error)
	DeleteOrderProducts(orderId, productId string) error
	AddOrderProducts(orderId, productId string) error
}

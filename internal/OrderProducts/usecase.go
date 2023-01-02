package orderProducts

type UseCase interface {
	AddOrderProducts(params *AddOrderProductsRequest) (*AddOrderProductsResponse, error)
	DeleteOrderProducts(params *DeleteOrderProductsRequest) (*DeleteOrderProductsResponse, error)
	GetOrderProducts(params *GetOrderProductsRequest) (*GetOrderProductsResponse, error)
}

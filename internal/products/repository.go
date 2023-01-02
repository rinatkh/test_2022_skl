package products

import "github.com/rinatkh/test_2022/internal/products/models/core"

type ProductRepository interface {
	GetProductById(id string) (*core.Product, error)
	GetProducts(limit, offset int64) (*[]core.Product, int64, error)
	DeleteProduct(id string) error
	CreateProduct(product *core.Product) (*core.Product, error)
	UpdateProduct(product *core.Product) (*core.Product, error)
}

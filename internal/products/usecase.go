package products

import "github.com/rinatkh/test_2022/internal/products/models/dto"

type UseCase interface {
	CreateProduct(params *dto.CreateProductRequest) (*dto.CreateProductResponse, error)
	UpdateProduct(params *dto.UpdateProductRequest) (*dto.UpdateProductResponse, error)
	DeleteProduct(params *dto.DeleteProductRequest) (*dto.DeleteProductResponse, error)
	GetProduct(params *dto.GetProductRequest) (*dto.GetProductResponse, error)
	GetProducts(params *dto.GetProductsRequest) (*dto.GetProductsResponse, error)
}

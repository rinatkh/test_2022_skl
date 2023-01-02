package usecase

import (
	"github.com/rinatkh/test_2022/config"
	consts "github.com/rinatkh/test_2022/internal/constants"
	"github.com/rinatkh/test_2022/internal/currency"
	"github.com/rinatkh/test_2022/internal/products"
	"github.com/rinatkh/test_2022/internal/products/models/convert"
	"github.com/rinatkh/test_2022/internal/products/models/core"
	"github.com/rinatkh/test_2022/internal/products/models/dto"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/sirupsen/logrus"
)

type ProductUseCase struct {
	cfg         *config.Config
	log         *logrus.Entry
	repoProduct products.ProductRepository
	currencyUC  currency.UseCase
}

func NewProductUC(cfg *config.Config, log *logrus.Entry, repoProduct products.ProductRepository, currencyUC currency.UseCase) products.UseCase {
	return &ProductUseCase{
		cfg:         cfg,
		log:         log,
		repoProduct: repoProduct,
		currencyUC:  currencyUC,
	}
}

func (u ProductUseCase) CreateProduct(params *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	product := core.Product{
		Description: params.Description,
		PriceInUsd:  params.PriceInUsd,
		LeftInStock: params.LeftInStock,
	}
	result, err := u.repoProduct.CreateProduct(&product)
	if err != nil {
		return nil, err
	}
	return &dto.CreateProductResponse{Product: convert.Product2DTO(result, consts.USD, 1)}, nil
}

func (u ProductUseCase) UpdateProduct(params *dto.UpdateProductRequest) (*dto.UpdateProductResponse, error) {
	check, err := u.repoProduct.GetProductById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrProductDBNotFound
	}

	product := core.Product{
		Id:          params.Id,
		Description: params.Description,
		PriceInUsd:  params.PriceInUsd,
		LeftInStock: params.LeftInStock,
	}
	result, err := u.repoProduct.UpdateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateProductResponse{Product: convert.Product2DTO(result, consts.USD, 1)}, nil
}

func (u ProductUseCase) DeleteProduct(params *dto.DeleteProductRequest) (*dto.DeleteProductResponse, error) {
	check, err := u.repoProduct.GetProductById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrProductDBNotFound
	}
	err = u.repoProduct.DeleteProduct(params.Id)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteProductResponse{}, nil
}

func (u ProductUseCase) GetProduct(params *dto.GetProductRequest) (*dto.GetProductResponse, error) {
	result, err := u.repoProduct.GetProductById(params.Id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, constants.ErrProductDBNotFound
	}
	value, err := u.currencyUC.GetCurrency(&currency.GetCurrencyRequest{Name: params.Currency})
	if err != nil {
		return nil, constants.ErrCurrencyDBNotFound
	}
	return &dto.GetProductResponse{Product: convert.Product2DTO(result, params.Currency, value.CourseToUsd)}, nil
}

func (u ProductUseCase) GetProducts(params *dto.GetProductsRequest) (*dto.GetProductsResponse, error) {
	value, err := u.currencyUC.GetCurrency(&currency.GetCurrencyRequest{Name: params.Currency})
	if err != nil {
		return nil, constants.ErrCurrencyDBNotFound
	}
	list, length, err := u.repoProduct.GetProducts(params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	var result []dto.Product
	for _, i := range *list {
		result = append(result, convert.Product2DTO(&i, params.Currency, value.CourseToUsd))
	}
	return &dto.GetProductsResponse{Products: result, Length: length}, nil
}

package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/config"
	"github.com/rinatkh/test_2022/internal/OrderProducts"
	consts "github.com/rinatkh/test_2022/internal/constants"
	"github.com/rinatkh/test_2022/internal/products"
	"github.com/rinatkh/test_2022/internal/products/models/dto"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/sirupsen/logrus"
)

type OrderProductsUseCase struct {
	cfg               *config.Config
	log               *logrus.Entry
	repoOrderProducts orderProducts.OrderProductsRepository
	productsUC        products.UseCase
}

func NewOrderProductsUC(cfg *config.Config, log *logrus.Entry, repoOrderProducts orderProducts.OrderProductsRepository, productsUC products.UseCase) orderProducts.UseCase {
	return &OrderProductsUseCase{
		cfg:               cfg,
		log:               log,
		repoOrderProducts: repoOrderProducts,
		productsUC:        productsUC,
	}
}

func (u OrderProductsUseCase) AddOrderProducts(params *orderProducts.AddOrderProductsRequest) (*orderProducts.AddOrderProductsResponse, error) {
	product, err := u.productsUC.GetProduct(&dto.GetProductRequest{
		Id:       params.ProductId,
		Currency: consts.USD,
	})
	if err != nil {
		return nil, constants.ErrProductDBNotFound
	}
	if product.LeftInStock <= 0 {
		return nil, constants.NewCodedError(fmt.Sprintf("Can not add product '%s'", product.Description), fiber.StatusConflict)
	}
	_, err = u.productsUC.UpdateProduct(&dto.UpdateProductRequest{
		Id:          product.Id,
		Description: product.Description,
		PriceInUsd:  product.Price,
		LeftInStock: product.LeftInStock - 1,
	})
	if err != nil {
		return nil, constants.NewCodedError(fmt.Sprintf("Can not reduce amount of product '%s'", product.Description), fiber.StatusConflict)
	}
	return &orderProducts.AddOrderProductsResponse{}, u.repoOrderProducts.AddOrderProducts(params.OrderId, params.ProductId)
}
func (u OrderProductsUseCase) DeleteOrderProducts(params *orderProducts.DeleteOrderProductsRequest) (*orderProducts.DeleteOrderProductsResponse, error) {
	product, err := u.productsUC.GetProduct(&dto.GetProductRequest{
		Id:       params.ProductId,
		Currency: consts.USD,
	})
	if err != nil {
		return nil, constants.ErrProductDBNotFound
	}
	_, err = u.productsUC.UpdateProduct(&dto.UpdateProductRequest{
		Id:          product.Id,
		Description: product.Description,
		PriceInUsd:  product.Price,
		LeftInStock: product.LeftInStock + 1,
	})
	if err != nil {
		return nil, constants.NewCodedError(fmt.Sprintf("Can not increase amount of product '%s'", product.Description), fiber.StatusConflict)
	}
	return &orderProducts.DeleteOrderProductsResponse{}, u.repoOrderProducts.AddOrderProducts(params.OrderId, params.ProductId)
}
func (u OrderProductsUseCase) GetOrderProducts(params *orderProducts.GetOrderProductsRequest) (*orderProducts.GetOrderProductsResponse, error) {
	res, length, err := u.repoOrderProducts.GetOrderProducts(params.OrderId, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	return &orderProducts.GetOrderProductsResponse{
		OrderProducts: res,
		Length:        length,
	}, nil
}

package convert

import (
	"github.com/rinatkh/test_2022/internal/products/models/core"
	"github.com/rinatkh/test_2022/internal/products/models/dto"
)

func Product2DTO(product *core.Product, currency string, course float64) dto.Product {
	result := dto.Product{
		Id:          product.Id,
		Description: product.Description,
		Price:       product.PriceInUsd * course,
		Currency:    currency,
		LeftInStock: product.LeftInStock,
	}
	return result
}

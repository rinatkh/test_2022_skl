package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/internal/products"
	"github.com/rinatkh/test_2022/internal/products/models/dto"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/rinatkh/test_2022/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ProductHandler struct {
	productUC products.UseCase
	log       *logrus.Entry
}

func NewProductHandler(productUC products.UseCase, log *logrus.Entry) *ProductHandler {
	return &ProductHandler{
		productUC: productUC,
		log:       log,
	}
}

func (u ProductHandler) DeleteProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.DeleteProductRequest
		params.Id = ctx.Params("product_id")
		data, err := u.productUC.DeleteProduct(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ProductHandler) UpdateProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.UpdateProductRequest
		params.Id = ctx.Params("product_id")
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.productUC.UpdateProduct(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ProductHandler) GetProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetProductRequest
		params.Id = ctx.Params("product_id")
		params.Currency = ctx.Query("currency")

		data, err := u.productUC.GetProduct(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ProductHandler) GetProducts() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetProductsRequest
		var err error
		params.Limit, err = strconv.ParseInt(ctx.Query("limit", "20"),
			10, 64)
		if err != nil {
			return err
		}
		params.Offset, err = strconv.ParseInt(ctx.Query("offset", "0"),
			10, 64)
		if err != nil {
			return err
		}
		params.Currency = ctx.Query("currency")
		data, err := u.productUC.GetProducts(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ProductHandler) CreateProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.CreateProductRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.productUC.CreateProduct(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/internal/orders"
	"github.com/rinatkh/test_2022/internal/orders/models/dto"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/rinatkh/test_2022/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type OrderHandler struct {
	orderUC orders.UseCase
	log     *logrus.Entry
}

func NewOrderHandler(orderUC orders.UseCase, log *logrus.Entry) *OrderHandler {
	return &OrderHandler{
		orderUC: orderUC,
		log:     log,
	}
}

func (u OrderHandler) DeleteOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.DeleteOrderRequest
		params.Id = ctx.Params("order_id")
		params.UserId = ctx.Query("user_id")
		data, err := u.orderUC.DeleteOrder(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u OrderHandler) AddOrderProducts() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.UpdateOrderRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}
		params.Id = ctx.Params("order_id")
		data, err := u.orderUC.AddOrderProducts(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
func (u OrderHandler) DeleteOrderProducts() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.UpdateOrderRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}
		params.Id = ctx.Params("order_id")
		data, err := u.orderUC.DeleteOrderProducts(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u OrderHandler) GetOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetOrderRequest
		var err error
		params.Id = ctx.Params("order_id")
		params.UserId = ctx.Query("user_id")
		params.Currency = ctx.Query("currency")
		params.LimitProducts, err = strconv.ParseInt(ctx.Query("limit_products", "20"),
			10, 64)
		if err != nil {
			return err
		}
		params.OffsetProducts, err = strconv.ParseInt(ctx.Query("offset_products", "0"),
			10, 64)
		if err != nil {
			return err
		}
		data, err := u.orderUC.GetOrder(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u OrderHandler) GetOrders() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetOrdersRequest
		var err error
		params.UserId = ctx.Query("user_id")
		params.Currency = ctx.Query("currency")
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
		params.LimitProducts, err = strconv.ParseInt(ctx.Query("limit_products", "20"),
			10, 64)
		if err != nil {
			return err
		}
		params.OffsetProducts, err = strconv.ParseInt(ctx.Query("offset_products", "0"),
			10, 64)
		if err != nil {
			return err
		}
		data, err := u.orderUC.GetOrders(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u OrderHandler) CreateOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.CreateOrderRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}
		data, err := u.orderUC.CreateOrder(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

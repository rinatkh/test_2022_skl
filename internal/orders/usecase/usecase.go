package usecase

import (
	"github.com/rinatkh/test_2022/config"
	orderProductDTO "github.com/rinatkh/test_2022/internal/OrderProducts"
	orderProducts "github.com/rinatkh/test_2022/internal/OrderProducts"
	"github.com/rinatkh/test_2022/internal/currency"
	"github.com/rinatkh/test_2022/internal/orders"
	"github.com/rinatkh/test_2022/internal/orders/models/convert"
	"github.com/rinatkh/test_2022/internal/orders/models/core"
	orderDTO "github.com/rinatkh/test_2022/internal/orders/models/dto"
	"github.com/rinatkh/test_2022/internal/products"
	productDTO "github.com/rinatkh/test_2022/internal/products/models/dto"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/sirupsen/logrus"
)

type OrderUseCase struct {
	cfg             *config.Config
	log             *logrus.Entry
	repoOrder       orders.OrderRepository
	currencyUC      currency.UseCase
	orderProductsUC orderProducts.UseCase
	productsUC      products.UseCase
}

func NewOrderUC(cfg *config.Config, log *logrus.Entry, repoOrder orders.OrderRepository, currencyUC currency.UseCase, orderProductsUC orderProducts.UseCase, productsUC products.UseCase) orders.UseCase {
	return &OrderUseCase{
		cfg:             cfg,
		log:             log,
		repoOrder:       repoOrder,
		currencyUC:      currencyUC,
		orderProductsUC: orderProductsUC,
		productsUC:      productsUC,
	}
}
func (u OrderUseCase) getProductsPrice(productIDs *[]orderProducts.OrderProducts, currency string) (*[]productDTO.Product, float64, error) {
	var elems []productDTO.Product
	var price float64
	for _, i := range *productIDs {
		product, err := u.productsUC.GetProduct(&productDTO.GetProductRequest{Id: i.ProductId, Currency: currency})
		if err != nil {
			return nil, 0, err
		}
		elems = append(elems, product.Product)
		price += product.Price
	}
	return &elems, price, nil
}

func (u OrderUseCase) CreateOrder(params *orderDTO.CreateOrderRequest) (*orderDTO.CreateOrderResponse, error) {
	order, err := u.repoOrder.CreateOrder(&core.Order{UserId: params.UserId})
	if err != nil {
		return nil, err
	}

	var elems []productDTO.Product
	var price float64
	for _, i := range params.ProductIDs {
		product, err := u.productsUC.GetProduct(&productDTO.GetProductRequest{Id: i, Currency: params.Currency})
		if err != nil {
			return nil, err
		}
		_, err = u.orderProductsUC.AddOrderProducts(&orderProducts.AddOrderProductsRequest{OrderId: order.Id, ProductId: i})
		if err != nil {
			return nil, err
		}
		elems = append(elems, product.Product)
		price += product.Price
	}
	return &orderDTO.CreateOrderResponse{Order: convert.Order2DTO(order, &elems, int64(len(elems)), price, params.Currency)}, nil
}

func (u OrderUseCase) AddOrderProducts(params *orderDTO.UpdateOrderRequest) (*orderDTO.UpdateOrderResponse, error) {
	order, err := u.repoOrder.GetOrder(params.Id, params.UserId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, constants.ErrOrderDBNotFound
	}
	_, err = u.orderProductsUC.AddOrderProducts(&orderProductDTO.AddOrderProductsRequest{OrderId: order.Id, ProductId: params.ProductID})
	if err != nil {
		return nil, err
	}

	elems, err := u.orderProductsUC.GetOrderProducts(&orderProductDTO.GetOrderProductsRequest{
		OrderId: order.Id,
		Limit:   0,
		Offset:  0,
	})
	list, price, err := u.getProductsPrice(elems.OrderProducts, params.Currency)
	if err != nil {
		return nil, err
	}

	return &orderDTO.UpdateOrderResponse{Order: convert.Order2DTO(order, list, elems.Length, price, params.Currency)}, nil
}

func (u OrderUseCase) DeleteOrderProducts(params *orderDTO.UpdateOrderRequest) (*orderDTO.UpdateOrderResponse, error) {
	order, err := u.repoOrder.GetOrder(params.Id, params.UserId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, constants.ErrOrderDBNotFound
	}
	_, err = u.orderProductsUC.DeleteOrderProducts(&orderProductDTO.DeleteOrderProductsRequest{OrderId: order.Id, ProductId: params.ProductID})
	if err != nil {
		return nil, err
	}

	elems, err := u.orderProductsUC.GetOrderProducts(&orderProductDTO.GetOrderProductsRequest{
		OrderId: order.Id,
		Limit:   0,
		Offset:  0,
	})
	list, price, err := u.getProductsPrice(elems.OrderProducts, params.Currency)
	if err != nil {
		return nil, err
	}

	return &orderDTO.UpdateOrderResponse{Order: convert.Order2DTO(order, list, elems.Length, price, params.Currency)}, nil
}

func (u OrderUseCase) DeleteOrder(params *orderDTO.DeleteOrderRequest) (*orderDTO.DeleteOrderResponse, error) {
	order, err := u.repoOrder.GetOrder(params.Id, params.UserId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, constants.ErrOrderDBNotFound
	}
	elems, err := u.orderProductsUC.GetOrderProducts(&orderProductDTO.GetOrderProductsRequest{
		OrderId: order.Id,
		Limit:   0,
		Offset:  0,
	})
	for _, i := range *elems.OrderProducts {
		_, err = u.orderProductsUC.DeleteOrderProducts(&orderProductDTO.DeleteOrderProductsRequest{OrderId: order.Id, ProductId: i.ProductId})
		if err != nil {
			return nil, err
		}
	}
	err = u.repoOrder.DeleteOrder(order)
	if err != nil {
		return nil, err
	}
	return &orderDTO.DeleteOrderResponse{}, nil
}

func (u OrderUseCase) GetOrder(params *orderDTO.GetOrderRequest) (*orderDTO.GetOrderResponse, error) {
	order, err := u.repoOrder.GetOrder(params.Id, params.UserId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, constants.ErrOrderDBNotFound
	}
	elems, err := u.orderProductsUC.GetOrderProducts(&orderProductDTO.GetOrderProductsRequest{
		OrderId: order.Id,
		Limit:   params.LimitProducts,
		Offset:  params.OffsetProducts,
	})
	list, price, err := u.getProductsPrice(elems.OrderProducts, params.Currency)
	if err != nil {
		return nil, err
	}
	return &orderDTO.GetOrderResponse{Order: convert.Order2DTO(order, list, elems.Length, price, params.Currency)}, nil
}

func (u OrderUseCase) GetOrders(params *orderDTO.GetOrdersRequest) (*orderDTO.GetOrdersResponse, error) {
	listOrders, lengthOrders, err := u.repoOrder.GetOrders(params.UserId, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	if listOrders == nil {
		return &orderDTO.GetOrdersResponse{}, nil
	}
	var result []orderDTO.Order
	for _, i := range *listOrders {
		elems, err := u.orderProductsUC.GetOrderProducts(&orderProductDTO.GetOrderProductsRequest{
			OrderId: i.Id,
			Limit:   params.LimitProducts,
			Offset:  params.OffsetProducts,
		})
		listProducts, price, err := u.getProductsPrice(elems.OrderProducts, params.Currency)
		if err != nil {
			return nil, err
		}
		result = append(result, convert.Order2DTO(&i, listProducts, elems.Length, price, params.Currency))
	}

	return &orderDTO.GetOrdersResponse{Orders: result, Length: lengthOrders}, nil

}

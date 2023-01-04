package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	orderProducts "github.com/rinatkh/test_2022/internal/OrderProducts"
	mock_orderProducts "github.com/rinatkh/test_2022/internal/OrderProducts/mocks"
	consts "github.com/rinatkh/test_2022/internal/constants"
	mock_currency "github.com/rinatkh/test_2022/internal/currency/mocks"
	"github.com/rinatkh/test_2022/internal/orders"
	mock_orders "github.com/rinatkh/test_2022/internal/orders/mocks"
	coreOrder "github.com/rinatkh/test_2022/internal/orders/models/core"
	orderDTO "github.com/rinatkh/test_2022/internal/orders/models/dto"
	productDTO "github.com/rinatkh/test_2022/internal/products/models/dto"

	"github.com/rinatkh/test_2022/internal/test"
	mock_products "github.com/rinatkh/test_2022/internal/users/products"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testRepo := mock_orders.NewMockOrderRepository(ctrl)
	testCurrencyUC := mock_currency.NewMockUseCase(ctrl)
	testOrderProductsUC := mock_orderProducts.NewMockUseCase(ctrl)
	testProductsUC := mock_products.NewMockUseCase(ctrl)
	dbOrderImpl := NewOrderUC(test.TestConfig(t), test.TestLogger(t), testRepo, testCurrencyUC, testOrderProductsUC, testProductsUC)

	type Input struct {
		Input *orderDTO.CreateOrderRequest
	}
	type InputCreateOrder struct {
		Order *coreOrder.Order
	}
	type OutputCreateOrder struct {
		Order *coreOrder.Order
		Err   error
	}
	type InputGetProduct struct {
		Input *productDTO.GetProductRequest
	}
	type OutputGetProduct struct {
		Res *productDTO.GetProductResponse
		Err error
	}
	type InputAddOrderProducts struct {
		Input *orderProducts.AddOrderProductsRequest
	}
	type OutputAddOrderProducts struct {
		Res *orderProducts.AddOrderProductsResponse
		Err error
	}
	type Output struct {
		Res *orderDTO.CreateOrderResponse
		Err error
	}

	tests := []struct {
		name                   string
		input                  Input
		inputCreateOrder       InputCreateOrder
		outputCreateOrder      OutputCreateOrder
		inputGetProduct        InputGetProduct
		outputGetProduct       OutputGetProduct
		inputAddOrderProducts  InputAddOrderProducts
		outputAddOrderProducts OutputAddOrderProducts
		output                 Output
	}{
		{
			name: "Error working with db creating order",
			input: Input{&orderDTO.CreateOrderRequest{
				UserId:     "1",
				ProductIDs: nil,
				Currency:   consts.USD,
			}},
			inputCreateOrder: InputCreateOrder{Order: &coreOrder.Order{
				UserId: "1",
			}},
			outputCreateOrder: OutputCreateOrder{
				Order: nil,
				Err:   constants.ErrDB,
			},
			output: Output{
				Res: nil,
				Err: constants.ErrDB,
			},
		},
		{
			name: "Don't found product in BD",
			input: Input{Input: &orderDTO.CreateOrderRequest{
				UserId:     "2",
				ProductIDs: []string{"1", "2"},
				Currency:   consts.USD,
			}},
			inputCreateOrder: InputCreateOrder{Order: &coreOrder.Order{
				UserId: "2",
			}},
			outputCreateOrder: OutputCreateOrder{
				Order: &coreOrder.Order{
					Id:        "2",
					UserId:    "2",
					CreatedAt: time.Time{},
				},
				Err: nil,
			},
			inputGetProduct: InputGetProduct{Input: &productDTO.GetProductRequest{
				Id:       "1",
				Currency: consts.USD,
			}},
			outputGetProduct: OutputGetProduct{
				Res: &productDTO.GetProductResponse{},
				Err: constants.ErrProductDBNotFound,
			},
			output: Output{
				Res: nil,
				Err: constants.ErrProductDBNotFound,
			},
		},
		{
			name: "can't add product in BD",
			input: Input{Input: &orderDTO.CreateOrderRequest{
				UserId:     "3",
				ProductIDs: []string{"3"},
				Currency:   consts.USD,
			}},
			inputCreateOrder: InputCreateOrder{Order: &coreOrder.Order{
				UserId: "3",
			}},
			outputCreateOrder: OutputCreateOrder{
				Order: &coreOrder.Order{
					Id:        "3",
					UserId:    "3",
					CreatedAt: time.Time{},
				},
				Err: nil,
			},
			inputGetProduct: InputGetProduct{Input: &productDTO.GetProductRequest{
				Id:       "3",
				Currency: consts.USD,
			}},
			outputGetProduct: OutputGetProduct{
				Res: &productDTO.GetProductResponse{Product: productDTO.Product{
					Id:          "3",
					Description: "APPLE",
					Price:       10,
					Currency:    consts.USD,
					LeftInStock: 1,
				}},
				Err: nil,
			},
			inputAddOrderProducts: InputAddOrderProducts{Input: &orderProducts.AddOrderProductsRequest{
				OrderId:   "3",
				ProductId: "3",
			}},
			outputAddOrderProducts: OutputAddOrderProducts{
				Res: nil,
				Err: constants.NewCodedError("Can not add product APPLE", fiber.StatusConflict),
			},
			output: Output{
				Res: nil,
				Err: constants.NewCodedError("Can not add product APPLE", fiber.StatusConflict),
			},
		},
		{
			name: "Success",
			input: Input{Input: &orderDTO.CreateOrderRequest{
				UserId:     "4",
				ProductIDs: []string{"3"},
				Currency:   consts.USD,
			}},
			inputCreateOrder: InputCreateOrder{Order: &coreOrder.Order{
				UserId: "4",
			}},
			outputCreateOrder: OutputCreateOrder{
				Order: &coreOrder.Order{
					Id:        "4",
					UserId:    "4",
					CreatedAt: time.Time{},
				},
				Err: nil,
			},
			inputGetProduct: InputGetProduct{Input: &productDTO.GetProductRequest{
				Id:       "3",
				Currency: consts.USD,
			}},
			outputGetProduct: OutputGetProduct{
				Res: &productDTO.GetProductResponse{Product: productDTO.Product{
					Id:          "3",
					Description: "APPLE",
					Price:       10,
					Currency:    consts.USD,
					LeftInStock: 1,
				}},
				Err: nil,
			},
			inputAddOrderProducts: InputAddOrderProducts{Input: &orderProducts.AddOrderProductsRequest{
				OrderId:   "4",
				ProductId: "3",
			}},
			outputAddOrderProducts: OutputAddOrderProducts{
				Res: &orderProducts.AddOrderProductsResponse{},
				Err: nil,
			},
			output: Output{
				Res: &orderDTO.CreateOrderResponse{Order: orderDTO.Order{
					Id: "4",
					Products: []productDTO.Product{
						{
							Id:          "3",
							Description: "APPLE",
							Price:       10,
							Currency:    consts.USD,
							LeftInStock: 1,
						},
					},
					Length:    1,
					Price:     10,
					Currency:  consts.USD,
					CreatedAt: time.Time{},
				}},
				Err: nil,
			},
		},
	}
	gomock.InOrder(
		testRepo.EXPECT().CreateOrder(tests[0].inputCreateOrder.Order).Return(tests[0].outputCreateOrder.Order, tests[0].outputCreateOrder.Err),

		testRepo.EXPECT().CreateOrder(tests[1].inputCreateOrder.Order).Return(tests[1].outputCreateOrder.Order, tests[1].outputCreateOrder.Err),
		testProductsUC.EXPECT().GetProduct(tests[1].inputGetProduct.Input).Return(tests[1].outputGetProduct.Res, tests[1].outputGetProduct.Err),

		testRepo.EXPECT().CreateOrder(tests[2].inputCreateOrder.Order).Return(tests[2].outputCreateOrder.Order, tests[2].outputCreateOrder.Err),
		testProductsUC.EXPECT().GetProduct(tests[2].inputGetProduct.Input).Return(tests[2].outputGetProduct.Res, tests[2].outputGetProduct.Err),
		testOrderProductsUC.EXPECT().AddOrderProducts(tests[2].inputAddOrderProducts.Input).Return(tests[2].outputAddOrderProducts.Res, tests[2].outputAddOrderProducts.Err),

		testRepo.EXPECT().CreateOrder(tests[3].inputCreateOrder.Order).Return(tests[3].outputCreateOrder.Order, tests[3].outputCreateOrder.Err),
		testProductsUC.EXPECT().GetProduct(tests[3].inputGetProduct.Input).Return(tests[3].outputGetProduct.Res, tests[3].outputGetProduct.Err),
		testOrderProductsUC.EXPECT().AddOrderProducts(tests[3].inputAddOrderProducts.Input).Return(tests[3].outputAddOrderProducts.Res, tests[3].outputAddOrderProducts.Err),
	)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			res, errRes := orders.UseCase.CreateOrder(dbOrderImpl, test.input.Input)
			if !assert.Equal(t, test.output.Res, res) {
				t.Error("got : ", res, " expected :", test.output.Res)
			}
			if !assert.Equal(t, test.output.Err, errRes) {
				t.Error("got : ", errRes, " expected :", test.output.Err)
			}
		})
	}
}

func TestAddOrderProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testRepo := mock_orders.NewMockOrderRepository(ctrl)
	testCurrencyUC := mock_currency.NewMockUseCase(ctrl)
	testOrderProductsUC := mock_orderProducts.NewMockUseCase(ctrl)
	testProductsUC := mock_products.NewMockUseCase(ctrl)
	dbOrderImpl := NewOrderUC(test.TestConfig(t), test.TestLogger(t), testRepo, testCurrencyUC, testOrderProductsUC, testProductsUC)

	type Input struct {
		Input *orderDTO.UpdateOrderRequest
	}
	type InputGetOrder struct {
		Id     string
		UserId string
	}
	type OutputGetOrder struct {
		Order *coreOrder.Order
		Err   error
	}
	type InputAddOrderProducts struct {
		Input *orderProducts.AddOrderProductsRequest
	}
	type OutputAddOrderProducts struct {
		Res *orderProducts.AddOrderProductsResponse
		Err error
	}
	type InputGetOrderProducts struct {
		Input *orderProducts.GetOrderProductsRequest
	}
	type OutputGetOrderProducts struct {
		Res *orderProducts.GetOrderProductsResponse
		Err error
	}
	type InputGetProduct struct {
		Input *productDTO.GetProductRequest
	}
	type OutputGetProduct struct {
		Res *productDTO.GetProductResponse
		Err error
	}
	type Output struct {
		Res *orderDTO.UpdateOrderResponse
		Err error
	}

	tests := []struct {
		name                   string
		input                  Input
		inputGetOrder          InputGetOrder
		outputGetOrder         OutputGetOrder
		inputAddOrderProducts  InputAddOrderProducts
		outputAddOrderProducts OutputAddOrderProducts
		inputGetOrderProducts  InputGetOrderProducts
		outputGetOrderProducts OutputGetOrderProducts
		inputGetProduct        InputGetProduct
		outputGetProduct       OutputGetProduct
		output                 Output
	}{
		{
			name: "Didn't found order in db",
			input: Input{Input: &orderDTO.UpdateOrderRequest{
				Id:        "1",
				UserId:    "1",
				ProductID: "1",
				Currency:  consts.USD,
			}},
			inputGetOrder: InputGetOrder{
				Id:     "1",
				UserId: "1",
			},
			outputGetOrder: OutputGetOrder{
				Order: nil,
				Err:   nil,
			},
			output: Output{
				Res: nil,
				Err: constants.ErrOrderDBNotFound,
			},
		},
		{
			name: "Error AddOrderProducts",
			input: Input{Input: &orderDTO.UpdateOrderRequest{
				Id:        "2",
				UserId:    "2",
				ProductID: "2",
				Currency:  consts.USD,
			}},
			inputGetOrder: InputGetOrder{
				Id:     "2",
				UserId: "2",
			},
			outputGetOrder: OutputGetOrder{
				Order: &coreOrder.Order{
					Id:        "2",
					UserId:    "2",
					CreatedAt: time.Time{},
				},
				Err: nil,
			},
			inputAddOrderProducts: InputAddOrderProducts{Input: &orderProducts.AddOrderProductsRequest{
				OrderId:   "2",
				ProductId: "2",
			}},
			outputAddOrderProducts: OutputAddOrderProducts{
				Res: nil,
				Err: constants.NewCodedError("Can not add product 2", fiber.StatusConflict),
			},
			output: Output{
				Res: nil,
				Err: constants.NewCodedError("Can not add product 2", fiber.StatusConflict),
			},
		},
		{
			name: "Error GetOrderProducts",
			input: Input{Input: &orderDTO.UpdateOrderRequest{
				Id:        "3",
				UserId:    "3",
				ProductID: "3",
				Currency:  consts.USD,
			}},
			inputGetOrder: InputGetOrder{
				Id:     "3",
				UserId: "3",
			},
			outputGetOrder: OutputGetOrder{
				Order: &coreOrder.Order{
					Id:        "3",
					UserId:    "3",
					CreatedAt: time.Time{},
				},
				Err: nil,
			},
			inputAddOrderProducts: InputAddOrderProducts{Input: &orderProducts.AddOrderProductsRequest{
				OrderId:   "3",
				ProductId: "3",
			}},
			outputAddOrderProducts: OutputAddOrderProducts{
				Res: &orderProducts.AddOrderProductsResponse{},
				Err: nil,
			},
			inputGetOrderProducts: InputGetOrderProducts{Input: &orderProducts.GetOrderProductsRequest{
				OrderId: "3",
				Limit:   0,
				Offset:  0,
			}},
			outputGetOrderProducts: OutputGetOrderProducts{
				Res: nil,
				Err: constants.NewCodedError("Can not add product '3'", fiber.StatusConflict),
			},
			output: Output{
				Res: nil,
				Err: constants.NewCodedError("Can not add product '3'", fiber.StatusConflict),
			},
		},
		{
			name: "Success",
			input: Input{Input: &orderDTO.UpdateOrderRequest{
				Id:        "4",
				UserId:    "4",
				ProductID: "4",
				Currency:  consts.USD,
			}},
			inputGetOrder: InputGetOrder{
				Id:     "4",
				UserId: "4",
			},
			outputGetOrder: OutputGetOrder{
				Order: &coreOrder.Order{
					Id:        "4",
					UserId:    "4",
					CreatedAt: time.Time{},
				},
				Err: nil,
			},
			inputAddOrderProducts: InputAddOrderProducts{Input: &orderProducts.AddOrderProductsRequest{
				OrderId:   "4",
				ProductId: "4",
			}},
			outputAddOrderProducts: OutputAddOrderProducts{
				Res: &orderProducts.AddOrderProductsResponse{},
				Err: nil,
			},
			inputGetOrderProducts: InputGetOrderProducts{Input: &orderProducts.GetOrderProductsRequest{
				OrderId: "4",
				Limit:   0,
				Offset:  0,
			}},
			outputGetOrderProducts: OutputGetOrderProducts{
				Res: &orderProducts.GetOrderProductsResponse{
					OrderProducts: &[]orderProducts.OrderProducts{{
						OrderId:   "4",
						ProductId: "4",
						CreatedAt: time.Time{},
					}},
					Length: 1,
				},
				Err: nil,
			},
			inputGetProduct: InputGetProduct{Input: &productDTO.GetProductRequest{
				Id:       "4",
				Currency: consts.USD,
			}},
			outputGetProduct: OutputGetProduct{
				Res: &productDTO.GetProductResponse{Product: productDTO.Product{
					Id:          "4",
					Description: "APPLE",
					Price:       10,
					Currency:    consts.USD,
					LeftInStock: 5,
				}},
				Err: nil,
			},
			output: Output{
				Res: &orderDTO.UpdateOrderResponse{Order: orderDTO.Order{
					Id: "4",
					Products: []productDTO.Product{{
						Id:          "4",
						Description: "APPLE",
						Price:       10,
						Currency:    consts.USD,
						LeftInStock: 5,
					}},
					Length:    1,
					Price:     10,
					Currency:  consts.USD,
					CreatedAt: time.Time{},
				}},
				Err: nil,
			},
		},
	}
	gomock.InOrder(
		testRepo.EXPECT().GetOrder(tests[0].inputGetOrder.Id, tests[0].inputGetOrder.UserId).Return(tests[0].outputGetOrder.Order, tests[0].outputGetOrder.Err),

		testRepo.EXPECT().GetOrder(tests[1].inputGetOrder.Id, tests[1].inputGetOrder.UserId).Return(tests[1].outputGetOrder.Order, tests[1].outputGetOrder.Err),
		testOrderProductsUC.EXPECT().AddOrderProducts(tests[1].inputAddOrderProducts.Input).Return(tests[1].outputAddOrderProducts.Res, tests[1].outputAddOrderProducts.Err),

		testRepo.EXPECT().GetOrder(tests[2].inputGetOrder.Id, tests[2].inputGetOrder.UserId).Return(tests[2].outputGetOrder.Order, tests[2].outputGetOrder.Err),
		testOrderProductsUC.EXPECT().AddOrderProducts(tests[2].inputAddOrderProducts.Input).Return(tests[2].outputAddOrderProducts.Res, tests[2].outputAddOrderProducts.Err),
		testOrderProductsUC.EXPECT().GetOrderProducts(tests[2].inputGetOrderProducts.Input).Return(tests[2].outputGetOrderProducts.Res, tests[2].outputGetOrderProducts.Err),

		testRepo.EXPECT().GetOrder(tests[3].inputGetOrder.Id, tests[3].inputGetOrder.UserId).Return(tests[3].outputGetOrder.Order, tests[3].outputGetOrder.Err),
		testOrderProductsUC.EXPECT().AddOrderProducts(tests[3].inputAddOrderProducts.Input).Return(tests[3].outputAddOrderProducts.Res, tests[3].outputAddOrderProducts.Err),
		testOrderProductsUC.EXPECT().GetOrderProducts(tests[3].inputGetOrderProducts.Input).Return(tests[3].outputGetOrderProducts.Res, tests[3].outputGetOrderProducts.Err),
		testProductsUC.EXPECT().GetProduct(tests[3].inputGetProduct.Input).Return(tests[3].outputGetProduct.Res, tests[3].outputGetProduct.Err),
	)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			res, errRes := orders.UseCase.AddOrderProducts(dbOrderImpl, test.input.Input)
			if !assert.Equal(t, test.output.Res, res) {
				t.Error("got : ", res, " expected :", test.output.Res)
			}
			if !assert.Equal(t, test.output.Err, errRes) {
				t.Error("got : ", errRes, " expected :", test.output.Err)
			}
		})
	}
}

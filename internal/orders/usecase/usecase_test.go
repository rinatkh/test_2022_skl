package usecase

import (
	"github.com/golang/mock/gomock"
	orderProducts "github.com/rinatkh/test_2022/internal/OrderProducts"
	mock_orderProducts "github.com/rinatkh/test_2022/internal/OrderProducts/mocks"
	mock_currency "github.com/rinatkh/test_2022/internal/currency/mocks"
	"github.com/rinatkh/test_2022/internal/orders"
	mock_orders "github.com/rinatkh/test_2022/internal/orders/mocks"
	coreOrder "github.com/rinatkh/test_2022/internal/orders/models/core"
	orderDTO "github.com/rinatkh/test_2022/internal/orders/models/dto"
	coreProduct "github.com/rinatkh/test_2022/internal/products/models/core"
	productDTO "github.com/rinatkh/test_2022/internal/products/models/dto"
	"github.com/rinatkh/test_2022/internal/test"
	mock_products "github.com/rinatkh/test_2022/internal/users/products"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/stretchr/testify/assert"
	"testing"
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
		Product *coreProduct.Product
		Err     error
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
			name: "Don't found in BD",
			input: Input{&orderDTO.CreateOrderRequest{
				UserId:     "1",
				ProductIDs: nil,
				Currency:   "USD",
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
	}
	gomock.InOrder(
		testRepo.EXPECT().CreateOrder(tests[0].inputCreateOrder.Order).Return(tests[0].outputCreateOrder.Order, tests[0].outputCreateOrder.Err),
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

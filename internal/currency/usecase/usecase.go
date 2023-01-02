package usecase

import (
	"github.com/rinatkh/test_2022/config"
	"github.com/rinatkh/test_2022/internal/currency"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/sirupsen/logrus"
)

type CurrencyUseCase struct {
	cfg          *config.Config
	log          *logrus.Entry
	repoCurrency currency.CurrencyRepository
}

func NewProductUC(cfg *config.Config, log *logrus.Entry, repoCurrency currency.CurrencyRepository) currency.UseCase {
	return &CurrencyUseCase{
		cfg:          cfg,
		log:          log,
		repoCurrency: repoCurrency,
	}
}

func (u CurrencyUseCase) GetCurrency(params *currency.GetCurrencyRequest) (*currency.GetCurrencyResponse, error) {
	result, err := u.repoCurrency.GetCurrencyByName(params.Name)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, constants.ErrCurrencyDBNotFound
	}

	return &currency.GetCurrencyResponse{Currency: *result}, nil
}

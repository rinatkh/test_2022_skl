package currency

type UseCase interface {
	GetCurrency(params *GetCurrencyRequest) (*GetCurrencyResponse, error)
}

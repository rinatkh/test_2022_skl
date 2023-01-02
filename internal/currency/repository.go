package currency

type CurrencyRepository interface {
	GetCurrencyByName(name string) (*Currency, error)
}

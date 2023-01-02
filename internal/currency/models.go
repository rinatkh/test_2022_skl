package currency

type Currency struct {
	Name        string  `json:"name" db:"name"`
	CourseToUsd float64 `json:"course_to_usd" db:"course_to_usd"`
}

type GetCurrencyRequest struct {
	Name string `json:"name"`
}
type GetCurrencyResponse struct {
	Currency
}

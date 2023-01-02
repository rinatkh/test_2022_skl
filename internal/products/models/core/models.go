package core

type Product struct {
	Id          string  `db:"uuid"`
	Description string  `db:"description"`
	PriceInUsd  float64 `db:"price_in_usd"`
	LeftInStock int     `db:"left_in_stock"`
}

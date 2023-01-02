package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/test_2022/internal/products"
	"github.com/rinatkh/test_2022/internal/products/models/core"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) products.ProductRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetProductById(id string) (*core.Product, error) {
	var data []core.Product
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM products WHERE uuid='%s'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) GetProducts(limit, offset int64) (*[]core.Product, int64, error) {
	var data []core.Product
	queryStr := "SELECT * FROM products WHERE uuid <> ''"
	if limit == 0 {
		queryStr += " LIMIT 1"
	} else {
		queryStr += fmt.Sprintf(" LIMIT %d", limit)
	}
	queryStr += fmt.Sprintf(" OFFSET %d", offset)

	err := p.db.Select(
		&data, queryStr)

	if err != nil {
		return nil, 0, err
	}
	if len(data) == 0 {
		return nil, 0, nil
	}
	var length []int64
	err = p.db.Select(&length, "SELECT count(*) FROM products")
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

func (p postgresRepository) CreateProduct(product *core.Product) (*core.Product, error) {
	query := fmt.Sprintf("INSERT INTO products (description, price_in_usd, left_in_stock) VALUES ('%s', '%f', '%d')", product.Description, product.PriceInUsd, product.LeftInStock)

	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getProduct(product)
}

func (p postgresRepository) UpdateProduct(product *core.Product) (*core.Product, error) {
	query := fmt.Sprintf("UPDATE Products SET description='%s', price_in_usd='%f', left_in_stock='%d' where uuid='%s'", product.Description, product.PriceInUsd, product.LeftInStock, product.Id)
	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getProduct(product)
}

func (p postgresRepository) DeleteProduct(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM products WHERE uuid='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getProduct(product *core.Product) (*core.Product, error) {
	var data []core.Product
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM products WHERE description='%s' AND price_in_usd='%f' AND left_in_stock='%d'", product.Description, product.PriceInUsd, product.LeftInStock))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}

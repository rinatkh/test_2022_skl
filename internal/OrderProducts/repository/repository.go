package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	orderProducts "github.com/rinatkh/test_2022/internal/OrderProducts"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) orderProducts.OrderProductsRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetOrderProducts(orderId string, limit, offset int64) (*[]orderProducts.OrderProducts, int64, error) {
	var data []orderProducts.OrderProducts
	queryStr := fmt.Sprintf("SELECT * FROM OrderProducts WHERE order_id='%s'", orderId)
	queryStr += " ORDER BY created_at DESC"
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
	err = p.db.Select(&length, fmt.Sprintf("SELECT count(*) FROM OrderProducts WHERE order_id='%s'", orderId))
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

func (p postgresRepository) AddOrderProducts(orderId, productId string) error {
	query := fmt.Sprintf("INSERT INTO OrderProducts (order_id, product_id) VALUES ('%s', '%s')", orderId, productId)

	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeleteOrderProducts(orderId, productId string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM OrderProducts WHERE order_id='%s' AND product_id='%s'", orderId, productId))

	if res != nil {
		_ = res.Close()
	}
	return err
}

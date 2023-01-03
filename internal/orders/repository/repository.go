package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/test_2022/internal/orders"
	"github.com/rinatkh/test_2022/internal/orders/models/core"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) orders.OrderRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetOrder(id, userId string) (*core.Order, error) {
	var data []core.Order
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM orders WHERE uuid='%s' AND user_id='%s'", id, userId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) GetOrders(userId string, limit, offset int64) (*[]core.Order, int64, error) {
	var data []core.Order
	queryStr := fmt.Sprintf("SELECT * FROM Orders WHERE user_id='%s'", userId)
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
	err = p.db.Select(&length, "SELECT count(*) FROM Orders")
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

func (p postgresRepository) CreateOrder(order *core.Order) (*core.Order, error) {
	query := fmt.Sprintf("INSERT INTO Orders (user_id) VALUES ('%s')", order.UserId)
	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getOrder(order)
}

func (p postgresRepository) DeleteOrder(order *core.Order) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM Orders WHERE uuid='%s' AND user_id='%s'", order.Id, order.UserId))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getOrder(order *core.Order) (*core.Order, error) {
	var data []core.Order
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM Orders WHERE user_id='%s' ORDER BY created_at DESC", order.UserId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}

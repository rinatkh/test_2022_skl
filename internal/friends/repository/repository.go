package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/test_2022/internal/friends"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) friends.FriendsRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}
func (p postgresRepository) IsFriends(first, second string) (bool, error) {
	var data []friends.Friends
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM friends WHERE first_user='%s' AND second_user='%s' OR first_user='%s' AND second_user='%s'", first, second, second, first))
	if err != nil {
		return false, err
	}
	if len(data) == 0 {
		return false, nil
	}
	return true, nil
}
func (p postgresRepository) AddFriend(first, second string) error {
	query := fmt.Sprintf("INSERT INTO friends (first_user, second_user) VALUES ('%s', '%s')", first, second)

	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeleteFriend(first, second string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM friends WHERE first_user='%s' AND second_user='%s'", first, second))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) GetFriends(id string, limit, offset int64) (*[]friends.Friends, int64, error) {
	var data []friends.Friends
	queryStr := fmt.Sprintf("SELECT * FROM friends WHERE first_user='%s' OR second_user='%s'", id, id)
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
	err = p.db.Select(&length, fmt.Sprintf("SELECT count(*) FROM friends WHERE first_user='%s' OR second_user='%s'", id, id))
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

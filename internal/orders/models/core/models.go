package core

import (
	"time"
)

type Order struct {
	Id        string    `db:"uuid"`
	UserId    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

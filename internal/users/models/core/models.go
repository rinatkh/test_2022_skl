package core

import "time"

type User struct {
	Id         string    `db:"uuid"`
	Firstname  string    `db:"firstname"`
	Surname    string    `db:"surname"`
	Middlename *string   `db:"middlename"`
	Fio        string    `db:"fio"`
	Sex        string    `db:"sex"`
	BirthDate  time.Time `db:"birth_date"`
}

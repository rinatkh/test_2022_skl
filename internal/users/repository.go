package users

import "github.com/rinatkh/test_2022/internal/users/models/core"

type UserRepository interface {
	GetUserById(id string) (*core.User, error)
	GetUsers(limit, offset int64) (*[]core.User, int64, error)
	DeleteUser(id string) error
	CreateUser(user *core.User) (*core.User, error)
	UpdateUser(user *core.User) (*core.User, error)
}
